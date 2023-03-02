import React from "react"
import './App.css';
import { BrowserRouter, Route, Routes } from "react-router-dom"

import LocationProvider from "./provider/location"

import Profile from './pages/profile'
import Main from './pages/main'
import { LoginFormFields } from "./types/profile";

function App() {

  return (
    <BrowserRouter>
      <LocationProvider>
        <Routes>
          <Route index path="/" element={<Main />} />
          <Route path="/profile" element={<Profile />} />
        </Routes>
      </LocationProvider>
    </BrowserRouter>
  )

}

export default App