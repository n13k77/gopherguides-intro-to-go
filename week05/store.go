package demo

type Model map[string]interface{} // collection of any type of thing, each one identified by a string

type Models []Model // slice of Model

type data map[string]Models // map of slice of Model, using string as a key 

type Store struct {
	data data
}

func (s *Store) db() data {
	if s.data == nil {
		s.data = data{}
	}

	return s.data
}

func (s *Store) All(tn string) (Models, error) {
	db := s.db()

	mods, ok := db[tn]
	if !ok {
		return nil, &ErrTableNotFound{table: tn}
	}

	return mods, nil
}

func (s Store) Len(tn string) (int, error) {
	rows, err := s.All(tn)
	if err != nil {
		return 0, err
	}

	return len(rows), nil
}

func (s *Store) Insert(tn string, mod ...Model) {
	db := s.db()
	db[tn] = append(db[tn], mod...)
}

func (s *Store) Select(tn string, query Clauses) (Models, error) {
	rows, err := s.All(tn)
	if err != nil {
		return nil, err
	}

	if len(query) == 0 {
		return rows, nil
	}

	res := make(Models, 0, len(rows))

	for _, m := range rows {
		if query.Match(m) {
			res = append(res, m)
		}
	}

	if len(res) == 0 {
		return nil, &errNoRows{
			clauses: query,
			table:   tn,
		}
	}

	return res, nil
}
