package solver

import "github.com/abdulrahmank/solver/tic_tac_toe/ttt"

type GameStatus int

const (
	LOSE           GameStatus = -1
	NEUTRAL        GameStatus = 0
	WIN            GameStatus = 1
	POTENTIAL_WIN  GameStatus = 2
	POTENTIAL_LOSE GameStatus = 3
)

type Analyser interface {
	GetCellWiseWinProbability(b ttt.Board, c ttt.BoardCharacter) map[ttt.Cell]GameStatus
}

type AnalyserImpl struct{}

func (a *AnalyserImpl) GetCellWiseWinProbability(b ttt.Board, c ttt.BoardCharacter) map[ttt.Cell]GameStatus {
	result := make(map[ttt.Cell]GameStatus)
	for _, row := range b.Cells {
		for _, cell := range row {
			result[*cell] = NEUTRAL
			result[*cell] = NEUTRAL
		}
	}

	emptyCells := b.GetEmptyCells()
	for i := 0; i < b.Rows; i++ {
		for j := 0; j < b.Cols; j++ {
			if !contains(emptyCells, b.Cells[i][j]) {
				continue
			}
			rowStatus := make(map[string]int)
			colStatus := make(map[string]int)
			diagonalStatus := make(map[string]int)

			for ti := 0; ti < b.Rows; ti++ {
				rowStatus[b.Cells[ti][j].Val] += 1
			}

			for ti := 0; ti < b.Cols; ti++ {
				colStatus[b.Cells[i][ti].Val] += 1
			}

			if i == j {
				for ti := 0; ti < b.Cols; ti++ {
					diagonalStatus[b.Cells[ti][ti].Val] += 1
				}
			}
			// 2 will only work for 3 X 3 board
			if rowStatus[string(c)]|colStatus[string(c)]|diagonalStatus[string(c)] >= 2 {
				result[*b.Cells[i][j]] = WIN
			} else {
				if rowStatus[string(c)]|colStatus[string(c)]|diagonalStatus[string(c)] >= 1 {
					result[*b.Cells[i][j]] = POTENTIAL_WIN
				}
				chars := []ttt.BoardCharacter{ttt.X, ttt.O}
				for _, ch := range chars {
					if ch != c {
						if rowStatus[string(ch)]|colStatus[string(ch)]|diagonalStatus[string(ch)] >= 2 {
							result[*b.Cells[i][j]] = LOSE
						} else if rowStatus[string(ch)]|colStatus[string(ch)]|diagonalStatus[string(ch)] >= 1 {
							result[*b.Cells[i][j]] = POTENTIAL_LOSE
						}
					}
				}
			}
		}
	}

	return result
}

func contains(cells []ttt.Cell, cell *ttt.Cell) bool {
	for _, c := range cells {
		if *cell == c {
			return true
		}
	}
	return false
}
