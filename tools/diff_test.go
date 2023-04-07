package tools

import (
	"reflect"
	"sort"
	"testing"
)

func TestDiffStr(t *testing.T) {
	type args struct {
		a []string
		b []string
	}
	tests := []struct {
		name           string
		args           args
		wantInAAndB    []string
		wantInAButNotB []string
		wantInBButNotA []string
	}{
		{
			"c1",
			args{a: []string{"aaa", "bbb", "ccc", "ddd"}, b: []string{"bbb", "ccc"}},
			[]string{"bbb", "ccc"},
			[]string{"aaa", "ddd"},
			nil,
		},
		{
			"c2",
			args{a: []string{"aaa", "bbb", "ccc"}, b: []string{"bbb", "ccc", "ddd"}},
			[]string{"bbb", "ccc"},
			[]string{"aaa"},
			[]string{"ddd"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotInAAndB, gotInAButNotB, gotInBButNotA := DiffStr(tt.args.a, tt.args.b)

			sort.Slice(gotInAAndB, func(i, j int) bool {
				return gotInAAndB[i] < gotInAAndB[j]
			})
			sort.Slice(gotInAButNotB, func(i, j int) bool {
				return gotInAButNotB[i] < gotInAButNotB[j]
			})
			sort.Slice(gotInBButNotA, func(i, j int) bool {
				return gotInBButNotA[i] < gotInBButNotA[j]
			})
			if !reflect.DeepEqual(gotInAAndB, tt.wantInAAndB) {
				t.Errorf("DiffStr() gotInAAndB = %v, want %v", gotInAAndB, tt.wantInAAndB)
			}
			if !reflect.DeepEqual(gotInAButNotB, tt.wantInAButNotB) {
				t.Errorf("DiffStr() gotInAButNotB = %v, want %v", gotInAButNotB, tt.wantInAButNotB)
			}
			if !reflect.DeepEqual(gotInBButNotA, tt.wantInBButNotA) {
				t.Errorf("DiffStr() gotInBButNotA = %v, want %v", gotInBButNotA, tt.wantInBButNotA)
			}
		})
	}
}
