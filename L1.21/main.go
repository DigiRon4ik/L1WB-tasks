package main

import "fmt"

// EuropeanPlug represents the European standard for plugs
type EuropeanPlug interface {
	InsertIntoEuropeanSocket()
}

// USPlug represents the American standard for plugs
type USPlug interface {
	InsertIntoUSSocket()
}

// USDevice is a device with an American plug
type USDevice struct{}

func (d *USDevice) InsertIntoUSSocket() {
	fmt.Println("American plug connected to an American socket.")
}

// Adapter adapts an American plug to work with a European socket
type Adapter struct {
	usPlug USPlug
}

func (a *Adapter) InsertIntoEuropeanSocket() {
	fmt.Println("Adapter converts the European socket for the American plug.")
	a.usPlug.InsertIntoUSSocket()
}

func main() {
	// Create an American device
	americanDevice := &USDevice{}
	// Use an adapter to connect the American plug device to a European socket
	adapter := &Adapter{usPlug: americanDevice}
	// Connect the adapter to the European socket
	adapter.InsertIntoEuropeanSocket()
}

/*
 - Output: -
Adapter converts the European socket for the American plug.
American plug connected to an American socket.
*/
