package main

import (
	"net/url"

	"github.com/kataras/iris"
	"github.com/kataras/iris/core/host"
)

func main() {
	app := iris.New()

	app.RegisterView(iris.HTML("../blog/public", ".html").Reload(true))

	app.Get("/", func(ctx iris.Context) {
		// ctx.Writef("Hello from the SECURE server")
		if err := ctx.View("index.html"); err != nil {
			ctx.Application().Logger().Infof(err.Error())
		}
	})

	app.Get("/mypath", func(ctx iris.Context) {
		ctx.Writef("Hello from the SECURE server on path /mypath")
	})

	target, _ := url.Parse("https://localhost:443")
	go host.NewProxy("localhost:80", target).ListenAndServe()

	app.Run(iris.TLS(":443", "server.cert", "private.key"))
}
