package math

import "testing"

func TestGCD(t *testing.T) {
	tests := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		{
			name:     "正常字符串",
			a:        12,
			b:        4,
			expected: 4,
		},
		{
			name:     "空字符串",
			a:        0,
			b:        5,
			expected: 5,
		},
		{
			name:     "中文字符串",
			a:        1,
			b:        7,
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GCD(tt.a, tt.b)
			if got != tt.expected {
				t.Errorf("ProcessString() = %v, want %v", got, tt.expected)
			}
		})
	}
}
