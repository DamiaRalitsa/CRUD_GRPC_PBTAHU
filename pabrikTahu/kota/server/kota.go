package server

import (
	"context"
)

//step 5
type kota struct {
	writer ReadWriter
}

func NewKota(writer ReadWriter) KotaService {
	return &kota{writer: writer}
}

//Methode pada interface CustomerService di service.go
func (c *kota) AddKotaService(ctx context.Context, kota Kota) error {
	//fmt.Println("customer")
	err := c.writer.AddKota(kota)
	if err != nil {
		return err
	}

	return nil
}

func (c *kota) ReadKotaByNamaService(ctx context.Context, mob string) (Kota, error) {
	kota, err := c.writer.ReadKotaByNama(mob)
	//fmt.Println(cus)
	if err != nil {
		return kota, err
	}
	return kota, nil
}

func (c *kota) ReadKotaService(ctx context.Context) (Kotaa, error) {
	kot, err := c.writer.ReadKota()
	//fmt.Println("customer", cus)
	if err != nil {
		return kot, err
	}
	return kot, nil
}

func (c *kota) UpdateKotaService(ctx context.Context, kot Kota) error {
	err := c.writer.UpdateKota(kot)
	if err != nil {
		return err
	}
	return nil
}
