package greetings

import "fmt"

// 주어진 이름에 대한 인사 메세지를 반환
// 대문자로 시작하는 함수명은 자동으로 Export 되기 때문에, 스크립트 내부에서만 사용되는 함수는 소문자로 시작한다.
func Hello(name string) string {
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}
