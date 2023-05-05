package main

import "fmt"

type CommentServiceInterface interface {
	AddComment(comment string)
}

type CommentService struct {
}

func NewCommentService() *CommentService {
	return &CommentService{}
}

func (ncs *CommentService) AddComment(comment string) {
	fmt.Println(comment)
}
