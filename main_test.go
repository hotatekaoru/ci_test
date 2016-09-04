package main

import (
	"fmt"
	"testing"
	"github.com/onsi/gomega"
)

func TestSum(t *testing.T) {
	str1 := "1"
	str2 := "2"
	result := sum(str1, str2)

	// 本来のテストコードの書き方（ここだけ）
	if 3 != result {
		t.Errorf("Sum Error %s + %s = %v", str1, str2, result)
	}
}

func TestConvStrtoInt1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()

	result := convStrtoInt("1")

	gomega.Expect(result).To(gomega.Equal(1))

	if 1 != result {
		t.Errorf("Convert Error %v", result)
	}
}

func TestConvStrtoInt2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()

	gomega.Expect(convStrtoInt("a")).To(gomega.Panic())
}