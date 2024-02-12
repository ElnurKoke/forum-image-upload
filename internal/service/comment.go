package service

import (
	"fmt"
	"forum/internal/models"
	"forum/internal/storage"
	"strings"
)

type CommentServiceIR interface {
	GetCommentsByIdPost(id int) ([]models.Comment, error)
	CreateComment(id int, username, commentText string) (int, error)
	DeleteComment(id int) error
	GetCommentsByIdComment(id int) (models.Comment, error)
	UpdateComment(comment models.Comment) error
}

type CommentService struct {
	storage storage.CommentIR
}

func newCommentServ(storage storage.CommentIR) CommentServiceIR {
	return &CommentService{
		storage: storage,
	}
}

func (c *CommentService) DeleteComment(id int) error {
	return c.storage.DeleteComment(id)
}

func (c *CommentService) UpdateComment(comment models.Comment) error {
	commentText := strings.TrimSpace(comment.Text)
	if len(commentText) == 0 {
		return fmt.Errorf("Empty comment")
	}
	return c.storage.UpdateComment(comment)
}

func (c *CommentService) GetCommentsByIdComment(id int) (models.Comment, error) {
	return c.storage.GetCommentsByIdComment(id)
}

func (c *CommentService) GetCommentsByIdPost(id int) ([]models.Comment, error) {
	return c.storage.GetCommentsByIdPost(id)
}

func (c *CommentService) CreateComment(id int, username, commentText string) (int, error) {
	commentText = strings.TrimSpace(commentText)
	if len(commentText) == 0 {
		return 0, fmt.Errorf("Empty comment")
	}
	return c.storage.CreateComment(id, username, commentText)
}
