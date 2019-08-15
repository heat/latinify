package latinify

import "testing"

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
			"slug a simple name",
			args{"onezino gabriel Moreira"},
			"onezino-gabriel-moreira",
			false,
		},
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
