package main

import (
	"strings"
)

type SpamFilterCommentService struct {
	CommentService CommentServiceInterface
}

func NewSpamFilterCommentService(service CommentServiceInterface) *SpamFilterCommentService {
	return &SpamFilterCommentService{
		CommentService: service,
	}
}

func (sfcs *SpamFilterCommentService) AddComment(comment string) {
	if sfcs.isSpam(comment) {
		return
	}

	sfcs.CommentService.AddComment(comment)
}

func (sfcs *SpamFilterCommentService) isSpam(comment string) bool {
	return strings.Contains(comment, "http")
}
