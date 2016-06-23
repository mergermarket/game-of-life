package game

type GameWorld interface {
	Copy() GameWorld
	GetCells() []cell
	IsAlive(cell) bool
	GetAliveNeighbours(cell) int
	KillCell(cell)
	ResurrectCell(cell)
	String() string
}

type Game struct {
	world GameWorld
}

func NewGame(world GameWorld) *Game {
	return &Game{world}
}

func (g *Game) Step() {
	nextWorld := g.world.Copy()

	for _, cell := range g.world.GetCells() {
		neighbors := g.world.GetAliveNeighbours(cell)

		if g.world.IsAlive(cell) && neighbors < 2 {
			nextWorld.KillCell(cell)
		}
		if g.world.IsAlive(cell) && neighbors > 3 {
			nextWorld.KillCell(cell)
		}
		if g.world.IsAlive(cell) && (neighbors == 2 || neighbors == 3) {
			// Survives
		}
		if !g.world.IsAlive(cell) && neighbors == 3 {
			nextWorld.ResurrectCell(cell)
		}
	}
	g.world = nextWorld
}
