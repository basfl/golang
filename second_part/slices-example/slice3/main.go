package main

import (
	s "github.com/inancgumus/prettyslice"
)

func main() {
	upTask := make([]string, 0, 3)
	s.Show("upTask", upTask)
	upTask = append(upTask, "a")
	s.Show("upTask", upTask)
	twdslices := [][]int{
		{1, 2, 3},
		{1},
		{4, 5, 6},
	}
	s.Show("2d", twdslices)
	twdslices = append(twdslices, []int{12, 4}, []int{8, 7, 5})
	s.Show("2d", twdslices)
	var name []byte
	s.Show("name", name)
	str := "sss"
	name = append(name, str...)
	s.Show("name", name)

}
