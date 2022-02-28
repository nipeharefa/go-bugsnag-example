package main

import (
	"errors"

	bugsnaggin "github.com/bugsnag/bugsnag-go-gin"
	"github.com/bugsnag/bugsnag-go/v2"
	"github.com/gin-gonic/gin"
)

var (
	ErrUserNotFound = errors.New("user not found")
)

func main() {
	g := gin.Default()

	g.Use(bugsnaggin.AutoNotify(bugsnag.Configuration{
		// Your Bugsnag project API key, required unless set as environment
		// variable $BUGSNAG_API_KEY
		APIKey: "",
		// The import paths for the Go packages containing your source files
		ProjectPackages: []string{"main", "github.com/org/myapp"},
	}))

	g.GET("/meta", func(c *gin.Context) {
		user := bugsnag.User{Id: "1234", Name: "Conrad", Email: "me@example.com"}
		class := bugsnag.ErrorClass{
			Name: "MetaData",
		}
		meta := bugsnag.MetaData{
			"Client": map[string]interface{}{
				"IP": c.ClientIP(),
			},
		}
		bugsnag.Notify(ErrUserNotFound, c.Request.Context(), user, class, meta)
		c.String(200, "Hello World")
	})

	g.GET("/", func(c *gin.Context) {
		user := bugsnag.User{Id: "1234", Name: "Conrad", Email: "me@example.com"}
		class := bugsnag.ErrorClass{
			Name: "UserError",
		}
		bugsnag.Notify(ErrUserNotFound, c.Request.Context(), user, class)
		c.String(200, "Hello World")
	})

	g.Run(":8080")
}
