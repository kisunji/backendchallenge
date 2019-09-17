package matrix

import (
	"reflect"
	"testing"
)

func Test_transpose(t *testing.T) {
	type args struct {
		records [][]string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "NxN matrix is transposed",
			args: args{
				[][]string{
					{"1", "2", "3"},
					{"4", "5", "6"},
					{"7", "8", "9"},
				},
			},
			want: [][]string{
				{"1", "4", "7"},
				{"2", "5", "8"},
				{"3", "6", "9"},
			},
		},
		{
			name: "NxM matrix is transposed",
			args: args{
				[][]string{
					{"1", "2", "3"},
				},
			},
			want: [][]string{
				{"1"},
				{"2"},
				{"3"},
			},
		},
		{
			name: "empty matrix returns nil",
			args: args{
				nil,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := transpose(tt.args.records); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("transpose() = %v, want %v", got, tt.want)
			}
		})
	}
}
