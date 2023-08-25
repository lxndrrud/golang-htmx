package auth

import (
	"fmt"
	"html/template"
)

type authTemplateService struct{}

func newAuthTemplateService() *authTemplateService {
	return &authTemplateService{}
}

func (ats authTemplateService) LoginPageTemplate() (*template.Template, error) {
	template, err := template.ParseFiles(
		"shared/templates/layouts/basic-layout.html",
		"shared/templates/auth/auth-header.html",
		"modules/auth/templates/pages/login.html")
	if err != nil {
		err = fmt.Errorf("authTemplateService - LoginPageTemplate: %v", err.Error())
		return template, err
	}
	return template, err
}
