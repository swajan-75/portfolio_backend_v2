package repository

import (
	"context"

	"firebase.google.com/go/v4/db"
)

type Project_repo struct {
	Client *db.Client
}

func New_Project_repo(client *db.Client) *Project_repo {
	return &Project_repo{Client: client}
}

func (r *Project_repo) Push_Project(ctx context.Context, project interface{}) error {
	ref := r.Client.NewRef("projects")
	_, err := ref.Push(ctx, project)
	return err
}

func (r *Project_repo) Check_Is_Exists(ctx context.Context, key string) (bool, error) {
	var data interface{}
	err := r.Client.NewRef("projects/" + key).Get(ctx, &data)
	if err != nil {
		return false, err
	}
	return data != nil, nil
}

func (r *Project_repo) Set_Project(ctx context.Context, slug string, project interface{}) error {
	return r.Client.NewRef("projects/" + slug).Set(ctx, project)
}

func (r *Project_repo) Get_All_Projects(ctx context.Context) (map[string]interface{}, error) {
	var projects map[string]interface{}
	ref := r.Client.NewRef("projects")
	if err := ref.Get(ctx, &projects); err != nil {
		return nil, err
	}
	return projects, nil
}

func (r *Project_repo) Get_Project_By_Slug(ctx context.Context, slug string) (map[string]interface{}, error) {
	var project map[string]interface{}
	err := r.Client.NewRef("projects/" + slug).Get(ctx, &project)
	if err != nil {
		return nil, err
	}
	return project, nil
}

func (r *Project_repo) Update_Project(ctx context.Context, slug string, updates map[string]interface{}) error {
	return r.Client.NewRef("projects/" + slug).Update(ctx, updates)
}

func (r *Project_repo) Delete_Project(ctx context.Context, slug string) error {
	return r.Client.NewRef("projects/" + slug).Delete(ctx)
}