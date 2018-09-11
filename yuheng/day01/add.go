package main

func add(x, y int) int {
	return x + y
}

func main() {
	x := 0x100
	y := 0x200
	z := x + y
	add(x, y)
	println(z)
}

//go build -gcflags "-m" -o test add.go
