package queueimpl

type queue struct {
	values [10]int
	begin  int
	end    int
	total  int
}

func NewQueue() *queue {
	return &queue{
		values: [10]int{},
		begin:  0,
		end:    0,
		total:  0,
	}
}

func (q *queue) Queue(element int) {
	q.values[q.end] = element
	q.end = (q.end + 1) % len(q.values)
	q.total++
}

func (q *queue) Dequeue() int {
	element := q.values[q.begin]
	q.begin = (q.begin + 1) % len(q.values)
	q.total--
	return element
}
func (q *queue) IsEmpty() bool {
	return q.total == 0
}
