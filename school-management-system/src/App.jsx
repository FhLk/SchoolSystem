import { useState } from 'react'
import reactLogo from './assets/react.svg'
import viteLogo from '/vite.svg'
import './index.css'
import TeacherPage from './Page/TeacherPage'
import ClassPage from './Page/ClassPage'
import Navbar from './components/Navbar/NavBar'


function App() {
  return (
    <div>
      <Navbar/>
      <TeacherPage/> 
    </div>
  )
}

export default App
