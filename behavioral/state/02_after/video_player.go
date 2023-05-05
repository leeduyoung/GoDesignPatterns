package main

type VideoPlayer struct {
	state State
}

func NewVideoPlayer() *VideoPlayer {
	return &VideoPlayer{
		state: new(VideoPlayerStateStop),
	}
}

func (vp *VideoPlayer) Play() {
	vp.state.Play(vp)
}

func (vp *VideoPlayer) Stop() {
	vp.state.Stop(vp)
}

func (vp *VideoPlayer) ChangeState(state State) {
	vp.state = state
}
