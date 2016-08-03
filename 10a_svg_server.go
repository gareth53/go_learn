// web server that generates an SVG image of a 3D surface

package main

import (
  "fmt"
  "math"
  "log"
  "net/http"
)

const (
  width, height = 600, 320 // canvas sze in px
  cells = 100 // grid cells
  xyrange = 30.0
  xyscale = width / 2 / xyrange // px per xy unit
  zscale = height * 0.4  // px per z unit
  angle = math.Pi / 6   // perspctive angle (30˚)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
  http.HandleFunc("/", svg)
  log.Fatal(http.ListenAndServe("localhost:8000", nil))
}


func svg(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "<html><body><svg xmlns='http://www.w3.org/2000/svg' " +
             "style='stroke: grey; fill: #fff, stroke-width:0.7' " +
             "width='%d' height='%d'>", width, height)
  for i:=0; i<cells; i++ {
    for j:=0; j<cells; j++ {
      ax, ay := corner(i+1, j)
      bx, by := corner(i, j)
      cx, cy := corner(i, j+1)
      dx, dy := corner(i+1, j+1)
      fmt.Fprintf(w, "<polygon points='%g,%g,%g,%g,%g,%g,%g,%g'/>\n",
          ax, ay, bx, by, cx, cy, dx, dy)
    }
  }
  fmt.Fprintf(w, "</svg></body></html>")
}

func corner(i, j int) (float64, float64) {
  x := xyrange * (float64(i)/cells - 0.5)
  y := xyrange * (float64(j)/cells - 0.5)
  z := f(x, y)

  // project x, y & z isometrically onto 2d canvas
  sx := width/2 + (x-y)*cos30*xyscale
  sy := height/2 + (x+y)*sin30*xyscale - z*zscale
  return sx, sy
}

func f(x, y float64) float64 {
  r := math.Hypot(x, y) // dist from (0, 0)
  return math.Sin(r)/r
}
