package gin

import (
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"

	cfg "github.com/deall-users/config"
	"github.com/deall-users/pkg/validation"
)

func NewGin() *gin.Engine {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		//set json as default message key field
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})

		// add any custom validations etc. here
		v.RegisterValidation("decryptiontext", validation.ValidateDecryptText)
	}

	//mode
	runMode := ""
	switch cfg.RUN_MODE {
	case "production":
		runMode = gin.ReleaseMode
	default:
		runMode = gin.DebugMode
	}
	gin.SetMode(runMode)

	app := gin.Default()

	// set default middleware
	app.Use(gin.ErrorLogger())

	// customize log
	// app.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {}))

	return app
}
