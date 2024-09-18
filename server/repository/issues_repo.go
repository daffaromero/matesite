package repository

import (
	"context"
	"fmt"

	api "github.com/daffaromero/matesite/server/protobuf"
	"github.com/daffaromero/matesite/server/repository/query"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type IssueRepository interface {
	GetIssue(ctx context.Context, req *api.GetIssueRequest) (*api.GetIssueResponse, error)
	ListIssues(ctx context.Context, req *api.ListIssuesRequest) (*api.ListIssuesResponse, error)
	CreateIssue(ctx context.Context, req *api.CreateIssueRequest) (*api.CreateIssueResponse, error)
	UpdateIssue(ctx context.Context, req *api.UpdateIssueRequest) (*api.UpdateIssueResponse, error)
	DeleteIssue(ctx context.Context, id *api.DeleteIssueRequest) (*api.DeleteIssueResponse, error)
}

type issueRepository struct {
	db         Store
	issueQuery query.IssueQuery
}

func NewIssueRepository(db Store, issueQuery query.IssueQuery) IssueRepository {
	return &issueRepository{
		db:         db,
		issueQuery: issueQuery,
	}
}

func (r *issueRepository) GetIssue(ctx context.Context, req *api.GetIssueRequest) (*api.GetIssueResponse, error) {
	var issue *api.GetIssueResponse

	err := r.db.WithoutTx(ctx, func(pool *pgxpool.Pool) error {
		var err error
		issue, err = r.issueQuery.GetIssue(ctx, req)
		return err
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get issue: %w", err)
	}
	return issue, nil
}

func (r *issueRepository) ListIssues(ctx context.Context, req *api.ListIssuesRequest) (*api.ListIssuesResponse, error) {
	var issues *api.ListIssuesResponse

	err := r.db.WithoutTx(ctx, func(pool *pgxpool.Pool) error {
		var err error
		issues, err = r.issueQuery.ListIssues(ctx, req)
		return err
	})
	if err != nil {
		return nil, fmt.Errorf("failed to list issues: %w", err)
	}
	return issues, nil
}

func (r *issueRepository) CreateIssue(ctx context.Context, req *api.CreateIssueRequest) (*api.CreateIssueResponse, error) {
	var issue *api.CreateIssueResponse

	err := r.db.WithTx(ctx, func(tx pgx.Tx) error {
		var err error
		issue, err = r.issueQuery.CreateIssue(ctx, tx, req)
		return err
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create issue: %w", err)
	}
	return issue, nil
}

func (r *issueRepository) UpdateIssue(ctx context.Context, req *api.UpdateIssueRequest) (*api.UpdateIssueResponse, error) {
	var issue *api.UpdateIssueResponse

	err := r.db.WithTx(ctx, func(tx pgx.Tx) error {
		var err error
		issue, err = r.issueQuery.UpdateIssue(ctx, tx, req)
		return err
	})
	if err != nil {
		return nil, fmt.Errorf("failed to update issue: %w", err)
	}
	return issue, nil
}

func (r *issueRepository) DeleteIssue(ctx context.Context, id *api.DeleteIssueRequest) (*api.DeleteIssueResponse, error) {
	var res *api.DeleteIssueResponse

	err := r.db.WithTx(ctx, func(tx pgx.Tx) error {
		var err error
		res, err = r.issueQuery.DeleteIssue(ctx, tx, id)
		return err
	})
	if err != nil {
		return nil, fmt.Errorf("failed to delete issue: %w", err)
	}
	return res, nil
}
