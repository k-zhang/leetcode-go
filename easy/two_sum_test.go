package easy

import "testing"

func Test_twoSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}

	tests := []struct {
		name    string
		args    args
		want    int
		want1   int
		wantErr bool
	}{
		{"", args{[]int{2, 7, 11, 5}, 9}, 0, 1, false},
		{"", args{[]int{3, 2, 4}, 6}, 1, 2, false},
		{"", args{[]int{3, 3}, 6}, 0, 1, false},
		{"", args{[]int{3, 6, -1}, 2}, 0, 2, false},
		{"", args{[]int{1, 5, 0, 0}, 0}, 2, 3, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := twoSum(tt.args.nums, tt.args.target)
			if (err != nil) != tt.wantErr {
				t.Errorf("twoSum() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("twoSum() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("twoSum() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
