package main

import (
	"reflect"
	"testing"
)

func TestCheckDuplicateIDs(t *testing.T) {
	type args struct {
		tree *Tree
	}
	tests := []struct {
		name      string
		args      args
		wantValue *int
		wantLevel int
	}{

		{"nil tree", args{nil}, nil, 0},
		{"one node tree", args{&Tree{3, nil, nil}}, nil, 0},
		{"no duplicates", args{
			&Tree{
				1,
				&Tree{2, nil, nil},
				&Tree{3, nil, nil},
			}}, nil, 0},

		{"one duplicate same level", args{
			&Tree{
				1,
				&Tree{
					2,
					&Tree{4, nil, nil},
					nil,
				},
				&Tree{
					8,
					&Tree{5, nil, nil},
					&Tree{4, nil, nil},
				},
			}}, NewNullableInt(4), 2},

		{"multiple duplicates different levels", args{
			&Tree{
				8,
				&Tree{
					4,
					&Tree{5, nil, nil},
					&Tree{4, nil, nil},
				},
				&Tree{
					3,
					&Tree{2, nil, nil},
					&Tree{2, nil, nil},
				},
			}}, NewNullableInt(4), 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotValue, gotLevel := CheckDuplicateIDs(tt.args.tree)
			if gotValue != tt.wantValue {
				if gotValue == nil {
					t.Errorf("CheckDuplicateIDs() gotValue = %v, want %v", gotValue, *tt.wantValue)
				} else if tt.wantValue == nil {
					t.Errorf("CheckDuplicateIDs() gotValue = %v, want %v", *gotValue, tt.wantValue)
				} else {
					// check value equality
					if *gotValue != *tt.wantValue {
						t.Errorf("CheckDuplicateIDs() gotValue = %v, want %v", *gotValue, *tt.wantValue)

					}
				}
			}
			if gotLevel != tt.wantLevel {
				t.Errorf("CheckDuplicateIDs() gotLevel = %v, want %v", gotLevel, tt.wantLevel)
			}
		})
	}
}

func Test_recordLevels(t *testing.T) {
	type args struct {
		tree *Tree
	}
	tests := []struct {
		name string
		args args
		want []valueLevel
	}{
		{"nil tree", args{nil}, []valueLevel{}},
		{"one node tree", args{&Tree{6, nil, nil}}, []valueLevel{{6, 0}}},
		{"every tree node one child", args{
			&Tree{
				6,
				&Tree{
					2,
					nil,
					&Tree{8, nil, nil}},
				nil,
			}}, []valueLevel{{6, 0}, {2, 1}, {8, 2}}},
		{"Complete tree", args{
			&Tree{
				1,
				&Tree{
					2,
					&Tree{3, nil, nil},
					&Tree{4, nil, nil},
				},
				&Tree{
					5,
					&Tree{6, nil, nil},
					&Tree{7, nil, nil},
				},
			}}, []valueLevel{{1, 0}, {2, 1}, {5, 1}, {3, 2}, {4, 2}, {6, 2}, {7, 2}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := recordLevels(tt.args.tree); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("recordLevels() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_scanDuplicates(t *testing.T) {
	type args struct {
		vl []valueLevel
	}
	tests := []struct {
		name string
		args args
		want []valueLevel
	}{
		{"empty list", args{[]valueLevel{}}, []valueLevel{}},
		{"one item list", args{[]valueLevel{{5, 0}}}, []valueLevel{}},
		{"no duplicates", args{[]valueLevel{{3, 0}, {4, 1}, {5, 1}, {9, 2}}}, []valueLevel{}},
		{"one duplicate same level", args{[]valueLevel{{2, 0}, {3, 1}, {3, 1}, {9, 2}}}, []valueLevel{{3, 1}}},
		{"one duplicate different levels", args{[]valueLevel{{3, 0}, {4, 1}, {4, 2}, {5, 2}}}, []valueLevel{{4, 1}}},
		{"multiple duplicates multiple values", args{[]valueLevel{{3, 0}, {4, 1}, {4, 2}, {4, 2}, {5, 3}, {6, 3}, {6, 4}}}, []valueLevel{{4, 1}, {6, 3}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := scanDuplicates(tt.args.vl); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("scanDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
