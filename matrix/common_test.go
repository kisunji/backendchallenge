package matrix

import (
	"reflect"
	"testing"
)

func Test_atoiMatrix(t *testing.T) {
	type args struct {
		records [][]string
	}
	tests := []struct {
		name    string
		args    args
		want    [][]int
		wantErr bool
	}{
		{
			name: "integer input returns valid int matrix",
			args: args{
				[][]string{
					{"1", "2"},
					{"3", "4"},
				}},
			want: [][]int{
				{1, 2},
				{3, 4},
			},
			wantErr: false,
		},
		{
			name: "empty input returns empty matrix",
			args: args{
				[][]string{},
			},
			want:    [][]int{},
			wantErr: false,
		},
		{
			name: "float returns error",
			args: args{
				[][]string{
					{"1", "2.0"},
					{"3", "4"},
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "string returns error",
			args: args{
				[][]string{
					{"1", "hello"},
					{"3", "4"},
				}},
			want:    nil,
			wantErr: true,
		},
		{
			name: "integer greater than max throws error",
			args: args{
				[][]string{
					{"1", "2"},
					{"3", "9999999999999999999999999999"},
				}},
			want:    nil,
			wantErr: true,
		},
		{
			name: "integer less than min throws error",
			args: args{
				[][]string{
					{"1", "2"},
					{"3", "-9999999999999999999999999999"},
				}},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AtoiMatrix(tt.args.records)
			if (err != nil) != tt.wantErr {
				t.Errorf("atoiMatrix() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("atoiMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_atoiSlice(t *testing.T) {
	const MaxInt = int(^uint(0) >> 1)
	const MinInt = -MaxInt - 1
	type args struct {
		slice []string
	}
	tests := []struct {
		name    string
		args    args
		want    []int
		wantErr bool
	}{
		{
			name: "integer input returns valid int slice",
			args: args{
				[]string{"1", "2", "3"},
			},
			want:    []int{1, 2, 3},
			wantErr: false,
		},
		{
			name: "empty input returns valid empty slice",
			args: args{
				[]string{},
			},
			want:    []int{},
			wantErr: false,
		},
		{
			name: "float returns error",
			args: args{
				[]string{"1", "2", "0.3"},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "string returns error",
			args: args{
				[]string{"1", "2", "hello"},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "integer greater than max throws error",
			args: args{
				[]string{"1", "2", "9999999999999999999999999999"},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "integer less than min throws error",
			args: args{
				[]string{"1", "2", "-9999999999999999999999999999"},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AtoiSlice(tt.args.slice)
			if (err != nil) != tt.wantErr {
				t.Errorf("atoiSlice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("atoiSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}
