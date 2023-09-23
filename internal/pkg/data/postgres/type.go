package postgres

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/tanveerprottoy/stdlib-go-template/internal/pkg/constant"
	"github.com/tanveerprottoy/stdlib-go-template/pkg/stringsext"
)

// compose sql.NullString in NullString
type NullString struct {
	sql.NullString
}

func MakeNullString(s string, v bool) NullString {
	return NullString{sql.NullString{String: s, Valid: v}}
}

// Scan implements the Scanner interface for NullString
/* func (n *NullString) Scan(val any) error {
	var str string
	switch v := val.(type) {
	case []byte:
		err := json.Unmarshal(v, &str)
		if err != nil {
			return err
		}
		n.String = str
		return nil
	}
	return fmt.Errorf("unsupported Scan, storing driver.Value type %T into type *UInt", val)
} */

// MarshalJSON for NullString
func (n NullString) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(n.String)
}

// UnmarshalJSON for NullString
func (n *NullString) UnmarshalJSON(bytes []byte) error {
	// Unmarshalling into a pointer will let us detect null
	var v *string
	if err := json.Unmarshal(bytes, &v); err != nil {
		return err
	}
	if v != nil {
		n.Valid = true
		n.String = *v
	} else {
		n.Valid = false
	}
	return nil
}

/* func (n *NullString) Value() (driver.Value, error) {
    if !n.NullString.Valid {
        return nil, nil
    }
    return n.NullString.String, nil
} */

// compose sql.NullInt64 in NullInt64
type NullInt64 struct {
	sql.NullInt64
}

// MarshalJSON for NullInt64
func (n NullInt64) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(n.Int64)
}

// UnmarshalJSON for NullInt64
func (n *NullInt64) UnmarshalJSON(bytes []byte) error {
	// Unmarshalling into a pointer will let us detect null
	var v *int64
	if err := json.Unmarshal(bytes, &v); err != nil {
		return err
	}
	if v != nil {
		n.Valid = true
		n.Int64 = *v
	} else {
		n.Valid = false
	}
	return nil
}

// compose sql.NullFloat64 in NullInt64
type NullFloat64 struct {
	sql.NullFloat64
}

// MarshalJSON for NullInt64
func (n NullFloat64) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(n.Float64)
}

// UnmarshalJSON for NullInt64
func (n *NullFloat64) UnmarshalJSON(bytes []byte) error {
	// Unmarshalling into a pointer will let us detect null
	var v *float64
	if err := json.Unmarshal(bytes, &v); err != nil {
		return err
	}
	if v != nil {
		n.Valid = true
		n.Float64 = *v
	} else {
		n.Valid = false
	}
	return nil
}

// compose sql.NullBool in NullBool
type NullBool struct {
	sql.NullBool
}

// MarshalJSON for NullBool
func (n NullBool) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(n.Bool)
}

// UnmarshalJSON for NullBool
func (n *NullBool) UnmarshalJSON(bytes []byte) error {
	// Unmarshalling into a pointer will let us detect null
	var v *bool
	if err := json.Unmarshal(bytes, &v); err != nil {
		return err
	}
	if v != nil {
		n.Valid = true
		n.Bool = *v
	} else {
		n.Valid = false
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
func (p *Point) Scan(val any) error {
	// assume lat lng is not given
	p.Lat = constant.InvalidLatLng
	p.Lng = constant.InvalidLatLng
	if val == nil {
		p = nil
		return nil
	}
	switch v := val.(type) {
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
		p.Lng = lng
		p.Lat = lat
		return nil
	}
	return fmt.Errorf("unsupported Scan, storing driver.Value type %T into type *UInt", val)
}

// MarshalJSON for Point
func (p Point) MarshalJSON() ([]byte, error) {
	if p.Lat == constant.InvalidLatLng && p.Lng == constant.InvalidLatLng {
		return []byte("null"), nil
	}
	return json.Marshal(struct {
		Lat float64 `json:"lat"`
		Lng float64 `json:"lng"`
	}{
		Lat: p.Lat,
		Lng: p.Lng,
	})
}

/* func (p Point) Value() (driver.Value, error) {
	log.Println("Value() called.point: ", p)
	if p.Lat == constant.InvalidLatLng && p.Lng == constant.InvalidLatLng {
		return nil, nil
	}
	b, err := json.Marshal(p)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal Point value: %v", err)
	}
	return b, nil
} */

// JsonArray is a type for DB json array type
type JsonArray []map[string]any

// Scan implements the Scanner interface for Point
func (j *JsonArray) Scan(val any) error {
	var jsonData []map[string]any
	switch v := val.(type) {
	case []byte:
		err := json.Unmarshal(v, &jsonData)
		if err != nil {
			return err
		}
		*j = jsonData
		return nil
	}
	return fmt.Errorf("unsupported Scan, storing driver.Value type %T into type *UInt", val)
}

// JsonObject is a type for DB json array type
type JsonObject map[string]any

// Scan implements the Scanner interface for JsonObject
func (j *JsonObject) Scan(val any) error {
	var jsonObj map[string]any
	switch v := val.(type) {
	case []byte:
		err := json.Unmarshal(v, &jsonObj)
		if err != nil {
			return err
		}
		*j = jsonObj
		return nil
	}
	return fmt.Errorf("unsupported Scan, storing driver.Value type %T into type *UInt", val)
}

// JsonTokenObject is a type for DB json array type
type JsonTokenObject struct {
	ExpiresIn   int64  `json:"expiresIn"`
	GeneratedAt int64  `json:"generatedAt"`
	Token       string `json:"token"`
}

// Scan implements the Scanner interface for JsonTokenObject
func (j *JsonTokenObject) Scan(val any) error {
	var jsonObj JsonTokenObject
	switch v := val.(type) {
	case []byte:
		err := json.Unmarshal(v, &jsonObj)
		if err != nil {
			return err
		}
		*j = jsonObj
		return nil
	}
	return fmt.Errorf("unsupported Scan, storing driver.Value type %T into type *UInt", val)
}

// JsonArray is a type for DB json array type
type JsonStringArray []string

// Scan implements the Scanner interface for Point
func (j *JsonStringArray) Scan(val any) error {
	var jsonData JsonStringArray
	switch v := val.(type) {
	case []byte:
		err := json.Unmarshal(v, &jsonData)
		if err != nil {
			return err
		}
		*j = jsonData
		return nil
	}
	return fmt.Errorf("unsupported Scan, storing driver.Value type %T into type *UInt", val)
}

// Json2dArray is a type for DB json array type
type Json2dArray [][]string

// Scan implements the Scanner interface for JsonTokenObject
func (j *Json2dArray) Scan(val any) error {
	var jsonData Json2dArray
	switch v := val.(type) {
	case []byte:
		err := json.Unmarshal(v, &jsonData)
		if err != nil {
			return err
		}
		*j = jsonData
		return nil
	}
	return fmt.Errorf("unsupported Scan, storing driver.Value type %T into type *UInt", val)
}
