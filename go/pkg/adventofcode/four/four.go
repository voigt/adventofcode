package four

import (
	"log"
	"strconv"
	"strings"
)

type Board struct {
	State  [][]int
	Winner bool
	Place  int
	WinNum int
}

type Game struct {
	Boards   []Board
	Sequence []int
	WinCount int
}

func PartOne(inputs []string) int {
	// b1 := Board{
	// 	State: [][]int{
	// 		{22, 13, 17, 11, 0},
	// 		{8, 2, 23, 4, 24},
	// 		{21, 9, 14, 16, 7},
	// 		{6, 10, 3, 18, 5},
	// 		{1, 12, 20, 15, 19},
	// 	},
	// 	Winner: false,
	// }
	// b1.HasColWin()
	// b2 := Board{
	// 	State: [][]int{
	// 		{3, 15, 0, 2, 22},
	// 		{9, 18, 13, 17, 5},
	// 		{19, 8, 7, 25, 23},
	// 		{20, 11, 10, 24, 4},
	// 		{14, 21, 16, 12, 6},
	// 	},
	// 	Winner: false,
	// }
	// b3 := Board{
	// 	State: [][]int{
	// 		{14, 21, 17, 24, 4},
	// 		{10, 16, 15, 9, 19},
	// 		{18, 8, 23, 26, 20},
	// 		{22, 11, 13, 6, 5},
	// 		{2, 0, 12, 3, 7},
	// 	},
	// 	Winner: false,
	// }

	// g := Game{
	// 	Boards:   []Board{b1, b2, b3},
	// 	Sequence: []int{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24, 10, 16, 13, 6, 15, 25, 12, 22, 18, 20, 8, 19, 3, 26, 1},
	// }

	g := Game{
		Boards:   []Board{},
		Sequence: []int{},
	}

	for _, s := range strings.Split(inputs[0], ",") {
		num, _ := strconv.Atoi(s)
		g.Sequence = append(g.Sequence, num)
	}

	var board = Board{State: [][]int{}, Winner: false}
	for i, l := range inputs {
		if i < 2 {
			continue
		}
		if l == "" {
			g.Boards = append(g.Boards, board)
			board = Board{State: [][]int{}, Winner: false}
			continue
		}

		nums := strings.Fields(l)
		n := make([]int, 0)
		for _, number := range nums {
			i, err := strconv.Atoi(number)
			if err != nil {
				log.Fatal(err)
			}
			n = append(n, i)
		}
		board.State = append(board.State, n)
	}

	w := Board{}
	// lastSq := 0
	for _, n := range g.Sequence {
		g.MarkBoards(n)
		_, w = g.GetWinner(n)
		// lastSq = n
		if w.Winner {
			break
		}
	}

	return calculateScore(w, w.WinNum)
}

func PartTwo(inputs []string) int {
	g := Game{
		Boards:   []Board{},
		Sequence: []int{},
	}

	for _, s := range strings.Split(inputs[0], ",") {
		num, _ := strconv.Atoi(s)
		g.Sequence = append(g.Sequence, num)
	}

	var board = Board{State: [][]int{}, Winner: false}
	for i, l := range inputs {
		if i < 2 {
			continue
		}
		if l == "" {
			g.Boards = append(g.Boards, board)
			board = Board{State: [][]int{}, Winner: false}
			continue
		}

		nums := strings.Fields(l)
		n := make([]int, 0)
		for _, number := range nums {
			i, err := strconv.Atoi(number)
			if err != nil {
				log.Fatal(err)
			}
			n = append(n, i)
		}
		board.State = append(board.State, n)
	}

	var winner Board
	sq := 0
	for len(g.Boards) > 0 {
		for _, n := range g.Sequence {
			g.MarkBoards(n)
			index, win := g.GetWinner(n)
			if win.Winner {
				// fmt.Printf("winner score: number: %d, score: %d, loc: %d\n", n, calculateScore(win, n), index)
				sq = n
				winner = win
				g.Boards = append(g.Boards[:index], g.Boards[index+1:]...)
				break
			}
		}
	}

	// for _, b := range g.Boards {
	// 	fmt.Printf("%v\n", b)
	// }

	// w := g.GetLastWinner()
	return calculateScore(winner, sq)
}

func calculateScore(b Board, i int) int {
	sumUnmarked := 0

	// fmt.Printf("%v\n", b.State)

	for _, row := range b.State {
		for _, cell := range row {
			if cell != -1 {
				sumUnmarked += cell
			}
		}
	}

	// fmt.Printf("sumUnmarked: %d, lastSQ: %d\n", sumUnmarked, i)

	return sumUnmarked * i
}

func (g *Game) MarkBoards(n int) {
	for _, b := range g.Boards {
		b.MarkBoard(n)
	}
}

func (b *Board) MarkBoard(n int) {
	for i, row := range b.State {
		for j, cell := range row {
			if cell == n {
				b.State[i][j] = -1
			}
		}
	}
}

func (b *Board) HasRowWin() bool {

	for _, row := range b.State {

		isRowWin := true

		for _, cell := range row {
			if cell != -1 {
				isRowWin = false
				break
			}
		}

		if isRowWin {
			return true
		}
	}
	return false
}

func (b *Board) HasColWin() bool {

	for col := 0; col < len(b.State[0]); col++ {
		isColWin := true
		for j := 0; j < len(b.State); j++ {
			if b.State[j][col] != -1 {
				isColWin = false
				break
			}
		}
		if isColWin {
			return true
		}
	}

	return false
}

func (g *Game) GetWinner(n int) (int, Board) {
	for index, b := range g.Boards {
		b.Winner = false
		if b.HasRowWin() {
			b.Winner = true
			g.WinCount++
			b.Place = g.WinCount
			b.WinNum = n
			return index, b
		}

		if b.HasColWin() {
			b.Winner = true
			g.WinCount++
			b.Place = g.WinCount
			b.WinNum = n
			return index, b
		}
	}

	return 0, Board{}
}

func (g *Game) GetLastWinner() Board {
	lastWinner := Board{Place: 0}
	for _, b := range g.Boards {
		if b.Place > lastWinner.Place {
			lastWinner = b
		}
	}
	return lastWinner
}
