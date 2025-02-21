package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/KatherineRincon/Pija_Music_Api_Crud/Pija_Music_Api_Crud/controllers:ArtistaController"] = append(beego.GlobalControllerRouter["github.com/KatherineRincon/Pija_Music_Api_Crud/Pija_Music_Api_Crud/controllers:ArtistaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/KatherineRincon/Pija_Music_Api_Crud/Pija_Music_Api_Crud/controllers:ArtistaController"] = append(beego.GlobalControllerRouter["github.com/KatherineRincon/Pija_Music_Api_Crud/Pija_Music_Api_Crud/controllers:ArtistaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/KatherineRincon/Pija_Music_Api_Crud/Pija_Music_Api_Crud/controllers:ArtistaController"] = append(beego.GlobalControllerRouter["github.com/KatherineRincon/Pija_Music_Api_Crud/Pija_Music_Api_Crud/controllers:ArtistaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/KatherineRincon/Pija_Music_Api_Crud/Pija_Music_Api_Crud/controllers:ArtistaController"] = append(beego.GlobalControllerRouter["github.com/KatherineRincon/Pija_Music_Api_Crud/Pija_Music_Api_Crud/controllers:ArtistaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/KatherineRincon/Pija_Music_Api_Crud/Pija_Music_Api_Crud/controllers:ArtistaController"] = append(beego.GlobalControllerRouter["github.com/KatherineRincon/Pija_Music_Api_Crud/Pija_Music_Api_Crud/controllers:ArtistaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/KatherineRincon/Pija_Music_Api_Crud/Pija_Music_Api_Crud/controllers:AutorCoplasController"] = append(beego.GlobalControllerRouter["github.com/KatherineRincon/Pija_Music_Api_Crud/Pija_Music_Api_Crud/controllers:AutorCoplasController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/KatherineRincon/Pija_Music_Api_Crud/Pija_Music_Api_Crud/controllers:AutorCoplasController"] = append(beego.GlobalControllerRouter["github.com/KatherineRincon/Pija_Music_Api_Crud/Pija_Music_Api_Crud/controllers:AutorCoplasController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/KatherineRincon/Pija_Music_Api_Crud/Pija_Music_Api_Crud/controllers:AutorCoplasController"] = append(beego.GlobalControllerRouter["github.com/KatherineRincon/Pija_Music_Api_Crud/Pija_Music_Api_Crud/controllers:AutorCoplasController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/KatherineRincon/Pija_Music_Api_Crud/Pija_Music_Api_Crud/controllers:AutorCoplasController"] = append(beego.GlobalControllerRouter["github.com/KatherineRincon/Pija_Music_Api_Crud/Pija_Music_Api_Crud/controllers:AutorCoplasController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/KatherineRincon/Pija_Music_Api_Crud/Pija_Music_Api_Crud/controllers:AutorCoplasController"] = append(beego.GlobalControllerRouter["github.com/KatherineRincon/Pija_Music_Api_Crud/Pija_Music_Api_Crud/controllers:AutorCoplasController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/KatherineRincon/Pija_Music_Api_Crud/Pija_Music_Api_Crud/controllers:CancionesController"] = append(beego.GlobalControllerRouter["github.com/KatherineRincon/Pija_Music_Api_Crud/Pija_Music_Api_Crud/controllers:CancionesController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/KatherineRincon/Pija_Music_Api_Crud/Pija_Music_Api_Crud/controllers:CancionesController"] = append(beego.GlobalControllerRouter["github.com/KatherineRincon/Pija_Music_Api_Crud/Pija_Music_Api_Crud/controllers:CancionesController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/KatherineRincon/Pija_Music_Api_Crud/Pija_Music_Api_Crud/controllers:CancionesController"] = append(beego.GlobalControllerRouter["github.com/KatherineRincon/Pija_Music_Api_Crud/Pija_Music_Api_Crud/controllers:CancionesController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/KatherineRincon/Pija_Music_Api_Crud/Pija_Music_Api_Crud/controllers:CancionesController"] = append(beego.GlobalControllerRouter["github.com/KatherineRincon/Pija_Music_Api_Crud/Pija_Music_Api_Crud/controllers:CancionesController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/KatherineRincon/Pija_Music_Api_Crud/Pija_Music_Api_Crud/controllers:CancionesController"] = append(beego.GlobalControllerRouter["github.com/KatherineRincon/Pija_Music_Api_Crud/Pija_Music_Api_Crud/controllers:CancionesController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/KatherineRincon/Pija_Music_Api_Crud/Pija_Music_Api_Crud/controllers:ComidaController"] = append(beego.GlobalControllerRouter["github.com/KatherineRincon/Pija_Music_Api_Crud/Pija_Music_Api_Crud/controllers:ComidaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/KatherineRincon/Pija_Music_Api_Crud/Pija_Music_Api_Crud/controllers:ComidaController"] = append(beego.GlobalControllerRouter["github.com/KatherineRincon/Pija_Music_Api_Crud/Pija_Music_Api_Crud/controllers:ComidaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/KatherineRincon/Pija_Music_Api_Crud/Pija_Music_Api_Crud/controllers:ComidaController"] = append(beego.GlobalControllerRouter["github.com/KatherineRincon/Pija_Music_Api_Crud/Pija_Music_Api_Crud/controllers:ComidaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/KatherineRincon/Pija_Music_Api_Crud/Pija_Music_Api_Crud/controllers:ComidaController"] = append(beego.GlobalControllerRouter["github.com/KatherineRincon/Pija_Music_Api_Crud/Pija_Music_Api_Crud/controllers:ComidaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/KatherineRincon/Pija_Music_Api_Crud/Pija_Music_Api_Crud/controllers:ComidaController"] = append(beego.GlobalControllerRouter["github.com/KatherineRincon/Pija_Music_Api_Crud/Pija_Music_Api_Crud/controllers:ComidaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/KatherineRincon/Pija_Music_Api_Crud/Pija_Music_Api_Crud/controllers:CoplasController"] = append(beego.GlobalControllerRouter["github.com/KatherineRincon/Pija_Music_Api_Crud/Pija_Music_Api_Crud/controllers:CoplasController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/KatherineRincon/Pija_Music_Api_Crud/Pija_Music_Api_Crud/controllers:CoplasController"] = append(beego.GlobalControllerRouter["github.com/KatherineRincon/Pija_Music_Api_Crud/Pija_Music_Api_Crud/controllers:CoplasController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/KatherineRincon/Pija_Music_Api_Crud/Pija_Music_Api_Crud/controllers:CoplasController"] = append(beego.GlobalControllerRouter["github.com/KatherineRincon/Pija_Music_Api_Crud/Pija_Music_Api_Crud/controllers:CoplasController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/KatherineRincon/Pija_Music_Api_Crud/Pija_Music_Api_Crud/controllers:CoplasController"] = append(beego.GlobalControllerRouter["github.com/KatherineRincon/Pija_Music_Api_Crud/Pija_Music_Api_Crud/controllers:CoplasController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/KatherineRincon/Pija_Music_Api_Crud/Pija_Music_Api_Crud/controllers:CoplasController"] = append(beego.GlobalControllerRouter["github.com/KatherineRincon/Pija_Music_Api_Crud/Pija_Music_Api_Crud/controllers:CoplasController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/KatherineRincon/Pija_Music_Api_Crud/Pija_Music_Api_Crud/controllers:CredencialesController"] = append(beego.GlobalControllerRouter["github.com/KatherineRincon/Pija_Music_Api_Crud/Pija_Music_Api_Crud/controllers:CredencialesController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/KatherineRincon/Pija_Music_Api_Crud/Pija_Music_Api_Crud/controllers:CredencialesController"] = append(beego.GlobalControllerRouter["github.com/KatherineRincon/Pija_Music_Api_Crud/Pija_Music_Api_Crud/controllers:CredencialesController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/KatherineRincon/Pija_Music_Api_Crud/Pija_Music_Api_Crud/controllers:CredencialesController"] = append(beego.GlobalControllerRouter["github.com/KatherineRincon/Pija_Music_Api_Crud/Pija_Music_Api_Crud/controllers:CredencialesController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/KatherineRincon/Pija_Music_Api_Crud/Pija_Music_Api_Crud/controllers:CredencialesController"] = append(beego.GlobalControllerRouter["github.com/KatherineRincon/Pija_Music_Api_Crud/Pija_Music_Api_Crud/controllers:CredencialesController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/KatherineRincon/Pija_Music_Api_Crud/Pija_Music_Api_Crud/controllers:CredencialesController"] = append(beego.GlobalControllerRouter["github.com/KatherineRincon/Pija_Music_Api_Crud/Pija_Music_Api_Crud/controllers:CredencialesController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/KatherineRincon/Pija_Music_Api_Crud/Pija_Music_Api_Crud/controllers:CulturaController"] = append(beego.GlobalControllerRouter["github.com/KatherineRincon/Pija_Music_Api_Crud/Pija_Music_Api_Crud/controllers:CulturaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/KatherineRincon/Pija_Music_Api_Crud/Pija_Music_Api_Crud/controllers:CulturaController"] = append(beego.GlobalControllerRouter["github.com/KatherineRincon/Pija_Music_Api_Crud/Pija_Music_Api_Crud/controllers:CulturaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/KatherineRincon/Pija_Music_Api_Crud/Pija_Music_Api_Crud/controllers:CulturaController"] = append(beego.GlobalControllerRouter["github.com/KatherineRincon/Pija_Music_Api_Crud/Pija_Music_Api_Crud/controllers:CulturaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/KatherineRincon/Pija_Music_Api_Crud/Pija_Music_Api_Crud/controllers:CulturaController"] = append(beego.GlobalControllerRouter["github.com/KatherineRincon/Pija_Music_Api_Crud/Pija_Music_Api_Crud/controllers:CulturaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/KatherineRincon/Pija_Music_Api_Crud/Pija_Music_Api_Crud/controllers:CulturaController"] = append(beego.GlobalControllerRouter["github.com/KatherineRincon/Pija_Music_Api_Crud/Pija_Music_Api_Crud/controllers:CulturaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/KatherineRincon/Pija_Music_Api_Crud/Pija_Music_Api_Crud/controllers:EstiloMusicalController"] = append(beego.GlobalControllerRouter["github.com/KatherineRincon/Pija_Music_Api_Crud/Pija_Music_Api_Crud/controllers:EstiloMusicalController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/KatherineRincon/Pija_Music_Api_Crud/Pija_Music_Api_Crud/controllers:EstiloMusicalController"] = append(beego.GlobalControllerRouter["github.com/KatherineRincon/Pija_Music_Api_Crud/Pija_Music_Api_Crud/controllers:EstiloMusicalController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/KatherineRincon/Pija_Music_Api_Crud/Pija_Music_Api_Crud/controllers:EstiloMusicalController"] = append(beego.GlobalControllerRouter["github.com/KatherineRincon/Pija_Music_Api_Crud/Pija_Music_Api_Crud/controllers:EstiloMusicalController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/KatherineRincon/Pija_Music_Api_Crud/Pija_Music_Api_Crud/controllers:EstiloMusicalController"] = append(beego.GlobalControllerRouter["github.com/KatherineRincon/Pija_Music_Api_Crud/Pija_Music_Api_Crud/controllers:EstiloMusicalController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/KatherineRincon/Pija_Music_Api_Crud/Pija_Music_Api_Crud/controllers:EstiloMusicalController"] = append(beego.GlobalControllerRouter["github.com/KatherineRincon/Pija_Music_Api_Crud/Pija_Music_Api_Crud/controllers:EstiloMusicalController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/KatherineRincon/Pija_Music_Api_Crud/Pija_Music_Api_Crud/controllers:FavoritosController"] = append(beego.GlobalControllerRouter["github.com/KatherineRincon/Pija_Music_Api_Crud/Pija_Music_Api_Crud/controllers:FavoritosController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/KatherineRincon/Pija_Music_Api_Crud/Pija_Music_Api_Crud/controllers:FavoritosController"] = append(beego.GlobalControllerRouter["github.com/KatherineRincon/Pija_Music_Api_Crud/Pija_Music_Api_Crud/controllers:FavoritosController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/KatherineRincon/Pija_Music_Api_Crud/Pija_Music_Api_Crud/controllers:FavoritosController"] = append(beego.GlobalControllerRouter["github.com/KatherineRincon/Pija_Music_Api_Crud/Pija_Music_Api_Crud/controllers:FavoritosController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/KatherineRincon/Pija_Music_Api_Crud/Pija_Music_Api_Crud/controllers:FavoritosController"] = append(beego.GlobalControllerRouter["github.com/KatherineRincon/Pija_Music_Api_Crud/Pija_Music_Api_Crud/controllers:FavoritosController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/KatherineRincon/Pija_Music_Api_Crud/Pija_Music_Api_Crud/controllers:FavoritosController"] = append(beego.GlobalControllerRouter["github.com/KatherineRincon/Pija_Music_Api_Crud/Pija_Music_Api_Crud/controllers:FavoritosController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/KatherineRincon/Pija_Music_Api_Crud/Pija_Music_Api_Crud/controllers:LugaresController"] = append(beego.GlobalControllerRouter["github.com/KatherineRincon/Pija_Music_Api_Crud/Pija_Music_Api_Crud/controllers:LugaresController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/KatherineRincon/Pija_Music_Api_Crud/Pija_Music_Api_Crud/controllers:LugaresController"] = append(beego.GlobalControllerRouter["github.com/KatherineRincon/Pija_Music_Api_Crud/Pija_Music_Api_Crud/controllers:LugaresController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/KatherineRincon/Pija_Music_Api_Crud/Pija_Music_Api_Crud/controllers:LugaresController"] = append(beego.GlobalControllerRouter["github.com/KatherineRincon/Pija_Music_Api_Crud/Pija_Music_Api_Crud/controllers:LugaresController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/KatherineRincon/Pija_Music_Api_Crud/Pija_Music_Api_Crud/controllers:LugaresController"] = append(beego.GlobalControllerRouter["github.com/KatherineRincon/Pija_Music_Api_Crud/Pija_Music_Api_Crud/controllers:LugaresController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/KatherineRincon/Pija_Music_Api_Crud/Pija_Music_Api_Crud/controllers:LugaresController"] = append(beego.GlobalControllerRouter["github.com/KatherineRincon/Pija_Music_Api_Crud/Pija_Music_Api_Crud/controllers:LugaresController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/KatherineRincon/Pija_Music_Api_Crud/Pija_Music_Api_Crud/controllers:TrajesTipicosController"] = append(beego.GlobalControllerRouter["github.com/KatherineRincon/Pija_Music_Api_Crud/Pija_Music_Api_Crud/controllers:TrajesTipicosController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/KatherineRincon/Pija_Music_Api_Crud/Pija_Music_Api_Crud/controllers:TrajesTipicosController"] = append(beego.GlobalControllerRouter["github.com/KatherineRincon/Pija_Music_Api_Crud/Pija_Music_Api_Crud/controllers:TrajesTipicosController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/KatherineRincon/Pija_Music_Api_Crud/Pija_Music_Api_Crud/controllers:TrajesTipicosController"] = append(beego.GlobalControllerRouter["github.com/KatherineRincon/Pija_Music_Api_Crud/Pija_Music_Api_Crud/controllers:TrajesTipicosController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/KatherineRincon/Pija_Music_Api_Crud/Pija_Music_Api_Crud/controllers:TrajesTipicosController"] = append(beego.GlobalControllerRouter["github.com/KatherineRincon/Pija_Music_Api_Crud/Pija_Music_Api_Crud/controllers:TrajesTipicosController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/KatherineRincon/Pija_Music_Api_Crud/Pija_Music_Api_Crud/controllers:TrajesTipicosController"] = append(beego.GlobalControllerRouter["github.com/KatherineRincon/Pija_Music_Api_Crud/Pija_Music_Api_Crud/controllers:TrajesTipicosController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/KatherineRincon/Pija_Music_Api_Crud/Pija_Music_Api_Crud/controllers:UsuarioController"] = append(beego.GlobalControllerRouter["github.com/KatherineRincon/Pija_Music_Api_Crud/Pija_Music_Api_Crud/controllers:UsuarioController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/KatherineRincon/Pija_Music_Api_Crud/Pija_Music_Api_Crud/controllers:UsuarioController"] = append(beego.GlobalControllerRouter["github.com/KatherineRincon/Pija_Music_Api_Crud/Pija_Music_Api_Crud/controllers:UsuarioController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/KatherineRincon/Pija_Music_Api_Crud/Pija_Music_Api_Crud/controllers:UsuarioController"] = append(beego.GlobalControllerRouter["github.com/KatherineRincon/Pija_Music_Api_Crud/Pija_Music_Api_Crud/controllers:UsuarioController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/KatherineRincon/Pija_Music_Api_Crud/Pija_Music_Api_Crud/controllers:UsuarioController"] = append(beego.GlobalControllerRouter["github.com/KatherineRincon/Pija_Music_Api_Crud/Pija_Music_Api_Crud/controllers:UsuarioController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/KatherineRincon/Pija_Music_Api_Crud/Pija_Music_Api_Crud/controllers:UsuarioController"] = append(beego.GlobalControllerRouter["github.com/KatherineRincon/Pija_Music_Api_Crud/Pija_Music_Api_Crud/controllers:UsuarioController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
