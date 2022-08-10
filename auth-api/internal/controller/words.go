package controller

import "encoding/json"

type (
	Words []string
)

func (c *Controller) SetUserWords(userId int, newWords Words, automatic bool) error {
	asJsonObject, marshalError := json.Marshal(newWords)
	if marshalError != nil {
		return marshalError
	}
	_, err := c.SQL.Exec("CALL actualizar_palabras_clave(?, ?, ?)", userId, asJsonObject, automatic)
	return err
}

func (c *Controller) GetUserWords(userId int) (words Words, auto bool, err error) {
	row := c.SQL.QueryRow("SELECT pc.palabras, pc.automatico FROM palabras_clave AS pc WHERE pc.id_usuario = ? LIMIT 1", userId)
	var rawWords string
	err = row.Scan(&rawWords, &auto)
	if err != nil {
		return nil, false, err
	}
	err = json.Unmarshal([]byte(rawWords), &words)
	return words, auto, err
}

func (c *Controller) Account(userId int) (username, email string, err error) {
	row := c.SQL.QueryRow("SELECT usuarios.usuario, usuarios.correo FROM usuarios WHERE usuarios.id = ? LIMIT 1", userId)
	scanError := row.Scan(&username, &email)
	return username, email, scanError
}
