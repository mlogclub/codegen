package codegen

import "html/template"

var serviceTmpl = template.Must(template.New("service").Parse(`package services

import (
	"{{.PkgName}}/model"
	"{{.PkgName}}/repositories"

	"github.com/mlogclub/simple/sqls"
	"github.com/mlogclub/simple/web/params"
)

var {{.Name}}Service = new{{.Name}}Service()

func new{{.Name}}Service() *{{.CamelName}}Service {
	return &{{.CamelName}}Service{}
}

type {{.CamelName}}Service struct {
}

func (s *{{.CamelName}}Service) Get(id int64) *model.{{.Name}} {
	return repositories.{{.Name}}Repository.Get(sqls.DB(), id)
}

func (s *{{.CamelName}}Service) Take(where ...interface{}) *model.{{.Name}} {
	return repositories.{{.Name}}Repository.Take(sqls.DB(), where...)
}

func (s *{{.CamelName}}Service) Find(cnd *sqls.Cnd) []model.{{.Name}} {
	return repositories.{{.Name}}Repository.Find(sqls.DB(), cnd)
}

func (s *{{.CamelName}}Service) FindOne(cnd *sqls.Cnd) *model.{{.Name}} {
	return repositories.{{.Name}}Repository.FindOne(sqls.DB(), cnd)
}

func (s *{{.CamelName}}Service) FindPageByParams(params *params.QueryParams) (list []model.{{.Name}}, paging *sqls.Paging) {
	return repositories.{{.Name}}Repository.FindPageByParams(sqls.DB(), params)
}

func (s *{{.CamelName}}Service) FindPageByCnd(cnd *sqls.Cnd) (list []model.{{.Name}}, paging *sqls.Paging) {
	return repositories.{{.Name}}Repository.FindPageByCnd(sqls.DB(), cnd)
}

func (s *{{.CamelName}}Service) Count(cnd *sqls.Cnd) int64 {
	return repositories.{{.Name}}Repository.Count(sqls.DB(), cnd)
}

func (s *{{.CamelName}}Service) Create(t *model.{{.Name}}) error {
	return repositories.{{.Name}}Repository.Create(sqls.DB(), t)
}

func (s *{{.CamelName}}Service) Update(t *model.{{.Name}}) error {
	return repositories.{{.Name}}Repository.Update(sqls.DB(), t)
}

func (s *{{.CamelName}}Service) Updates(id int64, columns map[string]interface{}) error {
	return repositories.{{.Name}}Repository.Updates(sqls.DB(), id, columns)
}

func (s *{{.CamelName}}Service) UpdateColumn(id int64, name string, value interface{}) error {
	return repositories.{{.Name}}Repository.UpdateColumn(sqls.DB(), id, name, value)
}

func (s *{{.CamelName}}Service) Delete(id int64) {
	repositories.{{.Name}}Repository.Delete(sqls.DB(), id)
}

`))
