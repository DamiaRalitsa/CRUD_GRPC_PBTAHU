package endpoint

import (
	"context"

	svc "git.bluebird.id/pabrikTahu/tunjangan/server"

	kit "github.com/go-kit/kit/endpoint"
)

type TunjanganEndpoint struct {
	AddTunjanganEndpoint        kit.Endpoint
	ReadTunjanganByNamaEndpoint kit.Endpoint
	ReadTunjanganEndpoint       kit.Endpoint
	UpdateTunjanganEndpoint     kit.Endpoint
}

func NewTunjanganEndpoint(service svc.TunjanganService) TunjanganEndpoint {
	addTunjanganEp := makeAddTunjanganEndpoint(service)
	readTunjanganByNamaEp := makeReadTunjanganByNamaEndpoint(service)
	readTunjanganEp := makeReadTunjanganEndpoint(service)
	updateTunjanganEp := makeUpdateTunjanganEndpoint(service)
	return TunjanganEndpoint{AddTunjanganEndpoint: addTunjanganEp,
		ReadTunjanganByNamaEndpoint: readTunjanganByNamaEp,
		ReadTunjanganEndpoint:       readTunjanganEp,
		UpdateTunjanganEndpoint:     updateTunjanganEp,
	}
}

func makeAddTunjanganEndpoint(service svc.TunjanganService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.Tunjangan)
		err := service.AddTunjanganService(ctx, req)
		return nil, err
	}
}

func makeReadTunjanganByNamaEndpoint(service svc.TunjanganService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.Tunjangan)
		result, err := service.ReadTunjanganByNamaService(ctx, req.DeskripsiTunjangan)
		return result, err
	}
}

func makeReadTunjanganEndpoint(service svc.TunjanganService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		result, err := service.ReadTunjanganService(ctx)
		return result, err
	}
}

func makeUpdateTunjanganEndpoint(service svc.TunjanganService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.Tunjangan)
		err := service.UpdateTunjanganService(ctx, req)
		return nil, err
	}
}
