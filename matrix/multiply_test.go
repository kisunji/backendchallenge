package matrix

import (
	"testing"
)

func Test_productSlice(t *testing.T) {
	const MaxInt = int(^uint(0) >> 1)
	const MinInt = -MaxInt - 1
	type args struct {
		arr []int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "Positive inputs return product",
			args: args{
				[]int{1, 2, 3, 4, 5},
			},
			want:    120,
			wantErr: false,
		},
		{
			name: "Negative inputs return product",
			args: args{
				[]int{-1, -2, -3, -4, -5},
			},
			want:    -120,
			wantErr: false,
		},
		{
			name: "Empty array returns zero",
			args: args{
				[]int{},
			},
			want:    0,
			wantErr: false,
		},
		{
			name: "Overflow throws error",
			args: args{
				[]int{MaxInt, MaxInt},
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "Underflow throws error",
			args: args{
				[]int{MinInt, MinInt},
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := productSlice(tt.args.arr)
			if (err != nil) != tt.wantErr {
				t.Errorf("productSlice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("productSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMultiply(t *testing.T) {
	type args struct {
		records [][]string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Valid input returns desired result",
			args: args{
				[][]string{
					{"1", "2", "3"},
					{"4", "5", "6"},
					{"7", "8", "9"},
				},
			},
			want:    "362880",
			wantErr: false,
		},
		{
			name: "Empty matrix returns zero",
			args: args{
				[][]string{},
			},
			want:    "0",
			wantErr: false,
		},
		{
			name: "Invalid input throws error",
			args: args{
				[][]string{
					{"1", "2", "test"},
					{"4", "5", "6"},
					{"7", "8", "9"},
				},
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Multiply(tt.args.records)
			if (err != nil) != tt.wantErr {
				t.Errorf("Multiply() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Multiply() = %v, want %v", got, tt.want)
			}
		})
	}
}
