package controller

func (c *Controller) RegisterUser(username, email, password string) error {
	_, execError := c.SQL.Exec("CALL registrar_usuario(?, ?, ?)", username, email, password)
	return execError
}

func (c *Controller) Login(username, password string) (int, error) {
	row := c.SQL.QueryRow("SELECT login(?, ?)", username, password)
	var id int
	scanError := row.Scan(&id)
	return id, scanError
}

func (c *Controller) UsernameAvailable(username string) (bool, error) {
	row := c.SQL.QueryRow("SELECT usuario_disponible(?)", username)
	var result bool
	scanError := row.Scan(&result)
	return result, scanError
}

func (c *Controller) EmailAvailable(email string) (bool, error) {
	row := c.SQL.QueryRow("SELECT correo_disponible(?)", email)
	var result bool
	scanError := row.Scan(&result)
	return result, scanError
}

func (c *Controller) ChangePassword(userId int, oldPassword, newPassword string) error {
	_, err := c.SQL.Exec("CALL usuarios_cambiar_contrasena(?, ?, ?)", userId, oldPassword, newPassword)
	return err
}

func (c *Controller) ChangeEmail(userId int, email, password string) error {
	_, err := c.SQL.Exec("CALL usuarios_cambiar_correo(?, ?, ?)", userId, email, password)
	return err
}

func (c *Controller) SQLResetPassword(userId int, newPassword string) error {
	_, err := c.SQL.Exec("CALL reset_contrasena(?, ?)", userId, newPassword)
	return err
}

func (c *Controller) QueryUserIdByEmail(email string) (int, error) {
	row := c.SQL.QueryRow("SELECT obtener_id_usuario_por_correo(?)", email)
	var result int
	scanError := row.Scan(&result)
	return result, scanError
}
