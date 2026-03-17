package service

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
)

type Admin_Service struct {
	OtpService *Otp_service	
}

func New_Admin__Service(Otp_service *Otp_service) *Admin_Service{
	return &Admin_Service{OtpService: Otp_service}
}


type FirebaseSignInResponse struct {
	IdToken      string `json:"idToken"`
	Email        string `json:"email"`
	RefreshToken string `json:"refreshToken"`
	ExpiresIn    string `json:"expiresIn"`
	LocalId      string `json:"localId"`
}


func (s *Admin_Service) Login(c context.Context,email string , password string)(error){
	api_key := os.Getenv("FIREBASE_API_KEY")
	url := fmt.Sprintf("https://identitytoolkit.googleapis.com/v1/accounts:signInWithPassword?key=%s",api_key)

	payload,_ :=json.Marshal(map[string]interface{}{
		"email" : email,
		"password" : password,
		"returnSecureToken": true,
	})

	res,err := http.Post(url,"application/json",bytes.NewBuffer(payload))
	if err !=nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK{
		return errors.New("invalid email or password")
	}

	var login_resp FirebaseSignInResponse

	if err := json.NewDecoder(res.Body).Decode(&login_resp); err !=nil {
		return errors.New("failed to parse authentication response")
	}

	// err = s.OtpService.Send_otp(c,email,login_resp.IdToken)
	// if err != nil {
	// 	return errors.New("Failed to Sent OTP")
	// }

	return  s.OtpService.Send_otp(c,email,login_resp.IdToken)


}