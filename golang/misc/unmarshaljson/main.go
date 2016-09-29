package main

import (
	"encoding/json"
	"time"
)

type req struct {
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

// UnmarshalJSON
func (r *req) UnmarshalJSON(data []byte) error {
	type Alias req
	aux := &struct {
		CreatedAt string `json:"created_at"`
		*Alias
	}{
		Alias: (*Alias)(r),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	t, err := time.Parse("2006-01-02 15:04:05 MST", aux.CreatedAt)
	if err != nil {
		return err
	}
	r.CreatedAt = t
	return nil
}
