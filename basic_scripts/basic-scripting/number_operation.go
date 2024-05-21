package main

import (
	"fmt"
	"log"
	"math"
	"math/cmplx"
	"strconv"
)

func main() {

	//Type conversion
	var a int = 64
	var f float64 = float64(a)
	fmt.Println(f)

	var s float64 = 99.99
	var p int = int(s)
	fmt.Println(p)

	// working with complex numbers in go
	var complex_1 complex128 = complex(1, 2)               // (1+2i)
	var complex_2 complex128 = complex(3, 4)               // (3+4i)
	fmt.Println("Complex Adddition:", complex_1+complex_2) // (1+2i) + (3+4i)
	fmt.Println("Complex Subtraction", complex_1-complex_2)
	fmt.Println("Complex Multiplication", complex_1*complex_2)
	fmt.Println("Complex Division", complex_1/complex_2)
	fmt.Println("Complex Absolute Value", cmplx.Abs(complex_1))
	fmt.Println("Complex Absolute Value", cmplx.Abs(complex_2))

	// Math functions
	fmt.Println("Power", math.Pow(2, 8))      // 2^8
	fmt.Println("Square Root", math.Sqrt(49)) // √49
	fmt.Println("Sin", math.Sin(math.Pi/2))   // sin(π/2))
	fmt.Println("cos", math.Cos(math.Pi/2))   // cos(π/2)
	fmt.Println("Log", math.Log(math.E))      // log(e) where e is a constant

	// string to number parsing
	numToStr := "3.1415"
	result, err := strconv.ParseFloat(numToStr, 64)
	if err != nil {
		log.Fatalf("Error parsing string to float: ", err)
	}
	fmt.Println(result)
}
