// generated by jsonenums -type=Notification; DO NOT EDIT

package types

import (
	"encoding/json"
	"fmt"
)

var (
	_NotificationNameToValue = map[string]Notification{
		"PreProviderSubmitNotification":   PreProviderSubmitNotification,
		"PostProviderSubmitNotification":  PostProviderSubmitNotification,
		"NotificationPreDeploymentUpdate": NotificationPreDeploymentUpdate,
		"NotificationDeploymentUpdate":    NotificationDeploymentUpdate,
	}

	_NotificationValueToName = map[Notification]string{
		PreProviderSubmitNotification:   "PreProviderSubmitNotification",
		PostProviderSubmitNotification:  "PostProviderSubmitNotification",
		NotificationPreDeploymentUpdate: "NotificationPreDeploymentUpdate",
		NotificationDeploymentUpdate:    "NotificationDeploymentUpdate",
	}
)

func init() {
	var v Notification
	if _, ok := interface{}(v).(fmt.Stringer); ok {
		_NotificationNameToValue = map[string]Notification{
			interface{}(PreProviderSubmitNotification).(fmt.Stringer).String():   PreProviderSubmitNotification,
			interface{}(PostProviderSubmitNotification).(fmt.Stringer).String():  PostProviderSubmitNotification,
			interface{}(NotificationPreDeploymentUpdate).(fmt.Stringer).String(): NotificationPreDeploymentUpdate,
			interface{}(NotificationDeploymentUpdate).(fmt.Stringer).String():    NotificationDeploymentUpdate,
		}
	}
}

// MarshalJSON is generated so Notification satisfies json.Marshaler.
func (r Notification) MarshalJSON() ([]byte, error) {
	if s, ok := interface{}(r).(fmt.Stringer); ok {
		return json.Marshal(s.String())
	}
	s, ok := _NotificationValueToName[r]
	if !ok {
		return nil, fmt.Errorf("invalid Notification: %d", r)
	}
	return json.Marshal(s)
}

// UnmarshalJSON is generated so Notification satisfies json.Unmarshaler.
func (r *Notification) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("Notification should be a string, got %s", data)
	}
	v, ok := _NotificationNameToValue[s]
	if !ok {
		return fmt.Errorf("invalid Notification %q", s)
	}
	*r = v
	return nil
}
