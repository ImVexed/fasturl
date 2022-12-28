package fasturl

import (
	"net/url"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// sample input line from the blog post
var _url = "https://www.google.com/dir/1/2/search.html?arg=0-a&arg1=1-b&arg3-c#hash"

// make sure the benchmark isn't optimized away
var hits int
var reSSHD = regexp.MustCompile(`^(([^:/?#.]+):)?(//)?(([^:/]*)?(\\:([^/]*))?\\@)?(([^/:]+)|\\[[^/\\]]+\\])?(:(\\d*))?(/[^?#]*)(\\?([^#]*))?(#(.*))?`)

func BenchmarkRegex(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if len(reSSHD.FindAllStringSubmatch(_url, -1)) > 0 {
			hits++
		}
	}
}

func BenchmarkRagel(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if _, err := ParseURL(_url); err == nil {
			hits++
		}
	}
}

func BenchmarkStd(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if _, err := url.Parse(_url); err == nil {
			hits++
		}
	}
}

func TestParseURL(t *testing.T) {
	t.Run("Fail", func(t *testing.T) {
		url, err := ParseURL("I'm not a url")

		assert.Error(t, err)
		assert.Nil(t, url)
	})

	t.Run("Without HTTP", func(t *testing.T) {
		url, err := ParseURL("stackoverflow.com/questions/3771081/proper-way-to-check-for-url-equality")

		assert.NoError(t, err)
		assert.Equal(t, "stackoverflow.com", url.Host)
		assert.Equal(t, "/questions/3771081/proper-way-to-check-for-url-equality", url.Path)
		assert.Empty(t, url.Protocol)
		assert.Empty(t, url.Port)
		assert.Empty(t, url.Query)
		assert.Empty(t, url.Fragment)
	})

	t.Run("With HTTP", func(t *testing.T) {
		url, err := ParseURL("http://stackoverflow.com/questions/3771081/proper-way-to-check-for-url-equality")

		assert.NoError(t, err)
		assert.Equal(t, "stackoverflow.com", url.Host)
		assert.Equal(t, "/questions/3771081/proper-way-to-check-for-url-equality", url.Path)
		assert.Equal(t, "http", url.Protocol)
		assert.Empty(t, url.Port)
		assert.Empty(t, url.Query)
		assert.Empty(t, url.Fragment)
	})

	t.Run("Equality", func(t *testing.T) {
		url1, err1 := ParseURL("http://stackoverflow.com/questions/3771081/proper-way-to-check-for-url-equality")
		url2, err2 := ParseURL("http://stackoverflow.com/questions/3771081/proper-way-to-check-for-url-equality")

		assert.NoError(t, err1)
		assert.NoError(t, err2)
		assert.Equal(t, url2, url1)
	})

	t.Run("Extra Slash", func(t *testing.T) {
		url1, err1 := ParseURL("http://stackoverflow.com/questions/3771081/proper-way-to-check-for-url-equality")
		url2, err2 := ParseURL("http://stackoverflow.com/questions/3771081/proper-way-to-check-for-url-equality/")

		assert.NoError(t, err1)
		assert.NoError(t, err2)
		assert.NotEqual(t, url2, url1)
	})

	t.Run("FTP", func(t *testing.T) {
		url, err := ParseURL("ftp://ftp.is.co.za/rfc/rfc1808.txt")

		assert.NoError(t, err)
		assert.Equal(t, "ftp.is.co.za", url.Host)
		assert.Equal(t, "/rfc/rfc1808.txt", url.Path)
		assert.Equal(t, "ftp", url.Protocol)
		assert.Empty(t, url.Port)
		assert.Empty(t, url.Query)
		assert.Empty(t, url.Fragment)
	})

	t.Run("HTTP", func(t *testing.T) {
		url, err := ParseURL("http://www.ietf.org/rfc/rfc2396.txt")

		assert.NoError(t, err)
		assert.Equal(t, "www.ietf.org", url.Host)
		assert.Equal(t, "/rfc/rfc2396.txt", url.Path)
		assert.Equal(t, "http", url.Protocol)
		assert.Empty(t, url.Port)
		assert.Empty(t, url.Query)
		assert.Empty(t, url.Fragment)
	})

	t.Run("LDAP", func(t *testing.T) {
		url, err := ParseURL("ldap://www.ldap.org/c=GB?objectClass")

		assert.NoError(t, err)
		assert.Equal(t, "www.ldap.org", url.Host)
		assert.Equal(t, "/c=GB", url.Path)
		assert.Equal(t, "ldap", url.Protocol)
		assert.Equal(t, "objectClass", url.Query)
		assert.Empty(t, url.Port)
		assert.Empty(t, url.Fragment)
	})

	t.Run("MailTo", func(t *testing.T) {
		url, err := ParseURL("mailto:John.Doe@example.com")

		assert.NoError(t, err)
		assert.Equal(t, "example.com", url.Host)
		assert.Equal(t, "mailto", url.Protocol)
		assert.Empty(t, url.Path)
		assert.Empty(t, url.Port)
		assert.Empty(t, url.Query)
		assert.Empty(t, url.Fragment)
	})

	t.Run("News", func(t *testing.T) {
		url, err := ParseURL("news:comp.infosystems.www.servers.unix")

		assert.NoError(t, err)
		assert.Equal(t, "comp.infosystems.www.servers.unix", url.Host)
		assert.Equal(t, "news", url.Protocol)
		assert.Empty(t, url.Path)
		assert.Empty(t, url.Port)
		assert.Empty(t, url.Query)
		assert.Empty(t, url.Fragment)
	})

	t.Run("Tel", func(t *testing.T) {
		url, err := ParseURL("tel:+1-816-555-1212")

		assert.NoError(t, err)
		assert.Equal(t, "+1-816-555-1212", url.Host)
		assert.Equal(t, "tel", url.Protocol)
		assert.Empty(t, url.Path)
		assert.Empty(t, url.Port)
		assert.Empty(t, url.Query)
		assert.Empty(t, url.Fragment)
	})

	t.Run("Telnet", func(t *testing.T) {
		url, err := ParseURL("telnet://192.0.2.16:80/")

		assert.NoError(t, err)
		assert.Equal(t, "192.0.2.16", url.Host)
		assert.Equal(t, "/", url.Path)
		assert.Equal(t, "telnet", url.Protocol)
		assert.Equal(t, "80", url.Port)
		assert.Empty(t, url.Query)
		assert.Empty(t, url.Fragment)
	})

	t.Run("With query only", func(t *testing.T) {
		url, err := ParseURL("http://example.com?foo=bar")

		require.NoError(t, err)
		assert.Equal(t, "example.com", url.Host)
		assert.Equal(t, "foo=bar", url.Query)
	})
}
