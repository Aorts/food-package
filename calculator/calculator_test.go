package calculator_test

import (
	"testing"

	"github.com/Aorts/food-package/calculator"
)

func TestOrderFood(t *testing.T) {
	tests := []struct {
		name      string
		orderName string
		quantity  int
		wantErr   bool
	}{
		{"Valid Order", "red", 2, false},
		{"Invalid Order", "invalid", 2, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := calculator.OrderFood(tt.orderName, tt.quantity)
			if (err != nil) != tt.wantErr {
				t.Errorf("OrderFood() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCheckOut(t *testing.T) {
	tests := []struct {
		name     string
		isMember bool
		want     float64
	}{
		{"Non-Member", false, 0},
		{"Member", true, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := calculator.CheckOut(tt.isMember)
			if got != tt.want {
				t.Errorf("CheckOut() = %v, want %v", got, tt.want)
			}
		})
	}
}
