package excelsearch

import (
	"fmt"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/fatih/color"
	"github.com/momiom/excelgrep/logger"
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

// xlsx ファイルの検索結果を表します。
type Book struct {

	// Book名（ファイル名）
	BookName string `json:"book_name"`

	// １つの Book に含まれるすべてのシートのスライス
	Sheets []sheet `json:"sheets"`
}

// シートの検索結果を表します。
type sheet struct {

	// シート名
	SheetName string `json:"sheet_name"`

	// セルの検索結果のスライス
	Founds []found `jsoon:"founds"`
}

// １つのシートに含まれる全セルの検索結果を表す
type found struct {

	// A1形式のセル番地
	CellName string `json:"cell_name"`

	// 検索にヒットしたセルの文字列（一致箇所をハイライトしたセルの全文）
	Found string `json:"found"`
}

// xlsxPath で指定された xlsx ファイルから sep に一致する箇所を検索
func Grep(xlsxPath string, sep string) (Book, error) {
	logger.Debugf("Exec grep %s %s", sep, xlsxPath)

	if len(xlsxPath) == 0 {
		logger.Debugln("Done. xlsxPath is empty.")
		return Book{}, nil
	}

	// ファイルオープン
	f, err := excelize.OpenFile(xlsxPath)
	if err != nil {
		logger.Debugln("Done. Has Error.")
		return Book{}, err
	}

	// 各シート毎に全セルを検索
	var sheets []sheet
	for _, sheetName := range f.GetSheetList() {
		// シートの全セルを取得
		cols, err := f.GetCols(sheetName)
		if err != nil {
			logger.Debugln(err)
			continue
		}

		// 文字列検索
		var founds []found
		for colNum, col := range cols {
			for rowNum, rowCell := range col {
				if len(rowCell) != 0 {
					cellName, err := excelize.CoordinatesToCellName(colNum+1, rowNum+1)
					if err != nil {
						logger.Debugln(err)
						continue
					}

					foundWord := ""
					if s := strings.TrimSpace(rowCell); s != "" {
						foundWord = search(s, sep)
					}

					if foundWord != "" {
						findResult := found{
							CellName: cellName,
							Found:    foundWord,
						}
						founds = append(founds, findResult)
					}
				}
			}
		}

		if len(founds) > 0 {
			s := sheet{
				SheetName: sheetName,
				Founds:    founds,
			}
			sheets = append(sheets, s)
		}
	}

	var book Book
	if len(sheets) > 0 {
		book = Book{
			BookName: xlsxPath,
			Sheets:   sheets,
		}
	}

	logger.Debugln("Done.")
	return book, nil
}

func search(s string, sep string) string {
	if res := strings.Index(s, sep); res > -1 {
		// 一致箇所のみハイライトする
		highlight := color.New(color.FgRed).Add(color.Bold).Sprint(s[res : res+len(sep)])
		before := string(s[:res])
		after := string(s[res+len(sep):])
		return fmt.Sprintf("%s%s%s\n", before, highlight, after)
	}
	return ""
}
