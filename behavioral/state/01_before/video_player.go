package main

import "fmt"

type VideoPlayerState string

var (
	VideoPlayerStatePlay VideoPlayerState = "play"
	VideoPlayerStateStop VideoPlayerState = "stop"
)

type VideoPlayer struct {
	state VideoPlayerState
}

func NewVideoPlayer() *VideoPlayer {
	return &VideoPlayer{
		state: VideoPlayerStateStop,
	}
}

func (vp *VideoPlayer) Play() {
	switch vp.state {
	case VideoPlayerStatePlay:
		fmt.Println("이미 비디오가 재생중입니다.")
	case VideoPlayerStateStop:
		fmt.Println("비디오를 재생합니다.")
		vp.state = VideoPlayerStatePlay
	}
}

func (vp *VideoPlayer) Stop() {
	switch vp.state {
	case VideoPlayerStatePlay:
		fmt.Println("비디오를 정지합니다.")
		vp.state = VideoPlayerStateStop
	case VideoPlayerStateStop:
		fmt.Println("이미 비디오는 정지중입니다.")
	}
}
