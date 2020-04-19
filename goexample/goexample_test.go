package goexample

import "fmt"

// 파일명: *_test.go
// 함수명: Example*()
// 검증: stdout 결과와 // Output: 이하 결과 비교

func ExamplePrint() {
	fruits := [3]string{"사과", "바나나", "토마토"}
	for _, fruit := range fruits {
		fmt.Printf("%s는 맛있다.\n", fruit)
	}
	// Output:
	// 사과는 맛있다.
	// 바나나는 맛있다.
	// 토마토는 맛있다.
}

func ExampleShufull() {
	x := map[string]int{"a": 1, "b": 2, "c": 3}
	for k, v := range x {
		fmt.Printf("k=%s v=%d\n", k, v)
	}
	// Unordered output:
	// k=b v=2
	// k=c v=3
	// k=a v=1
}
