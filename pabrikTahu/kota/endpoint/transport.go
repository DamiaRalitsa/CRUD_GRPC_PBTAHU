package endpoint

import (
	"context"

	scv "git.bluebird.id/pabrikTahu/kota/server"

	pb "git.bluebird.id/pabrikTahu/kota/grpc"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/tracing/opentracing"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	google_protobuf "github.com/golang/protobuf/ptypes/empty"
	stdopentracing "github.com/opentracing/opentracing-go"
	oldcontext "golang.org/x/net/context"
)

type grpcKotaServer struct {
	addKota        grpctransport.Handler
	readKotaByNama grpctransport.Handler
	readKota       grpctransport.Handler
	updateKota     grpctransport.Handler
}

func NewGRPCKotaServer(endpoints KotaEndpoint, tracer stdopentracing.Tracer,
	logger log.Logger) pb.KotaServiceServer {
	options := []grpctransport.ServerOption{
		grpctransport.ServerErrorLogger(logger),
	}
	return &grpcKotaServer{
		addKota: grpctransport.NewServer(endpoints.AddKotaEndpoint,
			decodeAddKotaRequest,
			encodeEmptyResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "AddKota", logger)))...),
		readKotaByNama: grpctransport.NewServer(endpoints.ReadKotaByNamaEndpoint,
			decodeReadKotaByNamaRequest,
			encodeReadKotaByNamaResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "ReadKotaByNama", logger)))...),
		readKota: grpctransport.NewServer(endpoints.ReadKotaEndpoint,
			decodeReadKotaRequest,
			encodeReadKotaResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "ReadKota", logger)))...),
		updateKota: grpctransport.NewServer(endpoints.UpdateKotaEndpoint,
			decodeUpdateKotaRequest,
			encodeEmptyResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "UpdateKota", logger)))...),
	}
}

func decodeAddKotaRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.AddKotaReq)
	return scv.Kota{IDKota: req.GetIDKota(), NamaKota: req.GetNamaKota(), Status: req.GetStatus(), CreatedBy: req.GetCreatedBy()}, nil
}

func encodeEmptyResponse(_ context.Context, response interface{}) (interface{}, error) {
	return &google_protobuf.Empty{}, nil
}

func (s *grpcKotaServer) AddKota(ctx oldcontext.Context, kota *pb.AddKotaReq) (*google_protobuf.Empty, error) {
	_, resp, err := s.addKota.ServeGRPC(ctx, kota)
	if err != nil {
		return nil, err
	}
	return resp.(*google_protobuf.Empty), nil
}

func decodeReadKotaByNamaRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.ReadKotaByNamaReq)
	return scv.Kota{NamaKota: req.NamaKota}, nil
}

func decodeReadKotaRequest(_ context.Context, request interface{}) (interface{}, error) {
	return nil, nil
}

func encodeReadKotaByNamaResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(scv.Kota)
	return &pb.ReadKotaByNamaResp{IDKota: resp.IDKota, NamaKota: resp.NamaKota, Status: resp.Status, CreatedBy: resp.CreatedBy}, nil
}

func encodeReadKotaResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(scv.Kotaa)

	rsp := &pb.ReadKotaResp{}

	for _, v := range resp {
		itm := &pb.ReadKotaByNamaResp{
			IDKota:    v.IDKota,
			NamaKota:  v.NamaKota,
			Status:    v.Status,
			CreatedBy: v.CreatedBy,
		}
		rsp.AllKota = append(rsp.AllKota, itm)
	}
	return rsp, nil
}

func (s *grpcKotaServer) ReadKotaByNama(ctx oldcontext.Context, idkota *pb.ReadKotaByNamaReq) (*pb.ReadKotaByNamaResp, error) {
	_, resp, err := s.readKotaByNama.ServeGRPC(ctx, idkota)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ReadKotaByNamaResp), nil
}

func (s *grpcKotaServer) ReadKota(ctx oldcontext.Context, e *google_protobuf.Empty) (*pb.ReadKotaResp, error) {
	_, resp, err := s.readKota.ServeGRPC(ctx, e)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ReadKotaResp), nil
}

func decodeUpdateKotaRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.UpdateKotaReq)
	return scv.Kota{IDKota: req.IDKota, NamaKota: req.NamaKota, Status: req.Status, UpdatedBy: req.UpdatedBy}, nil
}

func (s *grpcKotaServer) UpdateKota(ctx oldcontext.Context, cus *pb.UpdateKotaReq) (*google_protobuf.Empty, error) {
	_, resp, err := s.updateKota.ServeGRPC(ctx, cus)
	if err != nil {
		return &google_protobuf.Empty{}, err
	}
	return resp.(*google_protobuf.Empty), nil
}
