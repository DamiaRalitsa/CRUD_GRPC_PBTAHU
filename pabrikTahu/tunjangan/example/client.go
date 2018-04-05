package main

import (
	"context"
	"fmt"
	"time"

	cli "git.bluebird.id/pabrikTahu/tunjangan/endpoint"
	opt "git.bluebird.id/pabrikTahu/util/grpc"
	util "git.bluebird.id/pabrikTahu/util/microservice"
	tr "github.com/opentracing/opentracing-go"
)

func main() {
	logger := util.Logger()
	tracer := tr.GlobalTracer()
	option := opt.ClientOption{Retry: 3, RetryTimeout: 500 * time.Millisecond, Timeout: 30 * time.Second}

	client, err := cli.NewGRPCTunjanganClient([]string{"127.0.0.1:2181"}, nil, option, tracer, logger)
	if err != nil {
		logger.Log("error", err)
	}

	//Add Customer
	//client.AddTunjanganService(context.Background(), svc.Tunjangan{DeskripsiTunjangan: "Kecelakaan", BesaranTunjangan: 300000, CreatedBy: "Damia"})

	//Get Customer By nama kota
	tunjNama, _ := client.ReadTunjanganByNamaService(context.Background(), "BPJS")
	fmt.Println("kota based on city name:", tunjNama)

	//List Customer
	//tunj, _ := client.ReadTunjanganService(context.Background())
	//fmt.Println("all tunjangan:", tunj)

	//Update Customer
	//client.UpdateTunjanganService(context.Background(), svc.Tunjangan{NamaTunjangan: "Jakarta", Status: 1, UpdatedBy: "Dam", IDTunjangan: 6})
}
