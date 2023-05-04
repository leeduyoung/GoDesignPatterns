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

// FontFamily 변경을 방지하기위해서 Setter는 제공하지않고 Getter만 제공한다.
func (f Font) FontFamily() string {
	return f.fontFamily
}

// FontSize 변경을 방지하기위해서 Setter는 제공하지않고 Getter만 제공한다.
func (f Font) FontSize() int {
	return f.fontSize
}
