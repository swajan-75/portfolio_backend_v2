package repository

import (
	"context"
	"time"

	"firebase.google.com/go/v4/db"
	"golang.org/x/crypto/bcrypt"
)

type OTP_repo struct {
	Client *db.Client
}


func New_OTP_repo(client *db.Client) *OTP_repo {
	return &OTP_repo{Client: client}
}

func (r *OTP_repo) Save_otp(c context.Context,email string , otp string, IdToken string) error {

	hash_otp,_ := bcrypt.GenerateFromPassword([]byte(otp),bcrypt.DefaultCost)


	data := map[string]interface{}{
		"code" : string(hash_otp),
		"expires_at" : time.Now().Add(5*time.Minute).Unix(),
		"failed_count" : 0,
		"locked_time" : 0,
		"IdToken" : IdToken,
	}

	return  r.Client.NewRef("tmp_otp/"+email).Set(c,data)
}

func (r *OTP_repo) Update_Try(c context.Context, email string , count int64 , lockUntil int64) error {
	updates := map[string]interface{}{
		"failed_count" : count,
		"locaked_time" : lockUntil,
	}
	return r.Client.NewRef("tmp_otp/"+email).Update(c,updates)

}


func (r *OTP_repo) Get_otp(c context.Context , email string) (map[string]interface{},error){
	var data map[string]interface{}
	err := r.Client.NewRef("tmp_otp/"+email).Get(c,&data)

	return data,err
}

func (r *OTP_repo) Delete_otp(c context.Context, email string) error {
	return r.Client.NewRef("tmp_otp/"+email).Delete(c)
}
