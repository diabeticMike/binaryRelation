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

var w, h int = 1024, 1024

func reflecsive(r [5][5]int) bool {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == j {
				if r[i][j] != 1 {
					return false
				}
			}
		}
	}
	return true
}

func antyreflecsive(r [5][5]int) bool {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == j {
				if r[i][j] != 0 {
					return false
				}
			}
		}
	}
	return true
}

func symmetrical(r [5][5]int) bool {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if r[i][j] != r[j][i] {
				return false
			}
		}
	}
	return true
}

func asymmetrical(r [5][5]int) bool {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {

			if r[i][j] == 1 && r[j][i] == 1 {
				return false
			}
		}
	}
	return true
}

func antysymmetrical(r [5][5]int) bool {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i != j {
				if r[i][j] == 1 && r[j][i] == 1 {
					return false
				}
			}
		}
	}
	return true
}

func transit(r1, r2 [5][5]int) bool {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if r1[i][j] == 1 {
				if r1[i][j] != r2[i][j] {
					return false
				}
			}
		}
	}
	return true
}

func acyclic(r1, r2 [5][5]int) bool {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if r1[i][j] == 1 {
				if r1[i][j] == r2[i][j] {
					return false
				}
			}
		}
	}
	return true
}

func connected(r [5][5]int) bool {
	for i := 0; i < 5; i++ {
		flag := false
		for j := 0; j < 5; j++ {
			if r[j][i] == 1 {
				flag = true
				break
			}
		}
		if !flag {
			return false
		}
	}
	return true
}

// Maximum
func findMax(r [5][5]int) []int {
	res := []int{}
	for i := 0; i < 5; i++ {
		if r[i] == [5]int{1, 1, 1, 1, 1} {
			res = append(res, i)
		}
	}

	if len(res) == 0 {
		return nil
	}
	return res
}

// Minimum
func findMin(r [5][5]int) []int {
	res := []int{}
	for i := 0; i < 5; i++ {
		if r[:][i] == [5]int{1, 1, 1, 1, 1} {
			res = append(res, i)
		}
	}

	if len(res) == 0 {
		return nil
	}
	return res
}

// Minoranta
func findMinor(r [5][5]int) []int {
	res := []int{}
	for i := 0; i < 5; i++ {
		if r[i] == [5]int{0, 0, 0, 0, 0} {
			res = append(res, i)
		}
	}

	if len(res) == 0 {
		return nil
	}
	return res
}

// Majoranta
func findMajor(r [5][5]int) []int {
	res := []int{}
	for i := 0; i < 5; i++ {
		if r[:][i] == [5]int{0, 0, 0, 0, 0} {
			res = append(res, i)
		}
	}

	if len(res) == 0 {
		return nil
	}
	return res
}

func main() {
	r := [5][5]int{
		{0, 0, 1, 1, 0},
		{0, 0, 1, 1, 0},
		{1, 1, 0, 1, 1},
		{1, 1, 1, 0, 1},
		{0, 0, 1, 1, 0}}
	draw(r, "r")

	// fmt.Print("Відношення еквівалентності ")
	// fmt.Println(reflecsive(r) && symmetrical(r) && transit(comp(r, r), r))

	// fmt.Print("Відношення нестрогого порядку ")
	// fmt.Println(reflecsive(r) && asymmetrical(r) && transit(comp(r, r), r))

	// fmt.Print("Відношення строгого порядку ")
	// fmt.Println(reflecsive(r) && asymmetrical(r) && transit(comp(r, r), r))

	// fmt.Print("Відношення лінійного порядку ")
	// fmt.Println(reflecsive(r) && antysymmetrical(r) && transit(comp(r, r), r))

	// fmt.Print("Відношення домінування ")
	// fmt.Println(antyreflecsive(r) && asymmetrical(r))

	// fmt.Print("Відношення подібності ")
	// fmt.Println(reflecsive(r) && symmetrical(r))

	fmt.Print("Відношення еквівалентності ")
	c := comp(r, ober(r))
	for _, a := range c {
		for _, v := range a {
			fmt.Print(v, " ")
		}
		fmt.Println()
	}
	draw(c, "equivalent")
	fmt.Print("Відношення еквівалентності ")
	fmt.Println(reflecsive(c) && symmetrical(c) && transit(comp(c, c), c))

	// fmt.Print("Відношення домінування ")
	// c := comp(r, ober(r))
	// for _, a := range c {
	// 	for _, v := range a {
	// 		fmt.Print(v, " ")
	// 	}
	// 	fmt.Println()
	// }
	// fmt.Print("Відношення домінування ")
	// fmt.Println(antyreflecsive(c) && asymmetrical(c))

	fmt.Print("Відношення толерантності ")
	c = ob(ob(r, ober(r)), per(r, ober(r)))
	for _, a := range c {
		for _, v := range a {
			fmt.Print(v, " ")
		}
		fmt.Println()
	}
	draw(c, "tolerant")

	fmt.Print("Відношення строгої переваги ")
	c = riz(r, ober(r))
	for _, a := range c {
		for _, v := range a {
			fmt.Print(v, " ")
		}
		fmt.Println()
	}
	draw(c, "strong_advandage")

	fmt.Print("Максимум ")
	if m := findMax(r); m != nil {
		fmt.Println(m)
	} else {
		fmt.Println("відсутній")
	}

	fmt.Print("Мінімум ")
	if m := findMin(r); m != nil {
		fmt.Println(m)
	} else {
		fmt.Println("відсутній")
	}

	fmt.Print("Мажоранта ")
	if m := findMajor(r); m != nil {
		fmt.Println(m)
	} else {
		fmt.Println("відсутня")
	}

	fmt.Print("Міноранта ")
	if m := findMinor(r); m != nil {
		fmt.Println(m)
	} else {
		fmt.Println("відсутня")
	}

	draw(r, "r")
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

func dop(r [5][5]int) [5][5]int {
	// fmt.Println("Доповнення (r1,r2)")
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if r[i][j] == 1 {
				r[i][j] = 0
			} else {
				r[i][j] = 1
			}
		}
	}

	// for _, a := range r1 {
	// 	for _, v := range a {
	// 		fmt.Print(v, " ")
	// 	}
	// 	fmt.Println()
	// }
	return r
}

func per(r1, r2 [5][5]int) [5][5]int {
	// fmt.Println("Перетин")

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if r1[i][j] == r2[i][j] {
			} else {
				r1[i][j] = 0
			}
		}
	}

	// fmt.Println()

	// for _, a := range r1 {
	// 	for _, v := range a {
	// 		fmt.Print(v, " ")
	// 	}
	// 	fmt.Println()
	// }
	return r1
}

func ob(r1, r2 [5][5]int) [5][5]int {
	// fmt.Println("Об'єднання")
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if r1[i][j] == 1 || r2[i][j] == 1 {
				r1[i][j] = 1
			}
		}
	}

	// fmt.Println()

	// for _, a := range r1 {
	// 	for _, v := range a {
	// 		fmt.Print(v, " ")
	// 	}
	// 	fmt.Println()
	// }
	return r1
}

func riz(r1, r2 [5][5]int) [5][5]int {
	// fmt.Println("\nРізниця")
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if r2[i][j] == 1 {
				r1[i][j] = 0
			}
		}
	}

	// fmt.Println()

	// for _, a := range r1 {
	// 	for _, v := range a {
	// 		fmt.Print(v, " ")
	// 	}
	// 	fmt.Println()
	// }
	return r1
}

func ober(r1 [5][5]int) [5][5]int {
	// fmt.Println("\nОбернена")

	// fmt.Println("r1:")

	c := [5][5]int{}
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			c[i][j] = r1[j][i]
		}
	}
	// for _, a := range c {
	// 	for _, v := range a {
	// 		fmt.Print(v, " ")
	// 	}
	// 	fmt.Println()
	// }

	draw(c, "ober_r1")
	return c
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
	// fmt.Println("\n звуження r1")

	// for _, b := range r1 {
	// 	for _, v := range b {
	// 		fmt.Print(v, " ")
	// 	}
	// 	fmt.Println()
	// }

	// fmt.Println("\n звуження r2")
	// for _, b := range r2 {
	// 	for _, v := range b {
	// 		fmt.Print(v, " ")
	// 	}
	// 	fmt.Println()
	// }
}

func comp(r1, r2 [5][5]int) [5][5]int {
	// fmt.Println("\nКомпозиція")

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

	// for _, b := range a {
	// 	for _, v := range b {
	// 		fmt.Print(v, " ")
	// 	}
	// 	fmt.Println()
	// }
	// draw(a, "compoz")
	return a
}
