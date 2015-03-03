package queue

type Queue struct {
	queue  []interface{}
	length int
}

func NewQueue() *Queue {
	q := &Queue{}
	q.queue = make([]interface{}, 0)
	q.length = 0

	return q
}

func (q *Queue) Length() int {
	return q.length
}

func (q *Queue) IsEmpty() bool {
	return q.length == 0
}

func (q *Queue) Push(el interface{}) {
	q.queue = append(q.queue, el)
	q.length++
}

func (q *Queue) Pop() interface{} {
	el := q.queue[0]
	q.queue = q.queue[1:]
	q.length--

	return el
}

func (q *Queue) Peek() interface{} {
	return q.queue[0]
}
