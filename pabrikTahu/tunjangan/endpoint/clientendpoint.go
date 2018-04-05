package endpoint

import (
	"context"
	"fmt"

	sv "git.bluebird.id/pabrikTahu/tunjangan/server"
)

func (ke TunjanganEndpoint) AddTunjanganService(ctx context.Context, tunjangan sv.Tunjangan) error {
	_, err := ke.AddTunjanganEndpoint(ctx, tunjangan)
	return err
}

func (ke TunjanganEndpoint) ReadTunjanganByNamaService(ctx context.Context, deskripsitunjangan string) (sv.Tunjangan, error) {
	req := sv.Tunjangan{DeskripsiTunjangan: deskripsitunjangan}
	fmt.Println(req)
	resp, err := ke.ReadTunjanganByNamaEndpoint(ctx, req)
	if err != nil {
		fmt.Println("error pada endpoint: ", err)
	}
	cus := resp.(sv.Tunjangan)
	return cus, err
}

func (ke TunjanganEndpoint) ReadTunjanganService(ctx context.Context) (sv.Tunjangans, error) {
	resp, err := ke.ReadTunjanganEndpoint(ctx, nil)
	fmt.Println("ke resp", resp)
	if err != nil {
		fmt.Println("error pada endpoint: ", err)
	}
	return resp.(sv.Tunjangans), err
}

func (ke TunjanganEndpoint) UpdateTunjanganService(ctx context.Context, kar sv.Tunjangan) error {
	_, err := ke.UpdateTunjanganEndpoint(ctx, kar)
	if err != nil {
		fmt.Println("error pada endpoint:", err)
	}
	return err
}
