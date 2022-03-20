package github

type Storage struct {
	spec *spec
}

func NewStorage(m map[string]interface{}) (*Storage, error) {
	s, err := wrapSpec(m)
	if err != nil {
		return nil, err
	}
	return &Storage{s}, nil
}

func (s *Storage) Get() (string, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Storage) Set(dir string) error {
	//TODO implement me
	panic("implement me")
}

func (s *Storage) Delete(dir string) error {
	//TODO implement me
	panic("implement me")
}

func (s *Storage) List() (map[string]interface{}, error) {
	//TODO implement me
	panic("implement me")
}
