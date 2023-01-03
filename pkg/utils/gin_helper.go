package utils

import (
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetHeaderBearerToken(c *gin.Context) (authType string, token string) {
	bearerToken := c.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[0], strings.Split(bearerToken, " ")[1]
	}

	return "", ""
}

// Check body request json (application/json) has been set or not.
// It will set to context(gin) if body havent added yet.
func CheckSetBody(c *gin.Context) error {
	if c.Request.Body == http.NoBody {
		return nil
	}

	if _, ok := c.Get(gin.BodyBytesKey); !ok {
		body, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			return err
		}

		c.Set(gin.BodyBytesKey, body)
	}

	return nil
}
