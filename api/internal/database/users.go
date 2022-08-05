package database

func (db *Database) RegisterUser(email, username, password string) error {
	_, execError := db.SQL.Exec("CALL registrar_usuario(?, ?, ?)", username, email, password)
	return execError
}

func (db *Database) Login(username, password string) (int, error) {
	row := db.SQL.QueryRow("SELECT login(?, ?)", username, password)
	var id int
	scanError := row.Scan(&id)
	return id, scanError
}
