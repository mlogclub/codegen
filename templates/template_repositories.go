package templates

import "html/template"

var RepositoryTmpl = template.Must(template.New("repository").Parse(`package repositories

import (
	"{{.PkgName}}/model"

	"github.com/mlogclub/simple/sqls"
	"github.com/mlogclub/simple/web/params"
	"gorm.io/gorm"
)

var {{.Name}}Repository = new{{.Name}}Repository()

func new{{.Name}}Repository() *{{.CamelName}}Repository {
	return &{{.CamelName}}Repository{}
}

type {{.CamelName}}Repository struct {
}

func (r *{{.CamelName}}Repository) Get(db *gorm.DB, id int64) *model.{{.Name}} {
	ret := &model.{{.Name}}{}
	if err := db.First(ret, "id = ?", id).Error; err != nil {
		return nil
	}
	return ret
}

func (r *{{.CamelName}}Repository) Take(db *gorm.DB, where ...interface{}) *model.{{.Name}} {
	ret := &model.{{.Name}}{}
	if err := db.Take(ret, where...).Error; err != nil {
		return nil
	}
	return ret
}

func (r *{{.CamelName}}Repository) Find(db *gorm.DB, cnd *sqls.Cnd) (list []model.{{.Name}}) {
	cnd.Find(db, &list)
	return
}

func (r *{{.CamelName}}Repository) FindOne(db *gorm.DB, cnd *sqls.Cnd) *model.{{.Name}} {
	ret := &model.{{.Name}}{}
	if err := cnd.FindOne(db, &ret); err != nil {
		return nil
	}
	return ret
}

func (r *{{.CamelName}}Repository) FindPageByParams(db *gorm.DB, params *params.QueryParams) (list []model.{{.Name}}, paging *sqls.Paging) {
	return r.FindPageByCnd(db, &params.Cnd)
}

func (r *{{.CamelName}}Repository) FindPageByCnd(db *gorm.DB, cnd *sqls.Cnd) (list []model.{{.Name}}, paging *sqls.Paging) {
	cnd.Find(db, &list)
	count := cnd.Count(db, &model.{{.Name}}{})

	paging = &sqls.Paging{
		Page:  cnd.Paging.Page,
		Limit: cnd.Paging.Limit,
		Total: count,
	}
	return
}

func (r *{{.CamelName}}Repository) FindBySql(db *gorm.DB, sqlStr string, paramArr... interface{}) (list []model.{{.Name}}) {
	db.Raw(sqlStr, paramArr...).Scan(&list)
	return
}

func (r *{{.CamelName}}Repository) CountBySql(db *gorm.DB, sqlStr string, paramArr... interface{}) (count int64) {
	db.Raw(sqlStr, paramArr...).Count(&count)
	return
}

func (r *{{.CamelName}}Repository) Count(db *gorm.DB, cnd *sqls.Cnd) int64 {
	return cnd.Count(db, &model.{{.Name}}{})
}

func (r *{{.CamelName}}Repository) Create(db *gorm.DB, t *model.{{.Name}}) (err error) {
	err = db.Create(t).Error
	return
}

func (r *{{.CamelName}}Repository) Update(db *gorm.DB, t *model.{{.Name}}) (err error) {
	err = db.Save(t).Error
	return
}

func (r *{{.CamelName}}Repository) Updates(db *gorm.DB, id int64, columns map[string]interface{}) (err error) {
	err = db.Model(&model.{{.Name}}{}).Where("id = ?", id).Updates(columns).Error
	return
}

func (r *{{.CamelName}}Repository) UpdateColumn(db *gorm.DB, id int64, name string, value interface{}) (err error) {
	err = db.Model(&model.{{.Name}}{}).Where("id = ?", id).UpdateColumn(name, value).Error
	return
}

func (r *{{.CamelName}}Repository) Delete(db *gorm.DB, id int64) {
	db.Delete(&model.{{.Name}}{}, "id = ?", id)
}

`))
