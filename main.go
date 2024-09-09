package main

import (
	// _ "gdx2-beego/controllers"
	"fmt"
	_ "gdx2-beego/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

func main() {
	orm.Debug = true
	orm.RegisterDataBase("default", "postgres", beego.AppConfig.String("sqlconn"))
	if beego.BConfig.RunMode == "dev" {
		fmt.Print("DB: " + beego.AppConfig.String("sqlconn") + "\n")
		fmt.Print("Mode: " + beego.BConfig.RunMode + "\n")

		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	// web.BConfig.RouterCaseSensitive = false

	// web.AutoRouter(&StlController{})
	// tree := web.PrintTree()
	// methods := tree["Data"].(web.M)
	// for k, v := range methods {
	// 	fmt.Printf("%s => %v\n", k, v)
	// }
	beego.Run()
}
