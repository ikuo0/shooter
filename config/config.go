
package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"reflect"
	"strconv"
)


type StKeyConfig struct {
	Valid bool `json:"Valid"`
	AxisUp int `json:"AxisUp"`
	AxisRight int `json:"AxisRight"`
	AxisDown int `json:"AxisDown"`
	AxisLeft int `json:"AxisLeft"`
	ButtonShot int `json:"ButtonShot"`
	ButtonBoost int `json:"ButtonBoost"`
	ButtonLock int `json:"ButtonLock"`
}

var KeyConfig StKeyConfig

func Read(filename string) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Print(err)
	} else {
		if err := json.Unmarshal(bytes, &KeyConfig); err != nil {
			log.Print(err)
		}
	}
}

func Write(filename string) {
	jsonBytes, err := json.Marshal(KeyConfig)
	if err != nil {
	} else {
		ioutil.WriteFile(filename, jsonBytes, 0666)
	}
}

func Dump() {
    v := reflect.Indirect(reflect.ValueOf(KeyConfig))
    t := v.Type()

    for i := 0; i < t.NumField(); i++ {
        println("Field: " + t.Field(i).Name)

        f := v.Field(i)
        i := f.Interface()
        if value, ok := i.(int); ok {
            println("Value: " + strconv.Itoa(value))
        } else {
            println("Value: " + f.String())
        }
    }
}

