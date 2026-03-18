package service

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"

	//"net/smtp"
	"os"
	"portfolio_backend_go/internal/repository"

	"golang.org/x/crypto/bcrypt"
	"gopkg.in/gomail.v2"
)

type Otp_service struct {
	Repo *repository.OTP_repo
}

func New_Otp_service(repo *repository.OTP_repo) *Otp_service {
	return &Otp_service{Repo: repo}
}

func (s *Otp_service) Send_otp(c context.Context, email string, IdToken string) error {
	otp := make([]byte, 3)
	rand.Read(otp)
	code := fmt.Sprintf("%06d", int(otp[0])%10*100000+int(otp[1])%10*10000+int(otp[2])%10*1000)

	





	FromEmail := os.Getenv("SMTP_EMAIL")
	Password := os.Getenv("SMTP_PASSWORD")
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")
	smtp_Port_int,err := strconv.Atoi(smtpPort);
	if err != nil {
		smtp_Port_int = 587;
	}
	//address := smtpHost + ":" + smtpPort

	m := gomail.NewMessage()
	m.SetHeader("From",FromEmail)
	m.SetHeader("To",email)
	m.SetHeader("Subject", "Your Portfolio Admin OTP")
    m.SetBody("text/html", "Your OTP code is: <b>" + code + "</b>")

	d := gomail.NewDialer(
        smtpHost, 
        smtp_Port_int, 
        FromEmail, 
        Password,
    )

    // This is the part that prevents the "Hang"
    if err := d.DialAndSend(m); err != nil {
        log.Printf("❌ SMTP Error: %v", err)
        return err
    }
    return nil


	//fmt.Printf("DEBUG: Host is |%s|\n", smtpHost)
	// fmt.Printf("DEBUG: Port is |%s|\n", smtpPort)

	// auth := smtp.PlainAuth(
	// 	"",
	// 	FromEmail,
	// 	Password,
	// 	smtpHost,
	// )
	// msg := []byte("Subject: Portfolio Login OTP\r\n\r\nYour code is: " + code)

	// err := smtp.SendMail(address, auth, FromEmail, []string{email}, msg)
	// if err != nil {
	// 	return errors.New("unable to send OTP email. please check your network or configuration")

	// }
	//safeEmail := base64.StdEncoding.EncodeToString([]byte(email))

	safeEmail := safe_email(email)

	return s.Repo.Save_otp(c, safeEmail, code,IdToken)

}

func safe_email(email string) string {
	return base64.StdEncoding.EncodeToString([]byte(email))
}

func (s *Otp_service) Verify_otp(c context.Context, email string, input_otp string) (string , error) {
	SafeEmail := safe_email(email)
	data, err := s.Repo.Get_otp(c, SafeEmail)
	if err != nil {
		return "",errors.New("OTP Not Found")
	}

	locked_until := int64(data["locked_time"].(float64))
	if time.Now().Unix() < locked_until {
		remaining := (locked_until - time.Now().Unix()) / 60
		return "",errors.New("Try Again After " + fmt.Sprintf("%d", remaining) + " minutes")
	}

	expiresAt := int64(data["expires_at"].(float64))
	if expiresAt < time.Now().Unix() {
		s.Repo.Delete_otp(c,SafeEmail)
	}

	hash_otp := data["code"].(string)

	err = bcrypt.CompareHashAndPassword([]byte(hash_otp),[]byte(input_otp))

	if err != nil {
		current_try := int(data["failed_count"].(float64))+1
	
		if current_try > 5 {
			locked_until := time.Now().Add(30 * time.Minute).Unix()
			
			s.Repo.Update_Try(c,SafeEmail,int64(current_try),locked_until)

			return "",fmt.Errorf("too many failed attempts. you are locked for %d minutes", (locked_until-time.Now().Unix())/60)
		}
		s.Repo.Update_Try(c,SafeEmail,int64(current_try),0)
		return "",fmt.Errorf("invalid OTP. %d attempts remaining", 5-current_try)
	}

	idToken, ok := data["IdToken"].(string) 
    if !ok {
        return "", errors.New("failed to retrieve session token")
    }
	s.Repo.Delete_otp(c,SafeEmail)
	return idToken,nil


}
