package main

import (
  "log"
  "fmt"
  "strconv"
  "net/http"
  "math"
  "image"
  "image/color"
  "image/gif"
  "math/rand"
)

func main() {
  http.HandleFunc("/", lissajous)
  log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

const (
    blackIndex = 1
    whiteIndex = 0
)
var foreground = color.RGBA{0x00, 0xff, 0x00, 1}
var background = color.Black
var palette = []color.Color{background, foreground}

func lissajous(out http.ResponseWriter, r *http.Request) {
    // get cycle and delay from queryString
    err := r.ParseForm()
    if err != nil {
      fmt.Fprintf(out, "Form Error = %q", err)
      return
    }
    var cycle_count string
    getcycles := r.Form["cycles"]
    switch {
      case len(getcycles) > 0:
        cycle_count = getcycles[0]
      default:
        cycle_count = "4"
    }
    cycles, err := strconv.ParseFloat(cycle_count, 32) // returns a List
    if err != nil {
      fmt.Fprintf(out, "Form Error = %q", err)
      return
    }
    if cycles > 100 {
      cycles = 100
    }
    const (
        res = 0.001
        size = 100
        nFrames = 64
        delay = 2
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
