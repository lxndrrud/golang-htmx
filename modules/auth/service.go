package auth

import "html/template"

type authService struct {
	authDataService     *authDataService
	authTemplateService *authTemplateService
}

func newAuthService(authDataService *authDataService,
	authTemplateService *authTemplateService) *authService {
	return &authService{authDataService, authTemplateService}
}

func (as authService) LoginPage() (*template.Template, error) {
	return as.authTemplateService.LoginPageTemplate()
}

func (as authService) LoginAction(login, password string) (*user, error) {
	user, err := as.authDataService.LoginAction(login, password)
	return user, err
}
