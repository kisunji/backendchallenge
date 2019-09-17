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
				[][]string{},
			},
			want: nil,
		},
		{
			name: "nil matrix returns nil",
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

func TestInvert(t *testing.T) {
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
			name: "Inverts valid matrix",
			args: args{
				[][]string{
					{"1", "2", "3"},
					{"4", "5", "6"},
					{"7", "8", "9"},
				},
			},
			want:    "1,4,7\n2,5,8\n3,6,9",
			wantErr: false,
		},
		{
			name: "NxM matrix is transposed",
			args: args{
				[][]string{
					{"1", "2", "3"},
				},
			},
			want:    "1\n2\n3",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Invert(tt.args.records)
			if (err != nil) != tt.wantErr {
				t.Errorf("Invert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Invert() = %v, want %v", got, tt.want)
			}
		})
	}
}
