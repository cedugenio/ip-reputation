package api

import (
	"net/http"

	"github.com/cedugenio/ip-reputation/internal/store"
	"github.com/gin-gonic/gin"
)

func StartServer() {
    r := gin.Default()

    r.GET("/check/:ip", func(c *gin.Context) {
        ip := c.Param("ip")
        score, err := store.GetIPScore(ip)
        if err != nil {
            c.JSON(http.StatusNotFound, gin.H{
                "ip":         ip,
                "reputation": "unknown",
            })
            return
        }

        c.JSON(http.StatusOK, gin.H{
            "ip":         ip,
            "reputation": score,
        })
    })

    r.Run(":8080")
}