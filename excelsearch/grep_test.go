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
			`{"book_name":"../testdata/sample_files/sample.xlsx","sheets":[{"sheet_name":"Sheet1","Founds":[{"cell_name":"A1","found":"\u001b[31;1mabc\u001b[0m\n"}]}]}`,
			"",
		},
		{
			"case 2",
			args{"../testdata/sample_files/sample2.xlsx", "abc"},
			`{"book_name":"../testdata/sample_files/sample2.xlsx","sheets":[{"sheet_name":"Sheet1","Founds":[{"cell_name":"A1","found":"\u001b[31;1mabc\u001b[0m\n"}]},{"sheet_name":"しーとに","Founds":[{"cell_name":"C1","found":"\u001b[31;1mabc\u001b[0m\n"}]}]}`,
			"",
		},
		{
			"case 3",
			args{"../testdata/no_such_dir/sample.xlsx", "abc"},
			``,
			"open ../testdata/no_such_dir/sample.xlsx: no such file or directory",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Grep(tt.args.xlsxPath, tt.args.sep)
			if err != nil {
				if err.Error() != tt.wantErr {
					t.Errorf("Grep() error = %v, wantErr %v", err.Error(), tt.wantErr)
				}
				return
			}

			j, err := json.Marshal(got)
			if err != nil {
				t.Errorf("Failed marshal. error = %v", err)
			}
			gotJson := string(j)

			if gotJson != tt.want {
				t.Errorf("Grep() = %v, want %v", gotJson, tt.want)
			}
		})
	}
}
