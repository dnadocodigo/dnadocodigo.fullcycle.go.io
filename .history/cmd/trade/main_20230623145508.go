package main

import (
	"sync"

	entity "github.com/dnadocodigo/dnadocodigo.fullcycle.go.io/internal/market/entities"
)

func main(){
	ordersIn := make(chan *entity.Order)
	ordersOut := make(chan *entity.Order)
	wg := &sync.WaitGroup{}
}