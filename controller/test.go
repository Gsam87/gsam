package controller

import "fmt"

type TestController interface {
	TestCase22() error
}

type TestControllerImpl struct {
}

func NewTestController() TestController {
	return &TestControllerImpl{}
}

func (c *TestControllerImpl) TestCase22() error {
	in := 3
	var res []string
	var backtrack func(path string, left int, right int)

	backtrack = func(path string, left int, right int) {
		// 如果左右括號都用完了，加入結果
		if left == 0 && right == 0 {
			res = append(res, path)
			return
		}

		// 優先加入左括號（如果還有剩下）
		if left > 0 {
			backtrack(path+"(", left-1, right)
		}
		// 只有當右括號剩下的比左括號多時才可以加入右括號
		if right > left {
			backtrack(path+")", left, right-1)
		}
	}

	backtrack("", in, in)
	fmt.Println("#### LeedCode case 22:", res)

	return nil
}
