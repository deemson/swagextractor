package swagextractor

import (
	"github.com/deemson/swagextractor/internal/lib/openapi"
	"testing"
)

func TestS(t *testing.T) {
	testCases := map[string]struct {
		source []string
		spec   []string
	}{
		"asd": {
			source: []string{
				`package test`,
				``,
				`func f() {`,
				`	r := router()`,
				`	r.Methods("GET", "POST").Path("/path").Handle(nil)`,
				`	return r`,
				`}`,
			},
			spec: openapi.Spec{
				Paths: openapi.Paths{
					"/path": openapi.Path{
						Get:  openapi.Get{},
						Post: openapi.Post{},
					},
				},
			},
		},
	}
	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {

		})
	}
}
