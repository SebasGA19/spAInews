package controller

import "encoding/json"

type (
	Words []string
)

func (c *Controller) SetUserWords(userId int, newWords Words) error {
	asJsonObject, marshalError := json.Marshal(newWords)
	if marshalError != nil {
		return marshalError
	}
	_, err := c.SQL.Exec("CALL actualizar_palabras_clave(?, ?)", userId, asJsonObject)
	return err
}

func (c *Controller) GetUserWords(userId int) (Words, error) {
	row := c.SQL.QueryRow("SELECT palabras_clave_de_usuario(?)", userId)
	var rawWords string
	scanError := row.Scan(&rawWords)
	if scanError != nil {
		return nil, scanError
	}
	var words Words
	unmarshallError := json.Unmarshal([]byte(rawWords), &words)
	return words, unmarshallError
}
