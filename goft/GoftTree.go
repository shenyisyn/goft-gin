package goft

type RouteInfo struct {
	Method      string
	Path        string
	Handler     string
	HandlerFunc interface{}
}
type RoutesInfo []RouteInfo
type GoftTree struct {
	trees methodTrees
}

func NewGoftTree() *GoftTree {
	tree := &GoftTree{
		trees: make(methodTrees, 0, 9),
	}
	return tree
}
func (this *GoftTree) addRoute(method, path string, handlers interface{}) {
	root := this.trees.get(method)
	if root == nil {
		root = new(node)
		root.fullPath = "/"
		this.trees = append(this.trees, methodTree{method: method, root: root})
	}
	root.addRoute(path, handlers)
}
func (this *GoftTree) getRoute(httpMethod, path string) nodeValue {
	t := this.trees
	for i, tl := 0, len(t); i < tl; i++ {
		if t[i].method != httpMethod {
			continue
		}
		root := t[i].root
		// Find route in tree
		value := root.getValue(path, nil, false)
		if value.handlers != nil {
			return value
		}
	}
	return nodeValue{}
}
