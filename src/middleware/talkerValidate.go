package middleware

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"regexp"

	"github.com/gin-gonic/gin"
)

type talker struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Talk struct {
		WatchedAt string `json:"watchedAt"`
		Rate      int    `json:"rate"`
	} `json:"talk"`
}

func (t *talker) Validate() error {
	// Name
	if reflect.ValueOf(t.Name).IsZero() {
		return errors.New("O campo \"name\" é obrigatório")
	}
	if len(t.Name) < 3 {
		return errors.New("O \"name\" deve ter pelo menos 3 caracteres")
	}

	// Age
	if reflect.ValueOf(t.Age).IsZero() {
		return errors.New("O campo \"age\" é obrigatório")
	}
	if t.Age < 18 {
		return errors.New("A pessoa palestrante deve ser maior de idade")
	}

	// talk
	if reflect.ValueOf(t.Talk).IsZero() {
		return errors.New("O campo \"talk\" é obrigatório e \"watchedAt\" e \"rate\" não podem ser vazios")
	}
	if reflect.ValueOf(t.Talk.Rate).IsZero() {
		return errors.New("O campo \"rate\" é obrigatório")
	}
	if t.Talk.Rate < 0 || t.Talk.Rate > 5 {
		return errors.New("O campo \"rate\" deve ser um inteiro de 1 à 5")
	}
	if reflect.ValueOf(t.Talk.WatchedAt).IsZero() {
		return errors.New("O campo \"watchedAt\" é obrigatório")
	}

	REGEX := regexp.MustCompile(`^([0-2][0-9]|(3)[0-1])(\/)(((0)[0-9])|((1)[0-2]))(\/)\d{4}$`)
	if !REGEX.MatchString(t.Talk.WatchedAt) {
		return errors.New("O campo \"watchedAt\" deve ter o formato \"dd/mm/aaaa\"")
	}

	return nil
}

func TalkerValidate(c *gin.Context) {
	ByteBody, _ := ioutil.ReadAll(c.Request.Body)
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(ByteBody))

	var t talker
	err := json.Unmarshal(ByteBody, &t)
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"status":  http.StatusBadRequest,
				"message": fmt.Sprintf("Invalid JSON: %s", err.Error()),
			})
		return
	}

	if err := t.Validate(); err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"status":  http.StatusBadRequest,
				"message": err.Error(),
			})
		return
	}

	fmt.Println(t.Name)

	c.Next()
}
