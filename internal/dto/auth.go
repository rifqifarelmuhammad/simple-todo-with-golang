package dto

type (
	AuthRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	RegistrationRequest struct {
		AuthRequest
		ConfirmationPassword string `json:"confirmationPassword"`
	}
)
