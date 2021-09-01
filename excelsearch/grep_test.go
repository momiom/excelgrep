package excelsearch

import (
	"encoding/json"
	"testing"

	"github.com/momiom/excelgrep/logger"
)

func setup() {
	// logger.SetLogger(logger.Verbose)
	logger.SetLogger(logger.Silent)
}

func TestGrep(t *testing.T) {
	setup()

	type args struct {
		xlsxPath string
		sep      string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr string
	}{
		{
			"case 1",
			args{"../testdata/sample_files/sample.xlsx", "abc"},
			`{"book_name":"../testdata/sample_files/sample.xlsx","sheets":[{"sheet_name":"Sheet1","Founds":[{"cell_name":"A1","found":"abc\n"}]},{"sheet_name":"しーとに","Founds":[]}]}`,
			"",
		},
		{
			"case 2",
			args{"../testdata/sample_files/sample2.xlsx", "abc"},
			`{"book_name":"../testdata/sample_files/sample2.xlsx","sheets":[{"sheet_name":"Sheet1","Founds":[{"cell_name":"A1","found":"abc\n"}]},{"sheet_name":"しーとに","Founds":[{"cell_name":"C1","found":"abc\n"}]}]}`,
			"",
		},
		{
			"case 3",
			args{"../testdata/not_exists_dir/sample.xlsx", "abc"},
			"",
			"open ../testdata/not_exists_dir/sample.xlsx: no such file or directory",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Grep(tt.args.xlsxPath, tt.args.sep)
			if err != nil {
				if err.Error() != tt.wantErr {
					t.Errorf("Grep() error = %+v, wantErr %+v", err.Error(), tt.wantErr)
				}
				return
			}

			g, _ := json.Marshal(got)
			// w, _ := json.Marshal(tt.want)

			if string(g) != tt.want {
				// t.Errorf("\ngot = %+v \nwant = %+v", got, tt.want)
				t.Errorf("\ngot = %s \nwant = %s", string(g), tt.want)
			}
		})
	}
}
