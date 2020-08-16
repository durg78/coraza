package persistence

type MemoryEngine struct{
	data map[string]map[string][]string
}

func (r *MemoryEngine) Init() error{
	return nil
}

func (r *MemoryEngine) Get(key string) map[string][]string{
	return map[string][]string{}
}

func (r *MemoryEngine) Set(key string, data map[string][]string) error{
	return nil
}