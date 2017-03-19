package main

type state int

const (
	state0 state	= iota
	state1 state 	= iota
)

type cell struct {
	previous *cell
	next *cell

	point

	current state
	future state

}

type point struct {
	x int
	y int
}

type world struct {
	first *cell
	last  *cell
	sentinel *cell
}

func (w *world) init (c * cell){
	w.first = c
	w.last = c
	w.sentinel = c
}

func (w *world) addCell (c *cell){
	w.last.next = c
	c.previous = w.last
	w.last = c
}

//Add cell at the sentinell position
//Set the new cell as the sentinel
func (w *world) insertCell (c *cell){
	w.sentinel.previous.next = c
	w.sentinel.previous = c
	w.sentinel = c
}

//Remove cell at sentinell position
//Setting the previous as the new sentinel
func (w *world) removeCell(){

	var previous *cell = nil
	if(w.sentinel.next != nil){
		w.sentinel.next.previous = w.sentinel.previous

		if(w.sentinel == w.first){
			w.first = w.sentinel.next
		}

	}

	if(w.sentinel.previous != nil){
		w.sentinel.previous.next = w.sentinel.next

		if(w.sentinel == w.last){
			w.last = w.sentinel.previous
		}

		previous = w.sentinel.previous
	}


	w.sentinel.next = nil
	w.sentinel.previous = nil

	w.sentinel = previous
}

func main(){

	w := world{}


}

