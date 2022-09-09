package practices

type MergeNames struct {
}

func (mn MergeNames) UniqueNames(none, ntwo []string) []string {
	var res []string
	res = mn.removeDuplicates(none)
	ntwo = mn.removeDuplicates(ntwo)
	for _, n := range ntwo {
		if !mn.contain(n, none) {
			res = append(res, n)
		}
	}

	return res
}

func (mn MergeNames) removeDuplicates(names []string) []string {
	res := make([]string, 0)
	for _, n := range names {
		if !mn.contain(n, res) {
			res = append(res, n)
		}
	}
	return res
}

func (mn MergeNames) contain(name string, names []string) bool {
	for _, n := range names {
		if n == name {
			return true
		}
	}
	return false
}
