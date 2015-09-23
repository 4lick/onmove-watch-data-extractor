package main

import ( 
   "testing"          
   "extractor/parser"
) 

func TestOnMoveDateExtractorParseFilename(t *testing.T) {	
	cases := []struct {
		in, want string
	}{
		{"DCKD2647", "0013-Dec-20 13:26:47 13"},
	}
	for _, c := range cases {
		got := parser.ParseFilename(c.in).Format("2006-Jan-02 15:04:05 06") 
		if got != c.want {
			t.Errorf("Parse(%q) == %q, want %q", c.in, got, c.want)
		}
	}  
}  
