package gif

import (
	"image"
	"image/color"
	"image/gif"
	"io"
)

var palette = []color.Color{color.White, color.Black}

type gridGetter interface {
	Step() [][]bool
}

const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette”
)

func Lissajous(gg gridGetter, out io.Writer) {
	const (
		size    = 100 // image canvas covers [-size..+size]
		nframes = 64  // number of animation frames
		delay   = 8   // delay between frames in 10ms units
	)
	anim := gif.GIF{LoopCount: nframes}
	for i := 0; i < nframes; i++ {
		grid := gg.Step()
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)

		for x, _ := range grid {
			for y, filledIn := range grid[x] {
				if filledIn {
					img.SetColorIndex(x, y, blackIndex)
				}
			}
		}

		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) 
}
