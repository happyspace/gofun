package main

import (
	"fmt"
	"math"
	"os"
	"time"
)

type MyFloat float64

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprint(float64(e))
}

type Vertex struct {
	X, Y float64
}

type Abser interface {
	Abs() float64
}

type MyError struct {
	When time.Time
	What string
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	return math.Sqrt(x), nil
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
} 

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v Vertex) ScaleV(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func interfaceFun() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}
	
	a = f 
	a = &v
	
	// a = v -- error Vertex vs *Vertex
	
	
	fmt.Println(a.Abs())
}

type Reader interface {
	Read(b []byte) (n int, err error)
}

type Writer interface {
	Write(b []byte) (n int, err error)
}

type ReadWriter interface {
	Reader
	Writer
}

type Person struct {
	Name string 
	Age int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

type IPAddr [4]byte

func (ip IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3])
}

func main() {
	v := &Vertex{3, 4}
	fmt.Println(v.Abs())
	
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
	
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs())
	v.Scale(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.Abs())
	v.ScaleV(5)
	fmt.Printf("By value scaling: %+v, Abs: %v\n", v, v.Abs())
	
	interfaceFun()
	
	var w Writer
	w = os.Stdout
	
	fmt.Fprintf(w, "hello, writer\n")
	
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)
	
	addrs := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	
	for n, a := range addrs {
		fmt.Printf("%v %v\n", n, a)
	}
	
	if err := run(); err != nil {
		fmt.Println(err)
	}
}

