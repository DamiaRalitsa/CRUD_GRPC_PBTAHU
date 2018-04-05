package server

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	addKota           = `insert into Kota(namakota,status,createdby,createdon)values (?,?,?,?)`
	selectKotaByNama  = `select idkota, namakota, status, createdby from Kota where namakota = ?`
	selectKota        = `select idkota, namakota, status from Kota`
	updateKota        = `update Kota set namakota=?,status=?,updatedby=?,updatedon=? where idkota=?`
	selectKotaByEmail = `select namakota, status from Kota where email=?`
)

//step 4
type dbReadWriter struct {
	db *sql.DB
}

func NewDBReadWriter(url string, schema string, user string, password string) ReadWriter {
	schemaURL := fmt.Sprintf("%s:%s@tcp(%s)/%s", user, password, url, schema)
	db, err := sql.Open("mysql", schemaURL)
	if err != nil {
		panic(err)
	}
	return &dbReadWriter{db: db}
}

func (rw *dbReadWriter) AddKota(kota Kota) error {
	fmt.Println("add")
	tx, err := rw.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(addKota, kota.NamaKota, OnAdd, kota.CreatedBy, time.Now())
	//fmt.Println(err)
	if err != nil {
		return err

	}
	return tx.Commit()
}

func (rw *dbReadWriter) ReadKotaByNama(namakota string) (Kota, error) {
	kota := Kota{NamaKota: namakota}
	err := rw.db.QueryRow(selectKotaByNama, namakota).Scan(&kota.IDKota, &kota.NamaKota, &kota.Status, &kota.CreatedBy)

	if err != nil {
		return Kota{}, err
	}

	return kota, nil
}

func (rw *dbReadWriter) ReadKota() (Kotaa, error) {
	kota := Kotaa{}
	rows, _ := rw.db.Query(selectKota)
	defer rows.Close()
	for rows.Next() {
		var c Kota
		err := rows.Scan(&c.IDKota, &c.NamaKota, &c.Status)
		if err != nil {
			fmt.Println("error query:", err)
			return kota, err
		}
		kota = append(kota, c)
	}
	//fmt.Println("db nya:", customer)
	return kota, nil
}

func (rw *dbReadWriter) UpdateKota(kot Kota) error {
	//fmt.Println("update")
	tx, err := rw.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(updateKota, kot.NamaKota, kot.Status, kot.UpdatedBy, time.Now(), kot.IDKota)

	//fmt.Println("name:", cus.Name, cus.CustomerId)
	if err != nil {
		return err
	}

	return tx.Commit()
}

/* func (rw *dbReadWriter) ReadKotaByEmail(email string) (Kota, error) {
	kota := Kota{Email: email}
	err := rw.db.QueryRow(selectKotaByEmail, email).Scan(&kota.IDKota, &kota.NamaKota, &kota.Status)

	//fmt.Println("err db", err)
	if err != nil {
		return Kota{}, err
	}

	return kota, nil
}
*/
