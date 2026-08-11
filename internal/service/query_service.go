package service

import (
	"github.com/smustafa-tech/hrms-backend/internal/models"
	"github.com/smustafa-tech/hrms-backend/internal/repository"
)

type QueryService struct {
	repo *repository.QueryRepository
}

func NewQueryService(repo *repository.QueryRepository) *QueryService {
	return &QueryService{repo: repo}
}

func (s *QueryService) GetUnreadCount(userID string) (int64, error) {
	return s.repo.CountOpenByUser(userID)
}

func toQueryResponse(q models.Query, replies []models.QueryReply) map[string]interface{} {
	return map[string]interface{}{
		"id":         q.ID,
		"userId":     q.UserID,
		"subject":    q.Subject,
		"message":    q.Message,
		"priority":   q.Priority,
		"status":     q.Status,
		"attachment": q.Attachment,
		"createdAt":  q.CreatedAt,
		"updatedAt":  q.UpdatedAt,
		"replies":    replies,
	}
}

func (s *QueryService) GetMyQueries(userID string) ([]map[string]interface{}, error) {
	queries, err := s.repo.FindByUserID(userID)
	if err != nil {
		return nil, err
	}

	var ids []string
	for _, q := range queries {
		ids = append(ids, q.ID.String())
	}

	replies, err := s.repo.FindRepliesForQueryIDs(ids)
	if err != nil {
		return nil, err
	}

	replyMap := make(map[string][]models.QueryReply)
	for _, reply := range replies {
		replyMap[reply.QueryID] = append(replyMap[reply.QueryID], reply)
	}

	result := make([]map[string]interface{}, 0, len(queries))
	for _, q := range queries {
		result = append(result, toQueryResponse(q, replyMap[q.ID.String()]))
	}

	return result, nil
}

func (s *QueryService) GetAllQueries() ([]map[string]interface{}, error) {
	queries, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}

	var ids []string
	for _, q := range queries {
		ids = append(ids, q.ID.String())
	}

	replies, err := s.repo.FindRepliesForQueryIDs(ids)
	if err != nil {
		return nil, err
	}

	replyMap := make(map[string][]models.QueryReply)
	for _, reply := range replies {
		replyMap[reply.QueryID] = append(replyMap[reply.QueryID], reply)
	}

	result := make([]map[string]interface{}, 0, len(queries))
	for _, q := range queries {
		result = append(result, toQueryResponse(q, replyMap[q.ID.String()]))
	}

	return result, nil
}

func (s *QueryService) SubmitQuery(q *models.Query) (*models.Query, error) {
	if err := s.repo.Create(q); err != nil {
		return nil, err
	}
	return q, nil
}

func (s *QueryService) ReplyToQuery(reply *models.QueryReply) (*models.QueryReply, error) {
	if err := s.repo.CreateReply(reply); err != nil {
		return nil, err
	}
	return reply, nil
}

func (s *QueryService) CloseQuery(id string) error {
	return s.repo.CloseQuery(id)
}

func (s *QueryService) DeleteQuery(id string) error {
	return s.repo.Delete(id)
}
