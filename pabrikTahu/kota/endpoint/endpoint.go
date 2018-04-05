package endpoint

import (
	"context"

	svc "git.bluebird.id/pabrikTahu/kota/server"

	kit "github.com/go-kit/kit/endpoint"
)

type KotaEndpoint struct {
	AddKotaEndpoint        kit.Endpoint
	ReadKotaByNamaEndpoint kit.Endpoint
	ReadKotaEndpoint       kit.Endpoint
	UpdateKotaEndpoint     kit.Endpoint
}

func NewKotaEndpoint(service svc.KotaService) KotaEndpoint {
	addKotaEp := makeAddKotaEndpoint(service)
	readKotaByNamaEp := makeReadKotaByNamaEndpoint(service)
	readKotaEp := makeReadKotaEndpoint(service)
	updateKotaEp := makeUpdateKotaEndpoint(service)
	return KotaEndpoint{AddKotaEndpoint: addKotaEp,
		ReadKotaByNamaEndpoint: readKotaByNamaEp,
		ReadKotaEndpoint:       readKotaEp,
		UpdateKotaEndpoint:     updateKotaEp,
	}
}

func makeAddKotaEndpoint(service svc.KotaService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.Kota)
		err := service.AddKotaService(ctx, req)
		return nil, err
	}
}

func makeReadKotaByNamaEndpoint(service svc.KotaService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.Kota)
		result, err := service.ReadKotaByNamaService(ctx, req.NamaKota)
		return result, err
	}
}

func makeReadKotaEndpoint(service svc.KotaService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		result, err := service.ReadKotaService(ctx)
		return result, err
	}
}

func makeUpdateKotaEndpoint(service svc.KotaService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.Kota)
		err := service.UpdateKotaService(ctx, req)
		return nil, err
	}
}
