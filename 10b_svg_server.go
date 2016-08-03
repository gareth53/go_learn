package main
// web server that generates an SVG image of a 3D surface
// can accpet query str values: ?width=400&height=400&cells=5

import (
  "fmt"
  "math"
  "log"
  "net/http"
  "strconv"
)
var angle = math.Pi / 6   // perspective angle (30˚)
var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
  http.HandleFunc("/", process_svg)
  log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func process_svg(w http.ResponseWriter, r *http.Request) {
  err := r.ParseForm()
  if err != nil {
    fmt.Fprintf(w, "Form Error = %q", err)
    return
  }
  width := int(get_float_value(r, "width", 600.0, 100.0, 1000.0))
  height := int(get_float_value(r, "height", 320.0, 80.0, 800.0))
  cells := int(get_float_value(r, "cells", 100.0, 5.0, 200.0))
  svg(w, width, height, cells)
}

func getFloatValue(r *http.Request, key string, defaultval float64, min float64, max float64) float64 {
  keyvals := r.Form[key]
  strval := strconv.FormatFloat(defaultval, 'e', 10, 64)
  if len(keyvals) > 0 {
      strval = keyvals[0]
  }
  floatval, err := strconv.ParseFloat(strval, 32)
  if err != nil {
    return defaultval
  }
  return math.Max(math.Min(floatval, max), min)
}

func svg(w http.ResponseWriter, width, height, cells int) {
  xyrange := 30.0
  xyscale := float64(width) / 2 / xyrange // px per xy unit
  zscale := float64(height) * 0.4  // px per z unit

  fmt.Fprintf(w, "<html><body><p>width = %d</p><svg xmlns='http://www.w3.org/2000/svg' " +
             "style='stroke: grey; fill: #fff, stroke-width:0.7' " +
             "width='%d' height='%d'>", width, width, height)
  for i:=0; i<cells; i++ {
    for j:=0; j<cells; j++ {
      ax, ay := corner(i+1, j, cells, width, height, xyscale, xyrange, zscale)
      bx, by := corner(i, j, cells, width, height, xyscale, xyrange, zscale)
      cx, cy := corner(i, j+1, cells, width, height, xyscale, xyrange, zscale)
      dx, dy := corner(i+1, j+1, cells, width, height, xyscale, xyrange, zscale)
      fmt.Fprintf(w, "<polygon points='%g,%g,%g,%g,%g,%g,%g,%g'/>\n",
          ax, ay, bx, by, cx, cy, dx, dy)
    }
  }
  fmt.Fprintf(w, "</svg></body></html>")
}

func corner(i, j cells, width, height int, xyscale, xyrange, zscale float64) (float64, float64) {
  cell := float64(cells)
  x := xyrange * (float64(i)/cell - 0.5)
  y := xyrange * (float64(j)/cell - 0.5)
  z := f(x, y)

  // project x, y & z isometrically onto 2d canvas
  sx := float64(width)/2.0 + (x-y)*cos30*xyscale
  sy := float64(height)/2.0 + (x+y)*sin30*xyscale - z*zscale
  return sx, sy
}

func f(x, y float64) float64 {
  r := math.Hypot(x, y) // dist from (0, 0)
  return math.Sin(r)/r
}
