package error

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	resp "github.com/deall-users/pkg/response"
	"github.com/go-playground/validator/v10"
)

type errorValidation struct {
	Parameter string `json:"parameter"`
	Message   string `json:"message"`
}

type ErrValidations []*errorValidation

func ErrUnmarshal(err error) ErrValidations {
	var syntaxError *json.SyntaxError
	var unmarshalTypeError *json.UnmarshalTypeError

	ev := make(ErrValidations, 0)

	switch {
	case errors.As(err, &syntaxError):
		ev = append(ev, &errorValidation{
			Parameter: "json",
			Message:   fmt.Sprintf("Request body contains badly-formed JSON (at position %d)", syntaxError.Offset),
		})
		return ev
	case errors.Is(err, io.ErrUnexpectedEOF):
		ev = append(ev, &errorValidation{
			Parameter: "json",
			Message:   "Request body contains badly-formed JSON",
		})
		return ev
	case errors.As(err, &unmarshalTypeError):
		ev = append(ev, &errorValidation{
			Parameter: unmarshalTypeError.Field,
			Message:   fmt.Sprintf("%s must be (type:%s)", unmarshalTypeError.Field, unmarshalTypeError.Type),
		})
		return ev
	default:
		return nil
	}
}

// set custom error inside this func. Dont forget your all tags.
func ErrValidation(e validator.ValidationErrors) ErrValidations {
	ev := make(ErrValidations, 0)
	for _, err := range e {
		message := ""
		field := err.Field()

		switch err.Tag() {
		case "required":
			message = fmt.Sprintf("%s is required", field)
		case "alphanum":
			message = fmt.Sprintf("%s alpha numeric only", field)
		case "numeric":
			message = fmt.Sprintf("%s numeric only", field)
		default:
			message = fmt.Sprintf("%s is not valid", field)
		}

		ev = append(ev, &errorValidation{
			Parameter: field,
			Message:   message,
		})
	}

	return ev
}

func ErrBindingHandler(err error) ErrValidations {
	if errV, ok := err.(validator.ValidationErrors); ok {
		return ErrValidation(errV)
	} else {
		return ErrUnmarshal(err)
	}
}

func BindValidateJSON(c *gin.Context, obj interface{}) (ErrValidations, error) {
	if c.Request.Body == http.NoBody {
		return nil, NewErrorf(http.StatusInternalServerError, "No Body attached to request. Please, make sure your body is not empty!")
	}

	err := c.ShouldBindBodyWith(obj, binding.JSON)
	if err != nil {
		return ErrBindingHandler(err), err
	}

	return nil, nil
}

func BindValidateQuery(ctx *gin.Context, obj interface{}) (ErrValidations, error) {
	err := ctx.ShouldBindQuery(obj)
	if err != nil {
		return ErrBindingHandler(err), err
	}

	return nil, nil
}

func BindValidateURI(ctx *gin.Context, obj interface{}) (ErrValidations, error) {
	err := ctx.ShouldBindUri(obj)
	if err != nil {
		return ErrBindingHandler(err), err
	}

	return nil, nil
}

func RespErrValidation(ev ErrValidations) *resp.Response {
	return resp.DefaultResponse(ErrPayloadValidation.Error(), nil, ev, http.StatusBadRequest)
}
