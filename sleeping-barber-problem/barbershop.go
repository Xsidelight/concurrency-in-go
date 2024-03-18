package main

import "time"

type Barbershop struct {
	ShopCapacity    int
	CutDuration     time.Duration
	NumberOfBarbers int
	BarberDoneChan  chan bool
	ClientsChan     chan string
	Open            bool
}
