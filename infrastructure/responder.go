package infrastructure

type Responder struct {
	Meta map[string]any `json:"meta"`
	Data interface{}    `json:"data"`
}

func NewResponder() *Responder {
	return &Responder{}
}

func (s Responder) Success(data interface{}) Responder {
	s.Meta = make(map[string]any)
	s.Meta["success"] = true
	s.Data = data

	return s
}

func (s Responder) SuccessList(total int, limit int, offset int, data interface{}) Responder {
	s.Meta = make(map[string]any)
	s.Meta["success"] = true
	s.Meta["total"] = total
	s.Meta["limit"] = limit
	s.Meta["offset"] = offset
	s.Data = data

	return s
}

func (s Responder) Fail(data interface{}) Responder {
	s.Meta = make(map[string]any)
	s.Meta["success"] = false
	s.Data = data

	return s
}
