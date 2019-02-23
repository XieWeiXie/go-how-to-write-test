# How to write test with golang
[![Build Status](https://travis-ci.com/wuxiaoxiaoshen/go-how-to-write-test.svg?branch=master)](https://travis-ci.com/wuxiaoxiaoshen/go-how-to-write-test)
[![codecov](https://codecov.io/gh/wuxiaoxiaoshen/go-how-to-write-test/branch/master/graph/badge.svg)](https://codecov.io/gh/wuxiaoxiaoshen/go-how-to-write-test)
[![Go Report Card](https://goreportcard.com/badge/github.com/wuxiaoxiaoshen/go-how-to-write-test)](https://goreportcard.com/report/github.com/wuxiaoxiaoshen/go-how-to-write-test)

- TDD(Test-Driven development) 测试驱动开发
- 内置的 testing 库 、 表格驱动、样本测试、TestMain
- 第三方：goconvey
- Monkey 猴子补丁
- 数据库 mock
- travisCI
- 代码覆盖率



:fire::fire: TDD

- 快速实现功能
- 再设计和重构

:fire::fire: 软件测试

> 在指定的条件下，操作程序，发现程序错误

:fire::fire: 单元测试

对软件的组成单元进行测试，最小单位：函数

包含三个步骤：

- 指定输入
- 指定预期
- 函数结果和指定的预期比较

指标:

- 代码覆盖率：运行测试执行的代码占总代码的行数



:fire::fire: testing 库的使用

```
// Hello ...
func Hello() string {
	return "Hello World"
}
```

``` 
// 传统测试
func TestHello(t *testing.T) {
	result := Hello()
	want := "Hello World"
	if result == want {
		t.Logf("Hello() = %v, want %v", result, want)
	} else {
		t.Errorf("Hello() = %v, want %v", result, want)
	}

	want2 := "Hello world"
	if result == want2 {
		t.Logf("Hello() = %v, want %v", result, want)
	} else {
		t.Errorf("Hello() = %v, want %v", result, want)
	}

}

// 表格驱动测试： 使用匿名结构体，逻辑更清晰
func TestHelloWithTable(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
		{
			name: "test for hello",
			want: "Hello World",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Hello(); got != tt.want {
				t.Errorf("Hello() = %v, want %v", got, tt.want)
			}
		})
	}
}
```

运行：
```
// mode one 
go test  //  equal to : go test .  执行当前目录下的测试文件

// mode two 
go test ./..   // 加上路径参数，可以执行指定目录下的测试文件

```

样本测试：

```
func ExampleHello() {
	fmt.Println(Hello())
	// Output:
	// Hello World
} 

```

TestMain:
> 包的测试运行之前执行

``` 
func TestMain(m *testing.M) {
	fmt.Println("Before ====================")
	code := m.Run()
	fmt.Println("End ====================")
	os.Exit(code)
}
```


testing 包含下面几种方法：

- Log | Logf
- Error | ErrorF
- Fatal | FatalF


备注：

- 文件必须以 ...test.go 结尾
- 测试函数必须以 TestX... 开头, `X` 可以是 `_` 或者大写字母，不可以是小心字母或数字
- 参数：*testing.T
- 样本测试必须以 Example... 开头，输入使用注释的形式
- TestMain 每个包只有一个，参数为 *testing.M

覆盖率：

```
go test -cover

go test -coverprofile=cover.out
go tool cover -html=cover.out -o coverage.html
```




:fire::fire:第三方：goconvey

- 支持断言
- 支持嵌套
- 完全兼容内置 testing
- 提供 web UI

``` 
func TestAdd_Two(t *testing.T) {
	Convey("test add", t, func() {
		Convey("0 + 0", func() {
			So(Add(0, 0), ShouldEqual, 0)
		})
		Convey("-1 + 0", func() {
			So(Add(-1, 0), ShouldEqual, -1)
		})
	})
}

func TestFloatToString_Two(t *testing.T) {
	Convey("test float to string", t, func() {
		Convey("1.0/3.0", func() {
			result := FloatToString(1.0, 3.0)
			So(result, ShouldContainSubstring, "%")
			So(len(result), ShouldEqual, 6)
			So(result, ShouldEqual, "33.33%")
		})
	})

}

```

``` 
goconvey // 启动 web 界面
```



:fire::fire: Monkey 猴子补丁

- 函数打桩
- 过程打桩
- 方法打桩

``` 
// 函数
func main() {
	monkey.Patch(fmt.Println, func(a ...interface{}) (n int, err error) {
		s := make([]interface{}, len(a))
		for i, v := range a {
			s[i] = strings.Replace(fmt.Sprint(v), "hell", "*bleep*", -1)
		}
		return fmt.Fprintln(os.Stdout, s...)
	})
	fmt.Println("what the hell?") // what the *bleep*?
}
```
```
// 方法
func main() {
	var d *net.Dialer // Has to be a pointer to because `Dial` has a pointer receiver
	monkey.PatchInstanceMethod(reflect.TypeOf(d), "Dial", func(_ *net.Dialer, _, _ string) (net.Conn, error) {
		return nil, fmt.Errorf("no dialing allowed")
	})
	_, err := http.Get("http://google.com")
	fmt.Println(err) // Get http://google.com: no dialing allowed
}
```

``` 
// 过程
guard := Patch(DestroyResource, func(_ string) {
    
})
defer guard.Unpatch()
```

使用思路，被测函数中需要使用的其他依赖函数，进行打桩处理。

:fire::fire: sqlmock

对 sql 的执行过程进行打桩。

- 创建模拟链接
- 编写 原生 sql 语句
- 编写 返回值 或者 错误信息
- 判断执行结果和预设的返回值



:fire::fire: Reference

- [gotests](https://github.com/cweill/gotests) 自动生成测试代码，只需填写测试数据即可
- [goconvey](https://github.com/smartystreets/goconvey) 第三方测试库，兼容 testing 库
- [httpmock](https://github.com/jarcoal/httpmock) 接口模拟
- [how to test with Go](https://www.calhoun.io/how-to-test-with-go/) 参考文档
- [monkey](https://github.com/bouk/monkey) 猴子补丁
- [sqlmock](https://github.com/DATA-DOG/go-sqlmock) sqlmock