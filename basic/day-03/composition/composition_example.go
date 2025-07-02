package composition

type Engine struct {
	Power int
}

type Wheels struct {
	Count int
}

type Car struct {
	Engine Engine
	Wheels Wheels
	Color  string
}
