package gotest

import "testing"

// 파일명: *_test.go  (ex. sum.go -> sum_test.go)
// 함수명:Test 로 시작해야 go test 실행 대상이 됨

// 인수: *testing.T
// testing API
// t.Fatal: 테스트 실패 반환 API

// 실행: go test -v ./...
// 개별 실행: go test -v -run TestSum1

func TestSum1(t *testing.T) {
	if sum1(1, 2) != 3 {
		t.Fatal("sum(1,2) != 3")
	}
}
