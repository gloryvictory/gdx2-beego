package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["gdx2-beego/controllers:FieldController"] = append(beego.GlobalControllerRouter["gdx2-beego/controllers:FieldController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gdx2-beego/controllers:FieldController"] = append(beego.GlobalControllerRouter["gdx2-beego/controllers:FieldController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gdx2-beego/controllers:FieldController"] = append(beego.GlobalControllerRouter["gdx2-beego/controllers:FieldController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gdx2-beego/controllers:FieldController"] = append(beego.GlobalControllerRouter["gdx2-beego/controllers:FieldController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gdx2-beego/controllers:FieldController"] = append(beego.GlobalControllerRouter["gdx2-beego/controllers:FieldController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gdx2-beego/controllers:LuController"] = append(beego.GlobalControllerRouter["gdx2-beego/controllers:LuController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gdx2-beego/controllers:LuController"] = append(beego.GlobalControllerRouter["gdx2-beego/controllers:LuController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gdx2-beego/controllers:LuController"] = append(beego.GlobalControllerRouter["gdx2-beego/controllers:LuController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gdx2-beego/controllers:LuController"] = append(beego.GlobalControllerRouter["gdx2-beego/controllers:LuController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gdx2-beego/controllers:LuController"] = append(beego.GlobalControllerRouter["gdx2-beego/controllers:LuController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gdx2-beego/controllers:StaController"] = append(beego.GlobalControllerRouter["gdx2-beego/controllers:StaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gdx2-beego/controllers:StaController"] = append(beego.GlobalControllerRouter["gdx2-beego/controllers:StaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gdx2-beego/controllers:StaController"] = append(beego.GlobalControllerRouter["gdx2-beego/controllers:StaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gdx2-beego/controllers:StaController"] = append(beego.GlobalControllerRouter["gdx2-beego/controllers:StaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gdx2-beego/controllers:StaController"] = append(beego.GlobalControllerRouter["gdx2-beego/controllers:StaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gdx2-beego/controllers:StlController"] = append(beego.GlobalControllerRouter["gdx2-beego/controllers:StlController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gdx2-beego/controllers:StlController"] = append(beego.GlobalControllerRouter["gdx2-beego/controllers:StlController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gdx2-beego/controllers:StlController"] = append(beego.GlobalControllerRouter["gdx2-beego/controllers:StlController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gdx2-beego/controllers:StlController"] = append(beego.GlobalControllerRouter["gdx2-beego/controllers:StlController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gdx2-beego/controllers:StlController"] = append(beego.GlobalControllerRouter["gdx2-beego/controllers:StlController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gdx2-beego/controllers:StpController"] = append(beego.GlobalControllerRouter["gdx2-beego/controllers:StpController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gdx2-beego/controllers:StpController"] = append(beego.GlobalControllerRouter["gdx2-beego/controllers:StpController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gdx2-beego/controllers:StpController"] = append(beego.GlobalControllerRouter["gdx2-beego/controllers:StpController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gdx2-beego/controllers:StpController"] = append(beego.GlobalControllerRouter["gdx2-beego/controllers:StpController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gdx2-beego/controllers:StpController"] = append(beego.GlobalControllerRouter["gdx2-beego/controllers:StpController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
