package tcase

import (
	"regexp"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestReg(t *testing.T) {
	tests := []struct {
		ss   string
		reg  *regexp.Regexp
		want bool
	}{
		{
			ss:   "交易异常，商户号:Q0001507，订单号:217789763400192，错误码:3576，错误信息:3576(3576)实时支付邮路错误(对方未登录)",
			reg:  regexp.MustCompile(`交易异常，商户号:Q[\d]+，订单号:[\d]+，错误码:3576，错误信息:3576\(3576\)实时支付邮路错误`),
			want: true,
		},
	}
	for _, tc := range tests {
		got := tc.reg.MatchString(tc.ss)
		if diff := cmp.Diff(tc.want, got); diff != "" {
			t.Errorf("diff: (-want +got)\n%s", diff)
		}
	}
}

func TestMap(t *testing.T) {
	tests := []struct {
		key  string
		ss   []string
		want string
	}{
		{
			key:  "test",
			ss:   []string{"aa", "bb", "cc", "dd"},
			want: "bb",
		},
	}
	m := make(map[string]string)
	for _, tc := range tests {
		for _, v := range tc.ss {
			m[tc.key] = v
		}
		got := m[tc.key]
		if diff := cmp.Diff(tc.want, got); diff != "" {
			t.Errorf("diff: (-want +got)\n%s", diff)
		}
	}
}
