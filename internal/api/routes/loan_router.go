// routers/loan_router.go
package routers

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func SetupLoanRouter(r *gin.Engine) {
    r.GET("/track", func(c *gin.Context) {
        c.String(http.StatusOK, "Your loan will be fully paid by the end of this year!")
    })
}

