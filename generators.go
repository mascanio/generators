package generators

func Repeat[T any](done <-chan struct{}, values ...T) <-chan T {
	valueStream := make(chan T)

	go func() {
		defer close(valueStream)
		if len(values) == 0 {
			return
		}
		for {
			for _, v := range values {
				select {
				case <-done:
					return
				case valueStream <- v:
				}
			}
		}
	}()

	return valueStream
}

func Take[T any](done <-chan struct{}, valueStream <-chan T, numTake int) <-chan T {
	takeStream := make(chan T)

	go func() {
		defer close(takeStream)
		for i := 0; i < numTake; i++ {
			select {
			case <-done:
				return
			case takeStream <- <-valueStream:
			}
		}
	}()

	return takeStream
}

func RepeatFn[T any](done <-chan struct{}, fn func() T) <-chan T {
	valueStream := make(chan T)

	go func() {
		defer close(valueStream)
		for {
			select {
			case <-done:
				return
			case valueStream <- fn():
			}
		}
	}()

	return valueStream
}
