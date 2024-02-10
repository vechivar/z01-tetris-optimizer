package main

// structure utilisée pour construire le graph des pièces
type Graph struct {
	id         int
	neighbours []*Graph
	color      int
}

// Assigne un id de couleur à chacune des pièces.
func BuildColors(pieces []Piece, grid Grid) {
	if len(pieces) == 1 {
		return
	}

	var graph []Graph

	for _, x := range pieces {
		graph = append(graph, Graph{id: x.id, neighbours: nil, color: -1})
	}

	maxNeigh := 0
	maxNeighId := -1

	for _, x := range pieces {
		for _, y := range GetAdjacentPiecesId(x, grid) {
			graph[x.id].neighbours = append(graph[x.id].neighbours, &graph[y])
			if len(graph[x.id].neighbours) > maxNeigh {
				maxNeighId = x.id
			}
		}
	}

	usedColors := 1
	graph[maxNeighId].color = 0

	for i := 0; i < len(graph); i++ {
		if i != maxNeighId {
			colorCheck := make([]bool, usedColors)
			for _, x := range graph[i].neighbours {
				if x.color >= 0 {
					colorCheck[x.color] = true
				}
			}
			color := -1
			for i, x := range colorCheck {
				if !x {
					color = i
					break
				}
			}
			if color == -1 {
				color = usedColors
				usedColors++
			}
			graph[i].color = color
		}
	}

	for i := 0; i < len(pieces); i++ {
		pieces[i].color = graph[i].color
	}
}

// renvoie la liste des ids des pieces adjacentes à la pièce passée en paramètre
func GetAdjacentPiecesId(piece Piece, grid Grid) []int {
	var res []int

	for _, block := range piece.blocks {
		x := piece.anchor.x + block.x
		y := piece.anchor.y + block.y

		if x+1 < grid.size {
			i := grid.squares[x+1][y]
			if i >= 0 && i != piece.id && !ContainInt(res, i) {
				res = append(res, i)
			}
		}

		if y+1 < grid.size {
			i := grid.squares[x][y+1]
			if i >= 0 && i != piece.id && !ContainInt(res, i) {
				res = append(res, i)
			}
		}

		if x > 0 {
			i := grid.squares[x-1][y]
			if i >= 0 && i != piece.id && !ContainInt(res, i) {
				res = append(res, i)
			}
		}

		if y > 0 {
			i := grid.squares[x][y-1]
			if i >= 0 && i != piece.id && !ContainInt(res, i) {
				res = append(res, i)
			}
		}
	}

	return res
}

// renvoie true si y est présent dans x
func ContainInt(x []int, y int) bool {
	if x == nil {
		return false
	}

	for _, i := range x {
		if y == i {
			return true
		}
	}

	return false
}
