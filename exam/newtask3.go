package main

import "fmt"

type cuboid struct {
	len int
	wid int
	hig int
}

func (mycuboid *cuboid) Volume() int {

	return mycuboid.len * mycuboid.hig * mycuboid.wid
}

func (mycuboid *cuboid) Area() int {
	return (mycuboid.hig*mycuboid.len + mycuboid.hig*mycuboid.wid + mycuboid.len*mycuboid.wid) * 2
}

func main() {
	var mycuboid cuboid
	mycuboid.hig = 3
	mycuboid.len = 3
	mycuboid.wid = 3
	fmt.Println(mycuboid.Volume())
	fmt.Println(mycuboid.Area())
}
