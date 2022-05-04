package _2_switch

import (
	"fmt"
	"os"
	"os/exec"
	"testing"
)

// 不需要存在break 选择具备唯一性 重复会报错
func Test1(t *testing.T) {
	value := [...]int8{0, 1, 2, 3, 4, 5, 6}
	switch value[1] {

	case 0, 1:
		fmt.Println("0 or 1")
	case 2, 3:
		fmt.Println("2 or 3")
	case 4, 5, 6:
		fmt.Println("4 or 5 or 6")
	}
}

// 这里注意到switch的赋值使用
func underlyingError(err error) error {
	switch err := err.(type) {
	case *os.PathError:
		return err.Err
	case *os.LinkError:
		return err.Err
	case *os.SyscallError:
		return err.Err
	case *exec.Error:
		return err.Err
	}
	return err
}
