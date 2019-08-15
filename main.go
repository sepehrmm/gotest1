package gotest1

type Z int

func Test1(n int) Z {
	return Z(n)
}

func Test2(z Z) int {
	return int(z)
}
