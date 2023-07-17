package goLeetCode

func IsIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	s_to_t, t_to_s := map[byte]byte{}, map[byte]byte{}
	for i := range s {
		s_to_t_val, ok1 := s_to_t[s[i]]
		t_to_s_val, ok2 := t_to_s[t[i]]

		if ok1 != ok2 {
			return false
		} else if ok1 && ok2 {
			if s_to_t_val != t[i] || t_to_s_val != s[i] {
				return false
			}
		} else {
			s_to_t[s[i]] = t[i]
			t_to_s[t[i]] = s[i]
		}
	}
	return true
}
