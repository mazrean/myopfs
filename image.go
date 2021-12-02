package main

type Image struct {
	data []byte
}

func NewImage(data []byte) *Image {
	return &Image{data}
}
