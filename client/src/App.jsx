import React from 'react'
import {RouterProvider, createBrowserRouter} from 'react-router-dom'
import {About, Admin, AdminInsights, AdminLanding, Appointments, Availability, CreatePatient, ErrorPage, HomeLayOut, Landing, Login, Register} from './pages'
import { store } from './store'

// loaders
import {loader as landingLoader} from './pages/Landing'
import {loader as appointmentLoader} from './pages/Appointments'
import {loader as adminLandingLoader} from './pages/AdminLanding'

// actions
import {action as loginAction} from './pages/Login'
import {action as createAppointmentAction} from './pages/CreateAppointment'


const router = createBrowserRouter([
  {
    path: '/',
    element: <HomeLayOut/>,
    errorElement: <ErrorPage/>,
    children : [
      {
        index : true,
        element: <Landing/>,
        loader: landingLoader
      },
      {
        path: 'about',
        element: <About/>,
      },
      {
        path: 'availability',
        element: <Availability/>,
      },
      {
        path: 'createpatient',
        element: <CreatePatient/>,
      },
      {
        path: 'appointments',
        element: <Appointments/>,
        loader: appointmentLoader(store)
      },
      {
        path: 'createappointment/:id',
        action: createAppointmentAction(store),
      },
      {
        path: 'admin',
        element: <Admin/>,
        children: [
          {
            index : true,
            element: <AdminLanding/>,
            loader: adminLandingLoader(store),
          },
          {
            path: "insights",
            element: <AdminInsights/>
          },
        ]
      },
    ]
  },
  {
    path: '/login',
    element: <Login/>,
    action: loginAction(store)
  },
  {
    path: '/register',
    element: <Register/>,
  },
])

const App = () => {
  return (
    <RouterProvider router={router}/>
  )
}

export default App