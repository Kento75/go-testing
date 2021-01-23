# golang でのテスト

## 参考リンク

https://golang.org/pkg/testing/

## ファイルの命名規則

テストコードを作成する場合、 `aaa_test.go` という命名ルールに従う必要がある。  
`go test` コマンド実行時にテストコードとして認識するためのプレフィックスになっている。  
基本的には、テスト対象のファイル名 + \_test でテストコードは作成する。

## テスト用パッケージ

テストの作成には、 標準パッケージの `testing` と `testify` を使用することが多いらしい。  
`testify` には、assertion、Mocking、Test Suite と揃っている。

## テストファイルを作成する

対象ファイルは、以下のファイル。

```sort.go
package sort

import (
	"sort"
)

func BubbleSort(elements []int) {
	keepWorking := true

	for keepWorking {
		keepWorking = false

		for i := 0; i < len(elements)-1; i++ {
			if elements[i] < elements[i+1] {
				keepWorking = true
				elements[i], elements[i+1] = elements[i+1], elements[i]
			}
		}
	}
}

func Sort(elements []int) {
	sort.Ints(elements)
}
```

テストには、 `testing` `testify` を使用する。

```sort_test.go
package sort

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBubbleSortOrderDESC(t *testing.T) {
	// Arrange
	elements := []int{9, 7, 5, 3, 1, 2, 4, 6, 8, 0}
	expected := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}

	// Act
	BubbleSort(elements)

	// Assert
	assert.EqualValues(t, expected, elements, "ソート後、降順であること") // 同じ値かつ、同じ型
}

func TestBubbleSortAlreadySorted(t *testing.T) {
	// Arrange
	elements := []int{5, 4, 3, 2, 1, 0}
	expected := []int{5, 4, 3, 2, 1, 0}

	// Act
	BubbleSort(elements)

	// Assert
	assert.EqualValues(t, expected, elements, "ソート後、順番に変更がないこと") // 同じ値かつ、同じ型
}

```

## カバレッジレポートの出力

golang のちょっと便利な機能として、カバレッジが標準でサポートされているところ。  
`-cover` オプションをつけることでテスト実行結果にカバレッジ情報が付与される。  
`-coverprofile` オプションでレポートを出力できる。

```
[kento@KentonoMacBook-Pro tdd-go]$ go test -cover -coverprofile cover.out ./sort
ok      src/github.com/Kento75/tdd-go/sort      0.298s  coverage: 87.5% of statements
```

レポートの内容はわかりづらい。

```cover.out
mode: set
src/github.com/Kento75/tdd-go/sort/sort.go:7.33,10.18 2 1
src/github.com/Kento75/tdd-go/sort/sort.go:10.18,13.40 2 1
src/github.com/Kento75/tdd-go/sort/sort.go:13.40,14.35 1 1
src/github.com/Kento75/tdd-go/sort/sort.go:14.35,17.5 2 1
src/github.com/Kento75/tdd-go/sort/sort.go:22.27,24.2 1 0

```

なので、カバレッジ情報を HTML として出力する。
(出力内容は cover.png 参照)

```
$ go tool cover -html=cover.out -o cover.html
```

## おまけ

`testing` にはベンチマークを測れる機能がある。  
先ほどの sort.go のテストファイルを以下のものに修正。  
バブルソートとソートのベンチマークを計測。

```
package sort

import "testing"

// ベンチマーク
func BenchmarkBubbleSort(b *testing.B) {
	// Arrange
	elements := getElements(1000000)

	// Act
	for i := 0; i < b.N; i++ {
		BubbleSort(elements)
	}
}

// ベンチマーク
func BenchmarkSort(b *testing.B) {
	// Arrange
	elements := getElements(1000000)
	// Act
	for i := 0; i < b.N; i++ {
		Sort(elements)
	}
}

// test helper
func getElements(n int) []int {
	result := make([]int, n)
	j := 0
	for i := n; i > 0; i-- {
		result[j] = i
		j++
	}
	return result
}

func getElementsForASC(n int) []int {
	result := make([]int, n)
	for i := 0; i < n; i++ {
		result[i] = i + 1
	}
	return result
}

```

ループの回数とループ 1 回あたりの処理時間が出力される。

```
[kento@KentonoMacBook-Pro tdd-go]$ go test -v -bench=. ./sort/
goos: darwin
goarch: amd64
pkg: src/github.com/Kento75/tdd-go/sort
BenchmarkBubbleSort
BenchmarkBubbleSort-4               1146            963343 ns/op
BenchmarkSort
BenchmarkSort-4                       15          76824727 ns/op
PASS
ok      src/github.com/Kento75/tdd-go/sort      3.730s
```

`-benchmem` オプションをつけるとメモリに関するベンチマークも取得できる。

```
[kento@KentonoMacBook-Pro tdd-go]$ go test -v -bench=. -benchmem ./sort/
goos: darwin
goarch: amd64
pkg: src/github.com/Kento75/tdd-go/sort
BenchmarkBubbleSort
BenchmarkBubbleSort-4               1066            945447 ns/op            7508 B/op          0 allocs/op
BenchmarkSort
BenchmarkSort-4                       14          75791738 ns/op          571716 B/op          1 allocs/op
PASS
ok      src/github.com/Kento75/tdd-go/sort      3.321s
```
