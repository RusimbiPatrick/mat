package mat

import (
    "testing"
    "reflect"
)

func TestAdd(t *testing.T) {
    m1 := CreateMatrix(2, 2)
    m2 := CreateMatrix(2, 2)
    m1.data = [][]float64{{1, 2}, {3, 4}}
    m2.data = [][]float64{{5, 6}, {7, 8}}

    expected := CreateMatrix(2, 2)
    expected.data = [][]float64{{6, 8}, {10, 12}}

    result, err := m1.Add(m2)
    if err != nil {
        t.Errorf("Add() error = %v", err)
        return
    }

    if !reflect.DeepEqual(result.data, expected.data) {
        t.Errorf("Add() = %v, want %v", result.data, expected.data)
    }
}

func TestMultiply(t *testing.T) {
    m1 := CreateMatrix(2, 2)
    m2 := CreateMatrix(2, 2)
    m1.data = [][]float64{{1, 2}, {3, 4}}
    m2.data = [][]float64{{5, 6}, {7, 8}}

    expected := CreateMatrix(2, 2)
    expected.data = [][]float64{{19, 22}, {43, 50}}

    result, err := m1.Multiply(m2)
    if err != nil {
        t.Errorf("Multiply() error = %v", err)
        return
    }

    if !reflect.DeepEqual(result.data, expected.data) {
        t.Errorf("Multiply() = %v, want %v", result.data, expected.data)
    }
}
