
package main

import (
    "./application"
	"./device"
	"./draw"
    //"bytes"
    //"image"
    //"fmt"
    _ "image/jpeg"
    "log"
    //"math"
    //"math/rand"
	"os"

    "github.com/hajimehoshi/ebiten"
    "github.com/hajimehoshi/ebiten/ebitenutil"
    //"github.com/pa-m/numgo"
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

type SceneInterface interface {
	Name() (string)
    Calc([]int)
    PreDraw() ([]*ebiten.Image, []*ebiten.DrawImageOptions)
	IsActive() (bool)
}

var (
	rc *Resource
	app *application.Application
	scenes []SceneInterface
)

////////////////////////////////////////
// Menu
////////////////////////////////////////
type SceneMenu struct {
	app *application.Application
}
func (me *SceneMenu) Name() (string) {
	return "Menu"
}
func (me *SceneMenu) Calc(pressed []int) {
	return
}
func (me *SceneMenu) PreDraw() ([]*ebiten.Image, []*ebiten.DrawImageOptions) {
	return []*ebiten.Image{}, []*ebiten.DrawImageOptions{}
}
func (me *SceneMenu) IsActive() (bool) {
	return true
}
func NewMenu(app *application.Application) (*SceneMenu) {
	app.Post("apply", "config", "")
	return &SceneMenu{app}
}

////////////////////////////////////////
// Config
////////////////////////////////////////
type SceneConfig struct {
}
func (me *SceneConfig) Name() (string) {
	return "Config"
}
func (me *SceneConfig) Calc(pressed []int) {
	return
}
func (me *SceneConfig) PreDraw() ([]*ebiten.Image, []*ebiten.DrawImageOptions) {
	return []*ebiten.Image{}, []*ebiten.DrawImageOptions{}
}
func (me *SceneConfig) IsActive() (bool) {
	return true
}
func NewConfig(app *application.Application) (*SceneConfig) {
	return &SceneConfig{}
}

////////////////////////////////////////
// Main
////////////////////////////////////////
type SceneMain struct {
}
func (me *SceneMain) Name() (string) {
	return "Main"
}
func (me *SceneMain) Calc(pressed []int) {
	return
}
func (me *SceneMain) PreDraw() ([]*ebiten.Image, []*ebiten.DrawImageOptions) {
	return []*ebiten.Image{}, []*ebiten.DrawImageOptions{}
}
func (me *SceneMain) IsActive() (bool) {
	return true
}
func NewMain(app *application.Application) (*SceneMain) {
	return &SceneMain{}
}

func MessageProc() {
	for msg, b := app.Receive(); b; msg, b = app.Receive() {
        switch msg.Command {
            case "new":
                scenes = []SceneInterface{}
                fallthrough
            case "apply":
                newScene := (SceneInterface)(nil)
                switch msg.Param1 {
                    case "menu":
                        newScene = NewMenu(app)
                    case "config":
                        newScene = NewConfig(app)
                    case "main":
                        newScene = NewMain(app)
                }
                scenes = append(scenes, newScene)
            case "exit":
                os.Exit(0)
        }
    }
}

func Update(screen *ebiten.Image) (error) {
    // Global process
    MessageProc()

    // Input
    pressed := device.Get()

    // Calc
    for _, scene := range scenes {
        scene.Calc(pressed)
    }
    
    if ebiten.IsDrawingSkipped() {
        return nil
    }

    // Draw
    for i, scene := range scenes {
        imgs, options := scene.PreDraw()
        draw.DrawMulti(screen, imgs, options)
		//ebitenutil.DebugPrint(screen, scene.Name())
		ebitenutil.DebugPrintAt(screen, scene.Name(), 0, i * 16)
    }

    newScenes := []SceneInterface{}
    for _, scene := range scenes {
        if scene.IsActive() {
            newScenes = append(newScenes, scene)
        }
    }
    scenes = newScenes

    return nil
}

func main() {
    rc = NewReource()
	app = application.New()
	app.Post("new", "menu", "")
    if err := ebiten.Run(Update, screenWidth, screenHeight, 1, "Rotate (Ebiten Demo)"); err != nil {
        log.Fatal(err)
    }
}
