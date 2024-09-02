package versiautils

func Must[In any, Out any](fn func(In) (Out, error), v In) Out {
	out, err := fn(v)
	if err != nil {
		panic(err)
	}

	return out
}
