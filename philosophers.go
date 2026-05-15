package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
)

func main() {
	task13()
}

type Chops struct {
	sync.Mutex
}

type Philo struct {
	leftC  *Chops
	rightC *Chops
	number int
}

type Host struct {
	request chan chan bool
	done    chan bool
}

func (p Philo) eat(h *Host, wg *sync.WaitGroup) {
	var key int = 0
	var randNum int = 0
	for key < 3 {
		randNum = rand.IntN(2)

		permission := make(chan bool)
		h.request <- permission
		<-permission

		if randNum == 0 {
			p.leftC.Lock()
			p.rightC.Lock()
		} else {
			p.rightC.Lock()
			p.leftC.Lock()
		}

		fmt.Println("starting to eat", p.number)

		randNum = rand.IntN(2)

		fmt.Println("finishing eating", p.number)
		if randNum == 0 {
			p.leftC.Unlock()
			p.rightC.Unlock()
		} else {
			p.rightC.Unlock()
			p.leftC.Unlock()
		}

		h.done <- true

		key++
	}

	wg.Done()
}

func (host Host) manage() {
	eating := 0
	queue := []chan bool{}
	for {

		if eating < 2 && len(queue) > 0 {
			permission := queue[0]
			queue = queue[1:]
			eating++
			permission <- true
		}

		select {
		case permission := <-host.request:
			queue = append(queue, permission)

		case <-host.done:
			eating--
		}
	}
}

func task13() {
	csticks := make([]*Chops, 5)
	for i := 0; i < 5; i++ {
		csticks[i] = new(Chops)
	}

	philos := make([]*Philo, 5)
	for i := 0; i < 5; i++ {
		philos[i] = &Philo{csticks[i], csticks[(i+1)%5], i + 1}
	}

	host := &Host{make(chan chan bool), make(chan bool)}

	var wg sync.WaitGroup

	wg.Add(5)

	go host.manage()

	for i := 0; i < 5; i++ {
		go philos[i].eat(host, &wg)
	}

	wg.Wait()
}
