package null

type Generic[T any] struct {
	Value T
	Valid bool
}

func NewGeneric[T any](value T, valid bool) Generic[T] {
	return Generic[T]{
		Value: value, Valid: valid,
	}
}

func GenericFrom[T any](value T) Generic[T] {
	return NewGeneric[T](value, true)
}

func GenericFromPtr[T any](value *T, opts ...Opts) Generic[T] {
	if value == nil {
		var zero T
		valid := false
		if opts != nil && len(opts) > 0 {
			valid = bool(opts[0])
		}

		return NewGeneric[T](zero, valid)
	}

	return NewGeneric[T](*value, true)
}

func (t *Generic[T]) ValueOrZero() T {
	if !t.Valid {
		var zero T
		return zero
	}

	return t.Value
}

func (t *Generic[T]) SetValid(v T) {
	t.Valid = true
	t.Value = v
}

//func (t *Generic[T]) SetValueIfNilValue(v T) {
//	if !t.Valid {
//		t.Valid = true
//		t.Value = v
//	} else if t.Value == nil {
//		t.Value = v
//		t.Valid = true
//	}
//}

func (t *Generic[T]) Ptr() *T {
	if !t.Valid {
		return nil
	}
	return &t.Value
}

func (t *Generic[T]) IsZero() bool {
	return !t.Valid
}
