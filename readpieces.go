package main

import (
	"bufio"
	"os"
	"sort"
)

type Coord struct {
	x int
	y int
}

type Piece struct {
	id     int     // id de la pièce
	blocks []Coord // coordonnées relatives des éléments de la pièce. Commence toujours par (0,0)
	anchor Coord   // position de la première case (0,0) des blocks dans la grille. Contient -1 sur la pièce n'est pas placée
	color  int
}

func ReadPieces() []Piece {
	// Lecture du fichier des pièces
	file, err := os.Open(os.Args[1])
	if err != nil {
		Error()
	}

	var pieces []Piece
	pieceCount := 0

	fileScanner := bufio.NewScanner(file)

	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		var coords []Coord
		for i := 0; i < 4; i++ {
			// Lecture des quatre lignes de la pièce et de la ligne vide au bout
			line := fileScanner.Text()
			if len(line) != 4 {
				Error()
			}
			for j, x := range line {
				if x == '#' {
					coords = append(coords, Coord{x: j, y: 3 - i})
				} else if x != '.' {
					Error()
				}
			}
			fileScanner.Scan()
		}
		if IsValidPiece(coords) {
			SortCoords(coords)
			piece := Piece{blocks: coords, anchor: Coord{x: -1, y: 0}, id: pieceCount, color: -1}
			pieces = append(pieces, piece)
			pieceCount++
		} else {
			Error()
		}
		if fileScanner.Text() != "" {
			Error()
		}
	}
	return pieces
}

// Vérifie que les coordonnées récupérées constituent une pièce valide
func IsValidPiece(coords []Coord) bool {
	if len(coords) != 4 {
		return false
	}

	adjCount := 0
	// On compte le nombre de "contacts" entre blocs.
	// Une pièce valide contient au moins 3 contacts
	for i, a := range coords {
		for _, b := range coords[i:] {
			if IsAdj(a, b) {
				adjCount++
			}
		}
	}

	return adjCount > 2
}

// Renvoie true si les cases de coordonnées c1 et c2 sont adjacentes dans la grille
func IsAdj(c1 Coord, c2 Coord) bool {
	return IntAbs(c1.x-c2.x)+IntAbs(c1.y-c2.y) == 1
}

func IntAbs(x int) int {
	// Valeur absolue d'entier
	if x > 0 {
		return x
	} else {
		return -x
	}
}

func SortCoords(coords []Coord) {
	// Réarrange les coordonnées des blocs la pièce.
	// Commence toujours par (0,0)
	// Les blocs sont rangés en priorité de gauche à droite, puis de bas en haut
	sort.Slice(coords, func(i, j int) bool {
		if coords[i].x == coords[j].x {
			return coords[i].y < coords[j].y
		} else {
			return coords[i].x < coords[j].x
		}
	})
	for i := 3; i >= 0; i-- {
		coords[i].x -= coords[0].x
		coords[i].y -= coords[0].y
	}
}
