package main

import "fmt"

type VideoPlayerStatePlay struct {
}

func (vpss *VideoPlayerStatePlay) Play(videoPlayer *VideoPlayer) {
	fmt.Println("이미 비디오가 재생중입니다.")
}

func (vpss *VideoPlayerStatePlay) Stop(videoPlayer *VideoPlayer) {
	fmt.Println("비디오를 재생합니다.")
	videoPlayer.ChangeState(new(VideoPlayerStatePlay))
}
