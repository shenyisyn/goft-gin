package lib

import "testing"

func TestComparableExpr_filter(t *testing.T) {
	tests := []struct {
		name string
		this ComparableExpr
		want string
	}{
		{"w1", ComparableExpr("a>3"), "gt .a 3"},
		{"w2", ComparableExpr("a==3"), "eq .a 3"},
		{"w3", ComparableExpr("a>=3"), "ge .a 3"},
		{"w4", ComparableExpr("a<3"), "lt .a 3"},
		{"w5", ComparableExpr("a>3"), "gt .a 3"},
		{"w6", ComparableExpr("a<>3"), ""},
		{"w7", ComparableExpr("ab3"), ""},
		{"w8", ComparableExpr("7>8"), "gt 7 8"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.this.filter(); got != tt.want {
				t.Errorf("ComparableExpr.filter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getCompareToken(t *testing.T) {
	type args struct {
		sign string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getCompareToken(tt.args.sign); got != tt.want {
				t.Errorf("getCompareToken() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseToken(t *testing.T) {
	type args struct {
		token string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseToken(tt.args.token); got != tt.want {
				t.Errorf("parseToken() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsComparableExpr(t *testing.T) {
	type args struct {
		expr string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"w1", args{"a>3"}, true},
		{"w2", args{"a==3"}, true},
		{"w3", args{"a>=3"}, true},
		{"w4", args{"a<3"}, true},
		{"w5", args{"a>3"}, true},
		{"w6", args{"a<>3"}, false},
		{"w7", args{"ab3"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsComparableExpr(tt.args.expr); got != tt.want {
				t.Errorf("IsComparableExpr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExecExpr(t *testing.T) {
	type args struct {
		expr string
		data map[string]interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ExecExpr(tt.args.expr, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExecExpr() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ExecExpr() = %v, want %v", got, tt.want)
			}
		})
	}
}
