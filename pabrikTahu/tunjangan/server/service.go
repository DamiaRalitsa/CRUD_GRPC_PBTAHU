package server

import "context"

type Status int32

const (
	//ServiceID is dispatch service ID
	ServiceID        = "tunjangan.bluebird.id"
	OnAdd     Status = 1
)

type Tunjangan struct {
	IDTunjangan        int64
	DeskripsiTunjangan string
	BesaranTunjangan   int64
	Status             int32
	CreatedBy          string
	UpdatedBy          string
}
type Tunjangans []Tunjangan

/*type Location struct {
	customerID   int64
	label        []int32
	locationType []int32
	name         []string
	street       string
	village      string
	district     string
	city         string
	province     string
	latitude     float64
	longitude    float64
}*/

type ReadWriter interface {
	AddTunjangan(Tunjangan) error
	ReadTunjanganByNama(string) (Tunjangan, error)
	ReadTunjangan() (Tunjangans, error)
	UpdateTunjangan(Tunjangan) error
}

type TunjanganService interface {
	AddTunjanganService(context.Context, Tunjangan) error
	ReadTunjanganByNamaService(context.Context, string) (Tunjangan, error)
	ReadTunjanganService(context.Context) (Tunjangans, error)
	UpdateTunjanganService(context.Context, Tunjangan) error
}
