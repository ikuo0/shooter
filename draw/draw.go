
package draw

import (
	//"bytes"
	//"image"
	//"image/color"
	//_ "image/png"
	//"log"
	"math"
	//"math/rand"
	//"time"

	"github.com/hajimehoshi/ebiten"
	//"github.com/hajimehoshi/ebiten/ebitenutil"
	//"github.com/hajimehoshi/ebiten/examples/resources/images"
	//"github.com/hajimehoshi/ebiten/inpututil"
)

/*
*/

func Rotate(w, h, angle int, op *ebiten.DrawImageOptions) {
	//op := ebiten.DrawImageOptions{}
	op.GeoM.Translate(-float64(w) / 2, -float64(h) / 2)
	op.GeoM.Rotate(float64(angle%360) * 2 * math.Pi / 360)
	//op.GeoM.Translate(screenWidth/2, screenHeight/2)
}

func Option(w, h, x, y, angle int, op *ebiten.DrawImageOptions) {
	if angle != 0 {
		Rotate(w, h, angle, op)
	}
	op.GeoM.Translate(float64(x), float64(y))
}

func Draw(dst *ebiten.Image, src *ebiten.Image, op *ebiten.DrawImageOptions) {
	/*
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(s.x+dx), float64(s.y+dy))
	op.ColorM.Scale(1, 1, 1, alpha)
	*/
	/*
	op.GeoM.Translate(-float64(w)/2, -float64(h)/2)
	op.GeoM.Rotate(float64(count%360) * 2 * math.Pi / 360)
	op.GeoM.Translate(screenWidth/2, screenHeight/2)
	*/
	dst.DrawImage(src, op)
}

func DrawMulti(dst *ebiten.Image, src []*ebiten.Image, op []*ebiten.DrawImageOptions) {
	for i := 0; i < len(src); i++ {
		Draw(dst, src[i], op[i])
	}
}

