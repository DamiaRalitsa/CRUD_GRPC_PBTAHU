package server

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	addTunjangan           = `insert into Tunjangan(deskripsitunjangan,besarantunjangan,status,createdby,createdon)values (?,?,?,?,?)`
	selectTunjanganByNama  = `select idtunjangan, deskripsitunjangan, besarantunjangan, status, createdby from Tunjangan where deskripsitunjangan = ?`
	selectTunjangan        = `select idtunjangan, deskripsitunjangan, status from Tunjangan`
	updateTunjangan        = `update Tunjangan set deskripsitunjangan=?,status=?,updatedby=?,updatedon=? where idtunjangan=?`
	selectTunjanganByEmail = `select deskripsitunjangan, status from Tunjangan where email=?`
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

func (rw *dbReadWriter) AddTunjangan(tunjangan Tunjangan) error {
	fmt.Println("add")
	tx, err := rw.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(addTunjangan, tunjangan.DeskripsiTunjangan, OnAdd, tunjangan.CreatedBy, time.Now())
	//fmt.Println(err)
	if err != nil {
		return err

	}
	return tx.Commit()
}

func (rw *dbReadWriter) ReadTunjanganByNama(deskripsitunjangan string) (Tunjangan, error) {
	tunjangan := Tunjangan{DeskripsiTunjangan: deskripsitunjangan}
	err := rw.db.QueryRow(selectTunjanganByNama, deskripsitunjangan).Scan(&tunjangan.IDTunjangan, &tunjangan.DeskripsiTunjangan, &tunjangan.BesaranTunjangan, &tunjangan.Status, &tunjangan.CreatedBy)

	if err != nil {
		return Tunjangan{}, err
	}

	return tunjangan, nil
}

func (rw *dbReadWriter) ReadTunjangan() (Tunjangans, error) {
	tunjangan := Tunjangans{}
	rows, _ := rw.db.Query(selectTunjangan)
	defer rows.Close()
	for rows.Next() {
		var c Tunjangan
		err := rows.Scan(&c.IDTunjangan, &c.DeskripsiTunjangan, &c.Status)
		if err != nil {
			fmt.Println("error query:", err)
			return tunjangan, err
		}
		tunjangan = append(tunjangan, c)
	}
	//fmt.Println("db nya:", customer)
	return tunjangan, nil
}

func (rw *dbReadWriter) UpdateTunjangan(kot Tunjangan) error {
	//fmt.Println("update")
	tx, err := rw.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(updateTunjangan, kot.DeskripsiTunjangan, kot.Status, kot.UpdatedBy, time.Now(), kot.IDTunjangan)

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
