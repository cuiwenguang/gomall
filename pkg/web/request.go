package web

type RequestContext struct {
	Host  string
	Token string
	User  map[string]string
}
