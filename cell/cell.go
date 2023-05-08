package cell

type Cell interface{
	IsAlive() bool
	NextGeneration(int) (Cell, error)
}

