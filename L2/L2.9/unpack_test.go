package main

import "testing"

func TestUnpack(t *testing.T) {
	tests := []struct {
		input    string
		expected string
		wantErr  bool
	}{
		{`a4bc2d5e`, `aaaabccddddde`, false},
		{`abcd`, `abcd`, false},
		{``, ``, false},
		{`45`, "", true},

		{`qwe\4\5`, `qwe45`, false},
		{`qwe\45`, `qwe44444`, false},

		{`qwe\\5`, ``, true},
		{`qwe\`, ``, true},
		{`\`, ``, true},
		{`\\`, ``, true},
		{`qwe45`, ``, true},
		{`qwe\\\\4`, ``, true},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			got, err := unpack(test.input)
			if (err != nil) != test.wantErr {
				t.Errorf("expected err=%v, got err=%v", test.wantErr, err)
			}
			if got != test.expected {
				t.Errorf("expected=%q, got=%q\n", test.expected, got)
			}
		})
	}
}
