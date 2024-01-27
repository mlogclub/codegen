package templates

import (
	"html/template"

	v0 "github.com/mlogclub/codegen/templates/v0"
	v1 "github.com/mlogclub/codegen/templates/v1"
)

func GetRepositoryTemplate(version int) *template.Template {
	if version == 1 {
		return v1.RepositoryTemplate
	}
	return v0.RepositoryTemplate
}

func GetServiceTemplate(version int) *template.Template {
	if version == 1 {
		return v1.ServiceTemplate
	}
	return v0.ServiceTemplate
}

func GetControllerTemplate(version int) *template.Template {
	if version == 1 {
		return v1.ControllerTemplate
	}
	return v0.ControllerTemplate
}

func GetAdminIndexTemplate(version int) *template.Template {
	if version == 1 {
		return v1.AdminIndexTemplate
	}
	return v0.AdminIndexTemplate
}

func GetAdminEditTemplate(version int) *template.Template {
	if version == 1 {
		return v1.AdminEditTemplate
	}
	return v0.AdminEditTemplate
}
