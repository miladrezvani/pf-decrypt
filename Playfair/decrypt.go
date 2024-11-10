package playfairCipher

func (p *PlayFair) Decrypt() (result string) {
	genMtx := NewRows(p.Matrix)
	encword := WordToPairs(p.decrypt)
	for _, n := range encword {
		fstRI, sndRI, index1, index2 := Find(n, p.Matrix)
		DecryptPair(fstRI, sndRI, index1, index2, genMtx, p.Matrix, &result)
		Verify(&result)
	}
return
}

func DecryptPair(i, j, x, y int, mRows matrixRows, matrix [][]string, result *string) {
	if i == j && x != y {
		mRows.shiftToLeft(i, x, y, result)
	} else if i != j && x == y {
		mRows.shiftToUp(i, j, x, result)
	} else if i != j && x != y {
		mRows.getIntersection(matrix, i, j, x, y, result)
	}
}

func NewRows(matrix [][]string) matrixRows {
	var mRows matrixRows
	mtx := new(matrixR)
	for i, r := range matrix {
		m := make(map[int][]string)
		m[i] = r
		mtx.row = m
		mRows = append(mRows, *mtx)
	}
	return mRows
}

func (m matrixRows) shiftToLeft(rowIndex, ind1, ind2 int, result *string) {
	if ind1 == 0 {
		ind1 = LENMATRIX - 1
	} else {
		ind1--
	}
	if ind2 == 0 {
		ind2 = LENMATRIX - 1
	} else {
		ind2--
	}
	*result += m[rowIndex].row[rowIndex][ind1]
	*result += m[rowIndex].row[rowIndex][ind2]
}

func (m matrixRows) shiftToUp(fstRowIndex, sndRowIndex, index int, result *string) {
	if fstRowIndex == 0 {
		fstRowIndex = LENMATRIX - 1
	} else {
		fstRowIndex--
	}
	if sndRowIndex == 0 {
		sndRowIndex = LENMATRIX - 1
	} else {
		sndRowIndex--
	}
	*result += m[fstRowIndex].row[fstRowIndex][index]
	*result += m[sndRowIndex].row[sndRowIndex][index]
}

func (m matrixRows) getIntersection(matrix [][]string, fstRowIndex, sndRowIndex, indx1, indx2 int, result *string) {
	if len(endSpecialWords) == 0 {
		*result += m[fstRowIndex].row[fstRowIndex][indx2]
		*result += m[sndRowIndex].row[sndRowIndex][indx1]
	} else if len(endSpecialWords) == 2 {
		specRI1, specI1 := LetterIndex(endSpecialWords[0], matrix)
		specRI2, specI2 := LetterIndex(endSpecialWords[1], matrix)
		if specRI1 == fstRowIndex && specI1 == indx1 && specRI2 == sndRowIndex && specI2 == indx2 {
			*result += m[sndRowIndex].row[sndRowIndex][indx1]
			*result += m[fstRowIndex].row[fstRowIndex][indx2]
			endSpecialWords = endSpecialWords[:0]
		} else {
			*result += m[fstRowIndex].row[fstRowIndex][indx2]
			*result += m[sndRowIndex].row[sndRowIndex][indx1]
		}
	}
}
