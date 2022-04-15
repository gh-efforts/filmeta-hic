package paginator

import (
	"reflect"
	"testing"
)

func TestPage_Calculate(t *testing.T) {
	type fields struct {
		PageNum     int64
		PageSize    int64
		PageMaxSize int64
	}
	type args struct {
		total int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int64
	}{
		{
			"c1",
			fields{
				PageNum:  1,
				PageSize: 10,
			},
			args{total: 30},
			3,
		},
		{
			"c2",
			fields{
				PageNum:  1,
				PageSize: 10,
			},
			args{total: 27},
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Page{
				Page:        tt.fields.PageNum,
				PageSize:    tt.fields.PageSize,
				PageMaxSize: tt.fields.PageMaxSize,
			}
			if got := p.Calculate(tt.args.total); got != tt.want {
				t.Errorf("Calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPage_Limit(t *testing.T) {
	type fields struct {
		PageNum     int64
		PageSize    int64
		PageMaxSize int64
	}
	tests := []struct {
		name   string
		fields fields
		want   int64
	}{
		{
			"c1",
			fields{
				PageNum:     1,
				PageSize:    10,
				PageMaxSize: 3,
			},
			3,
		},
		{
			"c2",
			fields{
				PageNum:  1,
				PageSize: 0,
			},
			PageDefaultSize,
		},
		{
			"c3",
			fields{
				PageNum:     1,
				PageSize:    10,
				PageMaxSize: 20,
			},
			10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Page{
				Page:        tt.fields.PageNum,
				PageSize:    tt.fields.PageSize,
				PageMaxSize: tt.fields.PageMaxSize,
			}
			if got := p.Limit(); got != tt.want {
				t.Errorf("Limit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPage_Offset(t *testing.T) {
	type fields struct {
		PageNum     int64
		PageSize    int64
		PageMaxSize int64
	}
	tests := []struct {
		name   string
		fields fields
		want   int64
	}{
		{
			"c1",
			fields{
				PageNum:     0,
				PageSize:    0,
				PageMaxSize: 0,
			},
			0,
		},
		{
			"c2",
			fields{
				PageNum:  2,
				PageSize: 10,
			},
			10,
		},
		{
			"c3",
			fields{
				PageNum:  2,
				PageSize: 0,
			},
			20,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Page{
				Page:        tt.fields.PageNum,
				PageSize:    tt.fields.PageSize,
				PageMaxSize: tt.fields.PageMaxSize,
			}
			if got := p.Offset(); got != tt.want {
				t.Errorf("Offset() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPage_ToListPage(t *testing.T) {
	type fields struct {
		PageNum     int64
		PageSize    int64
		PageMaxSize int64
	}
	type args struct {
		total int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   ListPage
	}{
		{
			"c1",
			fields{
				PageNum:  1,
				PageSize: 10,
			},
			args{total: 30},
			ListPage{
				TotalPage: 3,
				Page:      1,
				Total:     30,
				PageSize:  10,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Page{
				Page:        tt.fields.PageNum,
				PageSize:    tt.fields.PageSize,
				PageMaxSize: tt.fields.PageMaxSize,
			}
			if got := p.ToListPage(tt.args.total); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToListPage() = %+v, want %+v", got, tt.want)
			}
		})
	}
}
