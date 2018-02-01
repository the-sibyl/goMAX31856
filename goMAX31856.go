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
//	"golang.org/x/exp/io/spi"
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

//	fmt.Println(dev.SetMaxSpeed(7629))

//	dev.SetBitOrder(spi.LSBFirst)
//	dev.SetBitsPerWord(8)

/*
	err = dev.Tx([]byte{
//		0x80, 0b10010000
		0x80, 0x90,
	//	0x82, 0x00,
	}, nil)

	if err != nil {
		panic(err)
	}
*/

	time.Sleep(time.Millisecond * 200)


	readValue := make([]byte, 2)

	dev.SetCSChange(false)

	for {
/*
//		dev.SetCSChange(true)
		dev.Tx([]byte{
		0x80, 0x99,
		}, readValue)

		fmt.Println("Write,", readValue)

*/
		time.Sleep(time.Millisecond * 100)
//		dev.SetCSChange(false)

		dev.Tx([]byte{
		0x01, 0x0,
		}, readValue)

		fmt.Println("Read,", readValue)

		time.Sleep(time.Millisecond * 100)


	}
}
