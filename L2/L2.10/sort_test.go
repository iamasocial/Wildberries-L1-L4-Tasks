package main

import (
	"reflect"
	"testing"
)

func TestCompareNumeric(t *testing.T) {
	tests := []struct {
		a, b string
		want int
	}{
		{"10", "20", -1},
		{"20", "10", 1},
		{"10", "10", 0},
		{"a", "b", -1},
	}

	for _, test := range tests {
		if got := compareNumeric(test.a, test.b); got != test.want {
			t.Errorf("compareNumeric(%q, %q) = %d, want %d", test.a, test.b, got, test.want)
		}
	}
}

func TestCompareMonth(t *testing.T) {
	tests := []struct {
		mon1, mon2 string
		want       int
	}{
		{"Jan", "feb", -1},
		{"fEb", "jaN", 1},
		{"FEB", "FEb", 0},
		{"a", "b", -1},
	}

	for _, test := range tests {
		if got := compareMonth(test.mon1, test.mon2); got != test.want {
			t.Errorf("compareMonth(%q, %q) = %d, want %d", test.mon1, test.mon2, got, test.want)
		}
	}
}

func TestCompareHumanNumeric(t *testing.T) {
	tests := []struct {
		a, b string
		want int
	}{
		{"", "", 0},
		{"a", "b", -1},

		{"10", "12", -1},
		{"12", "10", 1},
		{"10", "10", 0},

		{"10K", "12K", -1},
		{"12K", "10K", 1},
		{"12K", "12K", 0},

		{"10M", "12M", -1},
		{"12M", "10M", 1},
		{"12M", "12M", 0},

		{"10G", "12G", -1},
		{"12G", "10G", 1},
		{"12G", "12G", 0},

		{"10T", "12T", -1},
		{"12T", "10T", 1},
		{"12T", "12T", 0},

		{"1025K", "1M", 1},
		{"1M", "1025K", -1},

		{"1025M", "1G", 1},
		{"1G", "1025M", -1},

		{"1025G", "1T", 1},
		{"1T", "1025G", -1},
	}

	for _, test := range tests {
		if got := compareHumanNumeric(test.a, test.b); got != test.want {
			t.Errorf("compareHumanNumeric(%q, %q) = %d, want %d", test.a, test.b, got, test.want)
		}
	}
}

func TestRemoveDuplicates(t *testing.T) {
	input1 := []string{"a", "b", "c"}
	want1 := []string{"a", "b", "c"}

	if got1 := removeDuplicates(input1); !reflect.DeepEqual(got1, want1) {
		t.Errorf("removeDuplicates(%q) = %q, want %q", input1, got1, want1)
	}

	input2 := []string{"a", "a", "b"}
	want2 := []string{"a", "b"}

	if got2 := removeDuplicates(input2); !reflect.DeepEqual(got2, want2) {
		t.Errorf("removeDuplicates(%q) = %q, want %q", input2, got2, want2)
	}

	input3 := []string{}
	want3 := []string{}

	if got3 := removeDuplicates(input3); !reflect.DeepEqual(got3, want3) {
		t.Errorf("removeDuplicates(%q) = %q, want %q", input3, got3, want3)
	}
}

func TestGetColumn(t *testing.T) {
	tests := []struct {
		input string
		col   int
		want  string
	}{
		{"Name\tAge\tSex\t", 3, "Sex"},
		{"Name\tAge\tSex\t", 0, "Name\tAge\tSex\t"},
		{"Name\tAge\tSex", 4, ""},
	}

	for _, test := range tests {
		if got := getColumn(test.input, test.col); got != test.want {
			t.Errorf("getColumn(%q, %d) = %q, want %q", test.input, test.col, got, test.want)
		}
	}
}

func TestPostProcessLines(t *testing.T) {
	tests := []struct {
		input  []string
		unique bool
		want   []string
	}{
		{
			input:  []string{"a", "b", "b"},
			unique: true,
			want:   []string{"a", "b"},
		},

		{
			input:  []string{"a", "b", "b"},
			unique: false,
			want:   []string{"a", "b", "b"},
		},

		{
			input:  []string{},
			unique: true,
			want:   []string{},
		},

		{
			input:  []string{},
			unique: false,
			want:   []string{},
		},
	}

	for _, test := range tests {
		if got := postProcessLines(test.input, test.unique); !reflect.DeepEqual(got, test.want) {
			t.Errorf("postProcessLines(%v, %v) = %v, want %v", test.input, test.unique, got, test.want)
		}
	}
}

func TestTrimRightSpaces(t *testing.T) {
	tests := []struct {
		input []string
		col   int
		want  []string
	}{
		{
			input: []string{"a ", "b ", "c "},
			col:   0,
			want:  []string{"a", "b", "c"},
		},

		{
			input: []string{
				"a\tb \tc",
				"x\ty \tz",
			},
			col: 2,
			want: []string{
				"a\tb\tc",
				"x\ty\tz",
			},
		},

		{
			input: []string{"a\tb\tc"},
			col:   1,
			want:  []string{"a\tb\tc"},
		},

		{
			input: []string{"a \tb \tc "},
			col:   4,
			want:  []string{"a \tb \tc "},
		},
	}

	for _, test := range tests {
		if got := trimRightSpaces(test.input, test.col); !reflect.DeepEqual(got, test.want) {
			t.Errorf("trimRightSpaces(%v, %d) = %v, want %v", test.input, test.col, got, test.want)
		}
	}
}

func TestPreProcessLines(t *testing.T) {
	tests := []struct {
		input  []string
		col    int
		ignore bool
		want   []string
	}{
		{
			input:  []string{},
			col:    0,
			ignore: true,
			want:   []string{},
		},
	}

	for _, test := range tests {
		if got := preProcessLines(test.input, test.col, test.ignore); !reflect.DeepEqual(got, test.want) {
			t.Errorf("preProcessLines(%v, %d, %v) = %v, want %v", test.input, test.col, test.ignore, got, test.want)
		}
	}
}

func TestLess(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		opts  Options
		i, j  int
		want  bool
	}{
		{
			name:  "alphabetical ascending",
			input: []string{"apple", "banana"},
			opts:  Options{},
			i:     0,
			j:     1,
			want:  true,
		},

		{
			name:  "alphabetical reverse",
			input: []string{"apple", "banana"},
			opts:  Options{Reverse: true},
			i:     0,
			j:     1,
			want:  false,
		},

		{
			name: "alphabetical column",
			input: []string{
				"a\tb",
				"c\td",
			},
			opts: Options{Column: 2},
			i:    0,
			j:    1,
			want: true,
		},

		{
			name: "alphabetical column reverse",
			input: []string{
				"a\tb",
				"c\td",
			},
			opts: Options{Column: 2, Reverse: true},
			i:    0,
			j:    1,
			want: false,
		},

		{
			name:  "numeric ascending",
			input: []string{"10", "12"},
			opts:  Options{Numeric: true},
			i:     0,
			j:     1,
			want:  true,
		},

		{
			name:  "numeric reverse",
			input: []string{"10", "12"},
			opts:  Options{Numeric: true, Reverse: true},
			i:     0,
			j:     1,
			want:  false,
		},

		{
			name: "numeric column",
			input: []string{
				"10\t12",
				"14\t16",
			},
			opts: Options{Numeric: true, Column: 2},
			i:    0,
			j:    1,
			want: true,
		},

		{
			name: "numeric columnd reverse",
			input: []string{
				"10\t12",
				"14\t16",
			},
			opts: Options{Numeric: true, Column: 2, Reverse: true},
			i:    0,
			j:    1,
			want: false,
		},

		{
			name:  "month comparison",
			input: []string{"jan", "may"},
			opts:  Options{MonthSort: true},
			i:     0,
			j:     1,
			want:  true,
		},

		{
			name:  "month reverse",
			input: []string{"jan", "may"},
			opts:  Options{MonthSort: true, Reverse: true},
			i:     0,
			j:     1,
			want:  false,
		},

		{
			name: "month column",
			input: []string{
				"bla\tjan",
				"blabla\taug",
			},
			opts: Options{MonthSort: true, Column: 2},
			i:    0,
			j:    1,
			want: true,
		},

		{
			name: "month column reverse",
			input: []string{
				"bla\tjan",
				"blabla\taug",
			},
			opts: Options{MonthSort: true, Column: 2, Reverse: true},
			i:    0,
			j:    1,
			want: false,
		},

		{
			name:  "human readable numeric comparison",
			input: []string{"10K", "12K"},
			opts:  Options{HumanNumeric: true},
			i:     0,
			j:     1,
			want:  true,
		},

		{
			name:  "human readable numeric reverse",
			input: []string{"10K", "12K"},
			opts:  Options{HumanNumeric: true, Reverse: true},
			i:     0,
			j:     1,
			want:  false,
		},

		{
			name: "human readable numeric column",
			input: []string{
				"bla\t10K",
				"bl\t12K",
			},
			opts: Options{HumanNumeric: true, Column: 2},
			i:    0,
			j:    1,
			want: true,
		},

		{
			name: "human readable numeric column reverse",
			input: []string{
				"bl\t10K",
				"bla\t12K",
			},
			opts: Options{HumanNumeric: true, Column: 2, Reverse: true},
			i:    0,
			j:    1,
			want: false,
		},
	}

	for _, test := range tests {
		s := SortableLines{lines: test.input, options: test.opts}
		if got := s.Less(test.i, test.j); got != test.want {
			t.Errorf("%s: Less(%d, %d) = %v, want %v", test.name, test.i, test.j, got, test.want)
		}
	}
}

func TestCheckSorted(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  bool
	}{
		{
			name:  "sorted",
			input: []string{"apple", "banana"},
			want:  true,
		},

		{
			name:  "unsorted",
			input: []string{"banana", "apple"},
			want:  false,
		},
	}

	for _, test := range tests {
		s := SortableLines{lines: test.input}
		if got := checkSorted(s); got != test.want {
			t.Errorf("%s: checkSorted(s) = %v, want %v", test.name, got, test.want)
		}
	}
}
