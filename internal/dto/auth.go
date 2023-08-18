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
	ChangePasswordRequest struct {
		OldPassword             string `json:"oldPassword"`
		NewPassword             string `json:"newPassword"`
		ConfirmationNewPassword string `json:"confirmationNewPassword"`
	}
)
