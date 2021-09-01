package excelsearch

import (
	"encoding/json"
	"fmt"
)

/*
[
	{
		sheetName: "Sheet1",
		founds: {
			cellName: "A1",
			found: "検索文字列"
		}
	}
]
*/

// xlsx ファイルの検索結果
type Book struct {

	// Book名（ファイル名）
	BookName string `json:"book_name"`

	// １つの Book に含まれるすべてのシートのスライス
	Sheets []sheet `json:"sheets"`
}

// シートの検索結果
type sheet struct {

	// シート名
	SheetName string `json:"sheet_name"`

	// セルの検索結果のスライス
	Founds []found `jsoon:"founds"`
}

// １つのシートに含まれる全セルの検索結果
type found struct {

	// A1形式のセル番地
	CellName string `json:"cell_name"`

	// 検索にヒットしたセルの文字列（一致箇所をハイライトしたセルの全文）
	Found string `json:"found"`
}

func NewBook(bookName string) Book {
	return Book{BookName: bookName, Sheets: []sheet{}}
}

func (b *Book) AppendSheet(sheetName string) {
	b.Sheets = append(b.Sheets, sheet{SheetName: sheetName, Founds: []found{}})
}

func (b *Book) AppendFound(cellName string, foundWord string) {
	f := found{
		CellName: cellName,
		Found:    foundWord,
	}
	s := &b.Sheets[len(b.Sheets)-1]
	s.Founds = append(s.Founds, f)
}

func (b Book) ToString() string {
	var str string
	if len(b.Sheets) > 0 {
		str = fmt.Sprintf("%s\n", b.BookName)

		for _, s := range b.Sheets {
			if len(s.Founds) <= 0 {
				continue
			}

			str = fmt.Sprintf("%s\t%s\n", str, s.SheetName)
			for _, f := range s.Founds {
				str = fmt.Sprintf("%s\t\t%s : %s\n", str, f.CellName, f.Found)
			}
		}
	}

	return str
}

func (b Book) ToJson() (string, error) {
	s, err := json.Marshal(b)
	if err != nil {
		return "", err
	}
	return string(s), nil
}
