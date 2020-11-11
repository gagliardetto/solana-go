package ws

import "reflect"

type Subscription struct {
	req               *request
	subID             uint64
	stream            chan result
	err               chan error
	reflectType       reflect.Type
	closeFunc         func(err error)
	unsubscribeMethod string
}

func newSubscription(req *request, reflectType reflect.Type, closeFunc func(err error), unsubscribeMethod string) *Subscription {
	return &Subscription{
		req:               req,
		reflectType:       reflectType,
		stream:            make(chan result, 200),
		err:               make(chan error, 1),
		closeFunc:         closeFunc,
		unsubscribeMethod: unsubscribeMethod,
	}
}

func (s *Subscription) Recv() (interface{}, error) {
	select {
	case d := <-s.stream:
		return d, nil
	case err := <-s.err:
		return nil, err
	}
}

func (s *Subscription) Unsubscribe() {
	s.unsubscribe(nil)
}

func (s *Subscription) unsubscribe(err error) {
	s.closeFunc(err)

}
