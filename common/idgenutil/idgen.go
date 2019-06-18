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
