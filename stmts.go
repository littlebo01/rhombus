package rhombus

type (
	Condition interface {
		Value(c *Context) bool
	}

	ifStmt struct {
		cond Condition
		job Task
		value bool
	}
)

func If(cond Condition, job Task) Task {
	return &ifStmt{cond, job, true}
}

func IfNot(cond Condition, job Task) Task {
	return &ifStmt{cond, job, false}
}

func (t *ifStmt) Do(c *Context) {
	if t.cond.Value(c) == t.value {
		t.job.Do(c)
	}
}

func (t *ifStmt) Value() interface{} {
	return nil
}
