package tcc

// 创建tcc表，记录事务状态，确认操作或者取消操作不成功，继续重试。

type Tcc interface {
	Try() error
	Confirm() error
	Cancel() error
}

// TCC事务，确认与取消操作重试3次 ，返回异常代表预操作有异常
func Transaction(tcc Tcc) error {
	retry := 3
	err := tcc.Try()
	if err != nil {
		for i := 0; i < retry; i++ {
			e := tcc.Cancel()
			if e == nil {
				break
			}
		}
		return err
	}
	for i := 0; i < retry; i++ {
		e := tcc.Confirm()
		if e == nil {
			break
		}
	}

}
