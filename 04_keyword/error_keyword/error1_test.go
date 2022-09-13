package error_keyword

import (
	"errors"
	"fmt"
	"io"
	"testing"
)

// 这里唯一需要注意的就是error.new的使用
func Test1(t *testing.T) {
	for _, req := range []string{"", "hello!"} {
		fmt.Printf("request: %s\n", req)
		resp, err := echo(req)

		//必须检查 才能使用结果
		if err != nil {
			fmt.Printf("error: %s\n", err)
			continue
		}
		fmt.Printf("response: %s\n", resp)
	}
}

//如果我们需要抛出异常 我们需要做的事情是返回可能的error
func echo(request string) (response string, err error) {
	if request == "" {
		err = errors.New("empty request")
		return
	}
	response = fmt.Sprintf("echo: %s", request)
	return
}

//非二分法判断
type temporary interface {
	Temporary() bool
}

//断言一个行为
func IsTemporary(err error) bool {
	te, ok := err.(temporary)
	return ok && te.Temporary()
}

func Test3(t *testing.T) {
	IsTemporary(errors.New("ee"))
}

// 减少套路 减少判断
type errorWrite struct {
	io.Writer
	err error
}

func (e *errorWrite) Write() (int, error) {
	if e.err != nil {
		return 0, e.err
	}
	return 1, nil
}
