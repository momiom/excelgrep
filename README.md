# excelgrep (eg)

excelgrep は指定のディレクトリから再帰的に excel ファイル（.xlsx） を見つけ出し、文字列パターンを検索するツールです。  
簡単に言えば以下のようなコマンドを実行できます。

```:sh
find /some/path -type f -name "*.xlsx" -not -name "~\$*.xlsx" | grep "some word"
```

残念ながら現在は find コマンドや grep コマンドと同等のオプションがあるわけではありません。  

## DEMO

WIP
## Usage

WIP

## Install

[こちらから Windows、macOS 向けのコンパイル済みバイナリをダウンロードできます。](https://github.com/momiom/excelgrep/releases)

excelgrep のバイナリ名は eg です。
インストーラーはありません。必要に応じて任意の場所でパスを通してください。

## Build

```
$ go version
go version go1.16.6 darwin/amd64

$ git clone https://github.com/momiom/excelgrep.git
$ cd excelgrep
$ go mod tify
$ go build -ldflags "-s -w" -o eg
$ ./eg --version
version: unknown
revision: unknown
```

`--version` で出力される情報は以下のようにビルドすることで任意に設定できます。
```
go build -ldflags "-s -w -X github.com/momiom/excel_grep/cmd.Version=x.x.x -X github.com/momiom/excel_grep/cmd.Revision=xxx" -o eg
```
