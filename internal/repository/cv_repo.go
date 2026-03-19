package repository

import (
	"context"
	"firebase.google.com/go/v4/db"
)

type CV_repo struct {
	Client *db.Client
}

func New_CV_repo(client *db.Client) *CV_repo {
	return &CV_repo{Client: client}
}

func (r *CV_repo) Save_CV(ctx context.Context, cv interface{}) error {
	_, err := r.Client.NewRef("cvs").Push(ctx, cv)
	return err
}

func (r *CV_repo) Get_All_CVs(ctx context.Context) (map[string]interface{}, error) {
	var cvs map[string]interface{}
	if err := r.Client.NewRef("cvs").Get(ctx, &cvs); err != nil {
		return nil, err
	}
	return cvs, nil
}

func (r *CV_repo) Set_Active_CV(ctx context.Context, id string) error {
	// first set all to inactive
	cvs, err := r.Get_All_CVs(ctx)
	if err != nil {
		return err
	}
	for key := range cvs {
		if err := r.Client.NewRef("cvs/"+key+"/active").Set(ctx, false); err != nil {
			return err
		}
	}
	// then set selected one to active
	return r.Client.NewRef("cvs/" + id + "/active").Set(ctx, true)
}

func (r *CV_repo) Delete_CV(ctx context.Context, id string) error {
	return r.Client.NewRef("cvs/" + id).Delete(ctx)
}

func (r *CV_repo) Get_Active_CV(ctx context.Context) (map[string]interface{}, error) {
	var cvs map[string]interface{}
	if err := r.Client.NewRef("cvs").Get(ctx, &cvs); err != nil {
		return nil, err
	}
	for _, v := range cvs {
		cv, ok := v.(map[string]interface{})
		if !ok {
			continue
		}
		if active, ok := cv["active"].(bool); ok && active {
			return cv, nil
		}
	}
	return nil, nil
}