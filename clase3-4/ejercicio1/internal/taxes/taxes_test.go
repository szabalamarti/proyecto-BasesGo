package taxes

import (
	"testing"
)

func TestGetTax(t *testing.T) {
	t.Run("case 01: salary under 50000", func(t *testing.T) {
		// Arrange
		salary := 49000.0
		expected := 0.0

		// Act
		result := GetTax(salary)

		// Assert
		if result != expected {
			t.Errorf("Expected %f, got %f", expected, result)
			return
		}
		t.Log("success")
	})
	t.Run("case 02: salary over 50000", func(t *testing.T) {
		// Arrange
		salary := 100000.0
		expected := 17000.0

		// Act
		result := GetTax(salary)

		// Assert
		if result != expected {
			t.Errorf("Expected %f, got %f", expected, result)
			return
		}
		t.Log("success")
	})
	t.Run("case 03: salary over 150000", func(t *testing.T) {
		// Arrange
		salary := 151000.0
		expected := 40770.0

		// Act
		result := GetTax(salary)

		// Assert
		if result != expected {
			t.Errorf("Expected %f, got %f", expected, result)
			return
		}
		t.Log("success")
	})
}
