package main

import (
	"os"
	"time"

	svg "ryansobol.com/learn-go-with-tests/maths/svg"
)

func main() {
	t := time.Now()
	svg.SVGWriter(os.Stdout, t)
}
