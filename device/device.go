
package device

import (
	//"fmt"
	"log"
	"sync"
	"time"
	//"math"
    "github.com/hajimehoshi/ebiten"
)


func Keyboard() ([]int) {
    //pressed := []ebiten.Key{}
    pressed := []int{}
    for k := ebiten.Key(0); k <= ebiten.KeyMax; k++ {
        if ebiten.IsKeyPressed(k) {
            pressed = append(pressed, int(k))
        }
    }
    return pressed
}

// https://github.com/hajimehoshi/ebiten/blob/master/examples/gamepad/main.go
var padMutex = sync.Mutex{}
func PadLock() {
	padMutex.Lock()
}

func PadUnlock() {
	padMutex.Unlock()
}

type PadStatus struct {
	Id int
	MaxAxis int
	AxisAdjust []float64
}

var PadsStatus = map[int]PadStatus{}

func (me *PadStatus) GetButtonID() ([]int) {
	pressed := []int{}
	idBase := 1000 * (me.Id + 1)
	for i, axisBase := range(me.AxisAdjust) {
		v := ebiten.GamepadAxis(me.Id, i)
		v = v - axisBase
		if v < -0.5 {
			pressed = append(pressed, idBase + i * 2)
		} else if v > 0.5 {
			pressed = append(pressed, idBase + i * 2 + 1)
		}
	}
	maxButton := ebiten.GamepadButtonNum(me.Id)
	for i := 0; i < maxButton; i++ {
		if ebiten.IsGamepadButtonPressed(me.Id, ebiten.GamepadButton(i)) {
			pressed = append(pressed, idBase + 100 + i)
		}
	}
	return pressed
}

func PadRegist(id int) {
    // unlock
	PadLock()
	PadsStatus[id] = PadStatus{}
	PadUnlock()

	//
	maxAxis := ebiten.GamepadAxisNum(id)
	axis := []float64{}
	for a := 0; a < maxAxis; a++ {
		axis = append(axis, 0)
	}
	log.Printf("connect pad, id=%d, axis=%d", id, maxAxis)
	const maxCount = 10
	for i := 0; i < maxCount; i++ {
		for a := 0; a < maxAxis; a++ {
			axis[a] += ebiten.GamepadAxis(id, a)
		}
		time.Sleep(20 * time.Millisecond)
	}
	for a := 0; a < maxAxis; a++ {
		axis[a] /= maxCount
	}

	// set Pad
	PadLock()
	defer PadUnlock()
	PadsStatus[id] = PadStatus{
		Id: id,
		MaxAxis: maxAxis,
		AxisAdjust: axis,
	}
}

func PadsRegist(ids []int) {
	for _, id := range ids {
		log.Printf("unknown pad id=%d", id)
		go PadRegist(id)
	}
}

func Gamepad() ([]int) {
	PadLock()
    ids := ebiten.GamepadIDs()
	newPadsStatus := map[int]PadStatus{}
	unRegists := []int{}
	pressed := []int{}

	defer func() {
		PadUnlock()
		PadsRegist(unRegists)
	} ()

    for _, id := range ids {
		if padStat, ok := PadsStatus[id]; ok {
			pressed = append(pressed, padStat.GetButtonID()...)
			newPadsStatus[id] = padStat
		} else {
			unRegists = append(unRegists, id)
		}
    }
	PadsStatus = newPadsStatus
    return pressed
}

func Get() ([]int) {
    pressed := []int{}
    pressed = append(pressed, Keyboard()...)
    pressed = append(pressed, Gamepad()...)
    return pressed
}

func IsKB(n int) (bool) {
	return n >= 0 && n < 1000
}

func IsPad(n int) (bool) {
	return n >= 1000
}

func PadNum(n int) (int) {
	return int(n / 1000)
}

/*
パッドの十字キー（アナログスティック）
// 上から右回り 3, 1, 2, 0
// 上から右回り 7, 5, 6, 4

パッドの十字キー（ボタンタイプ）
// 上から右回り 10,11,12,13
// 上から右回り 14,15,16,17

キーボードの十字キー（十字）
// 上から右回り 99, 91, 44, 81

キーボードの十字キー（WASD）
// 上から右回り 32, 13, 28, 10


*/
