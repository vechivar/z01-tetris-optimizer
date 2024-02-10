package main

import (
	"fmt"
	"math"
	"os"
)

type Grid struct {
	// Grille de taille size à remplir.
	// Contient les id des pièces occupant les cases, contient -1 si la case est vide
	size    int
	squares [][]int
}

func BuildSolution(pieces []Piece) {
	// Trouve une solutione au problème

	// Récupération de la taille minimale de la grille à remplir
	size64 := math.Sqrt(float64(4 * len(pieces)))
	size := int(size64)

	if math.Trunc(size64) != size64 {
		size++
	}

	for {
		// On augmente la taille de la grille jusqu'à trouver une solution
		grid := InitGrid(size)
		for i := 0; i < len(pieces); i++ {
			pieces[i].anchor.x = -1
		}
		if InitiateFill(pieces, grid) {
			fmt.Println("Result :")
			BuildColors(pieces, grid)
			PrintGrid(grid, pieces)
			os.Exit(0)
		} else {
			size++
		}
	}
}

func InitiateFill(pieces []Piece, grid Grid) bool {
	// Place la première pièce en bas à gauche de la grille, puis tente de trouver une solution.
	// Si pas de solution, on essaye avec une autre pièce
	for i := 0; i < len(pieces); i++ {
		pieces[i].anchor.x = 0
		pieces[i].anchor.y = 0
		for !ValidPlace(pieces[i], grid) {
			pieces[i].anchor.y++
			if pieces[i].anchor.y >= grid.size {
				return false
			}
		}
		PlaceOnGrid(pieces[i], grid)
		if FillSquare(pieces, grid) {
			return true
		} else {
			RemoveFromGrid(pieces[i], grid)
			pieces[i].anchor.x = -1
		}
	}
	return false
}

func FillSquare(pieces []Piece, grid Grid) bool {
	// Tente de trouver une solution de manière récursive selon le principe suivant :
	// - on tente de placer une pièce non-utilisée.
	// - si aucune place ne convient, on renvoie false
	// - sinon, on tente de placer une nouvelle pièce. Si cette pièce permet de trouver une solution, on renvoie true
	// ie : renvoie true si le placement actuel des pièces permet de trouver une solution

	// On vérifie si toutes les pièces ont été placées.
	// Si c'est le cas, la solution est trouvée.
	flag := true
	for _, x := range pieces {
		flag = flag && x.anchor.x >= 0
	}
	if flag {
		return true
	}

	// On tente de placer successivement toutes les pièces non utilisées
	for i := 0; i < len(pieces); i++ {
		if pieces[i].anchor.x < 0 {
			// On cherche à placer la pièce le plus en haut à gauche possible
			// On évite ainsi de laisser inutilement des trous
			x := 0
			y := 0
			flag = true
			for flag {
				if grid.squares[x][y] == -1 {
					pieces[i].anchor.x = x
					pieces[i].anchor.y = y
					if ValidPlace(pieces[i], grid) {
						// On a trouvé une place valide pour la pièce. On tente de trouver une solution avec ce placement
						flag = false
						PlaceOnGrid(pieces[i], grid)
						if FillSquare(pieces, grid) {
							return true
						}
						RemoveFromGrid(pieces[i], grid)
						pieces[i].anchor.x = -1
					}
				}
				y++
				if y == grid.size {
					y = 0
					x++
					if x == grid.size {
						flag = false
					}
				}
			}
			pieces[i].anchor.x = -1
		}
	}
	// Aucune pièce n'a permis de trouver une solution. On renvoie false.
	return false
}

func ValidPlace(piece Piece, grid Grid) bool {
	// Renvoie true si la pièce peut être placée dans la grille à partir de son champ anchor
	for i := 0; i < 4; i++ {
		x := piece.anchor.x + piece.blocks[i].x
		y := piece.anchor.y + piece.blocks[i].y
		if x < 0 || y < 0 || x >= grid.size || y >= grid.size || grid.squares[x][y] >= 0 {
			return false
		}
	}
	return true
}

func PlaceOnGrid(piece Piece, grid Grid) {
	// Place la pièce sur la grille à partir du champ anchor de la pièce.
	anchor := piece.anchor

	for i := 0; i < 4; i++ {
		block := piece.blocks[i]
		if grid.squares[anchor.x+block.x][anchor.y+block.y] >= 0 {
			fmt.Printf("problem placing piece %v\n", piece.id)
			os.Exit(0)
		}
		grid.squares[anchor.x+block.x][anchor.y+block.y] = piece.id
	}
}

func RemoveFromGrid(piece Piece, grid Grid) {
	// Retire la pièce de la grille.
	anchor := piece.anchor

	for i := 0; i < 4; i++ {
		block := piece.blocks[i]
		if grid.squares[anchor.x+block.x][anchor.y+block.y] != piece.id {
			fmt.Printf("problem removing piece %v\n", piece.id)
			os.Exit(0)
		}
		grid.squares[anchor.x+block.x][anchor.y+block.y] = -1
	}
}

func InitGrid(size int) Grid {
	// Initialisation d'une grille vide.
	var square [][]int
	for i := 0; i < size; i++ {
		var line []int
		for j := 0; j < size; j++ {
			line = append(line, -1)
		}
		square = append(square, line)
	}

	return Grid{squares: square, size: size}
}

func PrintGrid(grid Grid, pieces []Piece) {
	// Affichage de la grille
	for j := grid.size - 1; j >= 0; j-- {
		for i := 0; i < grid.size; i++ {
			if grid.squares[i][j] == -1 {
				fmt.Print(".")
			} else {
				fmt.Print(GetPieceString(pieces[grid.squares[i][j]]))
			}
		}
		fmt.Print("\n")
	}
}

func GetPieceString(piece Piece) string {
	// Donne un charactère à afficher à la pièce

	color := ""
	switch piece.color % 4 {
	case 0:
		color = "\033[31m"
	case 1:
		color = "\033[32m"
	case 2:
		color = "\033[34m"
	case 3:
		color = "\033[97m"
	}

	return color + string('A'+rune(piece.id)) + "\033[0m"
}
