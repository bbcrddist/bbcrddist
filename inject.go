package inject

import (
	"errors"
	"fmt"
	"reflect"
	"slices"
	"sync"
)

type Injector interface {
	Get(reflect.Type) (reflect.Value, error)
}

type Injectable interface {
	Type() reflect.Type
	Inject(Injector) (reflect.Value, error)
}

func New(args ...any) (*Scope, error) {
	return (*Scope)(nil).Scope(args...)
}


func Get[T any](inj Injector) (t T, err error) {
	typ := reflect.TypeOf(t)
	var rt reflect.Value
	rt, err = inj.Get(typ)
	if err != nil {
		return
	}
	t = rt.Interface().(T)
	return
}
