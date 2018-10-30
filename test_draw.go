
package main

import (
    "./draw"
    //"bytes"
    //"image"
    "fmt"
    _ "image/jpeg"
    "log"
    "math"
    "math/rand"

    "github.com/hajimehoshi/ebiten"
    "github.com/hajimehoshi/ebiten/ebitenutil"
    "github.com/pa-m/numgo"
    //"github.com/hajimehoshi/ebiten/examples/resources/images"
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
    me.count++
    w, h := rc.Img1.Size()
    op := &ebiten.DrawImageOptions{}
    x, y := int(screenWidth/2), int(screenHeight/2)
    draw.Option(w, h, x, y, me.count, op)
    draw.Draw(screen, rc.Img1, op)
    return nil
}

////////////////////////////////////////
// Test2
////////////////////////////////////////
var np = numgo.NumGo{}
type Test2 struct {
    length int
    x []float64
    y []float64
    index int
}

func MinMaxTransform(x []float64, n float64) ([]float64) {
    xmin := np.Min(x)
    xmax := np.Max(x)
    d := xmax - xmin
    res := []float64{}
    for i := 0; i < len(x); i++ {
        xi := x[i]
        res = append(res, (xi - xmin) / d)
    }
    for i := 0; i < len(res); i++ {
        xi := res[i]
        res[i] = xi * n
    }
    return res
}

func NewTest2() (*Test2) {
    x := np.Linspace(0, 2 * math.Pi, 60, true)
    x = append(x, x...)
    y := []float64{}
    for i := 0; i < len(x); i++ {
        y = append(y, math.Sin(x[i]))
        x[i] = float64(i)
    }
    x = MinMaxTransform(x, screenWidth)
    y = MinMaxTransform(y, screenHeight)
    length := len(x)

    return &Test2 {
        length: length,
        x: x,
        y: y,
    }
}

func (me* Test2) Update(screen *ebiten.Image) (error) {
    if ebiten.IsDrawingSkipped() {
        return nil
    }
    idx := me.index
    x := me.x[idx]
    y := me.y[idx]
    me.index++
    if me.index >= me.length {
        me.index = 0
    }

    w, h := rc.Img1.Size()
    op := &ebiten.DrawImageOptions{}
    draw.Option(w, h, int(x), int(y), 0, op)
    draw.Draw(screen, rc.Img1, op)
    return nil
}

////////////////////////////////////////
// Test3
////////////////////////////////////////

type Instance struct {
    Shift int
    Counter int
}

type Test3 struct {
    length int
    Instances []Instance
    x []float64
    y []float64
}

func NewTest3() (*Test3) {
    x := np.Linspace(0, 2 * math.Pi, 120, true)
    x = append(x, x...)
    y := []float64{}
    for i := 0; i < len(x); i++ {
        y = append(y, math.Sin(x[i]))
        x[i] = float64(i)
    }
    x = MinMaxTransform(x, screenWidth)
    y = MinMaxTransform(y, screenHeight)
    length := len(x)

    return &Test3 {
        length: length,
        x: x,
        y: y,
    }
}

func (me* Test3) Update(screen *ebiten.Image) (error) {
    if ebiten.IsDrawingSkipped() {
        return nil
    }

    // input
    if rand.Intn(5) == 0 {
        v := Instance{
            Shift: rand.Intn(30),
        }
        me.Instances = append(me.Instances, v)
    }

    // proc
    valid := []bool{}
    for i := 0; i < len(me.Instances); i++ {
        ii := &me.Instances[i]
        ii.Counter++
        if (ii.Counter) >= me.length {
            valid = append(valid, false)
        } else {
            valid = append(valid, true)
        }
    }
    newInstances := []Instance{}
    for i := 0; i < len(me.Instances); i++ {
        if valid[i] {
            ii := me.Instances[i]
            newInstances = append(newInstances, ii)
        }
    }
    me.Instances = newInstances

    // setup draw
    options := []*ebiten.DrawImageOptions{}
    imgs := []*ebiten.Image{}
    w, h := rc.Img1.Size()
    for i := 0; i < len(me.Instances); i++ {
        ii := me.Instances[i]
        idx := (ii.Counter + ii.Shift) % me.length
        x := me.x[ii.Counter]
        y := me.y[idx]
        op := ebiten.DrawImageOptions{}
        draw.Option(w, h, int(x), int(y), 0, &op)
        options = append(options, &op)
        imgs = append(imgs, rc.Img1)
    }

    // draw
    draw.DrawMulti(screen, imgs, options)
    msg := fmt.Sprintf("%d", len(me.Instances))
    ebitenutil.DebugPrint(screen, msg)
    return nil
}


type Proc interface {
    Update(screen *ebiten.Image) (error)
}

func main() {
    rc = NewReource()
    var pc Proc
    //pc = NewTest1()
    //pc = NewTest2()
    pc = NewTest3()
    if err := ebiten.Run(pc.Update, screenWidth, screenHeight, 1, "Rotate (Ebiten Demo)"); err != nil {
        log.Fatal(err)
    }
}
