
package config

import (
	"testing"
	"fmt"
	//"github.com/ikuo0/shooter/config"
)

func TestAll(t *testing.T) {
	fmt.Println("start")
	filename := ".\\config.json"
	Read(filename)
	Dump()
	Write(filename)
}

