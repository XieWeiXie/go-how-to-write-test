package partone

import (
	"testing"
)

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
		t.Logf("Hello() = %v, want %v", result, want)
	}

}

func testHello(expected string) func(t *testing.T) {
	return func(t *testing.T) {
		if expected == "Hello World" {
			t.Logf("Hello() = %v, want = %v", expected, "Hello World")
		} else {
			t.Errorf("Hello() = %v, want = %v", expected, "Hello World")
		}
	}

}

func TestHello2(t *testing.T) {
	t.Run("test for hello with run", testHello(Hello()))
}

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

func Test_world(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
		{
			name: "test for world",
			want: "world",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := world(); got != tt.want {
				t.Errorf("world() = %v, want %v", got, tt.want)
			}
		})
	}
}
