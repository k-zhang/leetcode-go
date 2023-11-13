package easy

import "testing"

func TestRangeSumQuery_SumRange(t *testing.T) {
	type fields struct {
		accSum []int
	}
	type args struct {
		left  int
		right int
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{"", fields{[]int{-2, 0, 3, -5, 2, -1}}, args{0, 2}, 1},
		{"", fields{[]int{-2, 0, 3, -5, 2, -1}}, args{2, 5}, -1},
		{"", fields{[]int{-2, 0, 3, -5, 2, -1}}, args{0, 5}, -3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRangeSumQuery(tt.fields.accSum)
			if got := r.SumRange(tt.args.left, tt.args.right); got != tt.want {
				t.Errorf("SumRange() = %v, want %v", got, tt.want)
			}
		})
	}
}
