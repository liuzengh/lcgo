package strings

// IsIsomorphic return true if two strings s and t are isomorphic, or false otherwise.
// Two strings s and t are isomorphic if the characters in s can be replaced to get t
func IsIsomorphic(s string, t string) bool {
	s2t := make(map[byte]byte)
	for i := range s {
		if value, ok := s2t[s[i]]; ok && value != t[i] {
			return false
		}
		s2t[s[i]] = t[i]
	}

	t2s := make(map[byte]byte)
	for k, v := range s2t {
		if _, ok := t2s[v]; ok {
			return false
		}
		t2s[v] = k
	}
	return true
}

// IsSubsequence return true if s is a subsequence of t, or false otherwise.
// A subsequence of a string is a new string that is formed from the original string by deleting some (can be none)
// of the characters without disturbing the relative positions of the remaining characters.
func IsSubsequence(s string, t string) bool {
	var i int
	for j := range s {
		for i < len(t) && t[i] != s[j] {
			i += 1
		}
		if i == len(t) {
			return false
		}
	}
	return true
}
