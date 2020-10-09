package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["supervise/controllers:SysController"] = append(beego.GlobalControllerRouter["supervise/controllers:SysController"],
        beego.ControllerComments{
            Method: "Cmd",
            Router: "/cmd",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["supervise/controllers:SysController"] = append(beego.GlobalControllerRouter["supervise/controllers:SysController"],
        beego.ControllerComments{
            Method: "Heartbeat",
            Router: "/heartbeat",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["supervise/controllers:SysController"] = append(beego.GlobalControllerRouter["supervise/controllers:SysController"],
        beego.ControllerComments{
            Method: "Inspection",
            Router: "/inspection",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["supervise/controllers:SysController"] = append(beego.GlobalControllerRouter["supervise/controllers:SysController"],
        beego.ControllerComments{
            Method: "InspectionGet",
            Router: "/inspection/:taskId",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["supervise/controllers:SysController"] = append(beego.GlobalControllerRouter["supervise/controllers:SysController"],
        beego.ControllerComments{
            Method: "InspectionDelete",
            Router: "/inspection/:taskId",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["supervise/controllers:TransaferController"] = append(beego.GlobalControllerRouter["supervise/controllers:TransaferController"],
        beego.ControllerComments{
            Method: "SensitivWords",
            Router: "/sensitivWords",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
