package Transfer

import "witchmat/Types"

type Lab1ReturnDTO struct {
	Status                string
	Msg                   string
	Triangle              Types.Matrix
	Determinant           float64
	DeterminantCalculated bool
	X                     []float64
	Discrepancy           []float64
	SolutionsType         string
}
type Lab1GetDTO map[string]Types.Matrix
