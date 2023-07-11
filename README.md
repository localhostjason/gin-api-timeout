# gin_api_timeout

主要针对api，实现超时时，会断开连接，返回错误信息

### 例子
```go
package main

import (
	timeout "github.com/localhostjason/gin-api-timeout"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type TimeoutMsg map[string]interface{}

var timeoutMsg = TimeoutMsg{
	"code": -1,
	"msg":  "timeout !!",
}

func main() {
	r := gin.Default()

	r.Use(timeout.Timeout(
		timeout.WithTimeout(3*time.Second),
		timeout.WithDefaultMsg(timeoutMsg)),
	)
	r.GET("/ping", func(c *gin.Context) {
		time.Sleep(5 * time.Second)
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
```

