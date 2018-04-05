package endpoint

import (
	"context"
	"fmt"

	sv "git.bluebird.id/pabrikTahu/kota/server"
)

func (ke KotaEndpoint) AddKotaService(ctx context.Context, kota sv.Kota) error {
	_, err := ke.AddKotaEndpoint(ctx, kota)
	return err
}

func (ke KotaEndpoint) ReadKotaByNamaService(ctx context.Context, namakota string) (sv.Kota, error) {
	req := sv.Kota{NamaKota: namakota}
	fmt.Println(req)
	resp, err := ke.ReadKotaByNamaEndpoint(ctx, req)
	if err != nil {
		fmt.Println("error pada endpoint: ", err)
	}
	cus := resp.(sv.Kota)
	return cus, err
}

func (ke KotaEndpoint) ReadKotaService(ctx context.Context) (sv.Kotaa, error) {
	resp, err := ke.ReadKotaEndpoint(ctx, nil)
	fmt.Println("ke resp", resp)
	if err != nil {
		fmt.Println("error pada endpoint: ", err)
	}
	return resp.(sv.Kotaa), err
}

func (ke KotaEndpoint) UpdateKotaService(ctx context.Context, kar sv.Kota) error {
	_, err := ke.UpdateKotaEndpoint(ctx, kar)
	if err != nil {
		fmt.Println("error pada endpoint:", err)
	}
	return err
}
