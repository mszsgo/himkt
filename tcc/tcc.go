package tcc

// 创建tcc表，记录事务状态，确认操作或者取消操作不成功，继续重试。

type Tcc interface {
	Try() error
	Confirm() error
	Cancel() error
}

// TCC事务，确认与取消操作重试3次
func Transaction(tcc Tcc) {
	retry := 3
	err := tcc.Try()
	if err != nil {
		for i := 0; i < retry; i++ {
			err = tcc.Cancel()
			if err == nil {
				break
			}
		}
		return
	}
	for i := 0; i < retry; i++ {
		err = tcc.Confirm()
		if err == nil {
			break
		}
	}

}
