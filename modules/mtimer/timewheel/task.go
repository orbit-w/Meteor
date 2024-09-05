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

type Timer struct {
	id       uint64
	bIndex   int64 //bucket id
	twIndex  int64 //time wheel id
	delay    int64 //时间戳，单位是ms
	round    int64
	expireAt int64
	callback Callback
}

func newTimer(_id uint64, _delay time.Duration, cb Callback) *Timer {
	return &Timer{
		id:       _id,
		callback: cb,
		delay:    _delay.Milliseconds(),
		expireAt: time.Now().Add(_delay).UnixNano(),
	}
}

func (t *Timer) setBucketIndex(i int64) {
	t.bIndex = i
}

func (t *Timer) setTimeWheelIndex(i int64) {
	t.twIndex = i
}

func (t *Timer) Expired(cur time.Time) bool {
	return t.expireAt <= cur.UnixNano()
}

type Task struct {
	Id       uint64
	cb       Callback
	expireAt time.Time
}

func newTask(cb Callback) Task {
	return Task{
		cb:       cb,
		expireAt: time.Now().Add(taskTimeout),
	}
}

func (t *Task) Expired() bool {
	return time.Now().After(t.expireAt)
}

type Command func(*Timer) (success bool)
