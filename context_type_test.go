package rhombus

import (
	"errors"
	"strconv"
)

type setTaskParams struct {
	value interface{}
}

type setTask struct {
	params *setTaskParams
}

type abortTask struct{}
type panicTask struct{}
type nilTask struct{}
type errTask struct{}

func newSetTask(i int) Task {
	return Value(
		strconv.Itoa(i),
		&setTask{&setTaskParams{
			value: i,
		}},
	)
}

func newDiscardSetTask(i int) Task {
	return Value(
		"_",
		&setTask{&setTaskParams{
			value: i,
		}},
	)
}

func (s *nilTask) Do(c *Context) {}
func (s *nilTask) Value() interface{} {
	return nil
}

func (s *setTask) Do(c *Context) {}

func (s *setTask) Value() interface{} {
	return s.params.value
}

func (s *abortTask) Do(c *Context) {
	c.Abort("aborted")
}

func (s *abortTask) Value() interface{} {
	return nil
}

func (s *panicTask) Do(c *Context) {
	a := make([]int, 0, 0)
	_ = a[100]
}

func (s *panicTask) Value() interface{} {
	return nil
}

func (s *errTask) Do(c *Context) {}
func (s *errTask) Value() interface{} {
	return nil
}
func (s *errTask) Error() error {
	return errors.New("test error")
}
