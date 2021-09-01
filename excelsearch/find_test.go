package excelsearch

import (
	"reflect"
	"sort"
	"testing"
)

func TestFind(t *testing.T) {
	tests := []struct {
		name    string
		args    string
		want    []string
		wantErr string
	}{
		{
			name:    "存在しないパスを指定したとき、空の配列が返されること",
			args:    "../fakepath/",
			want:    []string{},
			wantErr: "",
		},
		{
			name:    "xlsxファイルを指定したとき、指定ファイル１つが入った配列が返されること",
			args:    "../testdata/sample_files/sample.xlsx",
			want:    []string{"../testdata/sample_files/sample.xlsx/"},
			wantErr: "",
		},
		{
			name: "ディレクトリで終わるパスを指定したとき、指定ディレクトリ以下の全ての*.xlsxファイルが返されること",
			args: "../testdata/",
			want: []string{
				"../testdata/sample_files/sample.xlsx",
				"../testdata/sample_files/sample2.xlsx",
				"../testdata/sample_files/sub_dir/sample.xlsx",
				"../testdata/sample_files/sub_dir/sample2.xlsx",
			},
			wantErr: "",
		},
		{
			name: "**をパスに指定したとき、指定ディレクトリ以下の全ての*.xlsxファイルが返されること",
			args: "../testdata/**",
			want: []string{
				"../testdata/sample_files/sample.xlsx",
				"../testdata/sample_files/sample2.xlsx",
				"../testdata/sample_files/sub_dir/sample.xlsx",
				"../testdata/sample_files/sub_dir/sample2.xlsx",
			},
			wantErr: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Find(tt.args)
			if err != nil {
				if err.Error() != tt.wantErr {
					t.Errorf("Find() error = %T : %v, wantErr %T : %v", err, err, tt.wantErr, tt.wantErr)
				}
				return
			}

			sort.Strings(got)
			sort.Strings(tt.want)

			if !reflect.DeepEqual(got, tt.want) {
				if len(got) != len(tt.want) {
					t.Errorf("\nFind()\t%T : %v\nwant\t%T : %v\n", got, got, tt.want, tt.want)
				}
			}
		})
	}
}
