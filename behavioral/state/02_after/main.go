package main

func main() {
	videoPlayer := NewVideoPlayer()
	videoPlayer.Stop()
	videoPlayer.Play()
	videoPlayer.Play()
	videoPlayer.Stop()
}
