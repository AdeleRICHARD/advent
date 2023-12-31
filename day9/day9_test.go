package day9

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var moves = []Move{
	{"R", 4},
	{"U", 3},
	{"R", 2},
	{"D", 2},
	{"L", 2},
}

func Test_moveRight(t *testing.T) {
	type args struct {
		rope   [][]string
		head   []int
		tail   []int
		nbMove int
		count  *int
	}

	tests := []struct {
		name     string
		args     args
		wantRope [][]string
		wantHead []int
		wantTail []int
	}{
		{
			name: "Move right at start of rope",
			args: args{
				rope: [][]string{
					{"H", "", "", ""},
				},
				head:   []int{0, 0},
				tail:   []int{0, 0},
				nbMove: 2,
				count:  new(int),
			},
			wantRope: [][]string{
				{"#", "T", "H", ""},
			},
			wantHead: []int{0, 2},
			wantTail: []int{0, 1},
		},
		{
			name: "Move right",
			args: args{
				rope: [][]string{
					{"#", "T", "H", "", ""},
				},
				head:   []int{0, 2},
				tail:   []int{0, 1},
				nbMove: 2,
				count:  new(int),
			},
			wantRope: [][]string{
				{"#", "#", "#", "T", "H"},
			},
			wantHead: []int{0, 4},
			wantTail: []int{0, 3},
		},
		{
			name: "Move right after left",
			args: args{
				rope: [][]string{
					{"#", "#", "H", "T", ""},
				},
				head:   []int{0, 2},
				tail:   []int{0, 3},
				nbMove: 2,
				count:  new(int),
			},
			wantRope: [][]string{
				{"#", "#", "", "T", "H"},
			},
			wantHead: []int{0, 4},
			wantTail: []int{0, 3},
		},
		{
			name: "Move right after up",
			args: args{
				rope: [][]string{
					{"#", "#", "#", "", ""},
					{"", "", "T", "", ""},
					{"", "", "H", "", ""},
				},
				head:   []int{2, 2},
				tail:   []int{1, 2},
				nbMove: 2,
				count:  new(int),
			},
			wantRope: [][]string{
				{"#", "#", "#", "", ""},
				{"", "", "#", "", ""},
				{"", "", "", "T", "H"},
			},
			wantHead: []int{2, 4},
			wantTail: []int{2, 3},
		},
		{
			name: "Move right after down",
			args: args{
				rope: [][]string{
					{"#", "#", "#", "", "", ""},
					{"", "", "", "H", "", ""},
					{"", "", "", "T", "", ""},
				},
				head:   []int{1, 3},
				tail:   []int{2, 3},
				nbMove: 2,
				count:  new(int),
			},
			wantRope: [][]string{
				{"#", "#", "#", "", "", ""},
				{"", "", "", "", "T", "H"},
				{"", "", "", "#", "", ""},
			},
			wantHead: []int{1, 5},
			wantTail: []int{1, 4},
		},
		{
			name: "Move right one time",
			args: args{
				rope: [][]string{
					{"#", "#", "#", "", "", ""},
					{"", "", "", "H", "", ""},
					{"", "", "", "T", "", ""},
				},
				head:   []int{1, 3},
				tail:   []int{2, 3},
				nbMove: 1,
				count:  new(int),
			},
			wantRope: [][]string{
				{"#", "#", "#", "", "", ""},
				{"", "", "", "", "H", ""},
				{"", "", "", "T", "", ""},
			},
			wantHead: []int{1, 4},
			wantTail: []int{2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRope, gotHead, gotTail := moveRight(tt.args.rope, tt.args.head, tt.args.tail, tt.args.nbMove, tt.args.count)
			assert.Equal(t, tt.wantRope, gotRope)
			assert.Equal(t, tt.wantHead, gotHead)
			assert.Equal(t, tt.wantTail, gotTail)
		})
	}
}

func Test_moveUp(t *testing.T) {
	type args struct {
		rope   [][]string
		head   []int
		tail   []int
		nbMove int
		count  *int
	}
	tests := []struct {
		name  string
		args  args
		want  [][]string
		want1 []int
		want2 []int
	}{
		{
			name: "Move up",
			args: args{
				rope: [][]string{
					{" ", "", "", ""},
					{"", "", "", ""},
					{"", "", "T", "H"},
					{"", "", "", ""},
					{"", "", "", ""},
				},
				head:   []int{2, 3},
				tail:   []int{2, 2},
				nbMove: 2,
			},
			want: [][]string{
				{" ", "", "", "H"},
				{"", "", "", "T"},
				{"", "", "#", ""},
				{"", "", "", ""},
				{"", "", "", ""},
			},
			want1: []int{0, 3},
			want2: []int{1, 3},
		},
		{
			name: "Move up",
			args: args{
				rope: [][]string{
					{" ", "", "", ""},
					{"", "", "", ""},
					{"", "", "", ""},
					{"", "", "", ""},
					{"", "", "T", "H"},
				},
				head:   []int{4, 3},
				tail:   []int{4, 2},
				nbMove: 4,
			},
			want: [][]string{
				{" ", "", "", "H"},
				{"", "", "", "T"},
				{"", "", "", "#"},
				{"", "", "", "#"},
				{"", "", "#", ""},
			},
			want1: []int{0, 3},
			want2: []int{1, 3},
		},
		{
			name: "Move up one",
			args: args{
				rope: [][]string{
					{" ", "", "", ""},
					{"", "", "", ""},
					{"", "", "", ""},
					{"", "", "", ""},
					{"", "", "T", "H"},
				},
				head:   []int{4, 3},
				tail:   []int{4, 2},
				nbMove: 1,
			},
			want: [][]string{
				{" ", "", "", ""},
				{"", "", "", ""},
				{"", "", "", ""},
				{"", "", "", "H"},
				{"", "", "T", ""},
			},
			want1: []int{3, 3},
			want2: []int{4, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRope, gotHead, gotTail := moveUp(tt.args.rope, tt.args.head, tt.args.tail, tt.args.nbMove, tt.args.count)
			assert.Equal(t, tt.want, gotRope)
			assert.Equal(t, tt.want1, gotHead)
			assert.Equal(t, tt.want2, gotTail)
		})
	}
}

func Test_moveLeft(t *testing.T) {
	type args struct {
		rope   [][]string
		head   []int
		tail   []int
		nbMove int
		count  *int
	}
	tests := []struct {
		name  string
		args  args
		want  [][]string
		want1 []int
		want2 []int
	}{
		{
			name: "Move left",
			args: args{
				rope: [][]string{
					{"#", "#", "#", ""},
					{"", "", "", "T"},
					{"", "", "", "H"},
				},
				head:   []int{2, 3},
				tail:   []int{1, 3},
				nbMove: 2,
			},
			want: [][]string{
				{"#", "#", "#", ""},
				{"", "", "", "#"},
				{"", "H", "T", ""},
			},
			want1: []int{2, 1},
			want2: []int{2, 2},
		},
		{
			name: "Move left at start of rope",
			args: args{
				rope: [][]string{
					{"", "", "", "H"},
				},
				head:   []int{0, 3},
				tail:   []int{0, 3},
				nbMove: 2,
			},
			want: [][]string{
				{"", "H", "T", "#"},
			},
			want1: []int{0, 1},
			want2: []int{0, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRope, gotHead, gotTail := moveLeft(tt.args.rope, tt.args.head, tt.args.tail, tt.args.nbMove, tt.args.count)
			assert.Equal(t, tt.want, gotRope)
			assert.Equal(t, tt.want1, gotHead)
			assert.Equal(t, tt.want2, gotTail)
		})
	}
}

func Test_moveDown(t *testing.T) {
	type args struct {
		rope   [][]string
		head   []int
		tail   []int
		nbMove int
		count  *int
	}
	tests := []struct {
		name  string
		args  args
		want  [][]string
		want1 []int
		want2 []int
	}{

		{
			name: "Move down",
			args: args{
				rope: [][]string{
					{"T", "H", "", ""},
					{"", "", "", ""},
					{"", "", "", ""},
					{"", "", "", ""},
				},
				head:   []int{0, 1},
				tail:   []int{0, 0},
				nbMove: 2,
			},
			want: [][]string{
				{"#", "", "", ""},
				{"", "T", "", ""},
				{"", "H", "", ""},
				{"", "", "", ""},
			},
			want1: []int{2, 1},
			want2: []int{1, 1},
		},
		{
			name: "Move down after up",
			args: args{
				rope: [][]string{
					{"", "H", "", ""},
					{"", "T", "", ""},
					{"", "", "", ""},
					{"", "", "", ""},
				},
				head:   []int{0, 1},
				tail:   []int{1, 1},
				nbMove: 2,
			},
			want: [][]string{
				{"", "", "", ""},
				{"", "T", "", ""},
				{"", "H", "", ""},
				{"", "", "", ""},
			},
			want1: []int{2, 1},
			want2: []int{1, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRope, gotHead, gotTail := moveDown(tt.args.rope, tt.args.head, tt.args.tail, tt.args.nbMove, tt.args.count)
			assert.Equal(t, tt.want, gotRope)
			assert.Equal(t, tt.want1, gotHead)
			assert.Equal(t, tt.want2, gotTail)
		})
	}
}
