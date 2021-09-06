package excelsearch

import (
	"fmt"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/fatih/color"
	"github.com/momiom/excelgrep/logger"
)

// xlsxPath で指定された xlsx ファイルから sep に一致する箇所を検索
func Grep(xlsxPath string, sep string) (Book, error) {
	logger.Debugf("Exec grep %s %s", sep, xlsxPath)

	if len(xlsxPath) == 0 {
		logger.Debugln("Done. xlsxPath is empty.")
		return Book{}, nil
	}

	result := NewBook(xlsxPath)

	// ファイルオープン
	f, err := excelize.OpenFile(xlsxPath)
	if err != nil {
		logger.Debugln("Done. Has Error.")
		return result, err
	}

	// 各シート毎に全セルを検索
	for _, sheetName := range f.GetSheetList() {
		// シートの全セルを取得
		rows, err := f.GetRows(sheetName)
		if err != nil {
			logger.Debugln(err)
			continue
		}

		result.AppendSheet(sheetName)

		// 文字列検索
		for rowNum, row := range rows {
			for colNum, colCell := range row {
				if len(colCell) != 0 {
					cellName, err := excelize.CoordinatesToCellName(colNum+1, rowNum+1)
					if err != nil {
						logger.Debugln(err)
						continue
					}

					if s := strings.TrimSpace(colCell); s != "" {
						foundWord := search(s, sep)
						if foundWord != "" {
							result.AppendFound(cellName, foundWord)
						}
					}
				}
			}
		}
	}

	logger.Debugln("Done.")
	return result, nil
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
