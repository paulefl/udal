// Package api defines the core domain types shared between gRPC handlers
// and the internal registry. These mirror the proto messages and will be
// replaced / augmented by generated proto structs once buf generate runs.
package api

import "time"

// DeviceStatus indicates the connectivity state of a registered device.
type DeviceStatus int

const (
	DeviceStatusUnknown DeviceStatus = iota
	DeviceStatusOnline
	DeviceStatusOffline
)

func (s DeviceStatus) String() string {
	switch s {
	case DeviceStatusOnline:
		return "online"
	case DeviceStatusOffline:
		return "offline"
	default:
		return "unknown"
	}
}

// Device represents a registered IoT device.
type Device struct {
	ID         string
	Name       string
	Capability string // capability schema name, e.g. "temperature-sensor"
	Transport  string // "mqtt" | "http" | "can"
	Status     DeviceStatus
	LastSeen   time.Time
	Labels     map[string]string
}

// PropertyValue is a discriminated union for typed property values.
type PropertyValue struct {
	BoolVal   *bool
	IntVal    *int64
	FloatVal  *float64
	StringVal *string
	BytesVal  []byte
	// JSONVal holds structured values (objects / arrays) as raw JSON bytes.
	JSONVal []byte
}

// PropertyUpdate is emitted by the Subscribe stream.
type PropertyUpdate struct {
	DeviceID     string
	PropertyPath string
	Value        PropertyValue
	Timestamp    time.Time
}
