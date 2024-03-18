package main

import (
	"time"

	"github.com/fatih/color"
)

type Barbershop struct {
	ShopCapacity    int
	CutDuration     time.Duration
	NumberOfBarbers int
	BarberDoneChan  chan bool
	ClientsChan     chan string
	Open            bool
}

func (shop *Barbershop) addBarber(barber string) {
	shop.NumberOfBarbers++

	go func() {

	}()
}

func (shop *Barbershop) cutHair(barber, client string) {
	color.Green("%s is cutting %s's hair.", barber, client)
	time.Sleep(shop.CutDuration)
	color.Green("%s is finished cutting %s", barber, client)
}

func (shop *Barbershop) sendBarberHome(barber string) {
	color.Cyan("%s is going home.", barber)
	shop.BarberDoneChan <- true
}
