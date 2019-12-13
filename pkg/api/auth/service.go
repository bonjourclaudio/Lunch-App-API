package auth

type AuthService interface {
	Login() *OAuth
}

type authService struct {
	OAuth *OAuth
}

func NewAuthService(oa *OAuth) AuthService {
	return  &authService {
		OAuth: oa,
	}
}

func (a *authService) Login() *OAuth{

	return a.OAuth
}