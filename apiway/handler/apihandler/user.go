package apihandler

import (
	"flag"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"zg6zk1/apiway/basic/pkg"
	"zg6zk1/apiway/handler/reqhandler"
	__ "zg6zk1/service/protobuf"
)

func Login(c *gin.Context) {
	var req reqhandler.LoginForm

	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  err.Error(),
		})
		return
	}
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := __.NewUserClient(conn)
	login, err := client.Login(c, &__.LoginRequest{
		Mobile:  req.Mobile,
		SendSms: req.SendSms,
	})
	if err != nil {
		return
	}

	token, err := pkg.NewJWT("zk").CreateToken(pkg.CustomClaims{
		ID: uint(login.Id),
	})
	if err != nil {
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "ok",
		"data": token,
	})

}

func List(c *gin.Context) {

}
