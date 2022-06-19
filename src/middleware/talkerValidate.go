package middleware

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type talker struct {
	Name string `json:"name"`
	// Age  int    `json:"age"`
	// Talk struct {
	// 	WatchedAt string `json:"watchedAt"`
	// 	Rate      int    `json:"rate"`
	// } `json:"talk"`
}

func (t *talker) Validate() error {
	// Name
	if t.Name == "" {
		return errors.New("O campo \"name\" é obrigatório")
	}
	if len(t.Name) < 3 {
		return errors.New("O \"name\" deve ter pelo menos 3 caracteres")
	}

	// // Age
	// if reflect.ValueOf(t.Age).IsZero() {
	// 	return errors.New("O campo \"age\" é obrigatório")
	// }
	// if t.Age < 18 {
	// 	return errors.New("A pessoa palestrante deve ser maior de idade")
	// }

	// if t.Talk.Rate < 0 || t.Talk.Rate > 5 {
	// 	return errors.New("Rate must be between 0 and 5")
	// }

	return nil
}

func TalkerValidate(c *gin.Context) {
	var json talker
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		c.Abort()
		return
	}

	c.Next()
}
