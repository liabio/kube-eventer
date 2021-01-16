package filters

func Skip(candidate string, options []string) (skip bool) {
	skip = true
	for _, op := range options {
		if op == candidate {
			return false
		}
		//exclude candidate
		if op == "-"+candidate {
			return true
		}
		if op == "*" {
			return false
		}
	}
	return
}
