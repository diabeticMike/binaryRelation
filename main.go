package main

import (
	"fmt"
	"log"
	"math"

	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font/gofont/goregular"
)

func minimization(c [][]int, n int) [][]int {
	cf := [][]int{}
	for i := 0; i < n; i++ {
		cf = append(cf, []int{})
		for j := 0; j < n; j++ {
			if i != j {
				mins := 100
				for k := 0; k < n; k++ {
					mm := c[i][k] + c[k][j]
					if mm == 1 {
						mins = 1
						break
					} else if mm < mins {
						mins = mm
					}
				}
				cf[i] = append(cf[i], mins)
			} else {
				cf[i] = append(cf[i], 0)
			}
		}
	}
	return cf
}

func path(copt, m [][]int, a, b, n int) []int {
	a--
	b--
	num := []int{}
	res := make([]int, 0, a+2)
	res = append(res, a)
	for i := 0; i < n; i++ {
		num = append(num, i)
	}
	num = remove(num, a)
	for i := 0; i < n; i++ {
		if a != b {
			mins := 100
			k := 0
			for _, j := range num {
				mm := m[a][j] + copt[j][b]
				if mm == 1 {
					k = j
					break
				} else if mm < mins {
					mins = mm
					k = j
				}
			}
			if k != 0 {
				res = append(res, k)
				num = remove(num, k)
				a = k
			}
		} else {
			return res
		}
	}
	res = append(res, b)
	return res
}

func remove(a []int, b int) []int {
	r := []int{}
	for _, v := range a {
		if v != b {
			r = append(r, v)
		}
	}
	return r
}

func main() {
	m := [][]int{
		{0, 10, 10, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100},   //1
		{10, 0, 10, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100},   //2
		{100, 10, 0, 10, 3, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100},     //3
		{10, 100, 10, 0, 100, 100, 100, 8, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100},     //4
		{100, 100, 3, 100, 0, 8, 100, 100, 7, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100},       //5
		{100, 100, 100, 100, 6, 0, 1, 100, 100, 7, 100, 100, 100, 100, 100, 100, 100, 100, 100},       //6
		{100, 100, 100, 100, 100, 100, 0, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100}, //7
		{100, 100, 100, 9, 100, 100, 4, 0, 100, 100, 100, 10, 100, 100, 100, 100, 100, 100, 100},      //8
		{100, 100, 100, 100, 7, 100, 100, 100, 0, 10, 100, 100, 100, 100, 14, 100, 100, 100, 100},     //9
		{100, 100, 100, 100, 100, 5, 100, 100, 10, 0, 3, 100, 100, 100, 100, 100, 100, 100, 100},      //10
		{100, 100, 100, 100, 100, 100, 100, 100, 100, 3, 0, 3, 4, 100, 100, 100, 100, 100, 100},       //11
		{100, 100, 100, 100, 100, 100, 100, 9, 100, 100, 3, 0, 100, 100, 100, 100, 100, 100, 18},      //12
		{100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 4, 100, 0, 4, 100, 100, 100, 100, 100},     //13
		{100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 4, 0, 13, 10, 100, 100, 100},     //14
		{100, 100, 100, 100, 100, 100, 100, 100, 13, 100, 100, 100, 100, 13, 0, 100, 14, 100, 100},    //15
		{100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 9, 100, 0, 13, 10, 100},     //16
		{100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 13, 13, 0, 100, 100},   //17
		{100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 10, 100, 0, 1},    //18
		{100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 14, 100, 100, 100, 100, 100, 1, 0},    //19
	}

	n := 19

	C := m
	for {
		Cf := minimization(C, n)
		Cs := minimization(Cf, n)
		if Equal(Cf, Cs) {
			C = Cf
			break
		} else {
			C = Cs
		}
	}

	p := path(C, m, 4, 19, n)
	draw(m, p, "graph")
	for i := 0; i < len(p); i++ {
		p[i]++
	}
	fmt.Println(p)
}

func Equal(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if !equal(v, b[i]) {
			return false
		}
	}
	return true
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

var w, h int = 2543, 1344

type node struct {
	x, y float64
}

func draw(r [][]int, path []int, filename string) {
	dc := gg.NewContext(w+100, h+100)
	m, _ := gg.LoadPNG("rivne.png")
	dc.DrawImage(m, 0, 0)
	// dc.DrawRectangle(0, 0, float64(w+100), float64(h+100))
	dc.SetRGB(102, 0, 102)
	dc.Fill()

	font, err := truetype.Parse(goregular.TTF)
	if err != nil {
		log.Fatal(err)
	}

	face := truetype.NewFace(font, &truetype.Options{Size: 24})
	dc.SetFontFace(face)

	drawNodes(r, path, dc)
	dc.SavePNG(fmt.Sprint(filename, ".png"))
}

func drawNodes(r [][]int, path []int, dc *gg.Context) {
	nodes := make([]node, 0, 0)
	// rand.Seed(time.Now().UnixNano())
	// nodes = append(nodes, node{x: float64(rand.Intn(w - 100)), y: float64(rand.Intn(h - 100))})
	// for len(nodes) < 6 {
	// 	for {
	// 		x, y := float64(rand.Intn(w-200)+100), float64(rand.Intn(h-200)+100)
	// 		if check(nodes, x, y) {
	// 			nodes = append(nodes, node{x, y})
	// 			break
	// 		}
	// 	}
	// }
	nodes = append(nodes, node{x: 580, y: 110})
	nodes = append(nodes, node{x: 460, y: 200})
	nodes = append(nodes, node{x: 580, y: 330})
	nodes = append(nodes, node{x: 700, y: 230})
	nodes = append(nodes, node{x: 590, y: 380})
	nodes = append(nodes, node{x: 710, y: 370})
	nodes = append(nodes, node{x: 705, y: 350})
	nodes = append(nodes, node{x: 780, y: 335})
	nodes = append(nodes, node{x: 613, y: 475})
	nodes = append(nodes, node{x: 750, y: 450})
	nodes = append(nodes, node{x: 820, y: 470})
	nodes = append(nodes, node{x: 880, y: 450})
	nodes = append(nodes, node{x: 870, y: 540})
	nodes = append(nodes, node{x: 890, y: 635})
	nodes = append(nodes, node{x: 660, y: 690})
	nodes = append(nodes, node{x: 925, y: 780})
	nodes = append(nodes, node{x: 720, y: 950})
	nodes = append(nodes, node{x: 1055, y: 670})
	nodes = append(nodes, node{x: 1050, y: 645})

	// drawDirections(r, nodes, dc)
	drawPath(r, nodes, path, dc)

	for k, v := range nodes {
		// if r[k][k] == 0 {
		dc.DrawCircle(v.x, v.y, 15)
		dc.SetRGB(0, 0, 0)
		dc.Fill()
		// } else {
		// 	dc.DrawCircle(v.x, v.y, 20)
		// 	dc.SetRGB(7, 45, 239)
		// 	dc.Fill()
		// }

		dc.SetRGB(1, 1, 1)
		dc.DrawStringAnchored(fmt.Sprint(k+1), v.x, v.y, 0.5, 0.5)
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

func drawDirections(r [][]int, nodes []node, dc *gg.Context) {
	for i := 0; i < 19; i++ {
		for j := 0; j < 19; j++ {
			if r[i][j] == 100 {
				continue
			}
			if r[i][j] == 0 {
				break
			}

			dc.SetRGBA(226, 106, 106, 1)
			dc.SetLineWidth(6)
			dc.DrawLine(nodes[i].x, nodes[i].y, nodes[j].x, nodes[j].y)
			dc.Stroke()
			// x := (nodes[i].x + nodes[j].x) / 2
			// y := (nodes[i].y + nodes[j].y) / 2

			// x = (x + nodes[j].x) / 2
			// y = (y + nodes[j].y) / 2

			// x = (x + nodes[j].x) / 2
			// y = (y + nodes[j].y) / 2

			// dc.SetRGBA(226, 106, 106, 1)
			// dc.SetLineWidth(15)
			// dc.DrawLine(x, y, nodes[j].x, nodes[j].y)
			// dc.Stroke()
		}
	}
}

func drawPath(r [][]int, nodes []node, path []int, dc *gg.Context) {
	if len(path) <= 1 {
		return
	}
	for i := 0; i < len(path)-1; i++ {
		second := path[i+1]
		dc.SetRGBA(0, 255, 0, 1)
		dc.SetLineWidth(6)
		dc.DrawLine(nodes[path[i]].x, nodes[path[i]].y, nodes[second].x, nodes[second].y)
		dc.Stroke()
		x := (nodes[path[i]].x + nodes[second].x) / 2
		y := (nodes[path[i]].y + nodes[second].y) / 2

		x = (x + nodes[second].x) / 2
		y = (y + nodes[second].y) / 2

		x = (x + nodes[second].x) / 2
		y = (y + nodes[second].y) / 2

		// dc.SetRGBA(226, 106, 106, 1)
		// dc.SetLineWidth(15)
		// dc.DrawLine(x, y, nodes[second].x, nodes[second].y)
		// dc.Stroke()
	}
}
