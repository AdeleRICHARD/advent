package day8

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var trees = [][]int{
	{1, 2, 3, 4, 5},
	{1, 5, 3, 4, 5},
	{1, 3, 3, 4, 5},
}

func Test_checkLeft(t *testing.T) {
	leftTrees := trees[1][:2]

	type args struct {
		trees       []int
		rowNb       int
		colNb       int
		currentTree int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "check left true",
			args: args{
				trees:       leftTrees,
				rowNb:       0,
				colNb:       2,
				currentTree: 6,
			},
			want: true,
		},
		{
			name: "check left false strict",
			args: args{
				trees:       leftTrees,
				rowNb:       0,
				colNb:       1,
				currentTree: 2,
			},
			want: false,
		},
		{
			name: "check left false equal",
			args: args{
				trees:       leftTrees,
				rowNb:       0,
				colNb:       1,
				currentTree: 1,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := checkLeft(tt.args.trees, tt.args.rowNb, tt.args.colNb, tt.args.currentTree)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_checkRight(t *testing.T) {
	rightTrees := trees[1][3:]
	type args struct {
		trees       []int
		rowNb       int
		colNb       int
		currentTree int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "check right true",
			args: args{
				trees:       rightTrees,
				rowNb:       0,
				colNb:       2,
				currentTree: 8,
			},
			want: true,
		},
		{
			name: "check right false strict",
			args: args{
				trees:       rightTrees,
				rowNb:       0,
				colNb:       1,
				currentTree: 4,
			},
			want: false,
		},
		{
			name: "check right false equal",
			args: args{
				trees:       rightTrees,
				rowNb:       0,
				colNb:       1,
				currentTree: 5,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := checkRight(tt.args.trees, tt.args.rowNb, tt.args.colNb, tt.args.currentTree)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_checkBottom(t *testing.T) {
	treesTop := trees[:2]
	type args struct {
		trees       [][]int
		colNb       int
		currentTree int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "check bottom true",
			args: args{
				trees:       treesTop,
				colNb:       2,
				currentTree: 4,
			},
			want: true,
		},
		{
			name: "check bottom false strict",
			args: args{
				trees:       treesTop,
				colNb:       2,
				currentTree: 2,
			},
			want: false,
		},
		{
			name: "check bottom false equal",
			args: args{
				trees:       treesTop,
				colNb:       2,
				currentTree: 3,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := checkBottom(tt.args.trees, tt.args.colNb, tt.args.currentTree)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_checkTop(t *testing.T) {
	treesBottom := trees[1:]
	type args struct {
		trees       [][]int
		colNb       int
		currentTree int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "check top true",
			args: args{
				trees:       treesBottom,
				colNb:       2,
				currentTree: 4,
			},
			want: true,
		},
		{
			name: "check top false strict",
			args: args{
				trees:       treesBottom,
				colNb:       2,
				currentTree: 2,
			},
			want: false,
		},
		{
			name: "check top false equal",
			args: args{
				trees:       treesBottom,
				colNb:       2,
				currentTree: 3,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := checkTop(tt.args.trees, tt.args.colNb, tt.args.currentTree)
			assert.Equal(t, tt.want, got)
		})
	}
}
