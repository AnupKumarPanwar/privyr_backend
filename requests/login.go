package requests

type Login struct {
	Email string `binding:"required"`
}
