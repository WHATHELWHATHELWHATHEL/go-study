package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

func main() {
	graphic(os.Stdout)
}

// here we receive the stdout as an output writer
func graphic(out io.Writer) {
	palette := []color.Color{color.Black, color.RGBA{0xFF, 0x52, 0x00, 0x50}, color.White}
	const (
		colorIndex = 1
		whiteIndex = 2
	)
	// to paint the sin/cos graph line, following mathematics params are required
	const (
		cycles  = 5
		res     = 0.001
		size    = 1000
		nframes = 64
		delay   = 8
	)
	freq := rand.Float64() // wave frequence, look at mathematics
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	// The following code is used only to handle the graphic animation
	for index := 0; index < nframes; index++ {
		// generate the canvas
		rectangle := image.Rect(0, 0, (2*size)+1, (2*size)+1)
		img := image.NewPaletted(rectangle, palette)
		for degree := 0.0; degree < cycles*2*math.Pi; degree += res {
			x := math.Sin(degree)
			y := math.Sin(degree*freq + phase)
			// The go strictly forbidden the use of triple expression
			// has to use if to handle this
			var choosenIndex uint8
			if rand.Float64() > 0.5 {
				choosenIndex = colorIndex
			} else {
				choosenIndex = whiteIndex
			}
			img.SetColorIndex(
				size+int(x*size+0.5),
				size+int(y*size+0.5),
				choosenIndex,
			)
		}
		phase += 0.1
		// Delay and Image is a array that decide the gap between each nframes
		// and the image each frame display
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	// the encodeAll accept only the object referrence
	gif.EncodeAll(out, &anim)
}
