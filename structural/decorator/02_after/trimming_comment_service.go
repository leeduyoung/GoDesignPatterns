package main

import (
	"strings"
)

type TrimmingCommentService struct {
	commentService CommentServiceInterface
}

func NewTrimmingCommentService(service CommentServiceInterface) *TrimmingCommentService {
	return &TrimmingCommentService{
		commentService: service,
	}
}

func (tcs *TrimmingCommentService) AddComment(comment string) {
	tcs.commentService.AddComment(tcs.trim(comment))
}

func (tcs *TrimmingCommentService) trim(comment string) string {
	return strings.ReplaceAll(comment, "...", "")
}
