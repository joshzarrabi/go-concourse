package concourse

import (
	"fmt"

	"github.com/concourse/atc"
	"github.com/concourse/go-concourse/concourse/internal"
	"github.com/tedsuo/rata"
)

type InvalidateCacheError struct {
	atc.InvalidateResponseBody
}

func (invalidateCacheError InvalidateCacheError) Error() string {
	return fmt.Sprintf("invalidate failed with exit status '%d':\n%s\n", invalidateCacheError.ExitStatus, invalidateCacheError.Stderr)
}

func (team *team) CheckResource(pipelineName string, resourceName string) (bool, error) {
	params := rata.Params{
		"pipeline_name": pipelineName,
		"resource_name": resourceName,
		"team_name":     team.name,
	}

	var resource atc.Resource
	err := team.connection.Send(internal.Request{
		RequestName: atc.InvalidateCache,
		Params:      params,
	}, &internal.Response{
		Result: &resource,
	})
	switch err.(type) {
	case nil:
		return resource, true, nil
	case internal.ResourceNotFoundError:
		return resource, false, nil
	default:
		return resource, false, err
	}
}
