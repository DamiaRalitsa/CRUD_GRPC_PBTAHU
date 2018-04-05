package main

import (
	"context"
	"fmt"
	"time"

	cli "git.bluebird.id/pabrikTahu/kota/endpoint"
	opt "git.bluebird.id/pabrikTahu/util/grpc"
	util "git.bluebird.id/pabrikTahu/util/microservice"
	tr "github.com/opentracing/opentracing-go"
)

func main() {
	logger := util.Logger()
	tracer := tr.GlobalTracer()
	option := opt.ClientOption{Retry: 3, RetryTimeout: 500 * time.Millisecond, Timeout: 30 * time.Second}

	client, err := cli.NewGRPCKotaClient([]string{"127.0.0.1:2181"}, nil, option, tracer, logger)
	if err != nil {
		logger.Log("error", err)
	}

	//Add Customer
	//client.AddKotaService(context.Background(), svc.Kota{NamaKota: "Depok", CreatedBy: "Damia"})

	//Get Customer By nama kota
	//kotNama, _ := client.ReadKotaByNamaService(context.Background(), "Depok")
	//fmt.Println("kota based on city name:", kotNama)

	//List Customer
	kota, _ := client.ReadKotaService(context.Background())
	fmt.Println("all kota:", kota)

	//Update Customer
	//client.UpdateKotaService(context.Background(), svc.Kota{NamaKota: "Jakarta", Status: 1, UpdatedBy: "Dam", IDKota: 6})

	//Get Customer By Emails
	// cusEmail, _ := client.ReadCustomerByEmailService(context.Background(), "joko@gmail.com")
	// fmt.Println("customer based on email:", cusEmail)
}
