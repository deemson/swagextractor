package openapi

type Paths map[string]Path

type Path struct {
	Get  Get
	Post Post
}

type Get struct {
}

type Post struct {
}
