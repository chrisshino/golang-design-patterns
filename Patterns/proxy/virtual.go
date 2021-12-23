package main

import "fmt"


type Image interface {
	Draw()
}

type Bitmap struct {
	filename string
}

func (b *Bitmap) Draw() {
	fmt.Println("Drawing image", b.filename)
}

func NewBitmap(filename string) *Bitmap {
	return &Bitmap{filename: filename}
}

func DrawImage(image Image) {

}

type LazyBitmap struct {
	filename string
	bitmap *Bitmap
}

func NewLazyBitmap(filename string) *LazyBitmap {
	return &LazyBitmap{filename: filename}
}

func (l *LazyBitmap) Draw() {
	if l.bitmap == nil {
		l.bitmap = NewBitmap(l.filename)
	}
	l.bitmap.Draw()
}

func main() {
 bmp := NewLazyBitmap("demo.png")
 DrawImage(bmp)

}