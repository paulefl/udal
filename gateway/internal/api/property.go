package api

import (
	"encoding/json"
	"fmt"
)

// PropertyStore stores current property values per device.
// In production this is backed by the transport adapter; the in-memory
// implementation is used by unit tests.
type PropertyStore interface {
	Get(deviceID, propertyPath string) (PropertyValue, error)
	Set(deviceID, propertyPath string, v PropertyValue) error
}

// MemoryPropertyStore is an in-memory PropertyStore for tests.
type MemoryPropertyStore struct {
	data map[string]PropertyValue // key: "deviceID/propertyPath"
}

func NewMemoryPropertyStore() *MemoryPropertyStore {
	return &MemoryPropertyStore{data: make(map[string]PropertyValue)}
}

func (s *MemoryPropertyStore) key(deviceID, propertyPath string) string {
	return deviceID + "/" + propertyPath
}

func (s *MemoryPropertyStore) Get(deviceID, propertyPath string) (PropertyValue, error) {
	v, ok := s.data[s.key(deviceID, propertyPath)]
	if !ok {
		return PropertyValue{}, fmt.Errorf("property %q not found on device %q", propertyPath, deviceID)
	}
	return v, nil
}

func (s *MemoryPropertyStore) Set(deviceID, propertyPath string, v PropertyValue) error {
	s.data[s.key(deviceID, propertyPath)] = v
	return nil
}

// FloatValue is a convenience constructor.
func FloatValue(f float64) PropertyValue { return PropertyValue{FloatVal: &f} }

// IntValue is a convenience constructor.
func IntValue(i int64) PropertyValue { return PropertyValue{IntVal: &i} }

// BoolValue is a convenience constructor.
func BoolValue(b bool) PropertyValue { return PropertyValue{BoolVal: &b} }

// StringValue is a convenience constructor.
func StringValue(s string) PropertyValue { return PropertyValue{StringVal: &s} }

// JSONValue is a convenience constructor for structured values.
func JSONValue(v any) (PropertyValue, error) {
	b, err := json.Marshal(v)
	if err != nil {
		return PropertyValue{}, err
	}
	return PropertyValue{JSONVal: b}, nil
}
