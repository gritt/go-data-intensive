package error

type (
	ErrLogger interface {
		LogError(err error)
	}

	ErrCapturer interface {
		CaptureError(err error)
	}

	ErrHandler struct {
		logger   ErrLogger
		capturer ErrCapturer
	}
)

func NewErrorHandler(errLogger ErrLogger, errCapturer ErrCapturer) ErrHandler {
	return ErrHandler{
		logger:   errLogger,
		capturer: errCapturer,
	}
}

func Handle(err error) {
	// is base error ?

	// -> log as error
	//    -> capture error with sentry

	// -> log as info

	// -> log as warning
}
