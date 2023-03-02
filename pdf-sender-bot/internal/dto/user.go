package dto

type User struct {
	UserID                    int64  // id в телеграм
	FirstName                 string // Имя
	SecondName                string // Фамилия
	ThirdName                 string // Отчество
	Organization              string // Организация
	Address                   string // Адрес
	Phone                     string // Телефон
	OGRN                      string // ОРГН / ОРГНИП
	VehicleModel              string // Модель ТС
	StateLicensePlate         string // Государственный номерной знак
	IDNumber                  string // Удостоверение №
	LicenseRegistrationNumber string // Лицензия # Регистрационный №
	LicenseSerial             string // Лицензия # Серия
	LicenseNumber             string // Лицензия # Номер
	GarageNumber              string // Гаражный номер
	PersonnelNumber           string // Табельный номер
}
