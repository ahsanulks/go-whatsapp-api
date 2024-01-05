package authentication

type LoginParam struct {
	ID string `validate:"required"`
}
