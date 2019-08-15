package gotest1

type z int

func Test1(n int) z {
	return z(n)
}

func Test2(z int) int {
	return int(z)
}
