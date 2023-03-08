import React from "react"
import { useNavigate } from "react-router-dom"
import { useEffect } from "react"

function Main() {
  const navigate = useNavigate()

  useEffect(() => {
    window.Telegram.WebApp.MainButton.hide()
    window.Telegram.WebApp.BackButton.hide()
  })

  return (
    <section id="container">

      <div className="flex justify-center flex-col items-center relative mb-2 flex-1">
        <div>svg</div>
        <div className="theme-text-color mb-3">Твой личный кабинет</div>
        <div className="theme-hint-color">Редактируй и просматривай профиль</div>
      </div>

      <div
        onClick={() => { navigate("/profile") }}
        className="w-full flex-initial justify-center cursor-pointer p-4 pt-3"
      >
        <section className="w-full flex justify-center theme-button-color p-4 rounded-lg">
          <button>
            <span className="theme-text-color">Посмотреть профиль</span>
          </button>
        </section>
      </div>

    </section>
  )
}

export default Main