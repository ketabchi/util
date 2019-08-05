package util

func CheckSliceEq(ss1, ss2 []string) bool {
	if len(ss1) != len(ss2) {
		return false
	}

	for _, s1 := range ss1 {
		found := false
		for _, s2 := range ss2 {
			if s1 == s2 {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func SliceContains(ss []string, s string) bool {
	for _, s1 := range ss {
		if s1 == s {
			return true
		}
	}
	return false
}
