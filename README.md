# excelgrep (eg)

excelgrep は Excel ファイルの横断検索ツールです。  
指定のディレクトリから再帰的に excel ファイル（\*.xlsx） を見つけ出し、文字列パターンを検索することができます。

以下のコマンドを実行できるイメージです。

```:sh
find /some/path -type f -name "*.xlsx" -not -name "~\$*.xlsx" | grep "some word"
```

ただし現在は find コマンドや grep コマンドと同等の機能はありません。  

## Screenshots

![usage](https://github.com/momiom/excelgrep/blob/main/images/usage.png?raw=true)

## Usage

```
eg [OPTIONS] PATTERN [PATH ...]
PATH | eg [OPTIONS] PATTERN
```
<dl>
  <dt>&lt;PATTERN&gt;</dt>
  <dd>検索する文字列を指定します。</dd>
  
  <dt>&lt;PATH&gt; (オプション)</dt>
  <dd>xlsx ファイルやディレクトリのパスを指定します。ワイルドカードとして <code>*</code> を利用できます。<br />
    スペース区切りで複数指定できます。</dd> 
</dl>

## Installation

[こちらから Windows、macOS 向けのコンパイル済みバイナリをダウンロードできます。](https://github.com/momiom/excelgrep/releases)

excelgrep のバイナリ名は eg または eg.exe です。  
インストーラーはありませんので、必要に応じて任意の場所でパスを通してください。

## Building

excelgrep は Go で記述されているため、コンパイルするには Go をインストールする必要があります。  
Go 1.15.0 以降でのビルドをお勧めします。

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
