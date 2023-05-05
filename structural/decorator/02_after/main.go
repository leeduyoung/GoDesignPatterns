package main

func main() {
	var (
		commentService CommentServiceInterface = NewCommentService()

		spamFilterON = true
		trimmingON   = true
	)

	if spamFilterON {
		commentService = NewSpamFilterCommentService(commentService)
	}

	if trimmingON {
		commentService = NewTrimmingCommentService(commentService)
	}

	commentService.AddComment("Gopher 여러분!")
	commentService.AddComment("반갑습니다...!")
	commentService.AddComment("https://ithub.tistory.com")
}
