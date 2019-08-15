package gotest1

type z int

func Test1(n int) z {
	return z(n)
}

func Test2(z z) int {
	return int(z)
}
