package request

func (r *Request) AsJson() ([]byte, error) {
	return json.Marshal(r)
}
