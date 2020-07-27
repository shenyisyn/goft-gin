package goft

import (
	"fmt"
	"strings"
)

type TreeNode struct {
	path       string
	part       string
	children   map[string]*TreeNode
	isWildcard bool
}
type TreeRouter struct {
	root  map[string]*TreeNode
	route map[string]interface{}
}

func NewTreeRouter() *TreeRouter {
	return &TreeRouter{root: make(map[string]*TreeNode), route: make(map[string]interface{})}
}
func pathes(path string) (ret []string) {
	pathList := strings.Split(path, "/")
	for _, p := range pathList {
		if p != "" {
			ret = append(ret, p)
			if p[0] == '*' {
				break
			}
		}
	}
	return
}
func (r *TreeRouter) addRoute(reqMethod string, reqPath string, handler interface{}) {
	pathList := pathes(reqPath)
	if _, ok := r.root[reqMethod]; !ok {
		r.root[reqMethod] = &TreeNode{children: make(map[string]*TreeNode)}
	}
	root := r.root[reqMethod]
	key := genRouteKey(reqMethod, reqPath)
	for _, p := range pathList {
		if root.children[p] == nil {
			root.children[p] = &TreeNode{
				part:       p,
				children:   make(map[string]*TreeNode),
				isWildcard: p[0] == ':' || p[0] == '*'}
		}
		root = root.children[p]
	}
	root.path = reqPath
	r.route[key] = handler
}
func genRouteKey(reqMethod string, reqPath string) string {
	return fmt.Sprintf("%s_%s", reqMethod, reqPath)
}
func (r *TreeRouter) getRoute(reqMethod string, reqPath string) (node *TreeNode, params map[string]string) {
	params = map[string]string{}
	searchPath := pathes(reqPath)
	var ok bool
	if node, ok = r.root[reqMethod]; !ok {
		return nil, nil
	}
	for i, part := range searchPath {
		var temp string
		for _, child := range node.children {
			if child.part == part || child.isWildcard {
				if child.part[0] == '*' {
					params[child.part[1:]] = strings.Join(searchPath[i:], "/")
				}
				if child.part[0] == ':' {
					params[child.part[1:]] = part
				}
				temp = child.part
			}
		}
		if temp[0] == '*' {
			return node.children[temp], params
		}
		node = node.children[temp]
	}
	return
}
func (r *TreeRouter) getHandler(reqMethod string, reqPath string) interface{} {
	if h, ok := r.route[genRouteKey(reqMethod, reqPath)]; ok {
		return h
	}
	return nil
}
