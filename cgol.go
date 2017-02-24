package main

import (
	"fmt"
	_ "io/ioutil"
	"os"
	"strconv"
	"time"
)

const (
	MaxX int = 10
	MaxY int = 10
)

type world struct {
	*grid
}

func (w *world) load(filename string){

	file,err := os.Open(filename)

	if err != nil {
		panic(err)
	}
	bytes := make([]byte, 1024)

	//init with emptry grid
	w.grid = &grid{}

	x,y := 0,0

	for {
		count, err := file.Read(bytes)
		if err != nil {
			fmt.Println(err)
		}

		if count == 0 {
			break
		}

		for _,byte := range bytes {
			switch byte {
			case 13:
				y++
				x = 0
				fmt.Println("13 found, increasing: ",y)
				continue
			case 10:
				continue
			default:
				if byte == 48 || byte == 49 {
					fmt.Println("Recognized byte: ", byte)
					n, err := strconv.Atoi(string(byte))

					if err != nil {
						panic(err)
					}

					c := cell{now: n, next: -1}
					w.grid[y][x] = &c

				} else {
					fmt.Println("Unrecogniwed byte: ", byte)
					break
				}
			}
			x++

		}
	}
	w.grid.show()
}

type grid [MaxX][MaxY] *cell
type cell struct {
	now int
	next int
}

func (c *cell) reset(){
	if c.next == -1 {
		return
	}
	c.now = c.next
	c.next = -1
}

func (g *grid) update(x int,y int) {
	c:=g[x][y]
	c.reset()
}

func (g *grid) prepare(x int, y int){
	cell := g[x][y]
	state := g.state(x,y)
	if (cell.now == 0 && state == 3) || (cell.now == 1 && state > 1 && state < 4){
		cell.next = 1
	} else {
		cell.next = 0
	}
}

func (g *grid) state(x int, y int) int{
	nX,pX,nY,pY := (x-1+MaxX)%MaxX,(x+1)%MaxX,(y-1+MaxY)%MaxY,(y+1)%MaxY
	return g[nX][y].now + g[nX][nY].now + g[nX][pY].now + g[pX][y].now + g[pX][nY].now + g[pX][pY].now + g[x][nY].now + g[x][pY].now
}

func (g *grid) show(){
	fmt.Println("Showing...")
	for y:=0; y<MaxY;y++{
		for x:=0; x<MaxX;x++{
			fmt.Print(g[y][x].now)
		}
		fmt.Println()
	}
}

func main(){

	w := world{}
	w.load("./w1")
	for {
		t := time.Now()
		for y:=0;y<MaxY;y++{
			for x:=0;x<MaxX;x++{
				w.prepare(x, y)
			}
		}
		for y:=0;y<MaxY;y++{
			for x:=0;x<MaxX;x++{
				w.update(x, y)
			}
		}
		w.show()
		elapse := float64((1000/10)) - time.Since(t).Seconds()
		time.Sleep(time.Duration(elapse) * time.Millisecond)
	}
}
