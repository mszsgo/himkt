package tcc

import (
	log "github.com/sirupsen/logrus"
	"sync"
	"time"
)

type TccInterface interface {
	Try() error
	Confirm() error
	Cancel() error
}

// TccImpl
type TccDefaultImpl struct {
	TryFun     func() error
	ConfirmFun func() error
	CancelFun  func() error
}

func NewTcc(try, confirm, cancel func() error) *TccDefaultImpl {
	return &TccDefaultImpl{
		TryFun:     try,
		ConfirmFun: confirm,
		CancelFun:  cancel,
	}
}

func (o *TccDefaultImpl) Try() error {
	return o.TryFun()
}
func (o *TccDefaultImpl) Confirm() error {
	return o.ConfirmFun()
}
func (o *TccDefaultImpl) Cancel() error {
	return o.CancelFun()
}

// 调用TCC处理，按顺序执行函数
func Tx(ts ...TccInterface) (err error) {
	var (
		chtcc = make(map[int]chan bool)
		cherr = make(chan error, 1)
		wg    = &sync.WaitGroup{}
	)

	// 任务顺序执行
	for i, job := range ts {
		wg.Add(1)
		chtcc[i] = make(chan bool, 1)
		go Job(chtcc[i], cherr, wg, job)
		wg.Wait()
		// 如果异常，结束任务
		if len(cherr) > 0 {
			// error
			err = <-cherr
			// cancel
			for _, v := range chtcc {
				v <- false
			}
			return err
		}
	}
	// confirm
	for _, v := range chtcc {
		v <- true
	}
	return
}

func Job(chtcc <-chan bool, cherr chan<- error, wg *sync.WaitGroup, tcc TccInterface) {
	var err error
	defer func() {
		if err != nil {
			cherr <- err
			wg.Done()

			//rollback
			Retry(tcc.Cancel)
			return
		}
		wg.Done()

		//处理成功，等待通道通知，等于true提交任务，等于false撤销任务
		if <-chtcc {
			//commit
			Retry(tcc.Confirm)
		} else {
			//rollback
			Retry(tcc.Cancel)
		}
	}()
	// 处理业务代码
	err = tcc.Try()
}

// 重试操作
func Retry(f func() error) {
	retry := 3
	for i := 0; i < retry; i++ {
		e := f()
		if e == nil {
			break
		}
		log.Error("tx.Retry Error => " + e.Error())
		time.Sleep(time.Duration(i+1) * time.Second)
	}
}
