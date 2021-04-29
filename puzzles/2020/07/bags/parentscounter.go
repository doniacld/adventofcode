package bags

// CountPossiblePaths returns the number of possible bags paths to get the target bag
func (r Bags) CountPossiblePaths(target string) int {
	var count int
	// maintain a list of the seen bags to avoid to count duplicates
	s := make(seen, 0)

	// loop until the visit of all paths ends
	for true {
		v, ok := r[target]
		if !ok {
			return count
		}
		// browse the parents of the target
		return r.browseParents(v.parents, s, count)
	}
	// should not happened ?
	return count
}

// browseParents browse the parents of a target bag
// and counts the number of ways to access this bag
func (r Bags) browseParents(v []bag, s seen, counter int) int {
	// browse the list of parents of the target bag
	for i := 0; i < len(v); i++ {
		// make the parent the new target
		target := v[i].name

		// has no parent and is not seen
		if !r.hasParent(target) && !s.isSeen(target) {
			s.insertSeen(target)
			counter++
			continue
		}

		// has parents and is not seen
		if !s.isSeen(target) {
			s.insertSeen(target)
			counter++
			counter = r.browseParents(r[target].parents, s, counter)
		}
		// if seen, we just continue and do nothing
	}
	return counter
}

// hasParent returns true if a target has parents or false if it is a root
func (r Bags) hasParent(target string) bool {
	if v, ok := r[target]; ok {
		return len(v.parents) != 0
	}
	return false
}

// seen holds a map of seen bags
type seen map[string]struct{}

// insertSeen inserts a bag into the seen map
func (s seen) insertSeen(bag string) {
	s[bag] = struct{}{}
}

// isSeen returns true if the target is already seen
func (s seen) isSeen(target string) bool {
	if _, ok := s[target]; ok {
		return true
	}
	return false
}

