package openapi

type Paths map[string]Path

type Path struct {
	Get  Get  `yaml:"get,omitempty"`
	Post Post `yaml:"post,omitempty"`
}

type Get struct {
}

type Post struct {
}
