package main

import (
	"./client"
	"fmt"
	"github.com/gorilla/websocket"
	"gopkg.in/gin-gonic/gin.v1"
)

var count = 0

func main() {
	router := gin.Default()
	gin.SetMode(gin.ReleaseMode)
	router.Static("/static", "./static")
	router.LoadHTMLGlob("template/*")

	v1 := router.Group("/v1")

	v1.GET("/get", func(c *gin.Context) {

		c.JSON(200, map[string]string{"version": "v1", "method": "get"})
	})

	v1.GET("/echo", func(c *gin.Context) {

		conn, err := websocket.Upgrade(c.Writer, c.Request, nil, 1024, 1024)
		count++
		userid := string(count)
		fmt.Println("websocket:", userid, " is comming")
		if err != nil {
			fmt.Println("webSocket 连接出错")
			return
		}

		client.Cache[userid] = conn
		fmt.Println(client.Cache)
		defer conn.Close()
		conn.SetCloseHandler(func(i int, s string) error {
			fmt.Println("i:", i, "websocket is close", "s:", s)
			return nil
		})
		conn.SetPingHandler(func(s string) error {
			fmt.Println("self is receive a ping message:", s)
			return nil
		})
		conn.SetPongHandler(func(s string) error {
			fmt.Println("self is receive a pong message:", s)
			return nil
		})

		for {
			mt, message, err := conn.ReadMessage()

			if err != nil {
				fmt.Println("read:", err)
				break
			}

			fmt.Printf("recv: %s\n", message)
			for _, v := range client.Cache {
				err = v.WriteMessage(mt, message)
				if err != nil {
					fmt.Println("write:", err)
					break
				}
			}

		}

	})

	v1.GET("/", func(c *gin.Context) {
		fmt.Println("somethins is comming")
		c.HTML(200, "index.tmpl", map[string]string{"websocetUrl": "wss://boqiao919.com/v1/echo"})
	})

	router.Run(":8080")

}

func log(v1 string) gin.HandlerFunc {
	fmt.Println(v1, "日志设置")
	return func(c *gin.Context) {
		method := c.Param("get")
		fmt.Println(v1, "日志开始。。。。。", method)
		c.Next()
		fmt.Println(v1, "日志结束。。。。。", method)
	}
}
