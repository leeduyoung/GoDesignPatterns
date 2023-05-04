package main

type Font struct {
	fontFamily string
	fontSize   int
}

func NewFont(fontFamily string, fontSize int) *Font {
	return &Font{
		fontFamily: fontFamily,
		fontSize:   fontSize,
	}
}
