package httpclientcache

import (
	"crypto/md5"
	"encoding/hex"
	"net/http"
	"sort"
	"strings"
)

type KeyGen struct {
	// Whether to use the protocol in the cache key
	IncludeProtocol bool
	// Whether to use the host in the cache key
	IncludeHost bool
	// Whether to use the path in the cache key
	IncludePath bool
	// Whether to use the query in the cache key
	IncludeQuery bool
	// Query keys to be included in the cache key (if IncludeQuery is true)
	QueryKeys []string
}

func (t *KeyGen) KeyStringHashHex(req *http.Request) string {
	h := md5.Sum([]byte(t.KeyString(req)))
	return hex.EncodeToString(h[:])
}

func (t *KeyGen) KeyString(req *http.Request) string {
	s := ""
	if t.IncludeProtocol {
		s += req.URL.Scheme
		s += "://"
	}
	if t.IncludeHost {
		s += req.URL.Host
	}
	if t.IncludePath {
		if !strings.HasPrefix(req.URL.Path, "/") {
			s += "/"
		}
		s += req.URL.Path
	}
	if t.IncludeQuery {
		query := req.URL.Query()
		if len(query) <= 0 {
			return s
		}
		queryKeys := []string{}
		for _, queryKey := range t.QueryKeys {
			queryKeys = append(queryKeys, queryKey)
		}
		sort.Strings(queryKeys)

		queries := []string{}
		for _, queryKey := range queryKeys {
			queryValues := make([]string, len(query[queryKey]))
			copy(queryValues, query[queryKey])
			sort.Strings(queryValues)
			for _, queryValue := range queryValues {
				queries = append(queries, queryKey+"="+queryValue)
			}
		}

		if len(queries) > 0 {
			s += "?"
			s += strings.Join(queries, "&")
		}
	}
	return s
}
