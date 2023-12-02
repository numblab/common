package domain

import (
	"sort"
	"unicode/utf8"
)

const (
	Dot     = "."
	DotByte = '.'
)

type Matcher struct {
	set        *succinctSet
	isAllowALl bool
}

func NewMatcher(domainSuffix []string) *Matcher {
	domainList := make([]string, 0, len(domainSuffix))
	seen := make(map[string]struct{}, len(domainList))
	for _, domain := range domainSuffix {
		// FQDN domain
		domain = fqdn(domain)
		// filter duplicate domain
		if _, ok := seen[domain]; ok {
			continue
		}
		// allow all domain
		if domain == "." {
			return &Matcher{isAllowALl: true}
		}
		seen[domain] = struct{}{}
		domainList = append(domainList, reverseDomainSuffix(domain))
	}
	sort.Strings(domainList)
	return &Matcher{set: newSuccinctSet(domainList)}
}

func (m *Matcher) Match(domain string) bool {
	if m.isAllowALl {
		return true
	}
	domain = fqdn(domain)
	return m.set.Has(reverseDomain(domain))
}

func reverseDomain(domain string) string {
	l := len(domain)
	b := make([]byte, l)
	for i := 0; i < l; {
		r, n := utf8.DecodeRuneInString(domain[i:])
		i += n
		utf8.EncodeRune(b[l-i:], r)
	}
	return string(b) + Dot
}

func reverseDomainSuffix(domain string) string {
	l := len(domain)
	b := make([]byte, l+2)
	for i := 0; i < l; {
		r, n := utf8.DecodeRuneInString(domain[i:])
		i += n
		utf8.EncodeRune(b[l-i:], r)
	}
	b[l] = DotByte
	b[l+1] = prefixLabel
	return string(b)
}

func fqdn(domain string) string {
	if domain == "" || domain == Dot {
		return domain
	}
	if domain[len(domain)-1] == '.' {
		return domain
	}
	return domain + Dot
}
