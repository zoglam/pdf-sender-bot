
import { useEffect } from "react"
import { useNavigate } from "react-router-dom"
import { LoginFormFields } from "./../types/profile"
import axios, {isCancel, AxiosError} from 'axios'

interface LoginPageProps {
  onSubmit: (data: LoginFormFields) => void
}

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

  const navigate = useNavigate()
  var formRef: HTMLFormElement | null

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
    const {
      Organization,
      Address,
      Phone,
      OGRN,
      ModelT,
      GovermentSign,
      FirstName,
      SecondName,
      ThirdName,
      СertificateNumber,
      LicenseRegistrationNumber,
      LicenseSerial,
      LicenseNumber,
      GarageNumber,
      PersonnelNumber
    } = formRef;

    axios.post("/api/profile")

    navigate('/')
  })

  useEffect(() => {
    window.Telegram.WebApp.MainButton.text = "Сохранить"

    window.Telegram.WebApp.BackButton.show()
    window.Telegram.WebApp.MainButton.show()
  })

  const SubmitForm: React.FormEventHandler<HTMLFormElement & FormFields> = (event) => {
    event.preventDefault();
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
          <input type="text" name="SecondName" id="SecondName" className="theme-button-border-color theme-secondary-bg-color border text-sm rounded-lg block w-full p-2.5" placeholder="Doe" required />
        </div>

        <div>
          <label htmlFor="FirstName" className="block mb-2 text-sm font-medium theme-text-color">Имя водителя</label>
          <input type="text" name="FirstName" id="FirstName" className="theme-button-border-color theme-secondary-bg-color border text-sm rounded-lg block w-full p-2.5" placeholder="Doe" required />
        </div>

        <div>
          <label htmlFor="ThirdName" className="block mb-2 text-sm font-medium theme-text-color">Отчество водителя</label>
          <input type="text" name="ThirdName" id="ThirdName" className="theme-button-border-color theme-secondary-bg-color border text-sm rounded-lg block w-full p-2.5" placeholder="Doe" required />
        </div>

        <div>
          <label htmlFor="Organization" className="block mb-2 text-sm font-medium theme-text-color">Организация</label>
          <input type="text" name="Organization" id="Organization" className="theme-button-border-color theme-secondary-bg-color border text-sm rounded-lg block w-full p-2.5" placeholder="John" required />
        </div>

        <div>
          <label htmlFor="Address" className="block mb-2 text-sm font-medium theme-text-color">Адрес</label>
          <input type="text" name="Organization" id="Address" className="theme-button-border-color theme-secondary-bg-color border text-sm rounded-lg block w-full p-2.5" placeholder="Doe" required />
        </div>

        <div>
          <label htmlFor="Phone" className="block mb-2 text-sm font-medium theme-text-color">Телефон</label>
          <input type="text" name="Phone" id="Phone" className="theme-button-border-color theme-secondary-bg-color border text-sm rounded-lg block w-full p-2.5" placeholder="7xxxxxxxxxx" required />
        </div>

        <div>
          <label htmlFor="OGRN" className="block mb-2 text-sm font-medium theme-text-color">ОГРН/ОГРНИП</label>
          <input type="text" name="OGRN" id="OGRN" className="theme-button-border-color theme-secondary-bg-color border text-sm rounded-lg block w-full p-2.5" placeholder="Doe" required />
        </div>

        <div>
          <label htmlFor="ModelT" className="block mb-2 text-sm font-medium theme-text-color">Марка и модель ТС</label>
          <input type="text" name="ModelT" id="ModelT" className="theme-button-border-color theme-secondary-bg-color border text-sm rounded-lg block w-full p-2.5" placeholder="Doe" required />
        </div>

        <div>
          <label htmlFor="GovermentSign" className="block mb-2 text-sm font-medium theme-text-color">Государственный номерной знак</label>
          <input type="text" name="GovermentSign" id="GovermentSign" className="theme-button-border-color theme-secondary-bg-color border text-sm rounded-lg block w-full p-2.5" placeholder="Doe" required />
        </div>

        <div>
          <label htmlFor="СertificateNumber" className="block mb-2 text-sm font-medium theme-text-color">Удостоверение №</label>
          <input type="text" name="СertificateNumber" id="СertificateNumber" className="theme-button-border-color theme-secondary-bg-color border text-sm rounded-lg block w-full p-2.5" placeholder="Doe" required />
        </div>

        <div>
          <label htmlFor="LicenseRegistrationNumber" className="block mb-2 text-sm font-medium theme-text-color">Лицензия / Регистрационный №</label>
          <input type="text" name="LicenseRegistrationNumber" id="LicenseRegistrationNumber" className="theme-button-border-color theme-secondary-bg-color border text-sm rounded-lg block w-full p-2.5" placeholder="Doe" required />
        </div>

        <div>
          <label htmlFor="LicenseSerial" className="block mb-2 text-sm font-medium theme-text-color">Лицензия / Серия №</label>
          <input type="text" name="LicenseSerial" id="LicenseSerial" className="theme-button-border-color theme-secondary-bg-color border text-sm rounded-lg block w-full p-2.5" placeholder="Doe" required />
        </div>

        <div>
          <label htmlFor="LicenseNumber" className="block mb-2 text-sm font-medium theme-text-color">Лицензия / Номер</label>
          <input type="text" name="LicenseNumber" id="LicenseNumber" className="theme-button-border-color theme-secondary-bg-color border text-sm rounded-lg block w-full p-2.5" placeholder="Doe" required />
        </div>

        <div>
          <label htmlFor="GarageNumber" className="block mb-2 text-sm font-medium theme-text-color">Гаражный номер</label>
          <input type="text" name="GarageNumber" id="GarageNumber" className="theme-button-border-color theme-secondary-bg-color border text-sm rounded-lg block w-full p-2.5" placeholder="1x" required />
        </div>

        <div>
          <label htmlFor="PersonnelNumber" className="block mb-2 text-sm font-medium theme-text-color">Табельный номер</label>
          <input type="text" name="PersonnelNumber" id="PersonnelNumber" className="theme-button-border-color theme-secondary-bg-color border text-sm rounded-lg block w-full p-2.5" placeholder="001" required />
        </div>
      </form>
    </section>
  )
}

export default Profile