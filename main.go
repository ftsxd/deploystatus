package main

import (
	"fmt"
	deploymentStatus "testapi/deploystatus"
	"time"

	"github.com/astaxie/beego"
)

type TaskController struct {
	beego.Controller
}

// 审批
func (c *TaskController) Approval() {
	//方式一
	c.Ctx.Request.ParseForm()
	fmt.Println(c.Ctx.Request.Form)
	c.Ctx.WriteString("")
	fmt.Println(c.Ctx.Request.FormValue("user"))

	//c.Ctx.Input.CopyBody(10 * 1024 * 1024)
	//var m map[string]interface{}
	//json.Unmarshal(c.Ctx.Input.RequestBody, &m)
	//fmt.Println(m)
	//fmt.Printf("%#v\n", m)
	//c.Ctx.WriteString("")
	time.Sleep(10 * time.Second)
	fmt.Println("取之")
	c.Ctx.WriteString("Approval")

}

// 查询
func (c *TaskController) Query() {
	c.Ctx.Request.ParseForm()
	fmt.Println(c.Ctx.Request.Form)
	c.Ctx.WriteString("")

	c.Ctx.WriteString("access")
}

// 查看集群状态
func (c *TaskController) Deployhealth() {
	aa := deploymentStatus.Testdeploy("default")
	fmt.Println(aa)
	c.Ctx.WriteString(aa)

}

// 自定义控制器方法和路由规则
func main() {
	//自动路由
	//url => 控制controller/action
	// task =》 Taskcontroller
	//add +  Add方法
	beego.AutoRouter(&TaskController{})
	beego.Run("0.0.0.0:8080")

}
