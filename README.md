# CRUD_GRPC_PBTAHU
This repo contains source code of CRUD to 3 Master Table

Container :

1. Kota (Master Table)
2. Tunjangan (Master Table)
3. User (Master table)

Instruction :

1. create database 'dbTahu'
2. create table Kota, Tunjangan, and User along with the complete fields
3. insert one row of data manually (recommended) for each master table for a head start to check the select query later on.
4. edit the service.conf (dbname, dbuid, dbpwd, port, etc.) according to your own computer
5. edit the grpc first, edit the struct in .proto with the same table name, and same variable struct that you inserted before
6. got to your terminal, go to grpc repository, and type in protoc *.proto --go_out=plugins=grpc:. to create the protobuff file.
7.
