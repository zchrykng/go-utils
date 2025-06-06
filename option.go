package util

import "errors"

type Option[T any] struct {
  value *T
}

var ErrOptionIsNone = errors.New("gonads: Option[T] has no value")

func NewOption[T any]() *Option[T] {
  return &Option[T]{}
}

func (o Option[T]) IsSome() bool {
  return o.value != nil
}

func (o Option[T]) IsNone() bool {
  return !o.IsSome()
}

func (o Option[T]) Take() (T, error) {
  if o.IsNone() {
    var zero T
    return zero, ErrOptionIsNone
  }

  return *o.value, nil
}

func (o *Option[T]) Set(val T) {
  o.value = &val
}

func (o *Option[T]) Clear() {
  o.value = nil
}

func (o Option[T]) Yank() T {
  if o.IsNone() {
    panic("Yank on None Option")
  }

  return *o.value
}

