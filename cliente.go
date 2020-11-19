package main

import (
	"encoding/gob"
	"fmt"
	"net"
	"time"
)

var proceso Proceso

type MSG struct {
	Nuevo   bool
	Process Proceso
}

type Proceso struct {
	Id       uint64
	I        uint64
	Terminar bool
}

func Procesar() {
	for {
		fmt.Printf("id %d: %d\n", proceso.Id, proceso.I)
		proceso.I = proceso.I + 1
		time.Sleep(time.Millisecond * 500)
	}
}

func cliente(msg MSG) {
	c, err := net.Dial("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = gob.NewEncoder(c).Encode(msg)
	if err != nil {
		fmt.Println(err)
	} else {
		err = gob.NewDecoder(c).Decode(&proceso)
		if err != nil {
			fmt.Println(err)
		} else {
			go Procesar()
		}
	}
	c.Close()
}

func desconectar(msg MSG) {
	c, err := net.Dial("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = gob.NewEncoder(c).Encode(msg)
	if err != nil {
		fmt.Println(err)
	}
	c.Close()
}

func main() {
	msg := MSG{Nuevo: true}
	go cliente(msg)
	var input string
	fmt.Scanln(&input)
	msg = MSG{
		Nuevo:   false,
		Process: proceso,
	}
	desconectar(msg)
}
