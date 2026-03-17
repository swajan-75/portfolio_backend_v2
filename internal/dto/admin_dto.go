package dto

type Admin_login_req struct {
	Email string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Login_response struct {
	Token string `json:"token"`
	Eamil string `json:"email"`
	ExpireIn string `json:"expirein"`
}

type Otp_req struct {
	Email string `json:"email" binding:"required"`

}

type Otp_verify struct {
	Email string `json:"email" binding:"required"`
	OTP string `json:"otp" binding:"required"`
}

type firebase_login_resp struct {
	IdToken      string `json:"idToken"`
	Email        string `json:"email"`
	RefreshToken string `json:"refreshToken"`
	ExpiresIn    string `json:"expiresIn"`
	LocalId      string `json:"localId"`
}