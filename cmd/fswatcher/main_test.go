package main

import "testing"

func TestDisplayName(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// Ordinary names are printed unchanged.
		{"", ""},
		{"/tmp/foo", "/tmp/foo"},
		{"/tmp/a b", "/tmp/a b"},
		{`C:\Users\foo\bar.txt`, `C:\Users\foo\bar.txt`},
		{`/tmp/"quoted"`, `/tmp/"quoted"`},
		{"/tmp/日本語.txt", "/tmp/日本語.txt"},
		{"/tmp/\uFFFD", "/tmp/\uFFFD"},
		{"/tmp/x\xe9y", "/tmp/x\xe9y"},

		// C0 controls and DEL.
		{"/tmp/a\nCREATE\t/etc/passwd", `"/tmp/a\nCREATE\t/etc/passwd"`},
		{"/tmp/a\tb", `"/tmp/a\tb"`},
		{"/tmp/a\rb", `"/tmp/a\rb"`},
		{"/tmp/\x1b[2J", `"/tmp/\x1b[2J"`},
		{"/tmp/\x1b]0;title\x07", `"/tmp/\x1b]0;title\a"`},
		{"/tmp/\x7f", `"/tmp/\x7f"`},
		{"C:\\a\nb", `"C:\\a\nb"`},

		// C1 controls encoded as UTF-8.
		{"/tmp/x\u009b2J", `"/tmp/x\u009b2J"`},

		// Raw 8-bit C1 bytes that are not valid UTF-8.
		{"/tmp/x\x9b2J", `"/tmp/x\x9b2J"`},
		{"/tmp/x\x9d0;t\x9c", `"/tmp/x\x9d0;t\x9c"`},
		{"/tmp/\x82\xa0", `"/tmp/\x82\xa0"`}, // Shift_JIS "あ"
	}
	for _, tt := range tests {
		if got := displayName(tt.name); got != tt.want {
			t.Errorf("displayName(%q) = %q, want %q", tt.name, got, tt.want)
		}
	}
}
