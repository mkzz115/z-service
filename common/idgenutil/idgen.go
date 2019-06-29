package idgenutil

import idworker "github.com/gitstliu/go-id-worker"

func CreateId() (int64, error) {
	currWoker := &idworker.IdWorker{}
	err := currWoker.InitIdWorker(1000, 1)
	if err != nil {
		return -1, err
	}
	newId, err := currWoker.NextId()
	if err != nil {
		return -1, err
	}
	return newId, nil
}

type IdGenerator interface {
	GenId() (int64, error)
}

type idgen struct {
	worker *idworker.IdWorker
}

func NewIdWorker() IdGenerator {
	currWoker := &idworker.IdWorker{}
	err := currWoker.InitIdWorker(1000, 1)
	if err != nil {
		panic(err.Error())
	}
	return &idgen{
		worker: currWoker,
	}
}

func (d *idgen) GenId() (int64, error) {
	newId, err := d.worker.NextId()
	if err != nil {
		return 0, err
	}
	return newId, nil
}
