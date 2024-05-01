package ctype

import (
	"database/sql/driver"
	"encoding/json"
)

type VerifyQuestion struct {
	Question1 *string `json:"question1"`
	Answer1   *string `json:"answer1"`
	Question2 *string `json:"question2"`
	Answer2   *string `json:"answer2"`
	Question3 *string `json:"question3"`
	Answer3   *string `json:"answer3"`
}

func (c *VerifyQuestion) Scan(val interface{}) error {
	return json.Unmarshal(val.([]byte), c)
}
func (c *VerifyQuestion) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}
