# Base de datos

El proyecto manejara diversos gestores de base de datos para manejar las cuentas de los usuarios, las sesiones de los usuarios, los artículos recolectados.

## SQL (MariaDB)

### Tablas

```mermaid
erDiagram
	usuarios {
		INTEGER id "NOT NULL, UNICO, AUTO INCREMENTAR, UNSIGNED"
		TIMESTAMP fecha_creacion "NOT NULL, DEFAULT CURRENT_TIMESTAMP"
		TIMESTAMP fecha_ultima_modificacion "NOT NULL, DEFAULT CURRENT_TIMESTAMP"
		BOOL activo "NOT NULL, DEFAULT TRUE"
		STRING correo "NOT NULL, UNIQUE, REGEXP"
		STRING usuario "NOT NULL, UNIQUE, REGEXP"
		STRING contrasena_salt "NOT NULL, REGEXP"
		STRING contrasena "NOT NULL, REGEXP"
	}
	palabras_clave {
		INTEGER id "NOT NULL, UNICO, AUTO INCREMENTAR, UNISGNED"
		TIMESTAMP fecha_creacion "NOT NULL, DEFAULT CURRENT_TIMESTAMP"
		TIMESTAMP fecha_modificacion "NOT NULL, DEFAULT CURRENT_TIMESTAMP"
		BOOL activo "NOT NULL, DEFAULT TRUE"
		INTEGER id_usuario "NOT NULL, FOREIGN KEY usuarios"
		BOOL automatico "NOT NULL, DEFAULT TRUE"
		JSON words "NOT NULL, DEFAULT '[]', REGEXP (CHECKS IF IT IS A JSON ARRAY)"
	}
	
```

### Entidad relación

```mermaid
erDiagram
	usuarios ||--o{ palabras_clave : tienen
```

### Funciones

- `login(usuario, contrasena) id`: Esta función recibe de entrada el usuario y contraseña, retorna el `id` en caso de éxito y `-1` en caso de fracaso.
- `usuario_disponible(usuario) BOOL`: Retorna verdadero si el usuario esta disponible para el registro.
- `correo_disponible(usuario) BOOL`: Retorna verdadero si el correo esta disponible para el registro.
- `palabras_clave(usuario_o_id) JSON`: Retorna las palabras claves del usuario.

### Procedimientos

- `registrar(usuario, correo, contrasena)`: Confirma que el correo, usuario y contraseña satisfacen las expresiones regulares, luego registra al usuario en la base de datos.
- `agregar_palabra_clave(usuario_o_id, palabra_clave)`: Agrega la palabra clave al arreglo de palabras claves del usuario.
- `remover_palabra_clave(usuario_o_id, palabra_clave)`: Remueve la palabra clave del arreglo de palabras claves del usuario.

## Redis

## MongoDB