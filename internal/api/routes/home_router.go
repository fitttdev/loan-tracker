// routers/home_router.go
package routers

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func SetupHomeRouter(r *gin.Engine) {
    r.LoadHTMLGlob("templates/*")
    r.GET("/", func(c *gin.Context) {
        c.HTML(http.StatusOK, "index.tmpl", gin.H{
            "title": "Main website",
        })
    })

    r.GET("/ping", func(c *gin.Context) {
        c.String(http.StatusOK, "pong")
    })
}

