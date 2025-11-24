package tests

import (
	"talenthouse/go-api/services"
	"testing"
)

// Función auxiliar para multiplicar matrices: C = A * B
func multiplyMatrices(A [][]float64, B [][]float64) [][]float64 {
	m := len(A)
	n := len(B[0])
	p := len(B)
	C := make([][]float64, m)
	for i := range C {
		C[i] = make([]float64, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			sum := 0.0
			for k := 0; k < p; k++ {
				sum += A[i][k] * B[k][j]
			}
			C[i][j] = sum
		}
	}
	return C
}

// Función auxiliar para comparar matrices con tolerancia
func matricesAlmostEqual(A, B [][]float64, tol float64) bool {
	if len(A) != len(B) {
		return false
	}
	for i := range A {
		if len(A[i]) != len(B[i]) {
			return false
		}
		for j := range A[i] {
			diff := A[i][j] - B[i][j]
			if diff < 0 {
				diff = -diff
			}
			if diff > tol {
				return false
			}
		}
	}
	return true
}

func TestQRFactorization(t *testing.T) {
	tol := 1e-9

	tests := []struct {
		name    string
		input   [][]float64
		wantErr bool
	}{
		{
			name: "matriz 2x3",
			input: [][]float64{
				{1, 2, 3},
				{4, 5, 6},
			},
			wantErr: false,
		},
		{
			name: "matriz cuadrada 3x3",
			input: [][]float64{
				{1, 2, 3},
				{0, 1, 4},
				{5, 6, 0},
			},
			wantErr: false,
		},
		{
			name:    "matriz vacía",
			input:   [][]float64{},
			wantErr: true,
		},
		{
			name: "columna linealmente dependiente",
			input: [][]float64{
				{1, 2},
				{2, 4},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Q, R, err := services.QRFactorization(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("QRFactorization() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				// Verificar que Q*R ≈ A
				Arec := multiplyMatrices(Q, R)
				if !matricesAlmostEqual(Arec, tt.input, tol) {
					t.Errorf("QRFactorization() Q*R != A\nQ*R = %v\nA = %v", Arec, tt.input)
				}
			}
		})
	}
}
