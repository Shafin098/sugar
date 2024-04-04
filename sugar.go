package sugar

// Sugar is a struct that holds the value and error channels
type Sugar[T any] struct {
	ValueChan chan T
	ErrChan   chan error
}

// Async provides a way to run a function asynchronously in a separate goroutine
func Async[K any, V any](f func(...K) (V, error), args ...K) Sugar[V] {
	c := make(chan V)
	e := make(chan error)
	go func() {
		val, err := f(args...)
		c <- val
		e <- err

	}()
	return Sugar[V]{ValueChan: c, ErrChan: nil}
}

// Await provides a way to wait for the result of an asynchronous function
func Await[K any](s Sugar[K]) (K, error) {
	select {
	case val := <-s.ValueChan:
		return val, nil
	case err := <-s.ErrChan:
		return *new(K), err
	}
}

// AwaitAll provides a way to wait for the results of multiple asynchronous functions
func AwaitAll[K any](s ...Sugar[K]) ([]K, []error) {
	var vals []K
	var errs []error
	for _, sugar := range s {
		select {
		case val := <-sugar.ValueChan:
			vals = append(vals, val)
			errs = append(errs, nil)
		case err := <-sugar.ErrChan:
			errs = append(errs, err)
			vals = append(vals, *new(K))
		}
	}
	return vals, errs
}
