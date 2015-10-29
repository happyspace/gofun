package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func sum() int {
	sum := 0
	for i :=0; i < 10; i++ {
		sum += i
	}
	return sum
}

func ppsum() int {
	sum := 1
	for ; sum < 1000; {
		sum += sum
	}
	return sum
}

func whilesum() int {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	return sum
}

func infsum() int {
	sum := 1
	for {
		sum += sum
		if sum > 1000 {
			break;
		}
	}
	return sum
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

func Sqrtnewton(x float64) float64 {
	z := float64(1)
	for i := 0; i < 10; i++ {
		z = z - (z * z - x)/(2 * z)
	}
	
	return z
}

const delta = 1e-8
func Sqrtnewtondelta(x float64) (float64, int) {
	z := float64(1)
	y := 0.0
	itr := 0
	for  {
		z = z - (z * z - x)/(2 * z)
		if math.Abs(z - y) < delta {
			fmt.Println("Minus", math.Abs(z - y))
			break
		}
		if itr > 100 {
			fmt.Println("Iteration", itr)
			break
		}
		itr = itr + 1
		y = z
	}
	
	return z, itr
}

const grn = "Go runs on "
func gorunson() string {
	var goos string
	switch os := runtime.GOOS; os {
		case "darwin":
			goos = "OS X."
		case "linux":
			goos = "Linux."
		default:
			// freebsd, openbsd,
			// plan9, windows..
			goos = os
					
	}
	return grn + goos
}

const whensat = "When's Saturday? "
func saturday() string {
	var days string
	today := time.Now().Weekday()
	switch time.Saturday {
		case today + 0:
			days = "Today."
		case today + 1:
			days = "Tomorrow."	
		case today + 2:
			days = "In two days."
		default:
			days = "Too far away."		
	}
	
	return whensat + days
}

func greeting() string {
	var greeting string
	t := time.Now()
	switch {
		case t.Hour() < 12:
			greeting = "Good morning!"
		case t.Hour() < 17:
			greeting = "Good afternoon."
		default:
			greeting = "Good evening."	
			
		
	}
	return greeting
}

func deferf() {
	defer fmt.Println("world")
	fmt.Println("hello")
}

func deferc() {
	fmt.Println("counting")
	
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	
	fmt.Println("done")
}

func main() {
	fmt.Println(sum())
	fmt.Println(ppsum())
	fmt.Println(whilesum())
	fmt.Println(infsum())
	fmt.Println(sqrt(2), sqrt(-4))
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
	
	for i := 0; i < 10; i++ {
		fmt.Println(Sqrtnewton(float64(i)), math.Sqrt(float64(i)))
	}
	fmt.Println()
	for i := 0; i < 10; i++ {
		sqrt, itr := Sqrtnewtondelta(float64(i))
		fmt.Println(itr, sqrt, math.Sqrt(float64(i)))
	}
	
	fmt.Println(gorunson())
	fmt.Println(saturday())
	fmt.Println(greeting())
	deferf()
	deferc()
}

