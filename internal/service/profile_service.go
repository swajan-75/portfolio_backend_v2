package service

import (
	"context"
	"fmt"
	"portfolio_backend_go/internal/models"
	"portfolio_backend_go/internal/repository"
)

type Profile_Service struct {
	Repo *repository.Profile_repo
}

func New_Profile_Service(repo *repository.Profile_repo) *Profile_Service {
	return &Profile_Service{Repo: repo}
}

func (s *Profile_Service) GetProfile(ctx context.Context) (*models.Profile, error) {
	return s.Repo.Get_Profile(ctx)
}

func (s *Profile_Service) UpdateProfile(ctx context.Context, profile models.Profile) error {
	return s.Repo.Update_Profile(ctx, profile)
}

// ─── Skill Categories ─────────────────────────────────────────────────────────

func (s *Profile_Service) AddSkillCategory(ctx context.Context, cat models.SkillCategory) error {
	profile, err := s.GetProfile(ctx)
	if err != nil {
		return err
	}
	profile.SkillCategories = append(profile.SkillCategories, cat)
	return s.UpdateProfile(ctx, *profile)
}

func (s *Profile_Service) UpdateSkillCategory(ctx context.Context, index int, cat models.SkillCategory) error {
	profile, err := s.GetProfile(ctx)
	if err != nil {
		return err
	}
	if index < 0 || index >= len(profile.SkillCategories) {
		return fmt.Errorf("category index out of bounds")
	}
	profile.SkillCategories[index] = cat
	return s.UpdateProfile(ctx, *profile)
}

func (s *Profile_Service) DeleteSkillCategory(ctx context.Context, index int) error {
	profile, err := s.GetProfile(ctx)
	if err != nil {
		return err
	}
	if index < 0 || index >= len(profile.SkillCategories) {
		return fmt.Errorf("category index out of bounds")
	}
	profile.SkillCategories = append(profile.SkillCategories[:index], profile.SkillCategories[index+1:]...)
	return s.UpdateProfile(ctx, *profile)
}

// ─── Skills inside a Category ─────────────────────────────────────────────────

func (s *Profile_Service) AddSkillToCategory(ctx context.Context, catIndex int, skill models.SkillItem) error {
	profile, err := s.GetProfile(ctx)
	if err != nil {
		return err
	}
	if catIndex < 0 || catIndex >= len(profile.SkillCategories) {
		return fmt.Errorf("category index out of bounds")
	}
	profile.SkillCategories[catIndex].Skills = append(profile.SkillCategories[catIndex].Skills, skill)
	return s.UpdateProfile(ctx, *profile)
}

func (s *Profile_Service) UpdateSkillInCategory(ctx context.Context, catIndex, skillIndex int, skill models.SkillItem) error {
	profile, err := s.GetProfile(ctx)
	if err != nil {
		return err
	}
	if catIndex < 0 || catIndex >= len(profile.SkillCategories) {
		return fmt.Errorf("category index out of bounds")
	}
	cat := &profile.SkillCategories[catIndex]
	if skillIndex < 0 || skillIndex >= len(cat.Skills) {
		return fmt.Errorf("skill index out of bounds")
	}
	cat.Skills[skillIndex] = skill
	return s.UpdateProfile(ctx, *profile)
}

func (s *Profile_Service) DeleteSkillFromCategory(ctx context.Context, catIndex, skillIndex int) error {
	profile, err := s.GetProfile(ctx)
	if err != nil {
		return err
	}
	if catIndex < 0 || catIndex >= len(profile.SkillCategories) {
		return fmt.Errorf("category index out of bounds")
	}
	cat := &profile.SkillCategories[catIndex]
	if skillIndex < 0 || skillIndex >= len(cat.Skills) {
		return fmt.Errorf("skill index out of bounds")
	}
	cat.Skills = append(cat.Skills[:skillIndex], cat.Skills[skillIndex+1:]...)
	return s.UpdateProfile(ctx, *profile)
}

// ─── Contacts / Socials ───────────────────────────────────────────────────────

func (s *Profile_Service) AddContact(ctx context.Context, contact models.Social) error {
	profile, err := s.GetProfile(ctx)
	if err != nil {
		return err
	}
	profile.Socials = append(profile.Socials, contact)
	return s.UpdateProfile(ctx, *profile)
}

func (s *Profile_Service) UpdateContact(ctx context.Context, index int, contact models.Social) error {
	profile, err := s.GetProfile(ctx)
	if err != nil {
		return err
	}
	if index < 0 || index >= len(profile.Socials) {
		return fmt.Errorf("contact index out of bounds")
	}
	profile.Socials[index] = contact
	return s.UpdateProfile(ctx, *profile)
}

func (s *Profile_Service) DeleteContact(ctx context.Context, index int) error {
	profile, err := s.GetProfile(ctx)
	if err != nil {
		return err
	}
	if index < 0 || index >= len(profile.Socials) {
		return fmt.Errorf("contact index out of bounds")
	}
	profile.Socials = append(profile.Socials[:index], profile.Socials[index+1:]...)
	return s.UpdateProfile(ctx, *profile)
}
