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

package main

import (
	"time"
	"github.com/the-sibyl/goMAX31856"
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
