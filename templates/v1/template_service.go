package v1

import "html/template"

var ServiceTemplate = template.Must(template.New("service").Parse(`package services

import (
	"{{.PkgName}}/internal/models"
	"{{.PkgName}}/internal/repositories"

	"github.com/mlogclub/simple/sqls"
	"github.com/mlogclub/simple/web/params"
)

var {{.Name}}Service = new{{.Name}}Service()

func new{{.Name}}Service() *{{.CamelName}}Service {
	return &{{.CamelName}}Service{}
}

type {{.CamelName}}Service struct {
}

func (s *{{.CamelName}}Service) Get(id int64) *models.{{.Name}} {
	return repositories.{{.Name}}Repository.Get(sqls.DB(), id)
}

func (s *{{.CamelName}}Service) Take(where ...interface{}) *models.{{.Name}} {
	return repositories.{{.Name}}Repository.Take(sqls.DB(), where...)
}

func (s *{{.CamelName}}Service) Find(cnd *sqls.Cnd) []models.{{.Name}} {
	return repositories.{{.Name}}Repository.Find(sqls.DB(), cnd)
}

func (s *{{.CamelName}}Service) FindOne(cnd *sqls.Cnd) *models.{{.Name}} {
	return repositories.{{.Name}}Repository.FindOne(sqls.DB(), cnd)
}

func (s *{{.CamelName}}Service) FindPageByParams(params *params.QueryParams) (list []models.{{.Name}}, paging *sqls.Paging) {
	return repositories.{{.Name}}Repository.FindPageByParams(sqls.DB(), params)
}

func (s *{{.CamelName}}Service) FindPageByCnd(cnd *sqls.Cnd) (list []models.{{.Name}}, paging *sqls.Paging) {
	return repositories.{{.Name}}Repository.FindPageByCnd(sqls.DB(), cnd)
}

func (s *{{.CamelName}}Service) Count(cnd *sqls.Cnd) int64 {
	return repositories.{{.Name}}Repository.Count(sqls.DB(), cnd)
}

func (s *{{.CamelName}}Service) Create(t *models.{{.Name}}) error {
	return repositories.{{.Name}}Repository.Create(sqls.DB(), t)
}

func (s *{{.CamelName}}Service) Update(t *models.{{.Name}}) error {
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
