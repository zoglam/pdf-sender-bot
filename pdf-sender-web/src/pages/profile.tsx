import React from "react"
import { useEffect, useState } from "react"
import { useNavigate } from "react-router-dom"
import { UserFields } from "./../types/profile"
import Spinner from "./spinner"
import axios, { isCancel, AxiosError } from 'axios'


type FormFields = {
  Organization: HTMLInputElement
  Address: HTMLInputElement
  Phone: HTMLInputElement
  OGRN: HTMLInputElement
  ModelT: HTMLInputElement
  GovermentSign: HTMLInputElement
  SecondName: HTMLInputElement
  FirstName: HTMLInputElement
  ThirdName: HTMLInputElement
  СertificateNumber: HTMLInputElement
  LicenseRegistrationNumber: HTMLInputElement
  LicenseSerial: HTMLInputElement
  LicenseNumber: HTMLInputElement
  GarageNumber: HTMLInputElement
  PersonnelNumber: HTMLInputElement
}

function Profile() {

  window.Telegram.WebApp.MainButton.text = "Сохранить"

  const navigate = useNavigate()
  var formRef: HTMLFormElement | null

  const [isLoading, setLoading] = useState(true);
  const [userData, setUserData] = useState({} as UserFields);

  const config = {
    headers: {
      "x-telegram-id": window.Telegram.WebApp.initDataUnsafe.user.id,
    }
  };

  window.Telegram.WebApp.onEvent('backButtonClicked', () => {
    navigate('/')
  })

  window.Telegram.WebApp.onEvent('mainButtonClicked', () => {
    if (formRef === null) {
      return
    }

    formRef.dispatchEvent(
      new Event("submit", { bubbles: true, cancelable: true })
    )

    const data = new FormData(formRef) as unknown as Iterable<
      [FormFields, FormDataEntryValue]
    >;
    const requestData: FormFields = Object.fromEntries(data);
    var json = JSON.stringify(requestData);


    axios.post(
      "/api/profile",
      json,
      config
    )

    navigate('/')
  })

  useEffect(() => {
    window.Telegram.WebApp.BackButton.show()
    window.Telegram.WebApp.MainButton.show()
  }, [])

  useEffect(() => {
    setLoading(true)

    axios.get(
      "/api/profile",
      config
    ).then(
      resp => {
        if (resp.status != 200) {
          console.log("ERROR resp.status", resp.status)
          return
        }
        var jsonData: UserFields = resp.data;
        setUserData({ ...userData, ...jsonData })
        setLoading(false)
      }
    )

  }, [])

  const SubmitForm: React.FormEventHandler<HTMLFormElement & FormFields> = (event) => {
    event.preventDefault();
  }

  if (isLoading) {
    return (
      <Spinner />
    )
  }

  return (
    <section className="flex justify-start flex-col p-8">

      <form
        className="space-y-4"
        onSubmit={SubmitForm}
        ref={ref => formRef = ref}
      >

        <div>
          <label htmlFor="SecondName" className="block mb-2 text-sm font-medium theme-text-color">Фамилия водителя</label>
          <input value={userData.SecondName} onChange={evt => setUserData({ ...userData, SecondName: evt.target.value })}
            type="text" name="SecondName" id="SecondName" className="theme-button-border-color theme-secondary-bg-color border text-sm rounded-lg block w-full p-2.5" placeholder="Doe" required />
        </div>

        <div>
          <label htmlFor="FirstName" className="block mb-2 text-sm font-medium theme-text-color">Имя водителя</label>
          <input value={userData.FirstName} onChange={evt => setUserData({ ...userData, FirstName: evt.target.value })}
            type="text" name="FirstName" id="FirstName" className="theme-button-border-color theme-secondary-bg-color border text-sm rounded-lg block w-full p-2.5" placeholder="Doe" required />
        </div>

        <div>
          <label htmlFor="ThirdName" className="block mb-2 text-sm font-medium theme-text-color">Отчество водителя</label>
          <input value={userData.ThirdName} onChange={evt => setUserData({ ...userData, ThirdName: evt.target.value })}
            type="text" name="ThirdName" id="ThirdName" className="theme-button-border-color theme-secondary-bg-color border text-sm rounded-lg block w-full p-2.5" placeholder="Doe" required />
        </div>

        <div>
          <label htmlFor="Organization" className="block mb-2 text-sm font-medium theme-text-color">Организация</label>
          <input value={userData.Organization} onChange={evt => setUserData({ ...userData, Organization: evt.target.value })}
            type="text" name="Organization" id="Organization" className="theme-button-border-color theme-secondary-bg-color border text-sm rounded-lg block w-full p-2.5" placeholder="John" required />
        </div>

        <div>
          <label htmlFor="Address" className="block mb-2 text-sm font-medium theme-text-color">Адрес</label>
          <input value={userData.Address} onChange={evt => setUserData({ ...userData, Address: evt.target.value })}
            type="text" name="Address" id="Address" className="theme-button-border-color theme-secondary-bg-color border text-sm rounded-lg block w-full p-2.5" placeholder="Doe" required />
        </div>

        <div>
          <label htmlFor="Phone" className="block mb-2 text-sm font-medium theme-text-color">Телефон</label>
          <input value={userData.Phone} onChange={evt => setUserData({ ...userData, Phone: evt.target.value })}
            type="text" name="Phone" id="Phone" className="theme-button-border-color theme-secondary-bg-color border text-sm rounded-lg block w-full p-2.5" placeholder="7xxxxxxxxxx" required />
        </div>

        <div>
          <label htmlFor="OGRN" className="block mb-2 text-sm font-medium theme-text-color">ОГРН/ОГРНИП</label>
          <input value={userData.OGRN} onChange={evt => setUserData({ ...userData, OGRN: evt.target.value })}
            type="text" name="OGRN" id="OGRN" className="theme-button-border-color theme-secondary-bg-color border text-sm rounded-lg block w-full p-2.5" placeholder="Doe" required />
        </div>

        <div>
          <label htmlFor="VehicleModel" className="block mb-2 text-sm font-medium theme-text-color">Марка и модель ТС</label>
          <input value={userData.VehicleModel} onChange={evt => setUserData({ ...userData, VehicleModel: evt.target.value })}
            type="text" name="VehicleModel" id="VehicleModel" className="theme-button-border-color theme-secondary-bg-color border text-sm rounded-lg block w-full p-2.5" placeholder="Doe" required />
        </div>

        <div>
          <label htmlFor="StateLicensePlate" className="block mb-2 text-sm font-medium theme-text-color">Государственный номерной знак</label>
          <input value={userData.StateLicensePlate} onChange={evt => setUserData({ ...userData, StateLicensePlate: evt.target.value })}
            type="text" name="StateLicensePlate" id="StateLicensePlate" className="theme-button-border-color theme-secondary-bg-color border text-sm rounded-lg block w-full p-2.5" placeholder="Doe" required />
        </div>

        <div>
          <label htmlFor="IDNumber" className="block mb-2 text-sm font-medium theme-text-color">Удостоверение №</label>
          <input value={userData.IDNumber} onChange={evt => setUserData({ ...userData, IDNumber: evt.target.value })}
            type="text" name="IDNumber" id="IDNumber" className="theme-button-border-color theme-secondary-bg-color border text-sm rounded-lg block w-full p-2.5" placeholder="Doe" required />
        </div>

        <div>
          <label htmlFor="LicenseRegistrationNumber" className="block mb-2 text-sm font-medium theme-text-color">Лицензия / Регистрационный №</label>
          <input value={userData.LicenseRegistrationNumber} onChange={evt => setUserData({ ...userData, LicenseRegistrationNumber: evt.target.value })}
            type="text" name="LicenseRegistrationNumber" id="LicenseRegistrationNumber" className="theme-button-border-color theme-secondary-bg-color border text-sm rounded-lg block w-full p-2.5" placeholder="Doe" required />
        </div>

        <div>
          <label htmlFor="LicenseSerial" className="block mb-2 text-sm font-medium theme-text-color">Лицензия / Серия №</label>
          <input value={userData.LicenseSerial} onChange={evt => setUserData({ ...userData, LicenseSerial: evt.target.value })}
            type="text" name="LicenseSerial" id="LicenseSerial" className="theme-button-border-color theme-secondary-bg-color border text-sm rounded-lg block w-full p-2.5" placeholder="Doe" required />
        </div>

        <div>
          <label htmlFor="LicenseNumber" className="block mb-2 text-sm font-medium theme-text-color">Лицензия / Номер</label>
          <input value={userData.LicenseNumber} onChange={evt => setUserData({ ...userData, LicenseNumber: evt.target.value })}
            type="text" name="LicenseNumber" id="LicenseNumber" className="theme-button-border-color theme-secondary-bg-color border text-sm rounded-lg block w-full p-2.5" placeholder="Doe" required />
        </div>

        <div>
          <label htmlFor="GarageNumber" className="block mb-2 text-sm font-medium theme-text-color">Гаражный номер</label>
          <input value={userData.GarageNumber} onChange={evt => setUserData({ ...userData, GarageNumber: evt.target.value })}
            type="text" name="GarageNumber" id="GarageNumber" className="theme-button-border-color theme-secondary-bg-color border text-sm rounded-lg block w-full p-2.5" placeholder="1x" required />
        </div>

        <div>
          <label htmlFor="PersonnelNumber" className="block mb-2 text-sm font-medium theme-text-color">Табельный номер</label>
          <input value={userData.PersonnelNumber} onChange={evt => setUserData({ ...userData, PersonnelNumber: evt.target.value })}
            type="text" name="PersonnelNumber" id="PersonnelNumber" className="theme-button-border-color theme-secondary-bg-color border text-sm rounded-lg block w-full p-2.5" placeholder="001" required />
        </div>
      </form>
    </section>
  )
}

export default Profile