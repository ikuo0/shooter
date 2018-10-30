
package application

import (
	//"fmt"
	//"log"
	//"sync"
	//"time"
	//"math"
    //"github.com/hajimehoshi/ebiten"
)

type Message struct {
	Command string
	Param1 string
	Param2 string
}

type Application struct {
	messages []Message
}

func (me *Application) Post(command, param1, param2 string) {
	me.messages = append(me.messages, Message{command, param1, param2})
}

func (me *Application) Receive() (Message, bool) {
	if len(me.messages) == 0 {
		return Message{}, false
	} else {
		res := me.messages[0]
		me.messages = me.messages[1:]
		return res, true
	}
}

func New() (*Application) {
	return &Application {
	}
}
