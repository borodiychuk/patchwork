package models

import (
	"testing"
)

func TestCanvas_XYforIndex(t *testing.T) {
	tests := []struct {
		name    string
		arg     int
		wantX   int
		wantY   int
		wantErr bool
	}{
		{arg: 7, wantX: 1, wantY: 2},
		{arg: 0, wantX: 0, wantY: 0},
		{arg: 11, wantX: 2, wantY: 3},
		{arg: 9, wantX: 0, wantY: 3},
		{arg: -1, wantX: -1, wantY: -1, wantErr: true},
		{arg: 12, wantX: -1, wantY: -1, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Canvas{
				Length: 4,
				Width:  3,
			}
			got, got1, err := c.XYforIndex(tt.arg)
			if (err != nil) != tt.wantErr {
				t.Errorf("Canvas.XYforIndex(%d) error = %v, wantErr %v", tt.arg, err, tt.wantErr)
				return
			}
			if got != tt.wantX {
				t.Errorf("Canvas.XYforIndex(%d) got = %v, want %v", tt.arg, got, tt.wantX)
			}
			if got1 != tt.wantY {
				t.Errorf("Canvas.XYforIndex(%d) got1 = %v, want %v", tt.arg, got1, tt.wantY)
			}
		})
	}
}

func TestCanvas_IndexForXY(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		args    args
		want    int
		wantErr bool
	}{
		{args: args{x: 1, y: 2}, want: 7},
		{args: args{x: 0, y: 0}, want: 0},
		{args: args{x: 2, y: 3}, want: 11},
		{args: args{x: 0, y: 3}, want: 9},
		{args: args{x: 0, y: 4}, want: -1, wantErr: true},
		{args: args{x: 3, y: 0}, want: -1, wantErr: true},
		{args: args{x: -1, y: 0}, want: -1, wantErr: true},
		{args: args{x: 0, y: -1}, want: -1, wantErr: true},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			c := &Canvas{
				Length: 4,
				Width:  3,
			}
			got, err := c.IndexForXY(tt.args.x, tt.args.y)
			if (err != nil) != tt.wantErr {
				t.Errorf("Canvas.IndexForXY(%d, %d) error = %v, wantErr %v", tt.args.x, tt.args.y, err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Canvas.IndexForXY(%d, %d) = %v, want %v", tt.args.x, tt.args.y, got, tt.want)
			}
		})
	}
}
