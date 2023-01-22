package set

type StringSet map[string]bool

func NewStringSet(stringSlice []string) StringSet {
	s := StringSet{}

	for _, v := range stringSlice {
		s[v] = true
	}

	return s
}

func (s StringSet) Has(val string) bool {
	_, has := s[val]
	return has
}

func (s StringSet) Add(val string) {
	s[val] = true;
}

func (s StringSet) Remove(val string) {
	delete(s, val)
}

// Keys returns a string slice of all set keys
func (s StringSet) Keys() []string {
	var keys []string

	for k := range s{
		keys = append(keys, k)
	}
	return keys
}


type IntSet map[int]bool

func NewIntSet(intSlice []int) IntSet {
	s := IntSet{}

	for _, v := range intSlice {
		s[v] = true
	}

	return s
}

func (s IntSet) Has(val int) bool {
	_, has := s[val]
	return has
}

func (s IntSet) Add(val int) {
	s[val] = true;
}

func (s IntSet) Remove(val int) {
	delete(s, val)
}

func (s IntSet) Keys() []int {
	var keys []int
	for k := range s {
		keys = append(keys, k)
	}
	return keys
}
