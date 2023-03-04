package fasturl

import (
	"net/url"
	"strings"
	"testing"
)

func FuzzNetURLDiff(f *testing.F) {
	f.Add("http://www.example.com:80/foo/bar?baz=qux#quux")

	f.Fuzz(func(t *testing.T, data string) {
		u, err := ParseURL(data)
		if err != nil {
			return
		}

		if !strings.HasPrefix(data, "http://") && !strings.HasPrefix(data, "https://") {
			return
		}
		if strings.Count(data, ":") > 2 {
			return
		}
		if strings.Contains(data, "///") {
			return
		}
		if strings.Contains(data, ":@") {
			return
		}
		if strings.HasSuffix(data, "@") {
			return
		}
		if strings.Contains(data, "%") {
			return
		}
		if strings.Contains(u.Path, ":/") {
			return
		}

		ug, err := url.Parse(data)
		if err != nil {
			return
		}

		if ug.Host == "" {
			return
		}

		if u.Host != ug.Hostname() {
			t.Errorf("host mismatch %q: %q != %q", data, u.Host, ug.Host)
		}
		if u.Path != ug.Path {
			t.Errorf("path mismatch %q: %q != %q", data, u.Path, ug.Path)
		}
		if u.Query != ug.RawQuery {
			t.Errorf("query mismatch %q: %q != %q", data, u.Query, ug.RawQuery)
		}
		if u.Fragment != ug.Fragment {
			t.Errorf("fragment mismatch %q: %q != %q", data, u.Fragment, ug.Fragment)
		}
	})
}
