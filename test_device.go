
package main

import (
    "./device"
    //"bytes"
    //"image"
    "fmt"
    _ "image/jpeg"
    "log"
	"strings"

    "github.com/hajimehoshi/ebiten"
    "github.com/hajimehoshi/ebiten/ebitenutil"
)

const (
    screenWidth  = 640
    screenHeight = 480
)

type Resource struct {
    Img1 *ebiten.Image
}

func NewReource() (*Resource) {
    img1, _, err := ebitenutil.NewImageFromFile("./resources/img1.png", ebiten.FilterNearest)
    if err != nil {
        log.Fatal(err)
    }
    return &Resource {
        Img1: img1,
    }
}
var rc *Resource

////////////////////////////////////////
// Test1
////////////////////////////////////////
type Test1 struct {
    count int
}

func NewTest1() (*Test1) {
    return &Test1 {
        count: 0,
    }
}

func (me* Test1) Update(screen *ebiten.Image) (error) {
    if ebiten.IsDrawingSkipped() {
        return nil
    }
	deviceInput := device.Get()
	sarr := []string{}
	for _, v := range(deviceInput) {
		sarr = append(sarr, fmt.Sprintf("%d", v))
	}
    msg := strings.Join(sarr, ", ")
    ebitenutil.DebugPrint(screen, msg)
    return nil
}

type Proc interface {
    Update(screen *ebiten.Image) (error)
}

func main() {
    rc = NewReource()
    var pc Proc
    pc = NewTest1()
    if err := ebiten.Run(pc.Update, screenWidth, screenHeight, 1, "Rotate (Ebiten Demo)"); err != nil {
        log.Fatal(err)
    }
}
