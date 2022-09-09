package practices_test

import (
	"fmt"
	"testing"

	"go-staffs/practices"
)

//go test -bench .|benchgraph -title="Fibonacci Graph"
func TestFibonacci_Common(t *testing.T) {
	practices.FibonacciIterative(8)

	fmt.Println("====== Normal Recursive =====")
	practices.FibonacciNormalRecursive(1, 0, 8)

	fmt.Println("====== Double Recursive =====")
	practices.FibonacciDoubleRecursive(8)

	fmt.Println("====== Currification =====")
	practices.FibonacciCurrification(8)
}

func benchmarkFibonacciIterative(value int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		practices.FibonacciIterative(value)
	}
}
func BenchmarkFibonacciIterative_8(b *testing.B) {
	benchmarkFibonacciIterative(8, b)
}
func BenchmarkFibonacciIterative_16(b *testing.B) {
	benchmarkFibonacciIterative(16, b)
}
func BenchmarkFibonacciIterative_32(b *testing.B) {
	benchmarkFibonacciIterative(32, b)
}
func BenchmarkFibonacciIterative_64(b *testing.B) {
	benchmarkFibonacciIterative(64, b)
}
func BenchmarkFibonacciIterative_128(b *testing.B) {
	benchmarkFibonacciIterative(128, b)
}

func BenchmarkFibonacciNormalRecursive_8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		practices.FibonacciNormalRecursive(1, 0, 8)
	}
}
func BenchmarkFibonacciNormalRecursive_16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		practices.FibonacciNormalRecursive(1, 0, 16)
	}
}
func BenchmarkFibonacciNormalRecursive_32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		practices.FibonacciNormalRecursive(1, 0, 32)
	}
}

func BenchmarkFibonacciNormalRecursive_64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		practices.FibonacciNormalRecursive(1, 0, 64)
	}
}
func BenchmarkFibonacciNormalRecursive_128(b *testing.B) {
	for i := 0; i < b.N; i++ {
		practices.FibonacciNormalRecursive(1, 0, 128)
	}
}

/*
func BenchmarkFibonacciDouble_Recursive1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		practices.FibonacciDoubleRecursive(8)
	}
}
func BenchmarkFibonacciDouble_Recursive2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		practices.FibonacciDoubleRecursive(16)
	}
}
func BenchmarkFibonacciDouble_Recursive3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		practices.FibonacciDoubleRecursive(32)
	}
}*/
func BenchmarkFibonacciCurrification_8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		practices.FibonacciCurrification(8)
	}
}
func BenchmarkFibonacciCurrification_16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		practices.FibonacciCurrification(16)
	}
}
func BenchmarkFibonacciCurrification_32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		practices.FibonacciCurrification(32)
	}
}
func BenchmarkFibonacciCurrification_64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		practices.FibonacciCurrification(64)
	}
}
func BenchmarkFibonacciCurrification_128(b *testing.B) {
	for i := 0; i < b.N; i++ {
		practices.FibonacciCurrification(128)
	}
}
