package usecase_intf

type (
	ParamsRegister struct {
		Name            string
		Email           string
		Password        string
		ConfirmPassword string
	}

	ParamsLogin struct {
		Email    string
		Password string
	}
)

type (
	ResultLogin struct {
		Name        string `json:"name"`
		Role        string `json:"role"`
		Email       string `json:"email"`
		AccessToken string `json:"access_token"`
	}
)

type AuthUsecase interface {
	Register(p *ParamsRegister) (int64, error)
	Login(p *ParamsLogin) (*ResultLogin, error)
}
