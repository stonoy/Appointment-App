import React from 'react'
import {RouterProvider, createBrowserRouter} from 'react-router-dom'
import {Appointments, ErrorPage, HomeLayOut, Landing, Login, Register, SingleAvailability} from './pages'

const router = createBrowserRouter([
  {
    path: '/',
    element: <HomeLayOut/>,
    errorElement: <ErrorPage/>,
    children : [
      {
        index : true,
        element: <Landing/>
      },
      {
        path: 'availability',
        element: <SingleAvailability/>,
      },
      {
        path: 'appointments',
        element: <Appointments/>,
      }
    ]
  },
  {
    path: '/login',
    element: <Login/>,
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