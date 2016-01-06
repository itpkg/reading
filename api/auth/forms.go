package auth

type SignInForm struct {
	RememberMe bool   `form:"remember_me"`
	Email      string `form:"email" binding:"email"`
	Password   string `form:"password" binding:"min=8"`
}
type SignUpForm struct {
	Username             string `form:"username" binding:"min=2,max=20"`
	Email                string `form:"email" binding:"email"`
	Password             string `form:"password" binding:"min=8"`
	PasswordConfirmation string `form:"password_confirmation" binding:"eqfield=Password"`
}

type PasswordForm struct {
	Token                string `form:"token" binding:"required"`
	Password             string `form:"password" binding:"min=8"`
	PasswordConfirmation string `form:"password_confirmation" binding:"eqfield=Password"`
}

type EmailForm struct {
	Email string `form:"email" binding:"email"`
}
