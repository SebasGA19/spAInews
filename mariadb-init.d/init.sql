# noinspection SqlNoDataSourceInspection

DELIMITER @@

CREATE OR REPLACE FUNCTION
    generar_salt()
    RETURNS VARCHAR(1024)
    LANGUAGE SQL
    NOT DETERMINISTIC
BEGIN
    RETURN SHA2(RAND(), 512);
END;
@@

DELIMITER ;

CREATE TABLE IF NOT EXISTS usuarios
(
    id                        INTEGER UNSIGNED     NOT NULL PRIMARY KEY AUTO_INCREMENT,
    fecha_creacion            TIMESTAMP            NOT NULL DEFAULT CURRENT_TIMESTAMP,
    fecha_ultima_modificacion TIMESTAMP            NOT NULL DEFAULT CURRENT_TIMESTAMP,
    activo                    BOOL                 NOT NULL DEFAULT TRUE,
    correo                    VARCHAR(1024) UNIQUE NOT NULL,
    usuario                   VARCHAR(1024) UNIQUE NOT NULL,
    contrasena_salt           VARCHAR(1024)        NOT NULL,
    contrasena                VARCHAR(1024)        NOT NULL,
    CONSTRAINT usuarios_usuario_valido CHECK ( usuario RLIKE '^[a-zA-Z0-9._-]+$' ),
    CONSTRAINT usuarios_correo_valido CHECK ( correo RLIKE
                                              '^[a-zA-Z0-9][a-zA-Z0-9._-]*[a-zA-Z0-9]@[a-zA-Z0-9][a-zA-Z0-9._-]*[a-zA-Z0-9]\\.[a-zA-Z]{2,63}$' )
);

CREATE TABLE IF NOT EXISTS palabras_clave
(
    id                        INTEGER UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
    fecha_creacion            TIMESTAMP        NOT NULL DEFAULT CURRENT_TIMESTAMP,
    fecha_ultima_modificacion TIMESTAMP        NOT NULL DEFAULT CURRENT_TIMESTAMP,
    activo                    BOOL             NOT NULL DEFAULT TRUE,
    id_usuario                INTEGER UNSIGNED NOT NULL,
    automatico                BOOL             NOT NULL DEFAULT FALSE,
    palabras                  JSON             NOT NULL DEFAULT '[]',
    CONSTRAINT fk_palabras_clave_id_usuario FOREIGN KEY (id_usuario) REFERENCES usuarios (id)
);

DELIMITER @@

-- - Related to usuarios - --

CREATE OR REPLACE TRIGGER before_update_usuarios
    BEFORE UPDATE
    ON usuarios
    FOR EACH ROW
BEGIN
    IF NOT OLD.activo AND NOT NEW.activo THEN
        SIGNAL SQLSTATE '45000' SET MYSQL_ERRNO = 30001, MESSAGE_TEXT =
                'El usuario esta desactivo, reactivar primero antes de actualizar';
    END IF;
END;
@@

CREATE OR REPLACE TRIGGER after_insert_usuarios
    AFTER INSERT
    ON usuarios
    FOR EACH ROW
BEGIN
    INSERT INTO palabras_clave (id_usuario) VALUES (NEW.id);
END;
@@

CREATE OR REPLACE PROCEDURE
    registrar_usuario(
    v_usuario VARCHAR(1024),
    v_correo VARCHAR(1024),
    v_contrasena VARCHAR(1024)
)
    LANGUAGE SQL
    NOT DETERMINISTIC
BEGIN
    DECLARE salt VARCHAR(1024);
    IF LENGTH(v_contrasena) = 0 THEN
        SIGNAL SQLSTATE '45000' SET MYSQL_ERRNO = 30001, MESSAGE_TEXT = 'La contrasena no puede estar vacia';
    END IF;
    SET salt = generar_salt();
    INSERT INTO usuarios (correo, usuario, contrasena_salt, contrasena)
    VALUES (v_correo, v_usuario, salt, ENCRYPT(v_contrasena, salt));
END;
@@

CREATE OR REPLACE FUNCTION
    obtener_id_usuario_por_correo(
    v_correo VARCHAR(1024)
)
    RETURNS INTEGER UNSIGNED
    LANGUAGE SQL
    NOT DETERMINISTIC
BEGIN
    DECLARE id_usuario INTEGER UNSIGNED;
    SET id_usuario = (SELECT id
                      FROM usuarios
                      WHERE usuarios.activo
                        AND usuarios.correo = v_correo
                      LIMIT 1);
    IF id_usuario IS NULL THEN
        SIGNAL SQLSTATE '45000' SET MYSQL_ERRNO = 30001, MESSAGE_TEXT = 'Usuario no encontrado';
    END IF;
    RETURN id_usuario;
END;
@@

CREATE OR REPLACE FUNCTION
    login(
    v_usuario VARCHAR(1024),
    v_contrasena VARCHAR(1024)
)
    RETURNS INTEGER UNSIGNED
    LANGUAGE SQL
    NOT DETERMINISTIC
BEGIN
    DECLARE id_usuario INTEGER UNSIGNED;
    SET id_usuario = (SELECT id
                      FROM usuarios
                      WHERE usuarios.activo
                        AND usuarios.usuario = v_usuario
                        AND usuarios.contrasena = ENCRYPT(v_contrasena, usuarios.contrasena_salt)
                      LIMIT 1);
    IF id_usuario IS NULL THEN
        SIGNAL SQLSTATE '45000' SET MYSQL_ERRNO = 30001, MESSAGE_TEXT = 'Usuario o contrasena invalidos';
    END IF;
    RETURN id_usuario;
END;
@@

CREATE OR REPLACE FUNCTION
    usuario_disponible(
    v_usuario VARCHAR(1024)
)
    RETURNS BOOL
    LANGUAGE SQL
    NOT DETERMINISTIC
BEGIN
    DECLARE id_usuario INTEGER UNSIGNED;
    SET id_usuario = (SELECT id
                      FROM usuarios
                      WHERE usuarios.activo
                        AND usuarios.usuario = v_usuario
                      LIMIT 1);
    RETURN id_usuario IS NULL;
END;
@@

CREATE OR REPLACE FUNCTION
    correo_disponible(
    v_correo VARCHAR(1024)
)
    RETURNS BOOL
    LANGUAGE SQL
    NOT DETERMINISTIC
BEGIN
    DECLARE id_usuario INTEGER UNSIGNED;
    SET id_usuario = (SELECT id
                      FROM usuarios
                      WHERE usuarios.activo
                        AND usuarios.correo = v_correo
                      LIMIT 1);
    RETURN id_usuario IS NULL;
END;
@@

CREATE OR REPLACE PROCEDURE
    usuarios_cambiar_contrasena(
    v_id_usuario INTEGER UNSIGNED,
    v_contrasena VARCHAR(1024),
    v_nueva_contrasena VARCHAR(1024)
)
    LANGUAGE SQL
    NOT DETERMINISTIC
BEGIN
    DECLARE salt VARCHAR(1024);
    IF LENGTH(v_nueva_contrasena) = 0 THEN
        SIGNAL SQLSTATE '45000' SET MYSQL_ERRNO = 30001, MESSAGE_TEXT = 'La contrasena no puede estar vacia';
    END IF;
    SET salt = generar_salt();
    UPDATE
        usuarios
    SET usuarios.contrasena_salt = salt,
        usuarios.contrasena      = ENCRYPT(v_nueva_contrasena, salt)
    WHERE usuarios.id = v_id_usuario
      AND usuarios.contrasena = ENCRYPT(v_contrasena, usuarios.contrasena_salt)
    LIMIT 1;
    IF ROW_COUNT() != 1 THEN
        SIGNAL SQLSTATE '45000' SET MYSQL_ERRNO = 30001, MESSAGE_TEXT = 'Credenciales invalidas';
    END IF;
END;
@@

CREATE OR REPLACE PROCEDURE
    usuarios_cambiar_correo(
    v_id_usuario INTEGER UNSIGNED,
    v_nuevo_correo VARCHAR(1024),
    v_contrasena VARCHAR(1024)
)
    LANGUAGE SQL
    NOT DETERMINISTIC
BEGIN
    UPDATE
        usuarios
    SET usuarios.correo = v_nuevo_correo
    WHERE usuarios.id = v_id_usuario
      AND usuarios.contrasena = ENCRYPT(v_contrasena, usuarios.contrasena_salt);
    IF ROW_COUNT() != 1 THEN
        SIGNAL SQLSTATE '45000' SET MYSQL_ERRNO = 30001, MESSAGE_TEXT = 'Credenciales invalidas';
    END IF;
END;
@@

CREATE OR REPLACE PROCEDURE
    reset_contrasena(
    v_id_usuario INTEGER UNSIGNED,
    v_nueva_contrasena VARCHAR(1024)
)
    LANGUAGE SQL
    NOT DETERMINISTIC
BEGIN
    DECLARE salt VARCHAR(1024);
    SET salt = generar_salt();
    UPDATE
        usuarios
    SET usuarios.contrasena_salt = salt,
        usuarios.contrasena      = ENCRYPT(v_nueva_contrasena, salt)
    WHERE usuarios.id = v_id_usuario;
    IF ROW_COUNT() != 1 THEN
        SIGNAL SQLSTATE '45000' SET MYSQL_ERRNO = 30001, MESSAGE_TEXT = 'Credenciales invalidas';
    END IF;
END;
@@

-- - Related to palabras_clave - --

CREATE OR REPLACE TRIGGER before_update_palabras_clave
    BEFORE UPDATE
    ON palabras_clave
    FOR EACH ROW
BEGIN
    DECLARE usuario_activo BOOL;
    IF NOT OLD.activo AND NOT NEW.activo THEN
        SIGNAL SQLSTATE '45000' SET MYSQL_ERRNO = 30001, MESSAGE_TEXT =
                'El registro esta desactivo, reactivar primero antes de actualizar';
    END IF;
    -- - El id_usuario es inmutable
    IF OLD.id_usuario != NEW.id_usuario THEN
        SIGNAL SQLSTATE '45000' SET MYSQL_ERRNO = 30001, MESSAGE_TEXT =
                'id_usuario es inmutable';
    END IF;
    -- - Si el usuario esta desactivado esto no se puede modificar
    SET usuario_activo = (
        SELECT usuarios.activo
        FROM usuarios
        WHERE usuarios.id = NEW.id_usuario
        LIMIT 1
    );
    IF NOT usuario_activo THEN
        SIGNAL SQLSTATE '45000' SET MYSQL_ERRNO = 30001, MESSAGE_TEXT =
                'El usuario esta desactivo, reactivar primero antes de actualizar';
    END IF;
END;
@@

CREATE OR REPLACE PROCEDURE
    actualizar_palabras_clave(
    v_id_usuario INTEGER UNSIGNED,
    v_nuevas_palabras JSON,
    v_automatico BOOL
)
    LANGUAGE SQL
    NOT DETERMINISTIC
BEGIN
    UPDATE
        palabras_clave AS pc
    SET pc.palabras   = v_nuevas_palabras,
        pc.automatico = v_automatico
    WHERE pc.id_usuario = v_id_usuario
    LIMIT 1;
END;
@@

DELIMITER ;