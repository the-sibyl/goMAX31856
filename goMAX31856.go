/*
Copyright (c) 2018 Forrest Sibley <My^Name^Without^The^Surname@ieee.org>

Permission is hereby granted, free of charge, to any person obtaining
a copy of this software and associated documentation files (the
"Software"), to deal in the Software without restriction, including
without limitation the rights to use, copy, modify, merge, publish,
distribute, sublicense, and/or sell copies of the Software, and to
permit persons to whom the Software is furnished to do so, subject to
the following conditions:

The above copyright notice and this permission notice shall be
included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/

//package goMAX31856
package main

import (
	"fmt"
	"time"
	"github.com/the-sibyl/piSPI"
)

func main() {

	o := spi.Devfs{
		Dev:	"/dev/spidev0.0",
		Mode:	spi.Mode1,
		MaxSpeed: 100000,
	}

	dev, err := spi.Open(&o)

	if err != nil {
		panic(err)
	}

	defer dev.Close()

	dev.SetCSChange(false)
	time.Sleep(time.Millisecond * 200)


	//readValue := make([]byte, 2)

	err = dev.Tx([]byte{
		0x80, 0xB4,
	}, nil)

	if err != nil {
		panic(err)
	}

	err = dev.Tx([]byte{
		0x82, 0x0,
	}, nil)

	if err != nil {
		panic(err)
	}

	getTemp(dev)

}

type MAX31856 struct {
	spidevPath string
	spiClockSpeed int
}

// Add functionality for DRDY pin
func (m *MAX31856) Setup() {

}

// Reset the faults register
func (m *MAX31856) ResetFaults() {

}

// Intended to be placed into a Goroutine
// Singleton behavior
func (m *MAX31856) GetTempAuto() {

}

// Intended to be called once per measurement
func (m *MAX31856) GetTempOnce() {

}

func getTemp(dev *spi.Device) {

	readValue := make([]byte, 4)

	// Read 0xC, 0xD, 0xE. The address auto-increments on the chip.
	dev.Tx([]byte{
	0xC, 0x0, 0x0, 0x0,
	}, readValue)

	// Discard the first byte, save the rest, and shift them to their proper positions
	var temp := uint32(readValue[1]) << 16 | uint32(readValue[2]) << 8 | uint32(readValue[3])

	fmt.Println(readValue)
	fmt.Println(temp)

}
