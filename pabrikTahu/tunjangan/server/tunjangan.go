package server

import (
	"context"
)

//step 5
type tunjangan struct {
	writer ReadWriter
}

func NewTunjangan(writer ReadWriter) TunjanganService {
	return &tunjangan{writer: writer}
}

//Methode pada interface CustomerService di service.go
func (c *tunjangan) AddTunjanganService(ctx context.Context, tunjangan Tunjangan) error {
	//fmt.Println("customer")
	err := c.writer.AddTunjangan(tunjangan)
	if err != nil {
		return err
	}

	return nil
}

func (c *tunjangan) ReadTunjanganByNamaService(ctx context.Context, mob string) (Tunjangan, error) {
	tunjangan, err := c.writer.ReadTunjanganByNama(mob)
	//fmt.Println(cus)
	if err != nil {
		return tunjangan, err
	}
	return tunjangan, nil
}

func (c *tunjangan) ReadTunjanganService(ctx context.Context) (Tunjangans, error) {
	kot, err := c.writer.ReadTunjangan()
	//fmt.Println("customer", cus)
	if err != nil {
		return kot, err
	}
	return kot, nil
}

func (c *tunjangan) UpdateTunjanganService(ctx context.Context, kot Tunjangan) error {
	err := c.writer.UpdateTunjangan(kot)
	if err != nil {
		return err
	}
	return nil
}
