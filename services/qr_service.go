package services

import (
	"errors"
	"talenthouse/go-api/utils"
)

func QRFactorization(A [][]float64) ([][]float64, [][]float64, error) {
	if len(A) == 0 || len(A[0]) == 0 {
		return nil, nil, errors.New("matriz vacía")
	}

	m := len(A)
	n := len(A[0])

	// Inicializar matrices Q y R
	Q := make([][]float64, m)
	for i := range Q {
		Q[i] = make([]float64, n)
	}

	R := make([][]float64, n)
	for i := range R {
		R[i] = make([]float64, n)
	}

	for j := 0; j < n; j++ {
		// Copiar columna j de A en v
		v := make([]float64, m)
		for i := 0; i < m; i++ {
			v[i] = A[i][j]
		}

		// Proceso de Gram-Schmidt
		for i := 0; i < j; i++ {
			R[i][j] = 0
			for k := 0; k < m; k++ {
				R[i][j] += Q[k][i] * A[k][j]
			}
			for k := 0; k < m; k++ {
				v[k] -= R[i][j] * Q[k][i]
			}
		}

		// Norma de v
		R[j][j] = utils.Sqrt(v)
		if R[j][j] != 0 {
			for i := 0; i < m; i++ {
				Q[i][j] = v[i] / R[j][j]
			}
		} else {
			// Si la norma es cero, columna linealmente dependiente
			for i := 0; i < m; i++ {
				Q[i][j] = 0
			}
		}
	}

	return Q, R, nil
}
