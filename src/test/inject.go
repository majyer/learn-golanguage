package main

import (
	"fmt"
	"github.com/codegangsta/inject"
	"reflect"
)

type SpecialString interface{}


func Say(name string, gender SpecialString, age int) {
	fmt.Printf("My name is %s, gender is %s, age is %d!\n", name, gender, age)
}

func main() {
	inj := inject.New()
	fmt.Println(inject.InterfaceOf((*interface{})(nil)))
	fmt.Println(inject.InterfaceOf((*SpecialString)(nil)))
	inj.Map("陈一回")
	inj.MapTo("男", (*SpecialString)(nil))
	inj.Map(20)
	fmt.Println("string is valid?", inj.Get(reflect.TypeOf("姓陈名一回")).IsValid())
	fmt.Println("SpecialString is valid?", inj.Get(inject.InterfaceOf((*SpecialString)(nil))).IsValid())
	fmt.Println("int is valid?", inj.Get(reflect.TypeOf(18)).IsValid())
	fmt.Println("[]byte is valid?", inj.Get(reflect.TypeOf([]byte("Golang"))).IsValid())
	inj2 := inject.New()
	inj2.Map([]byte("test"))
	inj.SetParent(inj2)
	fmt.Println("[]byte is valid?", inj.Get(reflect.TypeOf([]byte("Golang"))).IsValid())
	inj.Invoke(Say)

}