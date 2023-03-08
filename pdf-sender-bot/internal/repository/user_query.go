package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/zoglam/pdf-sender-bot/internal/dto"
)

type UserQuery interface {
	InsertUserData(data *dto.User) error
	UpdateUserData(data *dto.User) error
	GetUserData(id int64) (*dto.User, error)
}

type userQuery struct{}

const (
	InsertUserDataSQL = `
	insert into users values(
		@user_id,
		@firstName,
		@secondName,
		@thirdName,
		@organization,
		@address,
		@phone,
		@OGRN,
		@vehicleModel,
		@stateLicensePlate,
		@IDNumber,
		@licenseRegistrationNumber,
		@licenseSerial,
		@licenseNumber,
		@garageNumber,
		@personnelNumber
	)
	`

	UpdateUserDataSQL = `
	update users set
		firstName=@firstName,
		secondName=@secondName,
		thirdName=@thirdName,
		organization=@organization,
		address=@address,
		phone=@phone,
		OGRN=@OGRN,
		vehicleModel=@vehicleModel,
		stateLicensePlate=@stateLicensePlate,
		IDNumber=@IDNumber,
		licenseRegistrationNumber=@licenseRegistrationNumber,
		licenseSerial=@licenseSerial,
		licenseNumber=@licenseNumber,
		garageNumber=@garageNumber,
		personnelNumber=@personnelNumber
	where user_id=@user_id
	`

	GetUserDataSQL = `
	select
		user_id,
		firstName,
		secondName,
		thirdName,
		organization,
		address,
		phone,
		OGRN,
		vehicleModel,
		stateLicensePlate,
		IDNumber,
		licenseRegistrationNumber,
		licenseSerial,
		licenseNumber,
		garageNumber,
		personnelNumber
	from users
	where user_id=$1
	`
)

func (u *userQuery) InsertUserData(data *dto.User) error {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "query-name", "InsertUserData")
	args := pgx.NamedArgs{
		"user_id":                   data.UserID,
		"firstName":                 data.FirstName,
		"secondName":                data.SecondName,
		"thirdName":                 data.ThirdName,
		"organization":              data.Organization,
		"address":                   data.Address,
		"phone":                     data.Phone,
		"OGRN":                      data.OGRN,
		"vehicleModel":              data.VehicleModel,
		"stateLicensePlate":         data.StateLicensePlate,
		"IDNumber":                  data.IDNumber,
		"licenseRegistrationNumber": data.LicenseRegistrationNumber,
		"licenseSerial":             data.LicenseSerial,
		"licenseNumber":             data.LicenseNumber,
		"garageNumber":              data.GarageNumber,
		"personnelNumber":           data.PersonnelNumber,
	}

	_, err := DB.Exec(ctx, InsertUserDataSQL, args)
	if err != nil {
		return err
	}

	return nil
}
func (u *userQuery) UpdateUserData(data *dto.User) error {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "query-name", "UpdateUserData")
	args := pgx.NamedArgs{
		"user_id":                   data.UserID,
		"firstName":                 data.FirstName,
		"secondName":                data.SecondName,
		"thirdName":                 data.ThirdName,
		"organization":              data.Organization,
		"address":                   data.Address,
		"phone":                     data.Phone,
		"OGRN":                      data.OGRN,
		"vehicleModel":              data.VehicleModel,
		"stateLicensePlate":         data.StateLicensePlate,
		"IDNumber":                  data.IDNumber,
		"licenseRegistrationNumber": data.LicenseRegistrationNumber,
		"licenseSerial":             data.LicenseSerial,
		"licenseNumber":             data.LicenseNumber,
		"garageNumber":              data.GarageNumber,
		"personnelNumber":           data.PersonnelNumber,
	}
	_, err := DB.Exec(ctx, UpdateUserDataSQL, args)
	if err != nil {
		return err
	}
	return nil
}
func (u *userQuery) GetUserData(id int64) (*dto.User, error) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "query-name", "GetUserData")

	rows := DB.QueryRow(ctx, GetUserDataSQL, id)

	var user dto.User
	err := rows.Scan(
		&user.UserID,
		&user.FirstName,
		&user.SecondName,
		&user.ThirdName,
		&user.Organization,
		&user.Address,
		&user.Phone,
		&user.OGRN,
		&user.VehicleModel,
		&user.StateLicensePlate,
		&user.IDNumber,
		&user.LicenseRegistrationNumber,
		&user.LicenseSerial,
		&user.LicenseNumber,
		&user.GarageNumber,
		&user.PersonnelNumber,
	)
	if err != nil && err != pgx.ErrNoRows {
		return nil, err
	}

	return &user, nil
}
