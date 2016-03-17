package models

import "encoding/json"

//Cycle ...
type Cycle struct {
	Operation int
	Time      string
	Duration  int
}

//Chart model, provides JSON data to represent query
func (db *DB) Chart(v int, d int) ([]byte, error) {
	rows, err := db.Query("SELECT Operation, Time, Duration FROM Doors WHERE Vehicle = ? and Position = ?;", v, d)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cycles []*Cycle

	for rows.Next() {
		c := new(Cycle)
		err := rows.Scan(&c.Operation, &c.Time, &c.Duration)
		if err != nil {
			return nil, err
		}
		cycles = append(cycles, c)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return marshall(cycles)
}

func marshall(c []*Cycle) ([]byte, error) {
	b, err := json.Marshal(c)
	if err != nil {
		return nil, err
	}

	return b, nil
}
