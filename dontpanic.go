package dontpanic

import "log/slog"

type GoOptions struct {
	recoverFunc func(recoveredValue any)
}

type GoOption func(*GoOptions)

func WithRecover(f func(recoveredValue any)) GoOption {
	return func(o *GoOptions) { o.recoverFunc = f }
}

func Go(fn func(), opts ...GoOption) {
	o := &GoOptions{
		recoverFunc: func(recoveredValue any) {
			if recoveredValue != nil {
				slog.Error("DON'T PANIC! (but you might want a towel)", slog.Any("recovered", recoveredValue))
			}
		},
	}
	for _, opt := range opts {
		if opt != nil {
			opt(o)
		}
	}

	go func() {
		defer func() {
			if r := recover(); r != nil {
				o.recoverFunc(r)
			}
		}()
		fn()
	}()
}
