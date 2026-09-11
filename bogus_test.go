package main

import (
	"regexp"
	"strings"
	"testing"
)

func TestRandomStringLength(t *testing.T) {
	lengths := []int{1, 6, 10, 32}
	for _, l := range lengths {
		s := randomString(l)
		if len(s) != l {
			t.Errorf("expected length %d, got %d", l, len(s))
		}
	}
}

func TestRandomStringCharset(t *testing.T) {
	re := regexp.MustCompile("^[a-z0-9]{6}$")
	for i := 0; i < 50; i++ {
		s := randomString(6)
		if !re.MatchString(s) {
			t.Errorf("random string %q contains invalid characters", s)
		}
	}
}

func TestRandomStringUniqueness(t *testing.T) {
	seen := make(map[string]bool)
	count := 100
	for i := 0; i < count; i++ {
		s := randomString(6)
		seen[s] = true
	}
	if len(seen) < count {
		t.Errorf("expected %d unique strings, got %d", count, len(seen))
	}
}

func TestGenerateBranchName(t *testing.T) {
	branch := generateBranchName()
	parts := strings.Split(branch, "-")
	if len(parts) < 2 {
		t.Fatalf("expected branch name to contain at least one hyphen separator, got %q", branch)
	}

	suffix := parts[len(parts)-1]
	if len(suffix) != 6 {
		t.Errorf("expected suffix length 6, got %d in %q", len(suffix), branch)
	}

	re := regexp.MustCompile("^[a-z0-9]{6}$")
	if !re.MatchString(suffix) {
		t.Errorf("expected alphanumeric suffix, got %q in %q", suffix, branch)
	}
}
