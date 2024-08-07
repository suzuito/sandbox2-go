package entity

import "github.com/go-playground/validator/v10"

var validate *validator.Validate

func init() {
	initCELEnvAdminPrincipalAccessToken()
	initCELEnvUserPrincipalAccessToken()
	validate = validator.New()
}
