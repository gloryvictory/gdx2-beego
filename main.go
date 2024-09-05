package main

import (
	// _ "gdx2-beego/controllers"
	_ "gdx2-beego/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

func main() {
	orm.RegisterDataBase("default", "postgres", beego.AppConfig.String("sqlconn"))
	if beego.BConfig.RunMode == "dev" {
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
