import React from 'react'
import ReactDOM from 'react-dom/client'
import App from './App.jsx'
import './index.css'
import {
  createBrowserRouter,
  RouterProvider,
} from "react-router-dom";
import TeacherPage from './Page/TeacherPage.jsx';
import ClassPage from './Page/ClassPage.jsx';
import ErrorPage from './Page/ErrorPage.jsx';

const router = createBrowserRouter([
  {
    path :"/",
    element: <TeacherPage/>
  },
  {
    path :"/class",
    element: <ClassPage/>
  },
  {
    path: "/*",
    element: <ErrorPage/>
  }
]);

ReactDOM.createRoot(document.getElementById('root')).render(
  <React.StrictMode>
    <RouterProvider router={router} />
  </React.StrictMode>,
)
