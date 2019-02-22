# How to write test with golang


- TDD(Test-Driven development) 测试驱动开发
- 内置的 testing 库 和 表格驱动
- 第三方：goconvey


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

testing 包含下面几种方法：

- Log | Logf
- Error | ErrorF
- Fatal | FatalF


备注：

- 文件必须以 ...test.go 结尾
- 测试函数必须以 TestX... 开头, `X` 可以是 `_` 或者大写字母，不可以是小心字母或数字
- 参数：*testing.T

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



:fire::fire: Reference

- [gotests](https://github.com/cweill/gotests) 自动生成测试代码，只需填写测试数据即可
- [goconvey](https://github.com/smartystreets/goconvey) 第三方测试库，兼容testing 库