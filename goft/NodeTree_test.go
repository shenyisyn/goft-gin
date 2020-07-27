package goft

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"reflect"
	"testing"
)

func GenTestRouter() *TreeRouter {
	r := NewTreeRouter()
	r.addRoute("GET", "/", func(context *gin.Context) {
		context.String(200, "index")
	})
	r.addRoute("GET", "/user/:id", func(context *gin.Context) {
		context.String(200, "userdetail")
	})
	r.addRoute("GET", "/user/:id/profile", nil)
	r.addRoute("GET", "/news/kind/:kind", nil)
	r.addRoute("GET", "/news", nil)
	return r
}
func GenTestGoftRouter() *GoftTree {
	tmpHandler := func() {}
	r := NewGoftTree()
	r.addRoute("GET", "/", tmpHandler)
	r.addRoute("GET", "/v1/", tmpHandler)
	r.addRoute("GET", "/v1", tmpHandler)
	r.addRoute("GET", "/user/:id", tmpHandler)
	r.addRoute("GET", "/user/:id/profile", tmpHandler)
	r.addRoute("GET", "/news/kind/:kind", tmpHandler)
	r.addRoute("GET", "/news", tmpHandler)
	return r
}
func TestTreeRouter_getRoute(t *testing.T) {

	tests := []struct {
		name     string
		method   string
		path     string
		wantPath string
	}{
		{"r0", "GET", "/", "/"},
		{"r1", "GET", "/v1/", "/v1/"},
		{"r2", "GET", "/v1/?name=abc", "/v1/"},
		{"r3", "GET", "/?token=123&age=19", "/"},
		{"r4", "GET", "/user/123", "/user/:id"},
		{"r5", "GET", "/user/123?token=123", "/user/:id"},
		{"r6", "GET", "/user/zhangsan", "/user/:id"},
		{"r7", "GET", "/user/123/profile", "/user/:id/profile"},
		{"r8", "GET", "/user", ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := GenTestRouter()
			getNode, _ := r.getRoute(tt.method, tt.path)
			fmt.Println(r.getHandler(tt.method, getNode.path))

			if !reflect.DeepEqual(getNode.path, tt.wantPath) {
				t.Errorf("TreeRouter.getRoute() gotNode = %v, want %v", getNode.path, tt.wantPath)
			}
		})
	}
}

type testRequests []struct {
	path       string
	nilHandler bool
	route      string
	ps         Params
}

func checkRequests(t *testing.T, tree *node, requests testRequests, unescapes ...bool) {
	unescape := false
	if len(unescapes) >= 1 {
		unescape = unescapes[0]
	}

	for _, request := range requests {
		value := tree.getValue(request.path, nil, unescape)

		if value.handlers == nil {
			if !request.nilHandler {
				t.Errorf("handle mismatch for route '%s': Expected non-nil handle", request.path)
			}
		} else if request.nilHandler {
			t.Errorf("handle mismatch for route '%s': Expected nil handle", request.path)
		}

		if !reflect.DeepEqual(value.params, request.ps) {
			t.Errorf("Params mismatch for route '%s'", request.path)
		}
	}
}
func TestTreeWildcard(t *testing.T) {
	tree := &node{}
	routes := [...]string{
		"/",
		"/cmd/:tool/:sub",
		"/cmd/:tool/",
		"/src/*filepath",
		"/search/",
		"/search/:query",
		"/user_:name",
		"/user_:name/about",
		"/files/:dir/*filepath",
		"/doc/",
		"/doc/go_faq.html",
		"/doc/go1.html",
		"/info/:user/public",
		"/info/:user/project/:project",
	}
	for _, route := range routes {
		tree.addRoute(route, nil)
	}

	checkRequests(t, tree, testRequests{
		{"/", true, "/", nil},
		{"/cmd/test/", true, "/cmd/:tool/", Params{Param{Key: "tool", Value: "test"}}},
		{"/cmd/test", true, "", Params{Param{Key: "tool", Value: "test"}}},
		{"/cmd/test/3", true, "/cmd/:tool/:sub", Params{Param{Key: "tool", Value: "test"}, Param{Key: "sub", Value: "3"}}},
		{"/src/", true, "/src/*filepath", Params{Param{Key: "filepath", Value: "/"}}},
		{"/src/some/file.png", true, "/src/*filepath", Params{Param{Key: "filepath", Value: "/some/file.png"}}},
		{"/search/", true, "/search/", nil},
		{"/search/someth!ng+in+ünìcodé", true, "/search/:query", Params{Param{Key: "query", Value: "someth!ng+in+ünìcodé"}}},
		{"/search/someth!ng+in+ünìcodé/", true, "", Params{Param{Key: "query", Value: "someth!ng+in+ünìcodé"}}},
		{"/user_gopher", true, "/user_:name", Params{Param{Key: "name", Value: "gopher"}}},
		{"/user_gopher/about", true, "/user_:name/about", Params{Param{Key: "name", Value: "gopher"}}},
		{"/files/js/inc/framework.js", true, "/files/:dir/*filepath", Params{Param{Key: "dir", Value: "js"}, Param{Key: "filepath", Value: "/inc/framework.js"}}},
		{"/info/gordon/public", true, "/info/:user/public", Params{Param{Key: "user", Value: "gordon"}}},
		{"/info/gordon/project/go", true, "/info/:user/project/:project", Params{Param{Key: "user", Value: "gordon"}, Param{Key: "project", Value: "go"}}},
	})
}

func TestGoftTreeRouter_getRoute(t *testing.T) {
	tests := []struct {
		name     string
		method   string
		path     string
		wantPath string
	}{
		{"r0", "GET", "/", "/"},
		{"r1", "GET", "/v1", "/v1"},
		{"r2", "GET", "/v1/", "/v1/"},
		{"r3", "GET", "/?token=123&age=19", "/"},
		{"r4", "GET", "/user/123", "/user/:id"},
		{"r5", "GET", "/user/123", "/user/:id"},
		{"r6", "GET", "/user/zhangsan", "/user/:id"},
		{"r7", "GET", "/user/123/profile", "/user/:id/profile"},
		{"r8", "GET", "/user", ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := GenTestGoftRouter()
			getNode := r.getRoute(tt.method, tt.path)
			fmt.Println("show:", getNode.tsr)
			if !reflect.DeepEqual(getNode.fullPath, tt.wantPath) {
				t.Errorf("TreeRouter.getRoute() gotNode.path = %v, want %v", getNode.fullPath, tt.wantPath)
			}
		})
	}
}
