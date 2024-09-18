package controller

import (
	"github.com/daffaromero/matesite/server/config"
	api "github.com/daffaromero/matesite/server/protobuf"
	"github.com/daffaromero/matesite/server/service"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

type IssueController interface {
	Route(*fiber.App)
	GetIssue(ctx fiber.Ctx) error
	ListIssues(ctx fiber.Ctx) error
	CreateIssue(ctx fiber.Ctx) error
	UpdateIssue(ctx fiber.Ctx) error
	DeleteIssue(ctx fiber.Ctx) error
}

type issueController struct {
	validate *validator.Validate
	service  service.IssueService
}

func NewIssueController(validate *validator.Validate, service service.IssueService) IssueController {
	return &issueController{
		validate: validate,
		service:  service,
	}
}

func (c *issueController) Route(app *fiber.App) {
	api := app.Group(config.EndpointPrefix)
	api.Get("/:id", c.GetIssue)
	api.Get("/", c.ListIssues)
	api.Post("/new", c.CreateIssue)
	api.Put("/:id", c.UpdateIssue)
	api.Delete("/:id", c.DeleteIssue)
}

func (c *issueController) GetIssue(ctx fiber.Ctx) error {
	var req api.GetIssueRequest
	req.Id = ctx.Params("id")
	if req.Id == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "issue id not provided"})
	}

	res, err := c.service.GetIssue(ctx.Context(), &req)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(res)
}

func (c *issueController) ListIssues(ctx fiber.Ctx) error {
	var req *api.ListIssuesRequest

	res, err := c.service.ListIssues(ctx.Context(), req)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(res)
}

func (c *issueController) CreateIssue(ctx fiber.Ctx) error {
	var req api.CreateIssueRequest
	err := ctx.Bind().Body(&req)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := c.validate.Struct(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if req.Issue == nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "issue is required"})
	}

	title := req.Issue.Title
	description := req.Issue.Description

	res, err := c.service.CreateIssue(ctx.Context(), &req, title, description)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusCreated).JSON(res)
}

func (c *issueController) UpdateIssue(ctx fiber.Ctx) error {
	id := ctx.Params("id")
	if id == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "issue_id not provided"})
	}

	var req api.UpdateIssueRequest
	if err := ctx.Bind().Body(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if req.Issue == nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "issue cannot be empty"})
	}
	req.Issue.Id = id

	if err := c.validate.Struct(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if req.Issue.Title == "" && req.Issue.Description == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "At least one field (title or description) must be provided for update"})
	}

	res, err := c.service.UpdateIssue(ctx.Context(), &req, req.Issue.Title, req.Issue.Description)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(res)
}

func (c *issueController) DeleteIssue(ctx fiber.Ctx) error {
	var req api.DeleteIssueRequest
	req.Id = ctx.Params("id")
	if req.Id == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "issue id not provided"})
	}

	res, err := c.service.DeleteIssue(ctx.Context(), &req)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(res)
}
