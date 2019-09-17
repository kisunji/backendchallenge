package matrix

import "testing"

func TestFlatten(t *testing.T) {
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
			name: "string matrix",
			args: args{
				[][]string{
					{"1", "2", "3"},
					{"4", "5", "6"},
					{"7", "8", "9"},
				}},
			want:    "1,2,3,4,5,6,7,8,9",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Flatten(tt.args.records)
			if (err != nil) != tt.wantErr {
				t.Errorf("Flatten() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Flatten() = %v, want %v", got, tt.want)
			}
		})
	}
}
