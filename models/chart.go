package models

import "time"

//Cycle ...
type Cycle struct {
	Operation int
	Time      time.Time
	Duration  int
}

//Chart ...
func (db *DB) Chart(v int, d int) ([]*Cycle, error) {
	rows, err := db.Query("SELECT * FROM Doors WHERE Vehicle = 1304 and Position = 4;")
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
	return cycles, nil
}
