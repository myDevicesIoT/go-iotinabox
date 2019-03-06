package iotinabox

import (
	"context"
	"fmt"
	"log"
)

type LocationService service

type Location struct {
	Id        int64   `json:"id,omitempty"`
	Name      string  `json:name`
	Industry  string  `json:industry`
	CompanyId int64   `json:company_id`
	UserId    string  `json:user_id`
	Address   string  `json:address`
	City      string  `json:city`
	State     string  `json:state`
	Zip       string  `json:zip`
	Country   string  `json:country`
	Latitude  float64 `json:latitude`
	Longitude float64 `json:longitude`
	Timezone  string  `json:timezone`
}

func (d *LocationService) List(ctx context.Context) ([]*Location, error) {
	req, err := d.client.NewRequest("GET", "locations", nil)
	if err != nil {
		return nil, err
	}

	var locations []*Location
	_, err = d.client.Do(ctx, req, &locations)

	return locations, err
}

func (d *LocationService) Get(ctx context.Context, id int64) (*Location, error) {
	path := fmt.Sprintf("locations/%v", id)
	req, err := d.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	var location *Location
	_, err = d.client.Do(ctx, req, &location)

	return location, err
}

func (d *LocationService) Create(ctx context.Context, loc *Location) (*Location, error) {
	req, err := d.client.NewRequest("POST", "locations", loc)
	if err != nil {
		return nil, err
	}

	l := new(Location)
	resp, err := d.client.Do(ctx, req, l)
	if err != nil {
		log.Println(resp)
		return nil, err
	}

	log.Println(resp.StatusCode)
	return l, nil
}
