package main

import (
	"sync"

	entity "github.com/dnadocodigo/dnadocodigo.fullcycle.go.io/internal/market/entities"
	"golang.org/x/tools/go/analysis/passes/defers"
)

func main(){
	ordersIn := make(chan *entity.Order)
	ordersOut := make(chan *entity.Order)
	wg := &sync.WaitGroup{}
	defers wg.Wait{}
	kafkaMsgChan := make(chan, *c)
}