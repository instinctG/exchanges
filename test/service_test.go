package test

import (
	"github.com/instinctG/exchanges/internal/service"
	"reflect"
	"testing"
)

func TestExchangeCombinations(t *testing.T) {
	tests := []struct {
		amount    int
		banknotes []int
		expected  [][]int
	}{
		{
			amount:    400,
			banknotes: []int{5000, 2000, 1000, 200, 100, 50},
			expected: [][]int{
				{200, 200},
				{200, 100, 100},
				{200, 100, 50, 50},
				{200, 50, 50, 50, 50},
				{100, 100, 100, 100},
				{100, 100, 100, 50, 50},
				{100, 100, 50, 50, 50, 50},
				{100, 50, 50, 50, 50, 50, 50},
				{50, 50, 50, 50, 50, 50, 50, 50},
			},
		},
		{
			amount:    100,
			banknotes: []int{50, 20, 10},
			expected: [][]int{
				{50, 50},
				{50, 20, 20, 10},
				{50, 20, 10, 10, 10},
				{50, 10, 10, 10, 10, 10},
				{20, 20, 20, 20, 20},
				{20, 20, 20, 20, 10, 10},
				{20, 20, 20, 10, 10, 10, 10},
				{20, 20, 10, 10, 10, 10, 10, 10},
				{20, 10, 10, 10, 10, 10, 10, 10, 10},
				{10, 10, 10, 10, 10, 10, 10, 10, 10, 10},
			},
		},
		{
			amount:    50,
			banknotes: []int{50, 20, 10},
			expected: [][]int{
				{50},
				{20, 20, 10},
				{20, 10, 10, 10},
				{10, 10, 10, 10, 10},
			},
		},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			var result [][]int
			service.Exchange(tt.amount, tt.banknotes, []int{}, 0, &result)

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}
