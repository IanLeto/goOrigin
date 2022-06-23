package model

type Task interface {
}

type Pools struct {
	Limit int
}

func (tp *TaskPool) ReceiveTask(t Task) {

}