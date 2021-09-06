package excelsearch

import (
	"encoding/json"
	"testing"
)

func TestToString(t *testing.T) {
	tests := []struct {
		name string
		book Book
		want string
	}{
		{
			"複数のシートのうち、検索にヒットした文字列が含まれるシート名のみ出力されること",
			Book{
				BookName: "book1",
				Sheets: []sheet{
					// sheet1 には1件検索にヒットした文字列が含まれる
					{
						SheetName: "sheet1",
						Founds: []found{
							{CellName: "A1", Found: "abc"},
						},
					},
					// sheet2 には検索にヒットした文字列が存在しない
					{
						SheetName: "sheet2",
						Founds:    []found{},
					},
				},
			},
			"book1\n\tsheet1\n\t\tA1 : abc\n",
		},
		{
			"ブックに検索にヒットした文字列が含まれない場合はブック名が出力されないこと",
			Book{
				BookName: "book1",
				Sheets: []sheet{
					// sheet1 には1件検索にヒットした文字列が含まれる
					{
						SheetName: "sheet1",
						Founds:    []found{},
					},
					// sheet2 には検索にヒットした文字列が存在しない
					{
						SheetName: "sheet2",
						Founds:    []found{},
					},
				},
			},
			"",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.book.ToString()
			if got != tt.want {
				g, _ := json.Marshal(got)
				t.Errorf("expect: %s, got: %s", tt.want, g)
			}
		})
	}
}

func TestToJson(t *testing.T) {
	f := found{CellName: "A1", Found: "abc"}             // 検索パターンを含む文字列
	s1 := sheet{SheetName: "sheet1", Founds: []found{f}} // 検索にヒットしたシート
	s2 := sheet{SheetName: "sheet2", Founds: []found{}}  // 検索にヒットしなかったシート(出力されない)
	b := Book{BookName: "book1", Sheets: []sheet{s1, s2}}

	expect := `{"book_name":"book1","sheets":[{"sheet_name":"sheet1","Founds":[{"cell_name":"A1","found":"abc"}]},{"sheet_name":"sheet2","Founds":[]}]}`

	s, err := b.ToJson()

	if err != nil {
		t.Fatal(err)
	}

	if s != expect {
		t.Errorf("%s", s)
	}
}
