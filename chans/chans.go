package chans

import "runtime"

// Drain drains any elements remaining on the channel.
func Drain[T any](c <-chan T) {
	for range c {
	}
}

// Merge merges two channels of some element type into a single channel.
func Merge[T any](c1, c2 <-chan T) <-chan T {
	r := make(chan T)
	go func(c1, c2 <-chan T, r chan<- T) {
		defer close(r)
		for c1 != nil || c2 != nil {
			select {
			case v1, ok := <-c1:
				if ok {
					r <- v1
				} else {
					c1 = nil
				}
			case v2, ok := <-c2:
				if ok {
					r <- v2
				} else {
					c2 = nil
				}
			}
		}
	}(c1, c2, r)
	return r
}

// Ranger provides a convenient way to exit a goroutine sending values
// when the receiver stops reading them.
//
// Ranger returns a Sender and a Receiver. The Receiver provides a
// Next method to retrieve values. The Sender provides a Send method
// to send values and a Close method to stop sending values. The Next
// method indicates when the Sender has been closed, and the Send
// method indicates when the Receiver has been freed.
func Ranger[T any]() (*Sender[T], *Receiver[T]) {
	c := make(chan T)
	d := make(chan bool)
	s := &Sender[T]{values: c, done: d}
	r := &Receiver[T]{values: c, done: d}
	// The finalizer on the receiver will tell the sender
	// if the receiver stops listening.
	runtime.SetFinalizer(r, r.finalize)
	return s, r
}

// A Sender is used to send values to a Receiver.
type Sender[T any] struct {
	values chan<- T
	done   <-chan bool
}

// Send sends a value to the receiver. It reports whether any more
// values may be sent; if it returns false the value was not sent.
func (s *Sender[T]) Send(v T) bool {
	select {
	case s.values <- v:
		return true
	case <-s.done:
		// The receiver has stopped listening.
		return false
	}
}

// Close tells the receiver that no more values will arrive.
// After Close is called, the Sender may no longer be used.
func (s *Sender[T]) Close() {
	close(s.values)
}

// A Receiver receives values from a Sender.
type Receiver[T any] struct {
	values <-chan T
	done   chan<- bool
}

// Next returns the next value from the channel. The bool result
// reports whether the value is valid. If the value is not valid, the
// Sender has been closed and no more values will be received.
func (r *Receiver[T]) Next() (T, bool) {
	v, ok := <-r.values
	return v, ok
}

// finalize is a finalizer for the receiver.
// It tells the sender that the receiver has stopped listening.
func (r *Receiver[T]) finalize() {
	close(r.done)
}
