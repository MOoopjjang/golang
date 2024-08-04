package syncmanager

import (
	"goproject/demo/pb/pbtype"
	"sync"
)

type SyncManager struct {
	wg    sync.WaitGroup
	addCh chan *pbtype.User
}

var sm SyncManager

func Sm() *SyncManager {
	return &sm
}

func (s *SyncManager) Initialzie() {
	(*s).addCh = make(chan *pbtype.User)
}

func (s *SyncManager) Add(count int) {
	(*s).wg.Add(count)
}

func (s *SyncManager) Done() {
	(*s).wg.Done()
}

func (s *SyncManager) Close() {
	close((*s).addCh)
}

func (s *SyncManager) CloseAndWait() {
	(*s).Close()
	(*s).Wait()
}

func (s *SyncManager) Wait() {
	(*s).wg.Wait()
}

func (s *SyncManager) Ch() chan *pbtype.User {
	return (*s).addCh
}
