package main

type CommentServiceType string

var (
	CommentServiceTypeTrimming              CommentServiceType = "trimming"
	CommentServiceTypeSpamFilter            CommentServiceType = "spam_filter"
	CommentServiceTypeSpamFilterAndTrimming CommentServiceType = "spam_filter_and_trimming"
)

func main() {
	var commentServiceType = CommentServiceTypeSpamFilterAndTrimming
	var commentService CommentService

	switch commentServiceType {
	case CommentServiceTypeTrimming:
		commentService = new(TrimmingCommentService)
	case CommentServiceTypeSpamFilter:
		commentService = new(SpamFilterCommentService)
	case CommentServiceTypeSpamFilterAndTrimming:
		commentService = new(SpamFilterAndTrimmingCommentService)
	}

	commentService.AddComment("Gopher 여러분!")
	commentService.AddComment("반갑습니다...!")
	commentService.AddComment("https://ithub.tistory.com")
}
