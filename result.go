package util

type Result[T any] struct {
  value *T
  err error
}

func ResultOk[T any](value T) Result[T] {
  return Result[T]{value: &value, err: nil}
}

func ResultErr[T any](err error) Result[T] {
  return Result[T]{value: nil, err: err}
}

func (res Result[T]) IsOk() bool {
  return res.err == nil
}

func (res Result[T]) IsErr() bool {
  return !res.IsOk()
}

func (res Result[T]) Value() T {
  if !res.IsOk() {
    panic("Yank on Err Result")
  }

  return *res.value
}

func (res Result[T]) Error() error {
  if !res.IsErr() {
    panic("Yank on OK Result")
  }

  return res.err
}


