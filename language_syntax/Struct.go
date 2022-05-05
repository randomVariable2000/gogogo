package main

type foo struct {
	name string
	age  int8
	flag bool
}

type boo struct {
	name string
	age  int8
	flag bool
}

func main() {
	e1 := struct {
		name string
		age  int8
		flag bool
	}{
		name: "huang",
		age:  20,
		flag: false,
	}

	println(e1.name)
	e2 := boo{}
	e3 := foo{}

	// cannot use type foo as the type boo
	//e2 = e3
	// meanings that go didn't support an implicit conversion.
	// we should show our intention in code like this:
	e2 = boo(e3)

	println(e2.name)
}
