package homework12

import "container/heap"

type Task struct {
	Identifier int
	Priority   int
}

type HTask struct {
	Task
	index int
}

type PriorityQueue []*HTask

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Priority > pq[j].Priority
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}
func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*HTask)
	item.index = n
	*pq = append(*pq, item)
}
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

type Scheduler struct {
	tasks map[int]*HTask
	queue *PriorityQueue
}

func NewScheduler() Scheduler {
	queue := PriorityQueue([]*HTask{})
	heap.Init(&queue)
	return Scheduler{
		tasks: make(map[int]*HTask),
		queue: &queue,
	}
}

func (s *Scheduler) AddTask(task Task) {
	if _, ok := s.tasks[task.Identifier]; ok {
		return
	}
	hTask := &HTask{task, 0}
	s.tasks[task.Identifier] = hTask
	heap.Push(s.queue, hTask)
}

func (s *Scheduler) ChangeTaskPriority(taskID int, newPriority int) {
	task, ok := s.tasks[taskID]
	if !ok || task.Priority == newPriority {
		return
	}
	task.Priority = newPriority
	heap.Fix(s.queue, task.index)
}

func (s *Scheduler) GetTask() Task {
	if s.queue.Len() == 0 {
		return Task{}
	}
	task := heap.Pop(s.queue).(*HTask)
	delete(s.tasks, task.Identifier)
	return task.Task
}
