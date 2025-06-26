package latinify

import (
	"reflect"
	"strings"
	"testing"
)

func TestSlugify(t *testing.T) {
	type args struct {
		source string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			"slug a simple com acentos",
			args{"onezino gá\"b®riel Moreira"},
			"onezino-gabriel-moreira",
			false,
		},
		{
			"slug a simple name",
			args{"onezino gabriel Moreira"},
			"onezino-gabriel-moreira",
			false,
		},
		{
			"collapse multiple spaces",
			args{"onezino   gabriel  Moreira"},
			"onezino-gabriel-moreira",
			false,
		},
		{
			"already_sluged",
			args{"already_sluged"},
			"already_sluged",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Slugify(tt.args.source)
			if (err != nil) != tt.wantErr {
				t.Errorf("Slugify() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Slugify() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_apply(t *testing.T) {
	type args struct {
		funcs []stringTransform
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"Lower case Text",
			args{
				[]stringTransform{
					strings.ToLower,
				},
			},
			"hello world",
		},
		{
			"Space removed",
			args{
				[]stringTransform{
					func(src string) string {
						return strings.ReplaceAll(src, " ", "")
					},
				},
			},
			"HelloWorld",
		},
	}
	for _, tt := range tests {
		text := "Hello World"
		t.Run(tt.name, func(t *testing.T) {

			if got := apply(tt.args.funcs...)(text); !reflect.DeepEqual(got, tt.want) {

				t.Errorf("apply() = %v, want %v", got, tt.want)
			}
		})
	}
}
