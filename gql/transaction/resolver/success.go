package resolver

type Success struct {
}

func (s *Success) Success() bool {
	return true
}
