package api

import (
	"fmt"
	"reflect"
	"sync"
)

// A container for automatic Dependency Injection
type Container interface {
	// Sets the value pointed by v to the instance
	// of the dependency that shares the same type.
	// If no instance is found, a new one will be created
	// from the type of v.
	// Returns an error if the underlying type is not a struct.
	Inject(v any) error

	// Calls [Container.Inject] for every value in s
	// Returns an error if any calls to [Container.Inject] fails.
	BulkInject(s []any) error
}

type container struct {
	deps   map[reflect.Type]*any
	mutex  sync.RWMutex
	logger *Logger
}

var onceContainer sync.Once
var singleContainer *container

func GetContainer() Container {
	onceContainer.Do(func() {
		singleContainer = &container{
			deps:   make(map[reflect.Type]*any),
			logger: &Logger{Context: "Container"},
		}
	})

	return singleContainer
}

func (c *container) Inject(v any) error {
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Pointer || rv.Elem().Kind() != reflect.Struct {
		t := rv.Type()
		return fmt.Errorf("Cannot resolve dependency for type \"%s\".", t.Name())
	}

	t := rv.Elem().Type()

	c.mutex.RLock()
	if inst, exists := c.deps[t]; exists {
		rv.Elem().Set(reflect.ValueOf(*inst).Elem())
		c.mutex.RUnlock()
		return nil
	}
	c.mutex.RUnlock()

	inst := reflect.New(t).Elem()
	errCh := make(chan error, t.NumField())
	var wg sync.WaitGroup

	for i := 0; i < t.NumField(); i++ {
		f := inst.Field(i)
		if !f.CanSet() || f.Kind() != reflect.Struct {
			continue
		}

		wg.Add(1)
		go func(f reflect.Value, errCh chan<- error) {
			defer wg.Done()

			dep := reflect.New(f.Type()).Interface()
			errCh <- c.Inject(dep)

			f.Set(reflect.ValueOf(dep).Elem())
		}(f, errCh)
	}

	wg.Wait()
	close(errCh)

	for err := range errCh {
		if err != nil {
			return err
		}
	}
	c.mutex.Lock()
	c.deps[t] = &v
	c.mutex.Unlock()

	rv.Elem().Set(inst)
	c.logger.Success(fmt.Sprintf("Loaded dependency for type \"%s\"", t.Name()))
	return nil
}

func (c *container) BulkInject(s []any) error {
	errCh := make(chan error, len(s))
	var wg sync.WaitGroup

	for i := 0; i < len(s); i++ {
		wg.Add(1)
		go func(v any, errCh chan<- error) {
			defer wg.Done()
			errCh <- c.Inject(v)
		}(s[i], errCh)
	}

	wg.Wait()
	close(errCh)

	for err := range errCh {
		if err != nil {
			return err
		}
	}

	return nil
}
