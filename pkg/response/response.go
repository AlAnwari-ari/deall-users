package response

import (
	"net/http"

	"github.com/deall-users/pkg/utils"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Messages string      `json:"messages"`
	Data     interface{} `json:"data"`
	Errors   interface{} `json:"errors"`
	Code     int         `json:"code"`
}

// New Constructor Response
func DefaultResponse(message string, data interface{}, errors interface{}, code int) *Response {
	return &Response{
		Messages: message,
		Data:     data,
		Errors:   errors,
		Code:     code,
	}
}

func (resp *Response) RespJSON(c *gin.Context) error {
	if c.Writer.Written() {
		return nil
	}

	//check body has been set yet
	utils.CheckSetBody(c)

	// c.Writer.Header().Add("Tag-Version", constanta.Config.AppVersion)
	if gin.Mode() == gin.ReleaseMode {
		c.JSON(resp.Code, resp) //prod
	} else {
		c.IndentedJSON(resp.Code, resp) //dev
	}

	c.Request.Body.Close()
	return nil
}

func RespSuccess(data interface{}) *Response {
	return DefaultResponse("OK", data, nil, http.StatusOK)
}
