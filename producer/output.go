package main

import "fmt"

type output interface {
    out()
}

type stdout struct {
    data string
}

type kafka struct {
    data string
}

func (s stdout) out() {
    fmt.Println(s.data)
}

func (s kafka) out() {
    fmt.Println("Not Implemented")
}