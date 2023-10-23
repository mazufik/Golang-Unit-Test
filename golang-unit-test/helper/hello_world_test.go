package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Table Benchmark
func BenchmarkHelloWorldTable(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{
			name:    "HelloWorld(Zaydan)",
			request: "Zaydan",
		},
		{
			name:    "HelloWorld(Rahman)",
			request: "Rahman",
		},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.name)
			}
		})
	}
}

// Sub Benchmark
func BenchmarkSub(b *testing.B) {
	b.Run("Zaydan", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Zaydan")
		}
	})
	b.Run("Rahman", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Rahman")
		}
	})
}

// Benchmark
func BenchmarkHelloWorldZaydan(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Zaydan")
	}
}

func BenchmarkHelloWorldRahman(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Rahman")
	}
}

// Table Test
func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Zaydan",
			request:  "Zaydan",
			expected: "Hello Zaydan",
		},
		{
			name:     "Athaya",
			request:  "Athaya",
			expected: "Hello Athaya",
		},
		{
			name:     "Rahman",
			request:  "Rahman",
			expected: "Hello Rahman",
		},
		{
			name:     "Budi",
			request:  "Budi",
			expected: "Hello Budi",
		},
		{
			name:     "Joko",
			request:  "Joko",
			expected: "Hello Joko",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

// Sub Test
func TestSubTest(t *testing.T) {
	t.Run("Zaydan", func(t *testing.T) {
		result := HelloWorld("Zaydan")
		require.Equal(t, "Hello Zaydan", result, "Result must be 'Hello Zaydan'")
	})

	t.Run("Rahman", func(t *testing.T) {
		result := HelloWorld("Rahman")
		require.Equal(t, "Hello Rahman", result, "Result must be 'Hello Rahman'")
	})
}

// testing.M
func TestMain(m *testing.M) {
	// before
	fmt.Println("BEFORE UNIT TEST")

	m.Run()

	fmt.Println("AFTER UNIT TEST")
}

// Skip test (Membatalkan test)
func TestSkip(t *testing.T) {
	if runtime.GOOS == "linux" {
		t.Skip("Can not run on Linux")
	}

	result := HelloWorld("Zaydan")
	require.Equal(t, "Hello Zaydan", result, "Result must be 'Hello Zaydan'")
}

// cara testing dengan menggunakan assertion
func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Zaydan")
	assert.Equal(t, "Hello Zaydan", result, "Result must be 'Hello Zaydan'")
	fmt.Println("TestHelloWorld with Assert done")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Zaydan")
	require.Equal(t, "Hello Zaydan", result, "Result must be 'Hello Zaydan'")
	fmt.Println("TestHelloWorld with Require done")
}

// cara testing manual
func TestHelloWorldZaydan(t *testing.T) {
	result := HelloWorld("Zaydan")
	if result != "Hello Zaydan" {
		// error
		t.Error("Result must be 'Hello Zaydan")
	}

	fmt.Println("TestHelloWorldzaydan done")
}

func TestHelloWorldRahman(t *testing.T) {
	result := HelloWorld("Rahman")
	if result != "Hello Rahman" {
		// error
		t.Fatal("Result must be 'Hello Rahman")
	}

	fmt.Println("TestHelloWorldRahman done")
}
