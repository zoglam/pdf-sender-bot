package app

import (
	"bytes"
	"time"

	"github.com/zoglam/pdf-sender-bot/internal/dto"
	"github.com/zoglam/pdf-sender-bot/pkg/logg"
	"gopkg.in/telebot.v3"
)

func (t *App) GetPDF(c telebot.Context) error {

	userHasRegistrated, err := t.userService.DidUserFilledData(c.Sender().ID)

	logg.Info().Msgf("%+v %+v", userHasRegistrated, err)
	if err != nil {
		logg.Error().Err(err).Msg("Ошибка при генерации pdf")
		return err
	}
	if !userHasRegistrated {
		return c.Send("Заполните все поля в личном кабинете")
	}

	user, err := t.userService.GetUserProfile(c.Sender().ID)
	if err != nil {
		logg.Error().Err(err).Msgf("Ошибка при GetUserProfile")
		return err
	}

	data := &dto.PDFParams{
		QR:                        "string",
		Organization:              user.Organization,
		Address:                   user.Address,
		Phone:                     user.Phone,
		OGRN:                      user.OGRN,
		ModelT:                    user.VehicleModel,
		GovermentSign:             user.StateLicensePlate,
		FirstName:                 user.FirstName,
		SecondName:                user.SecondName,
		ThirdName:                 user.ThirdName,
		СertificateNumber:         user.IDNumber,
		LicenseRegistrationNumber: user.LicenseRegistrationNumber,
		LicenseSerial:             user.LicenseSerial,
		LicenseNumber:             user.LicenseNumber,
		GarageNumber:              user.GarageNumber,
		PersonnelNumber:           user.PersonnelNumber,
		ShortSignFIO:              "string",
	}
	pdf, err := t.PDFService.GeneratePDF(data)
	if err != nil {
		logg.Error().Msg("Ошибка при генерации pdf")
		return err
	}

	reader := bytes.NewReader(pdf)
	file := &telebot.Document{
		File:     telebot.FromReader(reader),
		FileName: time.Now().Format("2006-01-02") + "-путевой-лист.pdf",
	}
	return c.Send(file)
}
