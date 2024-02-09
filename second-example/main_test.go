package main

import "testing"

func Test_updateMessage(t *testing.T) {
	msg = "Hello, world!"

	wg.Add(2)
	go updateMessage("x")
	go updateMessage("Goodbye, cool world!")

	wg.Wait()

	if msg != "Goodbye, cool world!" {
		t.Errorf("Incorrect value in msg")
	}
}
