package main

import (
	"bufio"
	"bytes"
	"strings"
	"testing"
)

func TestEnum(t *testing.T) {
	for _, tt := range []struct {
		left      bool
		none      bool
		separator string
		in        string
		want      string
	}{
		{
			false,
			false,
			". ",
			"foo",
			"1. foo\n",
		},
		{
			false,
			false,
			". ",
			"foo\n",
			"1. foo\n",
		},
		{
			false,
			false,
			". ",
			"foo\nbar\n",
			"1. foo\n2. bar\n",
		},
		{
			false,
			false,
			". ",
			"foo\nbar\nfoo\nbar\nfoo\nbar\nfoo\nbar\nfoo\nbar\nfoo\nbar\n",
			" 1. foo\n 2. bar\n 3. foo\n 4. bar\n 5. foo\n 6. bar\n 7. foo\n 8. bar\n 9. foo\n10. bar\n11. foo\n12. bar\n",
		},
		{
			true,
			false,
			". ",
			"foo\nbar\nfoo\nbar\nfoo\nbar\nfoo\nbar\nfoo\nbar\nfoo\nbar\n",
			"1.  foo\n2.  bar\n3.  foo\n4.  bar\n5.  foo\n6.  bar\n7.  foo\n8.  bar\n9.  foo\n10. bar\n11. foo\n12. bar\n",
		},
		{
			false,
			true,
			". ",
			"foo\nbar\nfoo\nbar\nfoo\nbar\nfoo\nbar\nfoo\nbar\nfoo\nbar\n",
			"1. foo\n2. bar\n3. foo\n4. bar\n5. foo\n6. bar\n7. foo\n8. bar\n9. foo\n10. bar\n11. foo\n12. bar\n",
		},
		{
			false,
			false,
			"",
			"foo\nbar\n",
			"1foo\n2bar\n",
		},
	} {
		s := bufio.NewScanner(strings.NewReader(tt.in))
		var w bytes.Buffer

		enumerate(s, &w, tt.left, tt.none, tt.separator)

		have := w.String()

		if have != tt.want {
			t.Errorf("want: %s\nhave: %s", tt.want, have)
		}
	}
}

func TestValidOptions(t *testing.T) {
	if enumerate(nil, nil, true, true, "") == nil {
		t.Error("Missing expected error")
	}
}
