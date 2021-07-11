package my_slice

import (
	"reflect"
	"testing"
)

func Test_popFront(t *testing.T) {
	type args struct {
		input []int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 []int
	}{
		{
			name: "popFront",
			args: args{
				[]int{1, 2, 3, 4},
			},
			want:  1,
			want1: []int{2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := popFront(tt.args.input)
			if got != tt.want {
				t.Errorf("popFront() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("popFront() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_pushToFront(t *testing.T) {
	type args struct {
		e     int
		input []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "pushToFront",
			args: args{
				e:     0,
				input: []int{1, 2, 3, 4},
			},
			want: []int{0, 1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pushToFront(tt.args.e, tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pushToFront() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pushToEnd(t *testing.T) {
	type args struct {
		e     int
		input []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "pushToEnd",
			args: args{
				e:     5,
				input: []int{1, 2, 3, 4},
			},
			want: []int{1, 2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pushToEnd(tt.args.e, tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pushToEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPopEnd(t *testing.T) {
	type args struct {
		input []int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 []int
	}{
		{
			name: "PopEnd",
			args: args{
				input: []int{1, 2, 3, 4},
			},
			want:  4,
			want1: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := PopEnd(tt.args.input)
			if got != tt.want {
				t.Errorf("PopEnd() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("PopEnd() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
