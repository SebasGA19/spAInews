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

DELIMITER @@

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
    login(
    v_usuario VARCHAR(1024),
    v_contrasena VARCHAR(1024)
)
    RETURNS INT
    LANGUAGE SQL
    NOT DETERMINISTIC
BEGIN
    DECLARE id_usuario INT;
    SET id_usuario = (SELECT id
                      FROM usuarios
                      WHERE usuarios.activo
                        AND usuarios.usuario = v_usuario
                        AND usuarios.contrasena = ENCRYPT(v_contrasena, usuarios.contrasena_salt)
                      LIMIT 1);
    IF id_usuario IS NULL THEN
        RETURN -1;
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
    DECLARE id_usuario INT;
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
    DECLARE id_usuario INT;
    SET id_usuario = (SELECT id
                      FROM usuarios
                      WHERE usuarios.activo
                        AND usuarios.correo = v_correo
                      LIMIT 1);
    RETURN id_usuario IS NULL;
END;
@@

DELIMITER ;