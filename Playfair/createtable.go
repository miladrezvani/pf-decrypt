package playfairCipher

import "strings"

func NewMatrix(keyword string, word string) (pf *PlayFair) {
	mtx := [][]string{}
	pf = &PlayFair{keyword, mtx, strings.ToUpper(word)}
	return
}

func (p *PlayFair) CreateTable() {
	var check bool
	CheckKeyword(strings.ToUpper(p.Keyword), &p.Matrix)
	arr, index := MatrixEmptySpace(p.Matrix)
	if len(arr) != LENMATRIX {
		check = true
	} else {
		check = false
	}
	FillBlankSpace(index, check, strings.ToUpper(p.Keyword), &p.Matrix)
}
