package easy

import "testing"

func Test_reverseInteger(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{"", args{1}, 1, false},
		{"", args{0}, 0, false},
		{"", args{21}, 12, false},
		{"", args{321}, 123, false},
		{"", args{-321}, -123, false},
		{"", args{7463847412}, 2147483647, false},
		{"", args{-7463847412}, -2147483647, false},
		{"", args{8885774586302733229}, -1, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := reverseInteger(tt.args.x)
			if (err != nil) != tt.wantErr {
				t.Errorf("reverseInteger() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("reverseInteger() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reverseInteger2(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{1}, 1},
		{"", args{0}, 0},
		{"", args{21}, 12},
		{"", args{321}, 123},
		{"", args{-321}, -123},
		{"", args{7463847412}, 2147483647},
		{"", args{-7463847412}, -2147483647},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseInteger2(tt.args.x); got != tt.want {
				t.Errorf("reverseInteger2() = %v, want %v", got, tt.want)
			}
		})
	}
}
