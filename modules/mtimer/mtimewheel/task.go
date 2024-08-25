package mtimewheel

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

type Timer struct {
	id       uint64
	delay    time.Duration
	expireAt time.Time
	round    int
	callback Callback

	circle bool
}

func newTask(_id uint64, _delay time.Duration, cb Callback, _circle bool) *Timer {
	return &Timer{
		id:       _id,
		delay:    _delay,
		expireAt: time.Now().Add(_delay),
		callback: cb,
		circle:   _circle,
	}
}

func (t *Timer) Circle() bool {
	if t.circle {
		t.expireAt = time.Now().Add(t.delay)
	}
	return t.circle
}
