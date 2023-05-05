package main

import (
	"fmt"
	"strings"
)

type TrimmingCommentService struct {
}

func (tcs *TrimmingCommentService) AddComment(comment string) {
	fmt.Println(tcs.trim(comment))
}

func (tcs *TrimmingCommentService) trim(comment string) string {
	return strings.ReplaceAll(comment, "...", "")
}
