package tools

import "testing"

func TestIsNilFixed(t *testing.T) {
	type args struct {
		i interface{}
	}

	type user struct {
		Name string
	}

	var (
		ch chan struct{}
		u  *user
		u1 user
		s  []string
		m  map[string]string
	)

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "slice",
			args: args{i: s},
			want: true,
		},
		{
			name: "map",
			args: args{i: m},
			want: true,
		},
		{
			name: "chan",
			args: args{i: ch},
			want: true,
		},
		{
			name: "ptr",
			args: args{i: u},
			want: true,
		},
		{
			name: "struct",
			args: args{i: u1},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsNilFixed(tt.args.i); got != tt.want {
				t.Errorf("IsNilFixed() = %v, want %v", got, tt.want)
			}
		})
	}
}
