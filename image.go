// Property of Steven Gutzwiller (sgutzy@gmail.com)

package main

import(
	"flag"
	"image"
	_ "image/jpeg"
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
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, a := m.At(x, y).RGBA()
			glog.Errorf("(%v,%v): %v %v %v %v", y, x, r, g, b, a)
		}
	}
}