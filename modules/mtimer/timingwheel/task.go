package timewheel

import "time"

/*
   @Author: orbit-w
   @File: task
   @2024 8月 周四 23:31
*/

// Callback 延迟调用函数对象
type Callback struct {
	f    func(...any)  //f : 延迟函数调用原型
	args []interface{} //args: 延迟调用函数传递的形参
}

func newCallback(f func(...any), args ...any) Callback {
	return Callback{
		f:    f,
		args: args,
	}
}

func (cb *Callback) Exec() {
	defer func() {
		if err := recover(); err != nil {
			return
		}
	}()

	cb.f(cb.args...)
}

type TimerTask struct {
	id         uint64
	bIndex     int //bucket id
	expiration int64
	callback   Callback
}

func newTimer(_id uint64, _delay time.Duration, cb Callback) *TimerTask {
	return &TimerTask{
		id:         _id,
		callback:   cb,
		expiration: time.Now().UTC().Add(_delay).UnixMilli(),
	}
}

type Command func(*TimerTask) (success bool)
