package Source

import (
	"Source/evault/config"
	"Source/evault/handler"
	"context"

	"github.com/gin-gonic/gin"
)

func StartBackend() {

	config.Initdatabse()
	config.Migrateschema()

	ctx, cancel := context.WithCancel(context.Background())
	go handler.Managebg(ctx)
	defer cancel()
	r := gin.Default()
	r.GET("/pickuprequest", handler.Getallrequest)
	r.POST("/pickuprequest", handler.Createrequest)
	r.GET("/pickuprequest/:id", handler.Getrequest)
	r.PUT("/pickuprequest/:id", handler.Updaterequest)
	r.DELETE("/pickuprequest/:id", handler.Deleterequest)
	r.GET("/getwastetypes", handler.Getwastetypes)
	r.GET("/getcollectedlist", handler.Getcollected)
	r.GET("/gethistory", handler.Gethistory)
	r.GET("/gettotalweight", handler.GetTotalweight)
	r.Run(":8080")

}
