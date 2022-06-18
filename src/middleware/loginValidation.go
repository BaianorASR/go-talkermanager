package middleware

import (
	"errors"
	"net/http"

	"github.com/badoux/checkmail"
	"github.com/gin-gonic/gin"
)

type login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (l *login) Validate() error {
	// Check if email is valid
	if l.Email == "" {
		return errors.New("O campo \"email\" é obrigatório")
	}

	if err := checkmail.ValidateFormat(l.Email); err != nil {
		return errors.New("O \"email\" deve ter o formato \"email@email.com\"")
	}

	// Check if password is valid
	if l.Password == "" {
		return errors.New("O campo \"password\" é obrigatório")
	}

	if len(l.Password) < 6 {
		return errors.New("O \"password\" deve ter pelo menos 6 caracteres")
	}

	// No Errors
	return nil
}

func LoginValidation(c *gin.Context) {
	var json login
	if err := c.BindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		c.Abort()
		return
	}

	if err := json.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		c.Abort()
		return
	}

	c.Next()
}
