package main

import (
	"fmt"
	"strings"
	"golang.org/x/tour/pic"
	"golang.org/x/tour/wc"
	"math"
)

func pointer() {
	i, j := 42, 2701
	
	p := &i
	fmt.Println(*p)
	*p = 21
	fmt.Println(i)
	
	p = &j
	*p = *p / 37
	fmt.Println(j)
}

type Vertex struct {
	X int
	Y int
}

func structLiterals() {
	var (
		v1 = Vertex{1,2}
		v2 = Vertex{X: 1}
		v3 = Vertex{}
		p = &Vertex{1, 2}
		f = "%T\n"
		
	)

	
	fmt.Println(v1, p, v2, v3)
	fmt.Printf(f, p)
}

func array() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)
}

func slice() {
	s := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("s ==", s)
	
	for i := 0; i < len(s); i++ {
		fmt.Printf("s[%d] == %d\n", i, s[i])
	}
}

func tictactoe() {
	game := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	game[0][0] = "X"
	game[2][2] = "O"
	game[2][0] = "X"
	game[1][0] = "O"
	game[0][2] = "X"
	
	printBoard(game)
}

func printBoard(s [][]string) {
	for i := 0; i < len(s); i++ {
		fmt.Printf("%s\n", strings.Join(s[i], " "))
	}
}

func slicing() {
	s := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("s ==", s)
	fmt.Println("s[1:4] ==", s[1:4])
	
	fmt.Println("s[:3] ==", s[:3])
	
	fmt.Println("s[4:] ==", s[4:])
}

func makeslice() {
	a := make([]int, 5)
	printSlice("a", a)
	b := make([]int, 0, 5)
	printSlice("b", b)
	c := b[:2]
	printSlice("c", c)
	d := c[2:5]
	printSlice("d", d)
}

func zeroSlice() {
	var z []int
	fmt.Println(z, len(z), cap(z))
	if z == nil {
		fmt.Println("nil!")
	}
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", 
		s, len(x), cap(x), x)
}

func appendSlice() {
	var a []int
	printSlice("a", a)
	
	a = append(a, 0)
	printSlice("a", a)
	
	a = append(a, 1)
	printSlice("a", a)
	
	a = append(a, 2, 3, 4)
	printSlice("a", a)
}

func ranges() {
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}

func moreRanges() {
	pow := make([]int, 10)
	
	for i := range pow {
		pow[i] = 1 << uint(i)
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}

func Pic(dx, dy int) [][]uint8 {
	var pic [][]uint8 = make([][]uint8, dx)
	
	for i := range pic {
		pic[i] = make([]uint8, dy)
		for j := range pic[i] {
			// pic[i][j] = uint8(math.Pow(float64(i), float64(j)))
			// pic[i][j] = uint8((i+j)/2)
			pic[i][j] = uint8(i*j)
		}
	}
	
	return pic
}

type VertexLL struct {
	Lat, Long float64
}

var m = map[string]VertexLL{
	"Bell Labs": {
		40.68433, -74.39967,
	},
	"Google": {
		37.42202, -122.08408,
	},
}

func geography(key string, vLL VertexLL) map[string]VertexLL {
	if m == nil {
		m = make(map[string]VertexLL)
	}
	m[key] = vLL
	return m
}

func mutateMap() {
	m := make(map[string]int)
	
	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])
	
	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])
	
	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	words := strings.Fields(s)
	for i := range words {
		word := words[i]
		elem, ok := m[word]
		if ok {
			m[word] = elem + 1
		} else {
			m[word] = 1
		}
	}
	
	return m
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

var hypot = func(x, y float64) float64 {
	return math.Sqrt(x*x + y*y)
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func exAdder() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}


func fibonacci() func() int {
	fib := -1
	ona := -1
	return func() int {
		if fib < 0 {
			fib = 0
			return fib
		} else if ona < 0 {
			ona = 1
			return ona	
		} else {
			acci := fib + ona
			fib = ona
			ona = acci
			return acci
		}
	}
}

func main() {
	pointer()
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v)
	
	p := &v
	p.X = 1e9
	fmt.Println(v)
	
	structLiterals()
	array()
	slice()
	tictactoe()
	slicing()
	
	makeslice()
	zeroSlice()
	appendSlice()
	ranges()
	
	pic.Show(Pic)
	
	var vLL = VertexLL{40.68433, -74.39967,}
	geography("Bell Labs", vLL)
	fmt.Println(m["Bell Labs"])
	vLL = VertexLL{37.4828, 122.2361,}
	geography("Redwood City", vLL)
	fmt.Println(m["Redwood City"])
	
	fmt.Println(m)
	
	mutateMap()
	wc.Test(WordCount)
	
	fmt.Println(hypot(5, 12))
	
	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
	
	exAdder()
	
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
	
}

