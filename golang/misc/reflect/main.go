package main

import (
	"log"
	"reflect"
	"time"
)

type st struct {
	tm    time.Time
	value string
}

func main() {
	li := []st{
		st{value: "test"},
		st{tm: time.Now()},
	}
	for _, d := range li {
		zt := time.Time{}
		if d.tm == zt {
			log.Println("zero")
		}
		t := reflect.ValueOf(d.tm)
		v := reflect.ValueOf(d.value)
		aa := reflect.TypeOf(1)
		log.Printf("aaaa %s", aa)
		log.Printf("v isvalid: %t", v.IsValid())
		log.Printf("v canset: %t", v.CanSet())
		log.Printf("t isvalid: %t", t.IsValid())
		log.Printf("t canset: %t", t.CanSet())
		log.Printf("%+v, %+v", v, v.Kind())
		log.Printf("%+v", d)

		log.Println(t)
		orgTime, ok := t.Interface().(time.Time)
		if ok {
			log.Printf("%s", orgTime)
		}
		log.Println(t.Type().Name())
		log.Println(t.Type().PkgPath())
	}
}
