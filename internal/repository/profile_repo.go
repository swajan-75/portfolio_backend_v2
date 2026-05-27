package repository

import (
	"context"
	"portfolio_backend_go/internal/models"

	"firebase.google.com/go/v4/db"
)

type Profile_repo struct {
	Client *db.Client
}

func New_Profile_repo(client *db.Client) *Profile_repo {
	return &Profile_repo{Client: client}
}

func (r *Profile_repo) Get_Profile(ctx context.Context) (*models.Profile, error) {
	var profile models.Profile
	err := r.Client.NewRef("profile").Get(ctx, &profile)
	if err != nil {
		return nil, err
	}
	return &profile, nil
}

func (r *Profile_repo) Update_Profile(ctx context.Context, profile models.Profile) error {
	return r.Client.NewRef("profile").Set(ctx, profile)
}
