package latinify

import (
	"testing"
)

func TestStringWithTable(t *testing.T) {
	type args struct {
		source string
		tables []Table
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"teste remove acento", args{"tbó", []Table{}}, "tbo", false},
		{"teste mapa", args{"casa", []Table{
			Table{
				's': 'r',
			},
		}}, "cara", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StringWithTable(tt.args.source, tt.args.tables...)
			if (err != nil) != tt.wantErr {
				t.Errorf("StringWithTable() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("StringWithTable() = %v, want %v", got, tt.want)
			}
		})
	}
}

// func TestGenerateTable(t *testing.T) {

// 	for i := 0x0400; i <= 0x052F; i++ {
// 		r := rune(i)
// 		v, _ := String(string(r))
// 		if len(v) > 0 && !unicode.IsLetter(rune(v[0])) {
// 			v = ""
// 		}
// 		fmt.Printf("%q: '%s'\n", r, v)
// 	}
// 	fmt.Println("// ----- greek table")
// 	for i := 0x0370; i <= 0x03FF; i++ {
// 		r := rune(i)
// 		v, _ := String(string(r))
// 		if len(v) > 0 && !unicode.IsLetter(rune(v[0])) {
// 			v = ""
// 		}

// 		fmt.Printf("%q: '%s'\n", r, v)
// 	}
// 	fmt.Println("// ----- Basic Latin")
// 	for i := 0x0000; i <= 0x007F; i++ {
// 		r := rune(i)
// 		v, _ := String(string(r))
// 		if len(v) > 0 && !unicode.IsLetter(rune(v[0])) {
// 			v = ""
// 		}
// 		fmt.Printf("%q: '%s'\n", r, v)
// 	}
// 	fmt.Println("// ----- Latin-1 Supplement")
// 	for i := 0x0080; i <= 0x00FF; i++ {
// 		r := rune(i)
// 		v, _ := String(string(r))
// 		if len(v) > 0 && !unicode.IsLetter(rune(v[0])) {
// 			v = ""
// 		}
// 		fmt.Printf("%q: '%s'\n", r, v)
// 	}
// 	fmt.Println("// ----- Latin Extended-A")
// 	for i := 0x0100; i <= 0x017F; i++ {
// 		r := rune(i)
// 		v, _ := String(string(r))
// 		if len(v) > 0 && !unicode.IsLetter(rune(v[0])) {
// 			v = ""
// 		}
// 		fmt.Printf("%q: '%s'\n", r, v)
// 	}
// 	fmt.Println("// ----- Latin Extended-B")
// 	for i := 0x0180; i <= 0x024F; i++ {
// 		r := rune(i)
// 		v, _ := String(string(r))
// 		if len(v) > 0 && !unicode.IsLetter(rune(v[0])) {
// 			v = ""
// 		}
// 		fmt.Printf("%q: '%s'\n", r, v)
// 	}
// 	fmt.Println("// ----- IPA Extensions")
// 	for i := 0x0250; i <= 0x02AF; i++ {
// 		r := rune(i)
// 		v, _ := String(string(r))
// 		if len(v) > 0 && !unicode.IsLetter(rune(v[0])) {
// 			v = ""
// 		}
// 		fmt.Printf("%q: '%s'\n", r, v)
// 	}
// 	t.Fail()
// }
