package main

import (
	"strconv"
	"strings"
)

type FontFactory struct {
	cache map[string]*Font
}

func NewFontFactory() *FontFactory {
	return &FontFactory{
		cache: make(map[string]*Font),
	}
}

func (ff *FontFactory) getInstance(fontFamilyAndSize string) *Font {
	if font, exists := ff.cache[fontFamilyAndSize]; exists {
		return font
	}

	arr := strings.Split(fontFamilyAndSize, ":")
	fontFamily := arr[0]
	fontSize, _ := strconv.Atoi(arr[1])

	newFont := NewFont(fontFamily, fontSize)
	ff.cache[fontFamilyAndSize] = newFont
	return newFont
}
