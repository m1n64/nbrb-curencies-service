package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"gorm.io/gorm"
	"time"
)

type Rate struct {
	gorm.Model
	CurrencyID   uint       `json:"Cur_ID"`
	Date         CustomTime `gorm:"uniqueKey" json:"Date"`
	Scale        int        `json:"Cur_Scale"`
	OfficialRate float64    `json:"Cur_OfficialRate"`
	Currency     Currency   `gorm:"foreignKey:CurrencyID"`
}

// MarshalJSON generates a JSON representation of the Rate struct.
//
// Returns:
// - []byte: The JSON byte slice representing the Rate.
// - error: An error if the JSON marshaling fails.
func (r *Rate) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"id":            r.ID,
		"currency_id":   r.CurrencyID,
		"date":          r.Date,
		"scale":         r.Scale,
		"official_rate": r.OfficialRate,
		"currency": map[string]interface{}{
			"id":            r.Currency.ID,
			"currency_abbr": r.Currency.Abbreviation,
			"currency_name": r.Currency.Name,
		},
	})
}

type CustomTime struct {
	time.Time
}

// UnmarshalJSON parses the JSON-encoded data and stores the result in the CustomTime struct.
//
// Parameters:
// - b: A byte slice containing the JSON data to unmarshal.
// Returns:
// -
func (ct *CustomTime) UnmarshalJSON(b []byte) error {
	s := string(b)
	s = s[1 : len(s)-1]

	t, err := time.Parse("2006-01-02T15:04:05", s)
	if err != nil {
		return err
	}

	ct.Time = t
	return nil
}

// Value returns the custom time as a value suitable for database storage.
func (ct CustomTime) Value() (driver.Value, error) {
	return ct.Time, nil
}

// Scan converts a value from the database to a CustomTime struct.
func (ct *CustomTime) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	if t, ok := value.(time.Time); ok {
		ct.Time = t
		return nil
	}
	return errors.New("failed to scan CustomTime from database")
}
