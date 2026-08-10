package repository

import (
	"github.com/smustafa-tech/hrms-backend/internal/models"
	"gorm.io/gorm"
)

type AttendanceRepository struct {
	db *gorm.DB
}

func NewAttendanceRepository(db *gorm.DB) *AttendanceRepository {
	return &AttendanceRepository{db: db}
}

func (r *AttendanceRepository) Create(a *models.Attendance) error {
	return r.db.Create(a).Error
}

func (r *AttendanceRepository) Save(a *models.Attendance) error {
	return r.db.Save(a).Error
}

func (r *AttendanceRepository) FindByUserAndDate(userID, date string) (*models.Attendance, error) {
	var a models.Attendance
	err := r.db.Where("user_id = ? AND date = ?", userID, date).First(&a).Error
	if err != nil {
		return nil, err
	}
	return &a, nil
}

func (r *AttendanceRepository) FindByID(id string) (*models.Attendance, error) {
	var a models.Attendance
	err := r.db.Where("id = ?", id).First(&a).Error
	if err != nil {
		return nil, err
	}
	return &a, nil
}

func (r *AttendanceRepository) FindByUserID(userID string) ([]models.Attendance, error) {
	var records []models.Attendance
	err := r.db.Where("user_id = ?", userID).Order("date desc").Find(&records).Error
	return records, err
}

func (r *AttendanceRepository) FindTodayAll(date string) ([]models.Attendance, error) {
	var records []models.Attendance
	err := r.db.Where("date = ?", date).Find(&records).Error
	return records, err
}

func (r *AttendanceRepository) FindAll() ([]models.Attendance, error) {
	var records []models.Attendance
	err := r.db.Order("date desc").Find(&records).Error
	return records, err
}
