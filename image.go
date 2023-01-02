// Property of Steven Gutzwiller (sgutzy@gmail.com)

package main

import(
	"flag"
	"image"
	"image/color"
	"image/jpeg"
	"os"

	"github.com/golang/glog"
	)

var (
	file = flag.String("file", "", "Image file to use")
)

func main() {
	flag.Parse()
	if *file == "" {
		glog.Fatalf("File required")
	}
	reader, err := os.Open(*file)
	if err != nil {
		glog.Fatal(err)
	}
	defer reader.Close()
	m, _, err := image.Decode(reader)
	if err != nil {
		glog.Fatal(err)
	}
	bounds := m.Bounds()

	yLength := (bounds.Max.Y - bounds.Min.Y)/2
	xLength := (bounds.Max.X - bounds.Min.X)/2
	newImg := image.NewRGBA(image.Rectangle{
		image.Point{0,0},
		image.Point{xLength, yLength},
	})

	for y := bounds.Min.Y; y+1 < bounds.Max.Y; y+=2 {
		for x := bounds.Min.X; x+2 < bounds.Max.X; x+=2 {
			r11, g11, b11, a11 := m.At(x, y).RGBA()
			r12, g12, b12, a12 := m.At(x, y+1).RGBA()
			r21, g21, b21, a21 := m.At(x+1, y).RGBA()
			r22, g22, b22, a22 := m.At(x+1, y+1).RGBA()
			rNew := (r11 + r12 + r21 + r22) /4
			gNew := (g11 + g12 + g21 + g22) /4
			bNew := (b11 + b12 + b21 + b22) /4
			aNew := (a11 + a12 + a21 + a22) /4
			newImg.Set(bounds.Min.X - x, bounds.Min.Y - y, color.RGBA{uint8(rNew>>8),
				uint8(gNew>>8), uint8(bNew>>8), uint8(aNew>>8)})
		}
	}

	f, err := os.Create("image.jpeg")
	if err != nil {
		glog.Fatal(err)
	}
	err = jpeg.Encode(f, m, &jpeg.Options{Quality: 100})
	if err != nil {
		glog.Fatal(err)
	}
}