package matrix

import "testing"

func Test_sumSlice(t *testing.T) {
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
			name: "Positive inputs return sum",
			args: args{
				[]int{1, 2, 3, 4, 5},
			},
			want:    15,
			wantErr: false,
		},
		{
			name: "Negative inputs return sum",
			args: args{
				[]int{-1, -2, -3, -4, -5},
			},
			want:    -15,
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
				[]int{MaxInt, 1},
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "Underflow throws error",
			args: args{
				[]int{MinInt, -1},
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := sumSlice(tt.args.arr)
			if (err != nil) != tt.wantErr {
				t.Errorf("sumSlice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("sumSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}
