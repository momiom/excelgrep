package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"sync"

	"github.com/momiom/excelgrep/excelsearch"
	"github.com/momiom/excelgrep/logger"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh/terminal"
)

// メインコマンド
// コマンドに指定されたパスのバリデーション
// stdin(|) で渡されたパスのバリデーション
// find 処理
// grep 処理
func runRootCmd(cmd *cobra.Command, args []string) {

	// traceFile, err := os.Create("trace_proc2.out")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// defer traceFile.Close()

	// trace.Start(traceFile)
	// defer trace.Stop()

	// 引数パターン
	// ---- ↓↓↓↓コマンド引数 ----
	// eg word /any/path
	// eg word /any/path /any/path2
	// eg word ./some*/*some*.xlsx
	// ---- ↓↓↓pipe
	// find /some/path -name "*.xlsx" | eg word

	// コマンド引数やパイプで渡されたパターンとパスを取得する
	pattern, paths, err := getArgs(args)
	if err != nil {
		cmd.Println(err.Error())
	}
	logger.Debugf("pattern: %s", pattern)
	logger.Debugf("paths: %v", paths)

	// find
	logger.Debugln("find part START")
	var finds []string
	for _, p := range paths {
		find, err := excelsearch.Find(p)
		if err != nil {
			cmd.Println(err.Error())
		}
		finds = append(finds, find...)
	}
	logger.Debugln("find part DONE")

	// grep
	logger.Debugln("grep part START")
	wg := new(sync.WaitGroup)
	wg.Add(len(finds))

	// goroutine の上限を管理する channel
	sem := make(chan struct{}, 1)
	for i, f := range finds {
		logger.Debugf("Loop %d: file: %s", i, f)

		// channel に空データを入れてバッファを埋める
		sem <- struct{}{}
		go func(f string) {
			logger.Debugln("go func START")
			defer wg.Done()

			// grep 処理
			res, err := excelsearch.Grep(f, pattern)
			logger.Debugln("excelsearch.Grep done")
			if err != nil {
				cmd.Println(err)
			}

			// 結果表示
			cmd.Println(res.ToString())
			logger.Debugln("print result done")

			// 空データを取り出してバッファを空ける
			<-sem
			logger.Debugln("go func END")
		}(f)
	}
}

func getArgs(args []string) (string, []string, error) {
	var pattern string
	var paths []string

	// コマンドの引数の処理
	if terminal.IsTerminal(int(os.Stdin.Fd())) {
		// コマンドライン引数のみ指定された場合（パイプでない）
		// 検索文字列: 1個, 検索パス: 0個以上
		// 検索パスの指定がない場合は . （カレントディレクトリ）とする
		switch len(args) {
		case 0:
			return "", nil, fmt.Errorf("Invalid args.")
		case 1:
			pattern = args[0]
			paths = append(paths, ".")
		default:
			pattern = args[0]
			paths = append(paths, args[1:]...)
		}
	} else {
		// パイプ(stdin)の処理
		// コマンドライン引数には検索文字列のみが指定されている
		pattern = args[0]
		b, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			return "", nil, fmt.Errorf("Failed to read stdin.")
		}
		in := strings.Split(string(b), "\n")
		in = in[:len(in)-1] // 最後の改行で空の要素が入ってしまうため削除
		paths = append(paths, in...)
	}
	return pattern, paths, nil
}
