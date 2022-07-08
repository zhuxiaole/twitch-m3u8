package main
import (
    "net/http"
    "github.com/gin-gonic/gin"
	"github.com/wmw64/twitchpl"
)
func main() {
    router := gin.Default()
	router.GET("/", func(c *gin.Context) {
        r.LoadHTMLFiles("index.html")
		r.GET("/", func(c *gin.Context) {
			c.HTML(http.StatusOK, "index.html")
		})
    })
	router.GET("/api/:id", func(c *gin.Context) {
		pl, err := twitchpl.Get(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"message": "parse fail!",
			})
		} else {
			c.String(http.StatusOK, pl.AsJSON())
		}
    })
    router.Run(":8099")
}