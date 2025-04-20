package Services

import (
	"math"
	"witchmat/Transfer"
)

func SolveLab1(inp Transfer.Lab1GetDTO) Transfer.Lab1ReturnDTO {
	matrixInp := inp["matrix"]
	if matrixInp == nil {
		return Transfer.Lab1ReturnDTO{Status: "fail", Msg: "Request corrupted"}
	}
	xSize := len(matrixInp[0])
	ySize := len(matrixInp)
	matrix := matrixInp
	varind := 0
	var determinant float64 = 1
	var determinantCnt = 0
	for i := range ySize {
		if matrix[i][varind] == 0 {
			found := false
			for !found {
				if varind == xSize-1 {
					break
				}
				for i2 := i; i2 < ySize; i2++ {
					if matrix[i2][varind] != 0 {
						matrix[i2], matrix[i] = matrix[i], matrix[i2]
						found = true
						break
					}
				}
				if !found {
					varind++
				}
			}
		}
		if varind == xSize-1 {
			break
		}
		for j := varind + 1; j < xSize; j++ {
			matrix[i][j] /= matrix[i][varind]
		}
		determinant *= matrix[i][varind]
		determinantCnt++
		matrix[i][varind] = 1
		for i2 := i + 1; i2 < ySize; i2++ {
			if matrix[i2][varind] != 0 {
				for j := varind + 1; j < xSize; j++ {
					matrix[i2][j] -= matrix[i][j] * matrix[i2][varind]
				}
				matrix[i2][varind] = 0
			}
		}
		varind++
	}
	var determinantCalculated bool = false
	if xSize == ySize+1 {
		determinantCalculated = true
		if determinantCnt != ySize {
			determinant = 0
		}
	}

	var res = make([]float64, xSize-1)
	var resFound = make([]bool, xSize-1)
	for i := ySize - 1; i >= 0; i-- {
		var curVar = -1
		var curRes = matrix[i][xSize-1]
		for j := 0; j < xSize-1; j++ {
			if matrix[i][j] != 0 {
				if resFound[j] {
					curRes -= res[j] * matrix[i][j]
				} else {
					if curVar == -1 {
						curVar = j
					} else {
						return Transfer.Lab1ReturnDTO{Status: "ok", Triangle: matrix, Determinant: determinant, DeterminantCalculated: determinantCalculated, SolutionsType: "infinity"}
					}
				}
			}
		}
		if curVar == -1 {
			if matrix[i][xSize-1] == 0 {
				continue
			} else {
				return Transfer.Lab1ReturnDTO{Status: "ok", Triangle: matrix, Determinant: determinant, DeterminantCalculated: determinantCalculated, SolutionsType: "none"}
			}
		}
		resFound[curVar] = true
		res[curVar] = curRes / matrix[i][curVar]
	}
	var discrepancy = make([]float64, ySize)
	for i := range ySize {
		lSide := 0.0
		for j := range xSize - 1 {
			lSide += matrix[i][j] * res[j]
		}
		rSide := matrix[i][xSize-1]
		discrepancy[i] = math.Abs(lSide - rSide)
	}
	return Transfer.Lab1ReturnDTO{Status: "ok", Triangle: matrix, Determinant: determinant, DeterminantCalculated: determinantCalculated, X: res, SolutionsType: "one", Discrepancy: discrepancy}
}
