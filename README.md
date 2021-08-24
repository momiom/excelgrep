# excelgrep (eg)

[![go:1.15](https://img.shields.io/badge/GO-1.15-blue?style=flat-square&logo=go)](https://github.com/golang/go)

excelgrep は指定のディレクトリから再帰的に excel ファイル（.xlsx） を見つけ出し、文字列パターンを検索するツールです。  
簡単に言えば以下のようなコマンドを実行できます。

```:sh
find /some/path -type f -name "*.xlsx" -not -name "~\$*.xlsx" | grep "some word"
```

残念ながら現在は find コマンドや grep コマンドと同等のオプションがあるわけではありません。  
[今後のロードマップはこちら。](https://github.com/momiom/excelgrep/projects/1#column-12615916)


## DEMO


## Usage


## Install

excelgrep のバイナリ名は eg です。[こちらから Windows、macOS 向けのコンパイル済みバイナリをダウンロードできます。](https://github.com/momiom/excelgrep/releases)

インストーラーはありません。必要に応じて任意の場所でパスを通してください。

## Build
excelgrep は Go 1.15 で実装されているためビルドには同様の実行環境が必要です。

```
$ go version
go version go1.15.5 windows/amd64

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
