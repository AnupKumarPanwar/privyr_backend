package requests

type Signup struct {
	Email string `binding:"required"`
}
