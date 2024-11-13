package calculator

import "testing"

func TestAdd(t *testing.T) {
    result := Add(5, 3)
    if result != 8 {
        t.Errorf("Add(5, 3) = %d; want 8", result)
    }
}

func TestSubtract(t *testing.T) {
    result := Subtract(5, 3)
    if result != 2 {
        t.Errorf("Subtract(5, 3) = %d; want 2", result)
    }
}
