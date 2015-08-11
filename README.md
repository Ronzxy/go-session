Session
===========

Session handler usage 'beego session'.

### Usage

``` golang
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"openeasy.net/openeasy/session"
)

func main() {
	var (
		str  string = `{"Type": "memory", "cookieName": "gosessionsid", "gclifetime": 3600, "ProviderConfig": ""}`
		conf session.Config
		err  error
	)

	err = json.Unmarshal([]byte(str), &conf)
	if err != nil {
		// 解析配置文件
		fmt.Println("json.Unmarshal", err.Error())
		return
	}

	// 初始化sessions
	err = session.NewManager(conf)
	if err != nil {
		fmt.Println("Session", err.Error())
	}

	// 获取Gin默认实例
	r := gin.Default()
	// 添加路由
	r.GET("/", func(ctx *gin.Context) {
		s, _ := session.GetInst(ctx.Writer, ctx.Request)
		v := s.Get("userName")
		switch {
		case v != nil:
			{
				ctx.String(200, "UserName: %s", v.(string))
			}
		default:
			{
				ctx.String(200, "Nothing")
			}
		}

	})

	r.GET("/login", func(ctx *gin.Context) {
		s, _ := session.GetInst(ctx.Writer, ctx.Request)
		s.Set("userName", "I'm the userName.")

		ctx.String(200, "OK")
	})
	// 设置http服务
	s := &http.Server{
		Addr:           ":8080",
		Handler:        r,
		ReadTimeout:    30 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
```

### Session Manager: What providers are supported?
See <a href="https://github.com/astaxie/beego/tree/master/session">Beego Session</a>