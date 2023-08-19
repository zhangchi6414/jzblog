package service

import (
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	val "github.com/go-playground/validator/v10"
	"strings"
)

//接口校验

type ValidError struct {
	Key     string
	Message string
}

type ValidErrors []*ValidError

func (v *ValidError) Error() string {
	return v.Message
}

func (receiver ValidErrors) Error() string {
	return strings.Join(receiver.Errors(), ",")
}

func (receiver ValidErrors) Errors() []string {
	var errs []string
	for _, err := range receiver {
		errs = append(errs, err.Error())
	}
	return errs
}

func BindAndValid(c *gin.Context, v interface{}) (bool, ValidErrors) {
	var errs ValidErrors
	err := c.ShouldBindJSON(v)
	if err != nil {
		v := c.Value("trans")
		trans, _ := v.(ut.Translator)
		verrs, ok := err.(val.ValidationErrors)
		if !ok {
			return false, errs
		}
		for key, value := range verrs.Translate(trans) {
			errs = append(errs, &ValidError{
				Key:     key,
				Message: value,
			})
		}
		return false, errs
	}
	return true, nil
}
