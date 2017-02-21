package main

import (
	"time"
	"encoding/json"
	"fmt"
)

const (
	MaxX int = 10
	MaxY int = 10
)

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


func config(res *response) grid{
	world := grid{}
	for y:=0;y<MaxY;y++{
		for x:=0;x<MaxX;x++{
			world[x][y] = &cell{now:0,next:-1,}
		}
	}

	for _,e := range *res {
		world[e[0]][e[1]].next = 1
	}

	return world
}

type response [][2]int

func main(){
	res := response{}
	data := `[
		[5,5],
		[6,5],
		[7,5]
	]`

	json.Unmarshal([]byte(data), &res)
	world := config(&res)
	for {
		t := time.Now()
		for y:=0;y<MaxY;y++{
			for x:=0;x<MaxX;x++{
				world.prepare(x, y)
			}
		}
		for y:=0;y<MaxY;y++{
			for x:=0;x<MaxX;x++{
				world.update(x, y)
			}
		}
		fmt.Println(world)
		elapse := float64((1000/10)) - time.Since(t).Seconds()
		time.Sleep(time.Duration(elapse) * time.Millisecond)
	}
}
