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
