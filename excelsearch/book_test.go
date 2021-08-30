package excelsearch

import "testing"

func TestToString(t *testing.T) {
	f := found{CellName: "A1", Found: "abc"}             // 検索パターンを含む文字列
	s1 := sheet{SheetName: "sheet1", Founds: []found{f}} // 検索にヒットしたシート
	s2 := sheet{SheetName: "sheet2", Founds: []found{}}  // 検索にヒットしなかったシート(出力されない)
	b := Book{BookName: "book1", Sheets: []sheet{s1, s2}}

	expect := "book1\n\tsheet1\n\t\tA1 : abc\n"

	if b.ToString() != expect {
		t.Errorf("%s", b.ToString())
	}
}
