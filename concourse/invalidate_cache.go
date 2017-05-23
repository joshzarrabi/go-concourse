package concourse

import (
	"github.com/concourse/atc"
	"github.com/concourse/go-concourse/concourse/internal"
	"github.com/tedsuo/rata"
)

func (team *team) InvalidateCache(pipelineName string, resourceName string) (bool, error) {
	params := rata.Params{
		"pipeline_name": pipelineName,
		"resource_name": resourceName,
		"team_name":     team.name,
	}

	err := team.connection.Send(internal.Request{
		RequestName: atc.InvalidateCache,
		Params:      params,
	}, &internal.Response{})
	switch err.(type) {
	case nil:
		return true, nil
	case internal.ResourceNotFoundError:
		return false, nil
	default:
		return false, err
	}
}
