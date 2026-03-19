package service

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	//"log"
	//"strconv"
	"time"

	//"net/smtp"
	"os"
	"portfolio_backend_go/internal/repository"

	"github.com/resend/resend-go/v3"
	"golang.org/x/crypto/bcrypt"
	//"gopkg.in/gomail.v2"
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

htmlBody := `
<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <style>
        body { font-family: 'Inter', -apple-system, sans-serif; background-color: #050505; color: #ffffff; margin: 0; padding: 40px; }
        .wrapper { background-color: #050505; width: 100%; padding: 40px 0; }
        .container { max-width: 480px; margin: 0 auto; background: #111111; border: 1px solid #222222; border-radius: 24px; padding: 40px; }
        
        /* Brand Accent */
        .brand-accent { width: 40px; height: 4px; background: linear-gradient(90deg, #8b5cf6, #d946ef); border-radius: 2px; margin: 0 auto 24px; }
        
        .logo { font-size: 20px; font-weight: 800; color: #ffffff; text-align: center; letter-spacing: 2px; text-transform: uppercase; margin-bottom: 8px; }
        .title { font-size: 24px; font-weight: 700; color: #ffffff; margin-bottom: 12px; text-align: center; letter-spacing: -0.5px; }
        .subtitle { font-size: 15px; color: #a1a1aa; text-align: center; line-height: 1.5; margin-bottom: 32px; }
        
        /* OTP Box with subtle glow */
        .otp-box { 
            background: linear-gradient(145deg, #18181b, #09090b); 
            border: 1px solid #27272a;
            border-radius: 16px; 
            padding: 32px; 
            margin: 0 0 32px 0; 
            text-align: center;
            box-shadow: 0 4px 20px rgba(139, 92, 246, 0.1);
        }
        .otp-code { 
            font-family: 'Monaco', 'Consolas', monospace;
            font-size: 42px; 
            font-weight: 800; 
            color: #a78bfa; /* Soft Purple */
            letter-spacing: 12px; 
            margin: 0; 
            text-shadow: 0 0 15px rgba(139, 92, 246, 0.3);
        }
        
        .timer-row { display: flex; align-items: center; justify-content: center; gap: 8px; color: #f43f5e; font-size: 14px; font-weight: 600; margin-bottom: 32px; text-align:center; }
        
        .divider { height: 1px; background: linear-gradient(90deg, transparent, #27272a, transparent); margin: 32px 0; }
        
        .footer-text { font-size: 12px; color: #52525b; text-align: center; line-height: 1.6; }
        .link { color: #8b5cf6; text-decoration: none; }
    </style>
</head>
<body>
    <div class="wrapper">
        <div class="container">
            <div class="brand-accent"></div>
            <div class="logo">swajan.vercel.app</div>
            <h2 class="title">Verify your identity</h2>
            <p class="subtitle">Enter the following code to finish logging into your dashboard. This request was triggered from a new session.</p>
            
            <div class="otp-box">
                <h1 class="otp-code">` + code + `</h1>
            </div>

            <div class="timer-row">
                <span>⏱️ This code expires in 5 minutes</span>
            </div>

            <div class="divider"></div>
            
            <p class="footer-text">
                If you did not request this code, you can safely ignore this email. 
                For security, do not share this code with anyone.
            </p>
            
            <p class="footer-text" style="margin-top: 20px; font-weight: 600; letter-spacing: 0.5px;">
                &copy; 2026 SWAJAN BARUA PORTFOLIO
            </p>
        </div>
    </div>
</body>
</html>
`





	FromEmail := "onboarding@resend.dev"
	// Password := os.Getenv("SMTP_PASSWORD")
	// smtpHost := os.Getenv("SMTP_HOST")
	// smtpPort := os.Getenv("SMTP_PORT")
	 resend_api_key := os.Getenv("RESEND_API_KEY")

	// //smtp_Port_int,err := strconv.Atoi(smtpPort);
	// if err != nil {
	// 	smtp_Port_int = 587;
	// }
	//address := smtpHost + ":" + smtpPort

	resend_client := resend.NewClient(resend_api_key)

	params := &resend.SendEmailRequest{
		From: FromEmail,
		To: []string{email},
		Subject: "Portfolio Verification",
		Html: htmlBody,

	}

	sent, err := resend_client.Emails.Send(params)
	if err != nil {
		return fmt.Errorf("resend error: %v", err)
	}
	fmt.Printf("📧 OTP Sent Successfully! Message ID: %s\n", sent.Id)
    







	// m := gomail.NewMessage()
	// m.SetHeader("From",FromEmail)
	// m.SetHeader("To",email)
	// m.SetHeader("Subject", "Your Portfolio Admin OTP")
    // m.SetBody("text/html", "Your OTP code is: <b>" + code + "</b>")

	// d := gomail.NewDialer(
    //     smtpHost, 
    //     smtp_Port_int, 
    //     FromEmail, 
    //     Password,
    // )

    // // This is the part that prevents the "Hang"
    // if err := d.DialAndSend(m); err != nil {
    //     log.Printf("❌ SMTP Error: %v", err)
    //     return err
    // }
    // return nil


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
