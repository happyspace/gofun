package main

import (
	"fmt"
	"strings"
	"golang.org/x/tour/pic"
//	"math"
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

func geography(string, VertexLL) map[string]VertexLL {
	m = make(map[string]VertexLL)
	m["Bell Labs"] = VertexLL
	return m
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
}

