package dto

type User struct {
	UserID                    int64  `db:"user_id"`                   // id в телеграм
	FirstName                 string `db:"firstName"`                 // Имя
	SecondName                string `db:"secondName"`                // Фамилия
	ThirdName                 string `db:"thirdName"`                 // Отчество
	Organization              string `db:"organization"`              // Организация
	Address                   string `db:"address"`                   // Адрес
	Phone                     string `db:"phone"`                     // Телефон
	OGRN                      string `db:"OGRN"`                      // ОРГН / ОРГНИП
	VehicleModel              string `db:"vehicleModel"`              // Модель ТС
	StateLicensePlate         string `db:"stateLicensePlate"`         // Государственный номерной знак
	IDNumber                  string `db:"IDNumber"`                  // Удостоверение №
	LicenseRegistrationNumber string `db:"licenseRegistrationNumber"` // Лицензия # Регистрационный №
	LicenseSerial             string `db:"licenseSerial"`             // Лицензия # Серия
	LicenseNumber             string `db:"licenseNumber"`             // Лицензия # Номер
	GarageNumber              string `db:"garageNumber"`              // Гаражный номер
	PersonnelNumber           string `db:"personnelNumber"`           // Табельный номер
}
