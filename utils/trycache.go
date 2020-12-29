package utils


import (
	"reflect"
)

type TryCatch struct {
	errChan chan interface{}
	catches map[reflect.Type]func(err error)
	defaultCatch func(err error)
}

func (t TryCatch) Try(block func()) TryCatch {
	t.errChan = make(chan interface{})
	t.catches = map[reflect.Type]func(err error){}
	t.defaultCatch = func(err error) {}
	go func() {
		defer func() {
			t.errChan <- recover()
		}()
		block()
	}()
	return t
}

func (t TryCatch) CatchAll(block func(err error)) TryCatch {
	t.defaultCatch = block
	return t
}

func (t TryCatch) Catch(e error, block func(err error)) TryCatch {
	errorType := reflect.TypeOf(e)
	t.catches[errorType] = block
	return t
}

func (t TryCatch) Finally(block func()) TryCatch {
	err := <-t.errChan
	if err != nil {
		catch := t.catches[reflect.TypeOf(err)]
		if catch != nil {
			catch(err.(error))
		} else {
			t.defaultCatch(err.(error))
		}
	}
	block()
	return t
}


// Try catches exception from f
func Try(f func()) *tryStruct {
	return &tryStruct{
		catches: make(map[reflect.Type]ExeceptionHandler),
		hold:    f,
	}
}

// ExeceptionHandler handle exception
type ExeceptionHandler func(interface{})

type tryStruct struct {
	catches map[reflect.Type]ExeceptionHandler
	defaultCatch ExeceptionHandler
	hold    func()
}

func (t *tryStruct) Catch(e interface{}, f ExeceptionHandler) *tryStruct {
	t.catches[reflect.TypeOf(e)] = f
	return t
}

func (t *tryStruct) CacheAll(f ExeceptionHandler) *tryStruct {
	t.defaultCatch = f
	return t
}

func (t *tryStruct) Finally(f func()) {
	defer func() {
		if e := recover(); nil != e {
			if h, ok := t.catches[reflect.TypeOf(e)]; ok {
				h(e)
			} else {
				t.defaultCatch(e)
			}

			f()
		}
	}()

	t.hold()
}

