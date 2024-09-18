package query

import (
	"context"
	"errors"
	"fmt"
	"time"

	api "github.com/daffaromero/matesite/server/protobuf"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type IssueQuery interface {
	GetIssue(ctx context.Context, req *api.GetIssueRequest) (*api.GetIssueResponse, error)
	ListIssues(ctx context.Context, req *api.ListIssuesRequest) (*api.ListIssuesResponse, error)
	CreateIssue(ctx context.Context, tx pgx.Tx, req *api.CreateIssueRequest) (*api.CreateIssueResponse, error)
	UpdateIssue(ctx context.Context, tx pgx.Tx, req *api.UpdateIssueRequest) (*api.UpdateIssueResponse, error)
	DeleteIssue(ctx context.Context, tx pgx.Tx, id *api.DeleteIssueRequest) (*api.DeleteIssueResponse, error)
}

type issueQuery struct {
	db *pgxpool.Pool
}

func NewIssueQuery(db *pgxpool.Pool) IssueQuery {
	return &issueQuery{
		db: db,
	}
}

func (q *issueQuery) GetIssue(ctx context.Context, req *api.GetIssueRequest) (*api.GetIssueResponse, error) {
	if req == nil || req.Id == "" {
		return nil, errors.New("issue ID cannot be empty")
	}
	query := `SELECT id, title, description FROM issues WHERE id = $1 AND deleted_at IS NULL`

	row := q.db.QueryRow(ctx, query, req.Id)

	var issue api.Issue
	err := row.Scan(&issue.Id, &issue.Title, &issue.Description)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("issue with ID %s not found", req.Id)
		}
		return nil, err
	}

	return &api.GetIssueResponse{
		Issue: &issue,
	}, nil
}

func (q *issueQuery) ListIssues(ctx context.Context, req *api.ListIssuesRequest) (*api.ListIssuesResponse, error) {
	query := `SELECT id, title, description FROM issues WHERE deleted_at IS NULL`

	rows, err := q.db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to query Issues: %w", err)
	}
	defer rows.Close()

	var Issues []*api.Issue
	for rows.Next() {
		var issue api.Issue
		err := rows.Scan(&issue.Id, &issue.Title, &issue.Description)
		if err != nil {
			return nil, err
		}
		Issues = append(Issues, &issue)
	}

	return &api.ListIssuesResponse{
		Issues: Issues,
	}, nil
}

func (q *issueQuery) CreateIssue(ctx context.Context, tx pgx.Tx, req *api.CreateIssueRequest) (*api.CreateIssueResponse, error) {
	if req == nil || req.Issue == nil {
		return nil, errors.New("please provide a valid issue")
	}
	query := `INSERT INTO issues (id, title, description, created_at, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING id, title, description`

	createdAt := req.Issue.CreatedAt.AsTime()
	updatedAt := req.Issue.UpdatedAt.AsTime()

	var createdIssue api.Issue

	err := tx.QueryRow(ctx, query, req.Issue.Id, req.Issue.Title, req.Issue.Description, createdAt, updatedAt).Scan(&createdIssue.Id, &createdIssue.Title, &createdIssue.Description)
	if err != nil {
		return nil, err
	}

	return &api.CreateIssueResponse{
		Issue: &createdIssue,
	}, nil
}

func (q *issueQuery) UpdateIssue(ctx context.Context, tx pgx.Tx, req *api.UpdateIssueRequest) (*api.UpdateIssueResponse, error) {
	if req == nil || req.Issue == nil {
		return nil, errors.New("issue cannot be empty")
	}
	query := `UPDATE issues 
		SET 
			title = COALESCE($2, title), 
			description = COALESCE($3, description), 
			updated_at = $4
		WHERE id = $1 AND deleted_at IS NULL RETURNING id, title, description`

	updatedAt := time.Now()
	if req.Issue.UpdatedAt != nil {
		updatedAt = req.Issue.UpdatedAt.AsTime()
	}

	var updatedIssue api.Issue

	err := tx.QueryRow(ctx, query, req.Issue.Id, req.Issue.Title, req.Issue.Description, updatedAt).Scan(&updatedIssue.Id, &updatedIssue.Title, &updatedIssue.Description)
	if err != nil {
		return nil, err
	}

	return &api.UpdateIssueResponse{
		Issue: &updatedIssue,
	}, nil
}

func (q *issueQuery) DeleteIssue(ctx context.Context, tx pgx.Tx, req *api.DeleteIssueRequest) (*api.DeleteIssueResponse, error) {
	if req.Id == "" {
		return nil, errors.New("issue ID cannot be empty")
	}

	query := `UPDATE issues SET deleted_at = $2 WHERE id = $1 AND deleted_at IS NULL`

	_, err := tx.Exec(ctx, query, req.Id, time.Now())
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("issue with ID %s not found", req.Id)
		}
		return nil, err
	}

	return &api.DeleteIssueResponse{
		Success: true,
	}, nil
}
