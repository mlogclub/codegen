package codegen

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"path"
	"reflect"

	"github.com/mlogclub/codegen/templates"
	"github.com/mlogclub/simple/common/files"
	"github.com/mlogclub/simple/common/strs/strcase"
	"github.com/mlogclub/simple/common/structs"
)

type GenerateStruct struct {
	Name   string
	Fields []GenerateField
}

type GenerateField struct {
	CamelName   string
	NativeField reflect.StructField
}

type InputData struct {
	PkgName   string
	Name      string // FuckShit
	CamelName string // fuckShit
	KebabName string // FuckShit -> fuck-shit
	Fields    []GenerateField
}

// Options for generating code
type Options struct {
	BaseDir string
	PkgName string
	Version int

	Repository bool
	Service    bool
	Controller bool
	WebIndex   bool
	WebEdit    bool
}

// GenerateWithOption generate code with options
func GenerateWithOption(option Options, models ...GenerateStruct) {
	for _, model := range models {
		if option.Repository {
			if err := generateRepository(option.BaseDir, option.PkgName, option.Version, model); err != nil {
				fmt.Println(err)
			}
		}
		if option.Service {
			if err := generateService(option.BaseDir, option.PkgName, option.Version, model); err != nil {
				fmt.Println(err)
			}
		}
		if option.Controller {
			if err := generateController(option.BaseDir, option.PkgName, option.Version, model); err != nil {
				fmt.Println(err)
			}
		}
		if option.WebIndex {
			if err := generateWebIndex(option.BaseDir, option.PkgName, option.Version, model); err != nil {
				fmt.Println(err)
			}
		}
		if option.WebEdit {
			if err := generateWebEdit(option.BaseDir, option.PkgName, option.Version, model); err != nil {
				fmt.Println(err)
			}
		}
	}
}

func Generate(baseDir, pkgName string, version int, models ...GenerateStruct) {
	GenerateWithOption(Options{
		BaseDir:    baseDir,
		PkgName:    pkgName,
		Version:    version,
		Repository: true,
		Service:    true,
		Controller: true,
		WebIndex:   true,
		WebEdit:    true,
	}, models...)
}

func GetGenerateStruct(s interface{}) GenerateStruct {
	structName := structs.StructName(s)
	structFields := structs.StructFields(s)

	var fields []GenerateField
	for _, f := range structFields {
		if f.Anonymous {
			continue
		}
		fields = append(fields, GenerateField{
			CamelName:   strcase.ToLowerCamel(f.Name),
			NativeField: f,
		})
	}

	return GenerateStruct{
		Name:   structName,
		Fields: fields,
	}
}

func generateRepository(baseDir, pkgName string, version int, s GenerateStruct) error {
	var b bytes.Buffer
	err := templates.GetRepositoryTemplate(version).Execute(&b, &InputData{
		PkgName:   pkgName,
		Name:      s.Name,
		CamelName: strcase.ToLowerCamel(s.Name),
		KebabName: strcase.ToKebab(s.Name),
		Fields:    s.Fields,
	})
	if err != nil {
		return err
	}
	c := b.String()

	filename := "/repositories/" + strcase.ToSnake(s.Name+"_repository.go")
	if version > 0 {
		filename = "/internal/repositories/" + strcase.ToSnake(s.Name+"_repository.go")
	}
	p, err := getFilePath(baseDir, filename)
	if err != nil {
		return err
	}
	return writeFile(p, c)
}

func generateService(baseDir, pkgName string, version int, s GenerateStruct) error {
	var b bytes.Buffer
	err := templates.GetServiceTemplate(version).Execute(&b, &InputData{
		PkgName:   pkgName,
		Name:      s.Name,
		CamelName: strcase.ToLowerCamel(s.Name),
		KebabName: strcase.ToKebab(s.Name),
		Fields:    s.Fields,
	})
	if err != nil {
		return err
	}
	c := b.String()

	filename := "/services/" + strcase.ToSnake(s.Name+"_service.go")
	if version > 0 {
		filename = "/internal/services/" + strcase.ToSnake(s.Name+"_service.go")
	}

	p, err := getFilePath(baseDir, filename)
	if err != nil {
		return err
	}
	return writeFile(p, c)
}

func generateController(baseDir, pkgName string, version int, s GenerateStruct) error {
	var b bytes.Buffer
	err := templates.GetControllerTemplate(version).Execute(&b, &InputData{
		PkgName:   pkgName,
		Name:      s.Name,
		CamelName: strcase.ToLowerCamel(s.Name),
		KebabName: strcase.ToKebab(s.Name),
		Fields:    s.Fields,
	})
	if err != nil {
		return err
	}
	c := b.String()

	filename := "/controllers/admin/" + strcase.ToSnake(s.Name+"_controller.go")
	if version > 0 {
		filename = "/internal/controllers/admin/" + strcase.ToSnake(s.Name+"_controller.go")
	}
	p, err := getFilePath(baseDir, filename)
	if err != nil {
		return err
	}
	return writeFile(p, c)
}

func generateWebIndex(baseDir, pkgName string, version int, s GenerateStruct) error {
	var b bytes.Buffer
	err := templates.GetAdminIndexTemplate(version).Execute(&b, &InputData{
		PkgName:   pkgName,
		Name:      s.Name,
		KebabName: strcase.ToKebab(s.Name),
		Fields:    s.Fields,
	})
	if err != nil {
		return err
	}
	c := b.String()

	sub := path.Join("/web/admin/", strcase.ToKebab(s.Name), "index.vue")

	p, err := getFilePath(baseDir, sub)
	if err != nil {
		return err
	}
	return writeFile(p, c)
}

func generateWebEdit(baseDir, pkgName string, version int, s GenerateStruct) error {
	var b bytes.Buffer
	err := templates.GetAdminEditTemplate(version).Execute(&b, &InputData{
		PkgName:   pkgName,
		Name:      s.Name,
		KebabName: strcase.ToKebab(s.Name),
		Fields:    s.Fields,
	})
	if err != nil {
		return err
	}
	c := b.String()

	sub := path.Join("/web/admin/", strcase.ToKebab(s.Name), "components", "Edit.vue")

	p, err := getFilePath(baseDir, sub)
	if err != nil {
		return err
	}
	return writeFile(p, c)
}

func getFilePath(baseDir, sub string) (filepath string, err error) {
	filepath = path.Join(baseDir, sub)
	base := path.Dir(filepath)
	err = os.MkdirAll(base, os.ModePerm)
	return
}

func writeFile(filepath string, content string) error {
	exists, err := files.PathExists(filepath)
	if err != nil {
		return err
	}
	if exists {
		fmt.Println("文件已经存在...", filepath)
		return errors.New("文件已经存在..." + filepath)
	}
	return files.WriteString(filepath, content, true)
}
