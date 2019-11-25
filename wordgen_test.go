package wordgen

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	word := New(10)
	fmt.Printf("Word generated: %s\n", word)
	if len(word) != 10 {
		t.Errorf("Word length is incorrect: %d, want: %d.", len(word), 10)
	}
}

func BenchmarkNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		New(64)
	}
}

func BenchmarkOnlyLetters(b *testing.B) {
	for i := 0; i < b.N; i++ {
		OnlyLetters(64)
	}
}

func BenchmarkOnlyNumbers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		OnlyNumbers(64)
	}
}

func BenchmarkCoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		coin()
	}
}

func TestOnlyLetters(t *testing.T) {
	word := OnlyLetters(15)
	fmt.Printf("Word generated: %s\n", word)
	if len(word) != 15 {
		t.Errorf("Word length is incorrect: %d, want: %d.", len(word), 15)
	}
	for _, c := range word {
		if !((int(c) < 91 && int(c) > 64) || (int(c) < 123 && int(c) > 96)) {
			t.Error("Word only can contain letters.")
		}
	}
}

func TestOnlyNumbers(t *testing.T) {
	word := OnlyNumbers(12)
	fmt.Printf("Word generated: %s\n", word)
	if len(word) != 12 {
		t.Errorf("Word length is incorrect: %d, want: %d.", len(word), 12)
	}
	for _, c := range word {
		if !(int(c) < 58 && int(c) > 47) {
			t.Error("Word only can contain numbers.")
		}
	}
}

func TestNotSymbols(t *testing.T) {
	word := NotSymbols(20)
	fmt.Printf("Word generated: %s\n", word)
	if len(word) != 20 {
		t.Errorf("Word length is incorrect: %d, want: %d.", len(word), 20)
	}
}
