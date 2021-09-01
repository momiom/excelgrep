package excelsearch

import (
	pathutil "path"
	"strings"

	"github.com/mattn/go-zglob"
	"github.com/momiom/excelgrep/logger"
)

func Find(path string) ([]string, error) {
	logger.Debugln("START")
	logger.Debugf("arg: %s", path)

	var (
		res []string
		err error
	)

	// パスが渡された場合、末尾に *.xlsx がなければを結合して再起的に Glob
	// パスが渡されなかった場合、カレントディレクトリ配下を再帰的に Glob
	if len(path) > 0 {
		if !strings.HasSuffix(path, ".xlsx") {
			path = pathutil.Join(path, "/**/*.xlsx")
		}
		logger.Debugf("exec glob: %s", path)
		res, err = zglob.GlobFollowSymlinks(path)
	} else {
		logger.Debugln("exec glob: ./**/*.xlsx")
		res, err = zglob.GlobFollowSymlinks("./**/*.xlsx")
	}
	if err != nil {
		return nil, err
	}

	var files []string
	for _, p := range res {
		f := pathutil.Base(p)
		if !strings.HasPrefix(f, "~$") {
			files = append(files, p)
		}
	}

	logger.Debugln("DONE")
	return files, nil
}
