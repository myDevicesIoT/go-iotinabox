package iotinabox

import (
	"context"
	"fmt"
)

type DeviceService service

type Device struct {
	Id         int64  `json:"id,omitempty"`
	Name       string `json:thing_name`
	CompanyId  int64  `json:company_id`
	LocationId int64  `json:location_id`
	HardwareId string `json:hardware_id`
	SensorUse  string `json:sensor_use`
	SensorType string `json:sensor_type`
	TypeId     string `json:device_type_id`
}

func (d *DeviceService) GetById(ctx context.Context, id string) (*Device, error) {
	path := fmt.Sprintf("things/%v", id)
	req, err := d.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	var device *Device
	_, err = d.client.Do(ctx, req, &device)

	return device, err
}
