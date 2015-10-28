package main

import (
	"fmt"
	"time"
	
	"math/rand"
	"math/cmplx"
	"math"
)

var c, python, java bool

var ii, jj int = 1, 2
const Pi = 3.14

const (
	Big   = 1 << 100
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func printConversion() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z int = int(f)
	fmt.Println("Moo ", x, y, z, f)
}

var (
	ToBe bool = false
	MaxInt uint64 = 1<<64 - 1
	z complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
    fmt.Printf("hello, world\n")
    
    fmt.Println("The time is", time.Now())
    
    fmt.Println("My favorite number is", rand.Intn(10))
    
    fmt.Println("Pi ", math.Pi)
    
    fmt.Println(add(42, 13))
    
    a, b := swap("hello", "world")
    
    fmt.Println(a, b)
    
    fmt.Println(split(17))
    
    var i int
    fmt.Println(i, c, python, java)
    
    fmt.Println(ii, jj)
    
    const f = "%T(%v)\n"
    fmt.Printf(f, ToBe, ToBe)
    fmt.Printf(f, MaxInt, MaxInt)
	fmt.Printf(f, z, z)
	
	var iii int
	var fff float64
	var bbb bool
	var s string
	fmt.Printf("%v %v %v %q\n", iii, fff, bbb, s)
	
	printConversion()
	
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
	
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}

