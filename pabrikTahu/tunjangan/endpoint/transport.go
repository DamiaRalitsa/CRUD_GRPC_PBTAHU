package endpoint

import (
	"context"

	scv "git.bluebird.id/pabrikTahu/tunjangan/server"

	pb "git.bluebird.id/pabrikTahu/tunjangan/grpc"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/tracing/opentracing"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	google_protobuf "github.com/golang/protobuf/ptypes/empty"
	stdopentracing "github.com/opentracing/opentracing-go"
	oldcontext "golang.org/x/net/context"
)

type grpcTunjanganServer struct {
	addTunjangan        grpctransport.Handler
	readTunjanganByNama grpctransport.Handler
	readTunjangan       grpctransport.Handler
	updateTunjangan     grpctransport.Handler
}

func NewGRPCTunjanganServer(endpoints TunjanganEndpoint, tracer stdopentracing.Tracer,
	logger log.Logger) pb.TunjanganServiceServer {
	options := []grpctransport.ServerOption{
		grpctransport.ServerErrorLogger(logger),
	}
	return &grpcTunjanganServer{
		addTunjangan: grpctransport.NewServer(endpoints.AddTunjanganEndpoint,
			decodeAddTunjanganRequest,
			encodeEmptyResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "AddTunjangan", logger)))...),
		readTunjanganByNama: grpctransport.NewServer(endpoints.ReadTunjanganByNamaEndpoint,
			decodeReadTunjanganByNamaRequest,
			encodeReadTunjanganByNamaResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "ReadTunjanganByNama", logger)))...),
		readTunjangan: grpctransport.NewServer(endpoints.ReadTunjanganEndpoint,
			decodeReadTunjanganRequest,
			encodeReadTunjanganResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "ReadTunjangan", logger)))...),
		updateTunjangan: grpctransport.NewServer(endpoints.UpdateTunjanganEndpoint,
			decodeUpdateTunjanganRequest,
			encodeEmptyResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "UpdateTunjangan", logger)))...),
	}
}

func decodeAddTunjanganRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.AddTunjanganReq)
	return scv.Tunjangan{IDTunjangan: req.GetIDTunjangan(), DeskripsiTunjangan: req.GetDeskripsiTunjangan(), BesaranTunjangan: req.GetBesaranTunjangan(), Status: req.GetStatus(), CreatedBy: req.GetCreatedBy()}, nil
}

func encodeEmptyResponse(_ context.Context, response interface{}) (interface{}, error) {
	return &google_protobuf.Empty{}, nil
}

func (s *grpcTunjanganServer) AddTunjangan(ctx oldcontext.Context, tunjangan *pb.AddTunjanganReq) (*google_protobuf.Empty, error) {
	_, resp, err := s.addTunjangan.ServeGRPC(ctx, tunjangan)
	if err != nil {
		return nil, err
	}
	return resp.(*google_protobuf.Empty), nil
}

func decodeReadTunjanganByNamaRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.ReadTunjanganByNamaReq)
	return scv.Tunjangan{DeskripsiTunjangan: req.DeskripsiTunjangan}, nil
}

func decodeReadTunjanganRequest(_ context.Context, request interface{}) (interface{}, error) {
	return nil, nil
}

func encodeReadTunjanganByNamaResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(scv.Tunjangan)
	return &pb.ReadTunjanganByNamaResp{IDTunjangan: resp.IDTunjangan, DeskripsiTunjangan: resp.DeskripsiTunjangan, Status: resp.Status, CreatedBy: resp.CreatedBy}, nil
}

func encodeReadTunjanganResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(scv.Tunjangans)

	rsp := &pb.ReadTunjanganResp{}

	for _, v := range resp {
		itm := &pb.ReadTunjanganByNamaResp{
			IDTunjangan:        v.IDTunjangan,
			DeskripsiTunjangan: v.DeskripsiTunjangan,
			BesaranTunjangan:   v.BesaranTunjangan,
			Status:             v.Status,
			CreatedBy:          v.CreatedBy,
		}
		rsp.AllTunjangan = append(rsp.AllTunjangan, itm)
	}
	return rsp, nil
}

func (s *grpcTunjanganServer) ReadTunjanganByNama(ctx oldcontext.Context, idtunjangan *pb.ReadTunjanganByNamaReq) (*pb.ReadTunjanganByNamaResp, error) {
	_, resp, err := s.readTunjanganByNama.ServeGRPC(ctx, idtunjangan)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ReadTunjanganByNamaResp), nil
}

func (s *grpcTunjanganServer) ReadTunjangan(ctx oldcontext.Context, e *google_protobuf.Empty) (*pb.ReadTunjanganResp, error) {
	_, resp, err := s.readTunjangan.ServeGRPC(ctx, e)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ReadTunjanganResp), nil
}

func decodeUpdateTunjanganRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.UpdateTunjanganReq)
	return scv.Tunjangan{IDTunjangan: req.IDTunjangan, DeskripsiTunjangan: req.DeskripsiTunjangan, BesaranTunjangan: req.BesaranTunjangan, Status: req.Status, UpdatedBy: req.UpdatedBy}, nil
}

func (s *grpcTunjanganServer) UpdateTunjangan(ctx oldcontext.Context, cus *pb.UpdateTunjanganReq) (*google_protobuf.Empty, error) {
	_, resp, err := s.updateTunjangan.ServeGRPC(ctx, cus)
	if err != nil {
		return &google_protobuf.Empty{}, err
	}
	return resp.(*google_protobuf.Empty), nil
}
