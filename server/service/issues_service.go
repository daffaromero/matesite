package service

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/daffaromero/matesite/server/helper/logger"
	api "github.com/daffaromero/matesite/server/protobuf"
	"github.com/daffaromero/matesite/server/repository"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type IssueService interface {
	GetIssue(ctx context.Context, req *api.GetIssueRequest) (*api.GetIssueResponse, error)
	ListIssues(ctx context.Context, req *api.ListIssuesRequest) (*api.ListIssuesResponse, error)
	CreateIssue(ctx context.Context, req *api.CreateIssueRequest, title string, description string) (*api.CreateIssueResponse, error)
	UpdateIssue(ctx context.Context, req *api.UpdateIssueRequest, title string, description string) (*api.UpdateIssueResponse, error)
	DeleteIssue(ctx context.Context, req *api.DeleteIssueRequest) (*api.DeleteIssueResponse, error)
}

type issueService struct {
	repo   repository.IssueRepository
	logger *logger.Log
}

func NewIssueService(repo repository.IssueRepository, logger *logger.Log) IssueService {
	return &issueService{
		repo:   repo,
		logger: logger,
	}
}

func (s *issueService) GetIssue(ctx context.Context, req *api.GetIssueRequest) (*api.GetIssueResponse, error) {
	issue, err := s.repo.GetIssue(ctx, req)
	if err != nil {
		s.logger.Error(fmt.Sprintf("Failed to get issue: %v", err))
		return nil, err
	}
	return issue, nil
}

func (s *issueService) ListIssues(ctx context.Context, req *api.ListIssuesRequest) (*api.ListIssuesResponse, error) {
	issues, err := s.repo.ListIssues(ctx, req)
	if err != nil {
		s.logger.Error(fmt.Sprintf("Failed to list issues: %v", err))
		return nil, err
	}
	return issues, nil
}

func (s *issueService) CreateIssue(ctx context.Context, req *api.CreateIssueRequest, title string, description string) (*api.CreateIssueResponse, error) {
	now := &timestamppb.Timestamp{
		Seconds: time.Now().Unix(),
		Nanos:   int32(time.Now().Nanosecond()),
	}
	req.Issue.Id = uuid.New().String()
	req.Issue.Title = title
	req.Issue.Description = description
	req.Issue.CreatedAt = now
	req.Issue.UpdatedAt = now

	res, err := s.repo.CreateIssue(ctx, req)
	if err != nil {
		s.logger.Error(fmt.Sprintf("Failed to create issue: %v", err))
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			return nil, fiber.NewError(fiber.StatusBadRequest, "Issue already exists.")
		}
		return nil, err
	}

	return res, nil
}

func (s *issueService) UpdateIssue(ctx context.Context, req *api.UpdateIssueRequest, title string, description string) (*api.UpdateIssueResponse, error) {
	now := &timestamppb.Timestamp{
		Seconds: time.Now().Unix(),
		Nanos:   int32(time.Now().Nanosecond()),
	}
	req.Issue.Title = title
	req.Issue.Description = description
	req.Issue.UpdatedAt = now

	res, err := s.repo.UpdateIssue(ctx, req)
	if err != nil {
		s.logger.Error(fmt.Sprintf("Failed to update issue: %v", err))
		return nil, err
	}

	return res, nil
}

func (s *issueService) DeleteIssue(ctx context.Context, req *api.DeleteIssueRequest) (*api.DeleteIssueResponse, error) {
	res, err := s.repo.DeleteIssue(ctx, req)
	if err != nil {
		s.logger.Error(fmt.Sprintf("Failed to delete issue: %v", err))
		return nil, err
	}
	return res, nil
}
