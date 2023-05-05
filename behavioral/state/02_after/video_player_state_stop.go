package main

import "fmt"

type VideoPlayerStateStop struct {
}

func (vpss *VideoPlayerStateStop) Play(videoPlayer *VideoPlayer) {
	fmt.Println("비디오를 정지합니다.")
	videoPlayer.ChangeState(new(VideoPlayerStateStop))
}

func (vpss *VideoPlayerStateStop) Stop(videoPlayer *VideoPlayer) {
	fmt.Println("이미 비디오는 정지중입니다.")
}
