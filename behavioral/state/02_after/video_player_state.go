package main

type State interface {
	Play(videoPlayer *VideoPlayer)
	Stop(videoPlayer *VideoPlayer)
}
