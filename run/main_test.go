package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	result := add(2, 3)
	expected := 5
	if result != expected {
		t.Errorf("add(2, 3) = %d; want %d", result, expected)
	}
}

func TestSubtract(t *testing.T) {
	result := Subtract(5, 3)
	expected := 2
	if result != expected {
		t.Errorf("Subtract(5, 3) = %d; want %d", result, expected)
	}
}

func TestMultiply(t *testing.T) {
	result := Multiply(2, 3)
	expected := 6
	if result != expected {
		t.Errorf("Multiply(2, 3) = %d; want %d", result, expected)
	}
}

//func TestPersonGreet(t *testing.T) {
//	p := person{Name: "John", Age: 30}
//	expected := "Hello, my name is John and I am 30 years old.\n"
//	if got := captureOutput(p.greet); got != expected {
//		t.Errorf("greet() = %q; want %q", got, expected)
//	}
//}

// Helper function to capture output
//func captureOutput(f func()) string {
//	oldPrintln := fmt.Println
//	oldPrintf := fmt.Printf
//	defer func() {
//		fmt.Println = oldPrintln
//		fmt.Printf = oldPrintf
//	}()
//	var buf bytes.Buffer
//	fmt.Println = func(a ...interface{}) (n int, err error) {
//		return fmt.Fprint(&buf, a...)
//	}
//	fmt.Printf = func(format string, a ...interface{}) (n int, err error) {
//		return fmt.Fprintf(&buf, format, a...)
//	}
//	f()
//	return buf.String()
//}
