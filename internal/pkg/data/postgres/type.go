package postgres

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/tanveerprottoy/stdlib-go-template/pkg/stringsext"
)

// compose sql.NullString in NullString
type NullString struct {
	sql.NullString
}

// MarshalJSON for NullString
func (s NullString) MarshalJSON() ([]byte, error) {
	if !s.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(s.String)
}

// UnmarshalJSON for NullString
func (s *NullString) UnmarshalJSON(bytes []byte) error {
	// Unmarshalling into a pointer will let us detect null
	var v *string
	if err := json.Unmarshal(bytes, &v); err != nil {
		return err
	}
	if v != nil {
		s.Valid = true
		s.String = *v
	} else {
		s.Valid = false
	}
	return nil
}

// compose sql.NullInt64 in NullInt64
type NullInt64 struct {
	sql.NullInt64
}

// MarshalJSON for NullInt64
func (i NullInt64) MarshalJSON() ([]byte, error) {
	if !i.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(i.Int64)
}

// UnmarshalJSON for NullInt64
func (i *NullInt64) UnmarshalJSON(bytes []byte) error {
	// Unmarshalling into a pointer will let us detect null
	var v *int64
	if err := json.Unmarshal(bytes, &v); err != nil {
		return err
	}
	if v != nil {
		i.Valid = true
		i.Int64 = *v
	} else {
		i.Valid = false
	}
	return nil
}

// compose sql.NullFloat64 in NullInt64
type NullFloat64 struct {
	sql.NullFloat64
}

// MarshalJSON for NullInt64
func (f NullFloat64) MarshalJSON() ([]byte, error) {
	if !f.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(f.Float64)
}

// UnmarshalJSON for NullInt64
func (f *NullFloat64) UnmarshalJSON(bytes []byte) error {
	// Unmarshalling into a pointer will let us detect null
	var v *float64
	if err := json.Unmarshal(bytes, &v); err != nil {
		return err
	}
	if v != nil {
		f.Valid = true
		f.Float64 = *v
	} else {
		f.Valid = false
	}
	return nil
}

// compose sql.NullBool in NullBool
type NullBool struct {
	sql.NullBool
}

// MarshalJSON for NullBool
func (b NullBool) MarshalJSON() ([]byte, error) {
	if !b.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(b.Bool)
}

// UnmarshalJSON for NullBool
func (b *NullBool) UnmarshalJSON(bytes []byte) error {
	// Unmarshalling into a pointer will let us detect null
	var v *bool
	if err := json.Unmarshal(bytes, &v); err != nil {
		return err
	}
	if v != nil {
		b.Valid = true
		b.Bool = *v
	} else {
		b.Valid = false
	}
	return nil
}

// Point is a struct for DB point type
type Point struct {
	// Latitude is the Y axis, longitude is the X axis
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

// Scan implements the Scanner interface for Point
func (u *Point) Scan(value any) error {
	u.Lat = 0
	u.Lng = 0
	if value == nil {
		return nil
	}
	switch v := value.(type) {
	case []byte:
		// Unmarshalling into a pointer will let us detect null
		// need to manipulate the supplied value as it is a base64 encoded string
		s := string(v)
		// remove parentheses from the string
		s = s[1 : len(s)-1]
		// alternatively, we could use a regex to remove the parentheses
		// re := regexp.MustCompile(`\(([^)]+)\)`)
		// d = re.ReplaceAllString(string(bytes), "$1")
		// split the string by comma
		splits := stringsext.Split(s, ",")
		// first one is the longitude
		lngStr := splits[0]
		// second one is the latitude
		latStr := splits[1]
		lng, err := strconv.ParseFloat(lngStr, 64)
		if err != nil {
			return err
		}
		lat, err := strconv.ParseFloat(latStr, 64)
		if err != nil {
			return err
		}
		u.Lng = lng
		u.Lat = lat
		return nil
	}
	return fmt.Errorf("unsupported Scan, storing driver.Value type %T into type *UInt", value)
}
