/*
creates a lissajous aniated gif
to run:
  go run 04_lissajous.go > file.gif
*/

package main

import (
  "os"
  "io"
  "math"
  "image"
  "image/color"
  "image/gif"
  "math/rand"
)



const (
    blackIndex = 1
    whiteIndex = 0
)
var foreground = color.RGBA{0x00, 0xff, 0x00, 1}
var background = color.Black
var palette = []color.Color{background, foreground}

func main() {
    lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
    const (
        cycles = 5
        res = 0.001
        size = 100
        nFrames = 64
        delay = 8
    )
    freq := rand.Float64() * 3.0
    anim := gif.GIF{LoopCount: nFrames}
    phase := 0.
    for i := 0; i< nFrames; i++ {
        rect := image.Rect(0, 0, 2*size+1, 2*size+1)
        img := image.NewPaletted(rect, palette)
        for t := 0.; t < cycles*2*math.Pi; t += res {
            x := math.Sin(t)
            y := math.Sin(t*freq + phase)
            img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
        }
        phase += 0.1
        anim.Delay = append(anim.Delay, delay)
        anim.Image = append(anim.Image, img)
    }
    gif.EncodeAll(out, &anim)
}
