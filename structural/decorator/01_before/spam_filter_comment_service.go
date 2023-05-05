package main

import (
	"fmt"
	"strings"
)

type SpamFilterCommentService struct {
}

func (sfcs *SpamFilterCommentService) AddComment(comment string) {
	if sfcs.isSpam(comment) {
		return
	}

	fmt.Println(comment)
}

func (sfcs *SpamFilterCommentService) isSpam(comment string) bool {
	return strings.Contains(comment, "http")
}
