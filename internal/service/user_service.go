package service

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
	"github.com/smustafa-tech/hrms-backend/internal/dto"
	"github.com/smustafa-tech/hrms-backend/internal/models"
	"github.com/smustafa-tech/hrms-backend/internal/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetProfile(userID string) (*models.User, error) {
	return s.repo.FindByID(userID)
}

func (s *UserService) UpdateProfile(userID string, req dto.UpdateProfileRequest) (*models.User, error) {
	updates := map[string]interface{}{}
	if req.FirstName != "" {
		updates["first_name"] = req.FirstName
	}
	if req.MiddleName != "" {
		updates["middle_name"] = req.MiddleName
	}
	if req.LastName != "" {
		updates["last_name"] = req.LastName
	}
	if req.Email != "" {
		updates["email"] = req.Email
	}
	if req.Phone != "" {
		updates["phone"] = req.Phone
	}
	if req.AdharCard != "" {
		updates["adhar_card"] = req.AdharCard
	}
	if req.Designation != "" {
		updates["designation"] = req.Designation
	}
	if req.Bio != "" {
		updates["bio"] = req.Bio
	}

	if len(updates) == 0 {
		return nil, errors.New("no fields to update")
	}

	return s.repo.Update(userID, updates)
}

func (s *UserService) ChangePassword(userID string, req dto.ChangePasswordRequest) error {
	user, err := s.repo.FindByID(userID)
	if err != nil {
		return errors.New("user not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.OldPassword)); err != nil {
		return errors.New("old password is incorrect")
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("failed to hash new password")
	}

	_, err = s.repo.Update(userID, map[string]interface{}{"password": string(hashed)})
	return err
}

func (s *UserService) UploadProfilePhoto(userID, photoPath, photoURL string) (*models.User, error) {
	return s.repo.Update(userID, map[string]interface{}{"photo_path": photoPath, "photo_url": photoURL})
}
