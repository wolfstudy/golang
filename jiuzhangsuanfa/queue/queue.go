package queue

type Queue struct {
	Iterm []interface{}
}

func NewQueue() *Queue {
	iter := make([]interface{}, 0)
	return &Queue{iter}
}

func (q *Queue) Size() int {
	return len(q.Iterm)
}

func (q *Queue) EnQueue(value interface{}) {
	q.Iterm = append(q.Iterm, value)
}

func (q *Queue) DeQueue() interface{} {
	
}
