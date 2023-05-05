package main

import (
	"fmt"
	"strings"
)

type SpamFilterAndTrimmingCommentService struct {
}

func (sfatcs *SpamFilterAndTrimmingCommentService) AddComment(comment string) {
	if sfatcs.isSpam(comment) {
		return
	}

	fmt.Println(sfatcs.trim(comment))
}

func (sfatcs *SpamFilterAndTrimmingCommentService) trim(comment string) string {
	return strings.ReplaceAll(comment, "...", "")
}

func (sfatcs *SpamFilterAndTrimmingCommentService) isSpam(comment string) bool {
	return strings.Contains(comment, "http")
}
