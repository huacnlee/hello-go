# 入门 Go 编写应用

![Golang](https://user-images.githubusercontent.com/5518/28311187-6a2eeca8-6be1-11e7-84d7-8f4ed896b314.png)

## ☝🏾 为何会选用 Go

- 无依赖的绿色 Native 二进制可执行文件发布；
- 跨平台支持，一次写完 macOS, Linux, Windows 均可以使用；
- 编写的应用程序启动速度快！
- 编写简单，工具链完善，语法简单，执行效率高（例如多核并发）；
- 并发编写非常好使；
- 完善的标准库，例如图形图像的库不是什么语言都有的；
- 完善可用的三方库，例如 [GoQuery](https://github.com/PuerkitoBio/goquery) 用于解析 HTML 结构；
- 由于编译以及 Go 的一些强制约定，使得写好的东西很稳定，哪怕没测试（基本上编译通过就没什么大问题了）。
- Gofmt 强制语法标准，完美控制，再也不用纠结语法风格，也不需要 Lint；

## 🌈 那些场景可以用上它

- 需要跨平台，绿色发布；
- 需要一个无环境依赖的场景；
- 大量计算型的工作；
- 用 Go 编写 CLI 应用真的非常爽！

## 📖 如何学习 Go 语言

阅读官方文档... （我是有段时间睡觉前手机上看的，看到自己睡着 😴）

- [Effective Go](https://golang.org/doc/effective_go.html)
- [The Go Programming Language Specification](https://golang.org/ref/spec)

## 📦 开发工具

- Visual Studio Code + [Go 扩展](https://marketplace.visualstudio.com/items?itemName=lukehoban.Go)
- [Gocode](https://github.com/nsf/gocode) - 代码分析 Auto Complete 的实现，完美
- Gofmt - 内置，自动化格式代码（一般在保存的时候自动调用）


### 🎁 Auto Complete 功能

![2017-07-18 5 46 53](https://user-images.githubusercontent.com/5518/28311092-24c82166-6be1-11e7-88a9-26d0733fe987.png)

提示绝对准确，并且响应速度高(1ms 以内)！

### 🎛 对于测试的支持

自动在测试用例部分增加点击按钮，可以单独跑测试。

![2017-07-18 5 47 43](https://user-images.githubusercontent.com/5518/28311127-3c215b2a-6be1-11e7-9949-a78d40ccbf91.png)


顺便演示一下 Ruby 里面，我们用 TextMate 也是用这样的方式跑测试的 (光标放在一个段落，Command + Shift + R 运行)

![be0228859bb293f3](https://user-images.githubusercontent.com/5518/28311160-531881b4-6be1-11e7-8966-1b880dc5cb6c.png)

## 🎈 Hello world 开始

http://github.com/huacnlee/hello-go

任何语言都是从 Hello world 开始的：

```go
// ~/work/src/github.com/huacnlee/hello-go
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello world.")
}
```

然后运行:

```bash
$ cd ~/work/src/github.com/huacnlee/hello-go
$ go run main.go
```

## 🛎 理解 Go 的几个必要基础

### GOPATH

```bas
$ export GOPATH=~/work
$ mkdir $GOPATH/src
```

最重要的环境变量，告诉 Go 的各种工具，工作路径是哪里

- import 函数需要 GOPATH 才能知道哪里能找到三方库代码
- 编译工具也需要
- 自己的项目，一般我们也放在里面，便于 import 自己，例如 `github.com/huacnlee/hello-go`


```bash
$ cd $GOPATH/src
$ tree -L 3
├── github.com
│   ├── PuerkitoBio
│   │   └── goquery
│   ├── huacnlee
│   │   ├── mediom
```


### 📜 语言特性

- 静态语言，需要忘掉动态语言特性，需要类型转换 `strconv.Atoi("100")`；
- 值类型／指针；
- Array 类型（slice) 的操作很原始，非常原始：`userIds = append(userIds, 100)` 参见 [SliceTricks](https://github.com/golang/go/wiki/SliceTricks)
- 各种常见返回的 error 类型；
- 你可能找不到 Class 类型，在 Go 里面是 `struct` 实现简单结构体；
- 一个文件夹可以看作一个包，同一个文件夹里面 .go 文件的 `package` 必须是一样的；
- 如果是应用程序，项目根目录 `package` 得是 `main`；
- 函数、属性、变量等命名首字母大写是 `公开` 的，小写是 `私有` 的（相对于 `package`）；
- 用不上的变量要去掉，用不上的 `import` 也要去掉，否则会 `xxx declared and not used`;

### 理解 defer 关键字

defer 是一个非常巧妙的设计，可以让你提前准备好要在函数结束（return）是需要做的事情，这样一个开，一个关，连在一起。

```go
db, err := sql.Open("mysql", "user:password@/dbname")
if err != nil {
    panic(err.Error())
}
defer db.Close()

res, err = http.Get("https://github.com")
if err != nil {
  panic(err.Error())
}
defer res.Body.Close()
```

在线演示: https://tour.golang.org/flowcontrol/12

### JSON 序列化／反序列化

你不能像动态语言（Ruby, Node.js）那样自由的做 JSON 反序列化，必须有准备好的结构体来承载。

像这样的方式你就不行了：

```rb
b = '{ "name": "Foo", "body": "Hello" }'
m = JSON.parse(b)
```

因为在动态语言里面，对象结构是 Runtime 期间动态创建的。

你必须先告诉编译器，你需要的 JSON 结构：

```go
type Message struct {
    // `` 是 Go 对 struct 字段的特殊描述方式，用于很多的场景
    Name string `json:"name"`
    Body string `json:"raw"`
    Time int64
}

m := Message{"Alice", "Hello", 1294706395881547000}
// 序列化
b, err := json.Marshal(m)
// 然后 b 就等于 []byte(`{"name":"Alice","raw":"Hello","Time":1294706395881547000}`)
// 反序列化, b 要是 []byte 类型
err := json.Unmarshal(b, &m)
```

详见: https://blog.golang.org/json-and-go

### 🐵 演示编写一个常见的类

> $GOPATH/github.com/huacnlee/hello-go/monkey.go

```go
package main

// package 外部可访问
type Monkey struct {
	Name   string `json:"name"`
	age    int
	Gender int `json:"gender"`
}

// package 内部可访问
type gorilla struct {
	*Monkey
	Weight int
}

// package 外部可访问
// monkey 带 * 表示是一个指针
// err 没带 * 表示是值类型
func BuildMonkey(name string, age, gender int) (monkey *Monkey, err error) {
	err = nil
	// & 表示引用指针，别担心，用错了编译不过
	monkey = &Monkey{
		Name:   name,
		age:    age,
		Gender: gender,
	}
	return
}

// monkey.Age()
func (monkey *Monkey) Age() int {
	return monkey.age
}
```

### 编写测试

Go 内置简单的测试框架，据我了解大家都是用这套，只是用了一些辅助工具。

> $GOPATH/github.com/huacnlee/hello-go/monkey_test.go

```go
package main

import "testing"

func TestBuildMonkey(t *testing.T) {
	monkey, err := BuildMonkey("Foo", 5, 0)
	if (err != nil) {
		t.Error(err)
	}
	if monkey.Name != "Foo" {
		t.Error(monkey.Name)
	}
	if monkey.age != 5 {
		t.Error(monkey.age)
	}
	if monkey.Gender != 0 {
		t.Error(monkey.Gender)
	}
}
```

然后运行

```bash
$ go test ./...
ok  	github.com/huacnlee/hello-go	0.006s
?   	github.com/huacnlee/hello-go/utils	[no test files]
```

### 🏎 并发 Concurrency 编写

> NOTE: [并发不等于并行 Concurrency Is Not Parallelism](https://blog.golang.org/concurrency-is-not-parallelism) - “并发是一次处理（能力）很多事情，并行是一次执行（doing）很多事情”

Goroutine 是非常轻量级的线程，据我了解只会有内存 (2 KB 一个) 和 GC (1 µs 耗时) 的开销，可以把 Goroutine 看作一个异步调度器。

Goroutine 会自动根据情况分配 OS Thread 来实现多核运算（Go 1.5 以后，GOMAXPROCS 默认为 CPU 核数量）

你可以用 [sync.WaitGroup](https://golang.org/pkg/sync/) 来实现同步控制，使用 `chan` 来在不同的 Goroutine 之间共享值。

```go
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	rand.Seed(time.Now().UTC().UnixNano())

	numbers := []int{3, 4, 7, 10, 2, 5}

	for i := range numbers {
		number := numbers[i]
		wg.Add(1)

		go func() {
			defer wg.Done()
			sleepN := 100 + rand.Intn(100)
			time.Sleep(time.Duration(sleepN) * time.Millisecond)
			fmt.Println("number", number, "sleep", sleepN, "ms")
		}()
	}

	fmt.Println("Before wait()")

	wg.Wait()

	fmt.Println("After wait()")
}
```

在线演示：https://play.golang.org/p/OTT6EdqQz7

## ⛺️ 三方库依赖

### go get 内置工具

- 原理是基于 Git / SVN 或其它代码仓库的方式来管理三方库;
- 不会像 NPM / RubyGems 那样有一个包存储的中心服务器，而是直接用源代码来分发的；
- 安装三方库实际上是将远程代码拉到本地 `$GOPATH/src/github.com/foo/bar`;
- 引用三方库是写 URL 的方式，例如 `import "github.com/foo/bar"`;

### 🌰 来一包例子

```
.
├── main.go
└── utils
    └── format.go
```

```go
package main

import (
	"fmt"
	"github.com/huacnlee/hello-go/utils"
)

func main() {
	newName := utils.FormatName("自成")
	fmt.Println("Hello world.", newName)
}
```



## 🚀 发布你的应用

默认情况下 `go build` 会编译成当前平台，你可以使用 `GOOS` 和 `GOARCH` 来实现 **Cross Compile**

例如编译 Linux 64bit 支持:

```
$ GOOS=linux GOARCH=amd64 go build main.go
```

Windows 支持:

```
$ GOOS=windows GOARCH=amd64 go build main.go
```

或者在 Linux 编译为 macOS 支持

```
$ GOOS=darwin GOARCH=amd64 go build main.go
```

编译完以后，你会获取到一个大的二进制绿色文件 📟（内含 GC 的实现，所以起步文件比较大），给谁谁都可以执行 🎊 🎊 🎊。

```
go build -ldflags "-s -w"
```

- `-s` - Omit the symbol table and debug information.
- `-w` - Omit the DWARF symbol table.

https://golang.org/cmd/link/

## 🌏 如何编写 Web 应用

- 在 Go 的世界里面，都推荐用简单的 Web 框架
- 甚至有很多项目直接用标准库 `net/http` 来实现

```go
package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", func(ctx *gin.Context) {
		monkey, _ := BuildMonkey("Foo", 5, 0)
		ctx.JSON(200, monkey)
	})

	router.Run(":8080")
}
```

> NOTE! 在 Go 的世界，选择那个 Web 框架几乎不重要，因为最后不同的框架都可以整合在一块儿。


## 🍻 入门学习参考

- https://github.com/huacnlee/mediom - 复杂例子，包含 Web、数据库实现，早期代码完整的论坛项目
- https://github.com/jinzhu/gorm - 推荐使用 ORM
- https://github.com/gin-gonic/gin - 推荐使用 Web 框架
- http://divan.github.io/posts/go_concurrency_visualize - 用动画的方式演示并发的过程




