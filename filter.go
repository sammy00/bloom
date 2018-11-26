package bloom

type Filter struct{}

func (f *Filter) Add(data []byte) error { return nil }

func (f *Filter) Clear() {}

func (f *Filter) Loaded() bool { return false }

//func (f *Filter) MatchAndUpdate(data []byte)
func (f *Filter) Matches(data []byte) bool { return false }

func (f *Filter) Reload(msg *MessageLoad) {}

func (f *Filter) Snapshot() (*MessageLoad, error) {
	return nil, nil
}

func Load(msg *MessageLoad) *Filter {
	return nil
}

func New(N, tweak uint32, P float64, flag UpdateFlag) *Filter {
	return nil
}
