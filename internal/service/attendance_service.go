package service

import (
	"errors"
	"time"

	"github.com/smustafa-tech/hrms-backend/internal/dto"
	"github.com/smustafa-tech/hrms-backend/internal/models"
	"github.com/smustafa-tech/hrms-backend/internal/repository"
)

type AttendanceService struct {
	repo *repository.AttendanceRepository
}

func NewAttendanceService(repo *repository.AttendanceRepository) *AttendanceService {
	return &AttendanceService{repo: repo}
}

func todayDate() string {
	return time.Now().Format("2006-01-02")
}

func currentTime() string {
	return time.Now().Format("15:04:05")
}

// Employee self check-in / check-out
func (s *AttendanceService) MarkAttendance(req dto.MarkAttendanceRequest) (*models.Attendance, error) {
	date := todayDate()

	existing, err := s.repo.FindByUserAndDate(req.UserID, date)

	switch req.Action {
	case "checkIn":
		if err == nil && existing.CheckIn != "" {
			return nil, errors.New("already checked in today")
		}
		if err != nil {
			// create new record
			a := &models.Attendance{
				UserID:  req.UserID,
				Date:    date,
				Status:  "present",
				CheckIn: currentTime(),
			}
			if err := s.repo.Create(a); err != nil {
				return nil, err
			}
			return a, nil
		}
		existing.CheckIn = currentTime()
		existing.Status = "present"
		if err := s.repo.Save(existing); err != nil {
			return nil, err
		}
		return existing, nil

	case "checkOut":
		if err != nil {
			return nil, errors.New("no check-in record found for today")
		}
		existing.CheckOut = currentTime()
		if err := s.repo.Save(existing); err != nil {
			return nil, err
		}
		return existing, nil

	default:
		return nil, errors.New("invalid action, use checkIn or checkOut")
	}
}

// Admin / HR mark attendance for any employee
func (s *AttendanceService) AdminMarkAttendance(req dto.AdminMarkAttendanceRequest) (*models.Attendance, error) {
	date := todayDate()
	status := req.Status
	if status == "" {
		status = "present"
	}

	existing, err := s.repo.FindByUserAndDate(req.UserID, date)
	if err != nil {
		// create new
		a := &models.Attendance{
			UserID:  req.UserID,
			Date:    date,
			Status:  status,
			CheckIn: currentTime(),
		}
		if err := s.repo.Create(a); err != nil {
			return nil, err
		}
		return a, nil
	}

	existing.Status = status
	if err := s.repo.Save(existing); err != nil {
		return nil, err
	}
	return existing, nil
}

// Start break
func (s *AttendanceService) StartBreak(attendanceID string) (*models.Attendance, error) {
	a, err := s.repo.FindByID(attendanceID)
	if err != nil {
		return nil, errors.New("attendance record not found")
	}
	if a.BreakStart != "" && a.BreakEnd == "" {
		return nil, errors.New("already on break")
	}
	a.BreakStart = currentTime()
	a.BreakEnd = ""
	if err := s.repo.Save(a); err != nil {
		return nil, err
	}
	return a, nil
}

// End break — calculates break duration in minutes
func (s *AttendanceService) EndBreak(attendanceID string) (*models.Attendance, error) {
	a, err := s.repo.FindByID(attendanceID)
	if err != nil {
		return nil, errors.New("attendance record not found")
	}
	if a.BreakStart == "" {
		return nil, errors.New("break not started")
	}

	breakEnd := currentTime()
	a.BreakEnd = breakEnd

	// calculate break duration in minutes
	layout := "15:04:05"
	start, err1 := time.Parse(layout, a.BreakStart)
	end, err2 := time.Parse(layout, breakEnd)
	if err1 == nil && err2 == nil {
		diff := int(end.Sub(start).Minutes())
		if diff > 0 {
			a.TotalBreakTime += diff
		}
	}

	if err := s.repo.Save(a); err != nil {
		return nil, err
	}
	return a, nil
}

// Get own attendance records
func (s *AttendanceService) GetOwnAttendance(userID string) (map[string]interface{}, error) {
	records, err := s.repo.FindByUserID(userID)
	if err != nil {
		return nil, err
	}

	today := todayDate()
	var todayRecord *models.Attendance
	for i := range records {
		if records[i].Date == today {
			todayRecord = &records[i]
			break
		}
	}

	return map[string]interface{}{
		"attendanceRecords":      records,
		"todayRecord":            todayRecord,
		"employeeMonthlySummary": records,
	}, nil
}

// Get all employee attendance (admin/hr)
func (s *AttendanceService) GetAllAttendance() (map[string]interface{}, error) {
	today := todayDate()

	todayRecords, err := s.repo.FindTodayAll(today)
	if err != nil {
		return nil, err
	}

	allRecords, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}

	present := []models.Attendance{}
	for _, r := range todayRecords {
		if r.Status == "present" {
			present = append(present, r)
		}
	}

	return map[string]interface{}{
		"attendance":             allRecords,
		"todayPresentRecords":    present,
		"todayRecords":           todayRecords,
		"employeeMonthlySummary": allRecords,
	}, nil
}
