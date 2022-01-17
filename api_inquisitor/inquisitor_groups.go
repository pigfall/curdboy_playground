package api_inquisitor

import (
	"context"
	api "github.com/pigfall/curdboy_playground/http_api_contacts"
)

type ApiInquisitor interface {
	Check(ctx context.Context) error
}

func GetAllApiInquisitors(client api.ApiClientIfce) []ApiInquisitor {
	return []ApiInquisitor{

		CarServiceInquisitorNew(client),

		DeptServiceInquisitorNew(client),

		UserServiceInquisitorNew(client),
	}
}
