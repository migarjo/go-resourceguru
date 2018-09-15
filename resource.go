package resourceguru

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

type ResourcesService service

type Resource struct {
	ID               int         `json:"id"`
	Archived         bool        `json:"archived"`
	Bookable         bool        `json:"bookable"`
	Color            interface{} `json:"color"`
	Email            string      `json:"email"`
	Phone            string      `json:"phone"`
	Human            bool        `json:"human"`
	Image            string      `json:"image"`
	JobTitle         string      `json:"job_title"`
	Name             string      `json:"name"`
	Notes            string      `json:"notes"`
	CreatedAt        time.Time   `json:"created_at"`
	UpdatedAt        time.Time   `json:"updated_at"`
	LastUpdatedBy    int         `json:"last_updated_by"`
	BookedClientIds  []int       `json:"booked_client_ids"`
	BookedProjectIds []int       `json:"booked_project_ids"`
	Account          struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"account"`
	VacationAllowance float64 `json:"vacation_allowance"`
	AvailablePeriods  []struct {
		WeekDay    int         `json:"week_day"`
		StartTime  int         `json:"start_time"`
		EndTime    int         `json:"end_time"`
		ValidFrom  string      `json:"valid_from"`
		ValidUntil interface{} `json:"valid_until"`
	} `json:"available_periods"`
	CustomAvailablePeriods []interface{} `json:"custom_available_periods"`
	Overtimes              []interface{} `json:"overtimes"`
	ResourceType           struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"resource_type"`
	SelectedCustomFieldOptions []interface{} `json:"selected_custom_field_options"`
	CustomAttributes           struct {
		Phone string `json:"phone"`
	} `json:"custom_attributes"`
	Timezone struct {
		Name   string `json:"name"`
		Offset int    `json:"offset"`
	} `json:"timezone"`
}

func (r Resource) String() string {
	return Stringify(r)
}

// Get fetches a resource by id.
//
// ResourceGuru API docs: https://github.com/resourceguru/api-docs/blob/master/endpoints/resources.md#get-resource
func (s *ResourcesService) Get(ctx context.Context, owner string, resourceID int) (*Resource, *http.Response, error) {
	u := fmt.Sprintf("%v/resources/%v", owner, strconv.Itoa(resourceID))
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	resource := new(Resource)
	resp, err := s.client.Do(ctx, req, resource)
	if err != nil {
		return nil, resp, err
	}

	return resource, resp, nil
}
