package slice

type Identity[T any, V comparable] func(t T) V
