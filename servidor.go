package main

import (
	"encoding/gob"
	"fmt"
	"net"
	"time"
)

var procesos []*Proceso

type MSG struct {
	Nuevo   bool
	Process Proceso
}

type Proceso struct {
	Id       uint64
	I        uint64
}

func Procesar() {
	for {
		for _, p := range procesos {
			fmt.Printf("id %d: %d\n", p.Id, p.I)
			p.I = p.I + 1
		}
		fmt.Println("================")
		time.Sleep(time.Millisecond * 500)
	}
}

func servidor() {
	s, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		c, err := s.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handleCliente(c)
	}
}

func handleCliente(c net.Conn) {
	var msg MSG
	err := gob.NewDecoder(c).Decode(&msg)
	if err != nil {
		fmt.Println(err)
		return
	}
	if msg.Nuevo {
		enviarProceso(c)
	} else {
		recibirProceso(msg.Process)
	}
}

func enviarProceso(c net.Conn) {
	p := procesos[0]
	procesos = append(procesos[:0], procesos[1:]...)
	err := gob.NewEncoder(c).Encode(p)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func recibirProceso(p Proceso) {
	procesos = append(procesos, &p)
}

func main() {
	var i uint64
	for i = 0; i < 5; i++ {
		p := Proceso{
			Id:       i,
			I:        0,
		}
		procesos = append(procesos, &p)
	}
	go Procesar()

	go servidor()
	var input string
	fmt.Scanln(&input)
}
