package main

import (
	"time"

	"github.com/fatih/color"
)

// Barbershop represents a barber shop with its properties and methods.
type Barbershop struct {
	ShopCapacity    int           // Maximum capacity of the shop
	CutDuration     time.Duration // Duration it takes to cut hair
	NumberOfBarbers int           // Number of barbers available
	BarberDoneChan  chan bool     // Channel to indicate barber is done
	ClientsChan     chan string   // Channel for clients waiting for service
	Open            bool          // Flag indicating if the shop is open
}

// addBarber adds a new barber to the barbershop.
func (shop *Barbershop) addBarber(barber string) {
	shop.NumberOfBarbers++

	go func() {
		isSleeping := false
		color.Yellow("%s goes to the waiting room to check for clients.", barber)

		for {
			// Check if there are no clients, barber goes to sleep
			if len(shop.ClientsChan) == 0 {
				color.Yellow("There is nothing to do, so %s takes a nap.", barber)
				isSleeping = true
			}

			client, shopOpen := <-shop.ClientsChan

			if shopOpen {
				if isSleeping {
					color.Yellow("%s wakes %s up.", client, barber)
					isSleeping = false
				}
				// Perform hair cutting
				shop.cutHair(barber, client)
			} else {
				// Shop is closed, send the barber home and close the goroutine
				shop.sendBarberHome(barber)
				return
			}
		}
	}()
}

// cutHair simulates the process of a barber cutting hair for a client.
func (shop *Barbershop) cutHair(barber, client string) {
	color.Green("%s is cutting %s's hair.", barber, client)
	time.Sleep(shop.CutDuration)
	color.Green("%s is finished cutting %s", barber, client)
}

// sendBarberHome sends the barber home when the shop is closed.
func (shop *Barbershop) sendBarberHome(barber string) {
	color.Cyan("%s is going home.", barber)
	shop.BarberDoneChan <- true
}

// closeShopForDay closes the shop at the end of the day.
func (shop *Barbershop) closeShopForDay() {
	color.Cyan("Closing shop for the day.")

	close(shop.ClientsChan)
	shop.Open = false

	for a := 1; a <= shop.NumberOfBarbers; a++ {
		<-shop.BarberDoneChan
	}
	close(shop.BarberDoneChan)

	color.Green("The barbershop is now closed for the day, and everyone has gone home.")
	color.Green("---------------------------------------------------------------------")
}

// addClient adds a new client to the waiting room of the barbershop.
func (shop *Barbershop) addClient(client string) {
	color.Green("*** %s arrives!", client)

	if shop.Open {
		select {
		case shop.ClientsChan <- client:
			color.Yellow("%s takes a seat in the waiting room.", client)
		default:
			color.Red("The waiting room is full so %s leaves", client)
		}
	} else {
		color.Red("The shop is already closed, so %s leaves", client)
	}
}
