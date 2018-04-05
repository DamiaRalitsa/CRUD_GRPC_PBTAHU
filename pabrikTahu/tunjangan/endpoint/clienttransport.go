package endpoint

import (
	"context"
	"time"

	svc "git.bluebird.id/pabrikTahu/tunjangan/server"

	pb "git.bluebird.id/pabrikTahu/tunjangan/grpc"

	disc "git.bluebird.id/pabrikTahu/util/microservice"

	util "git.bluebird.id/pabrikTahu/util/grpc"

	"github.com/go-kit/kit/circuitbreaker"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/lb"
	"github.com/go-kit/kit/tracing/opentracing"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	google_protobuf "github.com/golang/protobuf/ptypes/empty"
	stdopentracing "github.com/opentracing/opentracing-go"
	"github.com/sony/gobreaker"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	grpcName = "grpc.TunjanganService"
)

func NewGRPCTunjanganClient(nodes []string, creds credentials.TransportCredentials, option util.ClientOption,
	tracer stdopentracing.Tracer, logger log.Logger) (svc.TunjanganService, error) {

	instancer, err := disc.ServiceDiscovery(nodes, svc.ServiceID, logger)
	if err != nil {
		return nil, err
	}

	retryMax := option.Retry
	retryTimeout := option.RetryTimeout
	timeout := option.Timeout

	var addTunjanganEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientAddTunjanganEndpoint, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		addTunjanganEp = retry
	}

	var readTunjanganByNamaEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientReadTunjanganByNamaEndpoint, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		readTunjanganByNamaEp = retry
	}

	var readTunjanganEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientReadTunjanganEndpoint, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		readTunjanganEp = retry
	}

	var updateTunjanganEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientUpdateTunjangan, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		updateTunjanganEp = retry
	}
	return TunjanganEndpoint{AddTunjanganEndpoint: addTunjanganEp, ReadTunjanganByNamaEndpoint: readTunjanganByNamaEp,
		ReadTunjanganEndpoint: readTunjanganEp, UpdateTunjanganEndpoint: updateTunjanganEp}, nil
}

func encodeAddTunjanganRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.Tunjangan)
	return &pb.AddTunjanganReq{
		IDTunjangan:        req.IDTunjangan,
		DeskripsiTunjangan: req.DeskripsiTunjangan,
		BesaranTunjangan:   req.BesaranTunjangan,
		Status:             req.Status,
		CreatedBy:          req.CreatedBy,
	}, nil
}

func encodeReadTunjanganByNamaRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.Tunjangan)
	return &pb.ReadTunjanganByNamaReq{DeskripsiTunjangan: req.DeskripsiTunjangan}, nil
}

func encodeReadTunjanganRequest(_ context.Context, request interface{}) (interface{}, error) {
	return &google_protobuf.Empty{}, nil
}

func encodeUpdateTunjanganRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.Tunjangan)
	return &pb.UpdateTunjanganReq{
		IDTunjangan:        req.IDTunjangan,
		DeskripsiTunjangan: req.DeskripsiTunjangan,
		BesaranTunjangan:   req.BesaranTunjangan,
		Status:             req.Status,
		CreatedBy:          req.CreatedBy,
	}, nil
}

func decodeTunjanganResponse(_ context.Context, response interface{}) (interface{}, error) {
	return nil, nil
}

func decodeReadTunjanganByNamaRespones(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.ReadTunjanganByNamaResp)
	return svc.Tunjangan{
		IDTunjangan:        resp.IDTunjangan,
		DeskripsiTunjangan: resp.DeskripsiTunjangan,
		BesaranTunjangan:   resp.BesaranTunjangan,
		Status:             resp.Status,
		CreatedBy:          resp.CreatedBy,
	}, nil
}

func decodeReadTunjanganResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.ReadTunjanganResp)
	var rsp svc.Tunjangans

	for _, v := range resp.AllTunjangan {
		itm := svc.Tunjangan{
			IDTunjangan:        v.IDTunjangan,
			DeskripsiTunjangan: v.DeskripsiTunjangan,
			BesaranTunjangan:   v.BesaranTunjangan,
			Status:             v.Status,
			CreatedBy:          v.CreatedBy,
		}
		rsp = append(rsp, itm)
	}
	return rsp, nil
}

func makeClientAddTunjanganEndpoint(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {

	endpoint := grpctransport.NewClient(
		conn,
		grpcName,
		"AddTunjangan",
		encodeAddTunjanganRequest,
		decodeTunjanganResponse,
		google_protobuf.Empty{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "AddTunjangan")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "AddTunjangan",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientReadTunjanganByNamaEndpoint(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {

	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"ReadTunjanganByNama",
		encodeReadTunjanganByNamaRequest,
		decodeReadTunjanganByNamaRespones,
		pb.ReadTunjanganByNamaResp{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "ReadTunjanganByNama")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "ReadTunjanganByNama",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientReadTunjanganEndpoint(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {

	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"ReadTunjangan",
		encodeReadTunjanganRequest,
		decodeReadTunjanganResponse,
		pb.ReadTunjanganResp{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "ReadTunjangan")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "ReadTunjangan",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientUpdateTunjangan(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {
	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"UpdateTunjangan",
		encodeUpdateTunjanganRequest,
		decodeTunjanganResponse,
		google_protobuf.Empty{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "UpdateTunjangan")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "UpdateTunjangan",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}
