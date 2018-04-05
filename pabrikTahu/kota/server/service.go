package server

import "context"

type Status int32

const (
	//ServiceID is dispatch service ID
	ServiceID        = "kota.bluebird.id"
	OnAdd     Status = 1
)

type Kota struct {
	IDKota    int64
	NamaKota  string
	Status    int32
	CreatedBy string
	UpdatedBy string
}
type Kotaa []Kota

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
	AddKota(Kota) error
	ReadKotaByNama(string) (Kota, error)
	ReadKota() (Kotaa, error)
	UpdateKota(Kota) error
}

type KotaService interface {
	AddKotaService(context.Context, Kota) error
	ReadKotaByNamaService(context.Context, string) (Kota, error)
	ReadKotaService(context.Context) (Kotaa, error)
	UpdateKotaService(context.Context, Kota) error
}
