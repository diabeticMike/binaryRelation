package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"time"

	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font/gofont/goregular"
)

func dop(r1, r2 [5][5]int) ([5][5]int, [5][5]int) {
	fmt.Println("Доповнення (r1,r2)")
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if r1[i][j] == 1 {
				r1[i][j] = 0
			} else {
				r1[i][j] = 1
			}
		}
	}

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if r2[i][j] == 1 {
				r2[i][j] = 0
			} else {
				r2[i][j] = 1
			}
		}
	}

	for _, a := range r1 {
		for _, v := range a {
			fmt.Print(v, " ")
		}
		fmt.Println()
	}

	fmt.Println()

	for _, a := range r2 {
		for _, v := range a {
			fmt.Print(v, " ")
		}
		fmt.Println()
	}
	return r1, r2
}

func per(r1, r2 [5][5]int) [5][5]int {
	fmt.Println("Перетин")

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if r1[i][j] == r2[i][j] {
			} else {
				r1[i][j] = 0
			}
		}
	}

	fmt.Println()

	for _, a := range r1 {
		for _, v := range a {
			fmt.Print(v, " ")
		}
		fmt.Println()
	}
	return r1
}

func ob(r1, r2 [5][5]int) [5][5]int {
	fmt.Println("Об'єднання")
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if r1[i][j] == 1 || r2[i][j] == 1 {
				r1[i][j] = 1
			}
		}
	}

	fmt.Println()

	for _, a := range r1 {
		for _, v := range a {
			fmt.Print(v, " ")
		}
		fmt.Println()
	}
	return r1
}

func riz(r1, r2 [5][5]int) [5][5]int {
	fmt.Println("\nРізниця")
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if r2[i][j] == 1 {
				r1[i][j] = 0
			}
		}
	}

	fmt.Println()

	for _, a := range r1 {
		for _, v := range a {
			fmt.Print(v, " ")
		}
		fmt.Println()
	}
	return r1
}

func ober(r1, r2 [5][5]int) {
	fmt.Println("\nОбернена")

	fmt.Println("r1:")

	c := [5][5]int{}
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			c[i][j] = r1[j][i]
		}
	}
	for _, a := range c {
		for _, v := range a {
			fmt.Print(v, " ")
		}
		fmt.Println()
	}

	draw(c, "ober_r1")

	fmt.Println("r2:")
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			c[i][j] = r2[j][i]
		}
	}
	for _, a := range c {
		for _, v := range a {
			fmt.Print(v, " ")
		}
		fmt.Println()
	}
	draw(c, "ober_r2")
}

func zvyg(r1, r2 [5][5]int) {
	for i := 0; i < 5; i++ {
		for j := 0; j < 3; j++ {
			r1[i][j] = 0
		}
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 5; j++ {
			r1[i][j] = 0
		}
	}
	draw(r1, "zvygennya_r1")

	for i := 0; i < 5; i++ {
		for j := 0; j < 3; j++ {
			r2[i][j] = 0
		}
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 5; j++ {
			r2[i][j] = 0
		}
	}
	draw(r2, "zvygennya_r2")
	fmt.Println("\n звуження r1")

	for _, b := range r1 {
		for _, v := range b {
			fmt.Print(v, " ")
		}
		fmt.Println()
	}

	fmt.Println("\n звуження r2")
	for _, b := range r2 {
		for _, v := range b {
			fmt.Print(v, " ")
		}
		fmt.Println()
	}
}

func comp(r1, r2 [5][5]int) {
	fmt.Println("\nКомпозиція")

	a := [5][5]int{}
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			a[i][j] = 0
			for k := 0; k < 5; k++ {
				a[i][j] += r1[i][k] * r2[k][j]
				if a[i][j] > 1 {
					a[i][j] = 1
				}
			}
		}
	}

	for _, b := range a {
		for _, v := range b {
			fmt.Print(v, " ")
		}
		fmt.Println()
	}
	draw(a, "compoz")
}

var w, h int = 1024, 1024

// x4 x5
func main() {
	r1 := [5][5]int{
		{0, 1, 0, 1, 0},
		{0, 0, 1, 1, 0},
		{0, 0, 0, 0, 1},
		{1, 0, 1, 0, 1},
		{0, 1, 0, 0, 0}}
	draw(r1, "r1")
	r2 := [5][5]int{
		{1, 0, 0, 1, 0},
		{0, 1, 1, 1, 0},
		{1, 0, 1, 1, 1},
		{1, 1, 1, 1, 1},
		{0, 1, 0, 0, 0}}
	draw(r2, "r2")
	a1, a2 := dop(r1, r2)
	draw(a1, "dopovn_r1")
	draw(a2, "dopovn_r2")

	p := per(r1, r2)
	draw(p, "peretyn")

	o := ob(r1, r2)
	draw(o, "obyednannya")
	ri := riz(r1, r2)
	draw(ri, "riznycya")
	//Симетрична різниця
	fmt.Print("Симетрична ")
	riz(o, p)
	draw(ri, "symm_riznycya")
	ober(r1, r2)

	comp(r1, r2)
	fmt.Print("Двоїсте відношення (обернена з доповнення)")
	ober(a1, a2)

	zvyg(r1, r2)
}

func draw(r [5][5]int, filename string) {
	dc := gg.NewContext(w+100, h+100)
	dc.DrawRectangle(0, 0, float64(w+100), float64(h+100))
	dc.SetRGB(102, 0, 102)
	dc.Fill()

	font, err := truetype.Parse(goregular.TTF)
	if err != nil {
		log.Fatal(err)
	}

	face := truetype.NewFace(font, &truetype.Options{Size: 48})
	dc.SetFontFace(face)

	drawNodes(r, dc)
	dc.SavePNG(fmt.Sprint(filename, ".png"))
}

type node struct {
	x, y float64
}

func drawNodes(r [5][5]int, dc *gg.Context) {
	nodes := make([]node, 0, 5)
	rand.Seed(time.Now().UnixNano())
	nodes = append(nodes, node{x: float64(rand.Intn(w - 100)), y: float64(rand.Intn(h - 100))})
	for len(nodes) < 5 {
		for {
			x, y := float64(rand.Intn(w-200)+100), float64(rand.Intn(h-200)+100)
			if check(nodes, x, y) {
				nodes = append(nodes, node{x, y})
				break
			}
		}
	}

	drawDirections(r, nodes, dc)

	for k, v := range nodes {
		if r[k][k] != 1 {
			dc.DrawCircle(v.x, v.y, 50)
			dc.SetRGB(0, 0, 0)
			dc.Fill()
		} else {
			dc.DrawCircle(v.x, v.y, 50)
			dc.SetRGB(7, 45, 239)
			dc.Fill()
		}

		dc.SetRGB(1, 1, 1)
		dc.DrawStringAnchored(fmt.Sprint(k), v.x, v.y, 0.5, 0.5)
		dc.Fill()

	}
}

func check(nodes []node, x, y float64) bool {
	for _, a := range nodes {
		if math.Abs(a.x-x) < float64(100) {
			return false
		} else if math.Abs(a.y-y) < float64(100) {
			return false
		} else if x < 50.0 || y < 50.0 {
			return false
		}
	}
	return true
}

func drawDirections(r [5][5]int, nodes []node, dc *gg.Context) {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if r[j][i] == 1 && i != j {
				dc.SetRGBA(226, 106, 106, 1)
				dc.SetLineWidth(5)
				dc.DrawLine(nodes[i].x, nodes[i].y, nodes[j].x, nodes[j].y)
				dc.Stroke()
				x := (nodes[i].x + nodes[j].x) / 2
				y := (nodes[i].y + nodes[j].y) / 2

				x = (x + nodes[j].x) / 2
				y = (y + nodes[j].y) / 2

				x = (x + nodes[j].x) / 2
				y = (y + nodes[j].y) / 2

				dc.SetRGBA(226, 106, 106, 1)
				dc.SetLineWidth(15)
				dc.DrawLine(x, y, nodes[j].x, nodes[j].y)
				dc.Stroke()
			}
		}
	}
}