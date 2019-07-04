package score

import (
	"fmt"
	"math"
)

func printPointer(s1, s2 []rune, mat [][]int) {
	fmt.Printf("%2s  |", " ")
	fmt.Printf("%5s", " ")
	for _, r := range s1 {
		fmt.Printf("%4c ", r)
	}
	fmt.Println()

	fmt.Print("----+")
	for i := 0; i < len(s1)*5+5; i++ {
		fmt.Print("-")
	}

	fmt.Println()
	for i, col := range mat {
		if i == 0 {
			fmt.Printf("%2s  |", " ")
		} else {
			fmt.Printf("%2c  |", s2[i-1])
		}

		for _, val := range col {
			var token string
			switch val {
			case 0:
				token = "DHV"
			case 1:
				token = "DH"
			case 2:
				token = "DV"
			case 3:
				token = "HV"
			case 4:
				token = "D"
			case 5:
				token = "H"
			case 6:
				token = "V"
			default:
				token = "0"
			}
			fmt.Printf("%4s ", token)
		}
		fmt.Println()
	}
	fmt.Println()
}

func printIntMat(s1, s2 []rune, mat [][]int) {
	fmt.Printf("%2s  |", " ")
	fmt.Printf("%5s", " ")
	for _, r := range s1 {
		fmt.Printf("%4c ", r)
	}
	fmt.Println()

	fmt.Print("----+")
	for i := 0; i < len(s1)*5+5; i++ {
		fmt.Print("-")
	}

	fmt.Println()
	for i, col := range mat {
		if i == 0 {
			fmt.Printf("%2s  |", " ")
		} else {
			fmt.Printf("%2c  |", s2[i-1])
		}

		for _, val := range col {
			fmt.Printf("%4d ", val)
		}
		fmt.Println()
	}
	fmt.Println()
}

func initNedlemanWunsch(row, column int) [][]int {
	mat := make([][]int, row)
	for i := 0; i < row; i++ {
		mat[i] = make([]int, column)
	}

	for i := 0; i < row; i++ {
		mat[i][0] = i * -2
	}

	for j := 0; j < column; j++ {
		mat[0][j] = j * -2
	}
	return mat
}

func diagonal(n1, n2 rune) int {
	if n1 == n2 {
		return 2
	}
	return -1
}

const (
	DHV = iota
	DH
	DV
	HV
	D
	H
	V
	NONE
)

func pointers(di, ho, ve int) int {
	pointer := max(di, ho, ve)
	switch {
	case di == pointer && ho == pointer && ve == pointer:
		return DHV
	case di == pointer && ho == pointer:
		return DH
	case di == pointer && ve == pointer:
		return DV
	case ho == pointer && ve == pointer:
		return HV
	case di == pointer:
		return D
	case ho == pointer:
		return H
	}
	return V
}

func max(x ...int) int {
	max := math.MinInt64
	for _, v := range x {
		if v > max {
			max = v
		}
	}
	return max
}

func initMatrix(n, m int) ([][]int, [][]int) {
	mat := make([][]int, m)
	cmat := make([][]int, m)
	for i := 0; i < m; i++ {
		mat[i] = make([]int, n)
		cmat[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		mat[i][0] = i * -2
		cmat[i][0] = V
	}

	for j := 0; j < n; j++ {
		mat[0][j] = j * -2
		cmat[0][j] = H
	}
	return mat, cmat
}

func drawResult(s1, s2 []rune, cmat [][]int) {
	n := len(s1) + 1
	m := len(s2) + 1
	x := m - 1
	y := n - 1

	if n > m {
		long := make([]rune, n)
		mid := make([]rune, n)
		short := make([]rune, n)
		pos := n - 2
		for x != 0 || y != 0 {
			switch cmat[x][y] {
			case D:
				x--
				y--
				long[pos] = s1[pos]
				mid[pos] = '|'
				short[pos] = s1[pos]
			case DHV, DH, HV, H:
				y--
				long[pos] = s1[pos]
				mid[pos] = ' '
				short[pos] = '-'
			case DV, V:
				x--
				long[pos] = s1[pos]
				mid[pos] = ' '
				short[pos] = '-'
			}
			pos--
		}
		fmt.Printf("%s\n%s\n%s\n", string(long), string(mid), string(short))
		fmt.Println()
		return
	}
	long := make([]rune, m)
	mid := make([]rune, m)
	short := make([]rune, m)
	pos := m - 2
	for x != 0 || y != 0 {
		switch cmat[x][y] {
		case D:
			x--
			y--
			long[pos] = s2[pos]
			mid[pos] = '|'
			short[pos] = s2[pos]
		case DH, H:
			y--
			long[pos] = s2[pos]
			mid[pos] = ' '
			short[pos] = '-'
		case DHV, HV, DV, V:
			x--
			long[pos] = s2[pos]
			mid[pos] = ' '
			short[pos] = '-'
		}
		pos--
	}

	fmt.Printf("%s\n%s\n%s\n", string(long), string(mid), string(short))
	fmt.Println()
}

func NeedlemanWunsch(str1, str2 string, draw bool) int {
	s1 := []rune(str1)
	s2 := []rune(str2)
	n := len(s1) + 1
	m := len(s2) + 1
	mat, cmat := initMatrix(n, m)
	mat[0][0] = 0
	cmat[0][0] = NONE
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			di := mat[i-1][j-1] + diagonal(s1[j-1], s2[i-1])
			ho := mat[i][j-1] - 2
			ve := mat[i-1][j] - 2
			mat[i][j] = max(di, ho, ve)
			cmat[i][j] = pointers(di, ho, ve)
		}
	}
	if draw {
		printIntMat(s1, s2, mat)
		printPointer(s1, s2, cmat)
		drawResult(s1, s2, cmat)
	}
	return mat[m-1][n-1]
}