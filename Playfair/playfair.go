package playfairCipher

import (
	"strings"
)

const (
	LENMATRIX int = 5
	ALPHABET string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	PLACEHOLDER = "X"
)

var (
	endSpecialWords []string
)

type PlayFair struct {
	Keyword string
	Matrix  [][]string
	decrypt string
}

type matrixR struct {
	row map[int][]string
}

type matrixRows []matrixR

func MatrixCheckLetters(keyword string) map[string]bool {
	m := make(map[string]bool)
	for _, k := range keyword {
		m[string(k)] = true
	}
	if !m["J"] {
		m["J"] = true
	}
	return m
}

func MatrixEmptySpace(matrix [][]string) (arr []string, index int) {
	for i, j := range matrix {
		if len(j) != LENMATRIX {
			arr, index = j, i
		}
	}
	return
}

func FillMatrix(s *[]string, matrix *[][]string, letterIndex int, letter string) {
	if len(*s) == LENMATRIX {
		(*matrix) = append((*matrix), (*s))
		*s = []string{}
	}
}

func FillBlankSpace(index int, check bool, keyword string, matrix *[][]string) {
	m := MatrixCheckLetters(keyword)
	s := []string{}
	passed := make(map[string]bool)
	for letterIndex, letter := range ALPHABET {
		if !m[string(letter)] {
			if check {
				if len((*matrix)[index]) != LENMATRIX {
					passed[string(letter)] = true
					(*matrix)[index] = append((*matrix)[index], string(letter))
				}
				if !passed[string(letter)] {
					s = append(s, string(letter))
				}
				FillMatrix(&s, matrix, letterIndex, string(letter))
			} else {
				s = append(s, string(letter))
				FillMatrix(&s, matrix, letterIndex, string(letter))
			}
		}
	}
}

func CheckKeyword(keyword string, matrix *[][]string) {
	tmp := []string{}
	passedL := make(map[string]bool)
	for i, n := range keyword {
		if !passedL[string(n)] {
			tmp = append(tmp, string(n))
		}
		passedL[string(n)] = true
		if len(tmp) == LENMATRIX {
			(*matrix) = append((*matrix), tmp)
			tmp = []string{}
		} else if len(tmp) != LENMATRIX && i == len(keyword)-1 {
			(*matrix) = append((*matrix), tmp)
		}
	}
}

func Replicated(str *[]string, fL *string) {
	if len(*str) == 2 {
		cmp := string((*str)[0])
		if string((*str)[1]) == cmp {
			tmp := []string{}
			tmp = append(tmp, cmp)
			ns := strings.Replace(string((*str)[1]), string((*str)[1]), PLACEHOLDER, 1)
			tmp = append(tmp, ns)
			*str = tmp
			*fL = string((*str)[0])
		}
	}
}

func WordToPairs(word string) (wList [][]string) {
	var str []string
	var fL string
	for i, w := range word {
		if fL != "" {
			str = append(str, fL, string(w))
			fL = ""
		} else {
			str = append(str, string(w))
		}
		if len(str) == 2 {
			Replicated(&str, &fL)
			wList = append(wList, str)
			str = []string{}
		} else if i == len(word)-1 && len(str) != 2 {
			tmp := str[0]
			str[0] = PLACEHOLDER
			str = append(str, tmp)
			wList = append(wList, str)
		}
		if i == len(word)-1 && wList[len(wList)-1][0] == fL && wList[len(wList)-1][1] == PLACEHOLDER {
			endSpecialWords = append(endSpecialWords, PLACEHOLDER, fL)
			wList = append(wList, []string{PLACEHOLDER, fL})
		}
	}
	return
}

func IsIn(w1, w2 string, list []string) (indx1, indx2 int, found1, found2 bool) {
	for i, j := range list {
		if w1 == j {
			indx1 = i
			found1 = true
		} else if w2 == j {
			indx2 = i
			found2 = true
		}
	}
	return
}

func LetterIndex(letter string, matrix [][]string) (rowIndex, index int) {
	for i, r := range matrix {
		for j, l := range r {
			if l == letter {
				rowIndex = i
				index = j
			}
		}
	}
	return
}

func Find(pair []string, matrix [][]string) (fstRowIndex, sndRowIndex int, indx1, indx2 int) {
	for b, w := range matrix {
		index1, index2, found1, found2 := IsIn(pair[0], pair[1], w)
		if found1 {
			indx1 = index1
			fstRowIndex = b
		} else if !found1 && pair[0] == "J" {
			x, y := LetterIndex("I", matrix)
			indx1 = y
			fstRowIndex = x
		}
		if found2 {
			indx2 = index2
			sndRowIndex = b
		} else if !found2 && pair[1] == "J" {
			x, y := LetterIndex("I", matrix)
			indx2 = y
			sndRowIndex = x
		}
	}
	return
}