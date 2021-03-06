// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// IMSAlert i m s alert
// swagger:model IMSAlert
type IMSAlert struct {

	// the alertInfo
	AlertInfo string `json:"alert_info,omitempty"`

	// the alertIp
	AlertIP string `json:"alert_ip,omitempty"`

	// the alertLevel
	AlertLevel int64 `json:"alert_level,omitempty"`

	// the alertObj
	AlertObj string `json:"alert_obj,omitempty"`

	// the alertReceiver
	AlertReceiver string `json:"alert_receiver,omitempty"`

	// the alertTitle
	AlertTitle string `json:"alert_title,omitempty"`

	// the alertWay
	AlertWay string `json:"alert_way,omitempty"`

	// the remarkInfo
	RemarkInfo string `json:"remark_info,omitempty"`

	// the subSystemId
	SubSystemID int64 `json:"sub_system_id,omitempty"`
}

// Validate validates this i m s alert
func (m *IMSAlert) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IMSAlert) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IMSAlert) UnmarshalBinary(b []byte) error {
	var res IMSAlert
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
