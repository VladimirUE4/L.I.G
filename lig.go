package main

import (
	"image/color"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	screenWidth  = 800
	screenHeight = 800
	cellSize     = 10
	rows         = screenHeight / cellSize
	cols         = screenWidth / cellSize
)

type Game struct {
	grid    [][]int
	newGrid [][]int
}

func NewGame() *Game {
	grid := make([][]int, rows)
	newGrid := make([][]int, rows)
	for i := range grid {
		grid[i] = make([]int, cols)
		newGrid[i] = make([]int, cols)
		for j := range grid[i] {
			grid[i][j] = rand.Intn(2)
		}
	}
	return &Game{grid: grid, newGrid: newGrid}
}

func (g *Game) countNeighbors(x, y int) int {
	count := 0
	directions := []struct{ dx, dy int }{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	for _, dir := range directions {
		nx, ny := x+dir.dx, y+dir.dy
		if nx >= 0 && nx < rows && ny >= 0 && ny < cols {
			count += g.grid[nx][ny]
		}
	}
	return count
}

func (g *Game) updateGrid() {
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			neighbors := g.countNeighbors(i, j)
			if g.grid[i][j] == 1 && (neighbors == 2 || neighbors == 3) {
				g.newGrid[i][j] = 1
			} else if g.grid[i][j] == 0 && neighbors == 3 {
				g.newGrid[i][j] = 1
			} else {
				g.newGrid[i][j] = 0
			}
		}
	}

	// Swap grids
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			g.grid[i][j] = g.newGrid[i][j]
		}
	}
}

func (g *Game) Update() error {
	g.updateGrid()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Clear()
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if g.grid[i][j] == 1 {
				x := float64(j * cellSize)
				y := float64(i * cellSize)
				ebitenutil.DrawRect(screen, x, y, cellSize, cellSize, color.RGBA{0, 255, 0, 255})
			}
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	rand.Seed(time.Now().UnixNano())
	game := NewGame()

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Vous savez, moi je ne crois pas qu’il y ait de bonne ou de mauvaise situation. Moi, si je devais résumer ma vie aujourd’hui avec vous, je dirais que c’est d’abord des rencontres. Des gens qui m’ont tendu la main, peut-être à un moment où je ne pouvais pas, où j’étais seul chez moi. Et c’est assez curieux de se dire que les hasards, les rencontres forgent une destinée... Parce que quand on a le goût de la chose, quand on a le goût de la chose bien faite, le beau geste, parfois on ne trouve pas l’interlocuteur en face je dirais, le miroir qui vous aide à avancer. Alors ça n’est pas mon cas, comme je disais là, puisque moi au contraire, j’ai pu ; et je dis merci à la vie, je lui dis merci, je chante la vie, je danse la vie... je ne suis qu’amour ! Et finalement, quand des gens me disent « Mais comment fais-tu pour avoir cette humanité ? », je leur réponds très simplement que c’est ce goût de l’amour, ce goût donc qui m’a poussé aujourd’hui à entreprendre une construction mécanique... mais demain qui sait ? Peut-être simplement à me mettre au service de la communauté, à faire le don, le don de soi.")

	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
