package v1

import "html/template"

var ControllerTemplate = template.Must(template.New("controller").Parse(`package admin

import (
	"{{.PkgName}}/model"
	"{{.PkgName}}/services"
	"strconv"

	"github.com/kataras/iris/v12"
	"github.com/mlogclub/simple/web"
	"github.com/mlogclub/simple/web/params"
)

type {{.Name}}Controller struct {
	Ctx iris.Context
}

func (c *{{.Name}}Controller) GetBy(id int64) *web.JsonResult {
	t := services.{{.Name}}Service.Get(id)
	if t == nil {
		return web.JsonErrorMsg("Not found, id=" + strconv.FormatInt(id, 10))
	}
	return web.JsonData(t)
}

func (c *{{.Name}}Controller) AnyList() *web.JsonResult {
	list, paging := services.{{.Name}}Service.FindPageByCnd(params.NewPagedSqlCnd(c.Ctx,
		params.QueryFilter{
			ParamName: "id",
		},
	).Desc("id"))
	return web.JsonData(&web.PageResult{Results: list, Page: paging})
}

func (c *{{.Name}}Controller) PostCreate() *web.JsonResult {
	t := &model.{{.Name}}{}
	if err := params.ReadForm(c.Ctx, t); err != nil {
		return web.JsonErrorMsg(err.Error())
	}

	if err := services.{{.Name}}Service.Create(t); err != nil {
		return web.JsonErrorMsg(err.Error())
	}
	return web.JsonData(t)
}

func (c *{{.Name}}Controller) PostUpdate() *web.JsonResult {
	id, err := params.FormValueInt64(c.Ctx, "id")
	if err != nil {
		return web.JsonErrorMsg(err.Error())
	}
	t := services.{{.Name}}Service.Get(id)
	if t == nil {
		return web.JsonErrorMsg("entity not found")
	}

	if err := params.ReadForm(c.Ctx, t); err != nil {
		return web.JsonErrorMsg(err.Error())
	}

	if err := services.{{.Name}}Service.Update(t); err != nil {
		return web.JsonErrorMsg(err.Error())
	}
	return web.JsonData(t)
}

func (c *{{.Name}}Controller) PostDelete() *web.JsonResult {
	ids := params.GetInt64Arr(c.Ctx, "ids")
	if len(ids) == 0 {
		return web.JsonErrorMsg("delete ids is empty")
	}
	for _, id := range ids {
		services.{{.Name}}Service.Delete(id)
	}
	return web.JsonSuccess()
}

`))
