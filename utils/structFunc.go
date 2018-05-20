package main

type User struct {
	findID func(int) (int, error)
	name   string
	age    int
}

func main() {
	n := NewUser()
	n.age ,_= n.findID(10)

}

func NewUser() *User{
	return &User{
		findID: func(i int) (int, error) {
			return i* 10, nil
		},
		name: "wolf",
	}
}
