package ws

type Subscription struct {
	req               *request
	subID             uint64
	stream            chan result
	err               chan error
	closeFunc         func(err error)
	unsubscribeMethod string
	decoderFunc       decoderFunc
}

type decoderFunc func([]byte) (interface{}, error)

func newSubscription(
	req *request,
	closeFunc func(err error),
	unsubscribeMethod string,
	decoderFunc decoderFunc,
) *Subscription {
	return &Subscription{
		req:               req,
		subID:             0,
		stream:            make(chan result, 200_000),
		err:               make(chan error, 100_000),
		closeFunc:         closeFunc,
		unsubscribeMethod: unsubscribeMethod,
		decoderFunc:       decoderFunc,
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
