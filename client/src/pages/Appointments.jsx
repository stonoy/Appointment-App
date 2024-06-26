import React from 'react'
import { customFetch } from '../utils'
import { redirect, useLoaderData } from 'react-router-dom'
import { Appointment, Pagination } from '../components'
import { toast } from 'react-toastify'

export const loader = (store) => async() => {
  const token = store.getState().user.token

  try {
    const resp = await customFetch('/getappointments', {
      headers : {
        "Authorization": `Bearer ${token}`
      }
    })
    // console.log(resp?.data)
    return resp?.data
  } catch (error) {
    console.log(error?.response?.status)
    if (error?.response?.status === 401 || error?.response?.status === 403){
        toast.error("Login to continue")
        return redirect("/")
    }

    if (error?.response?.status === 400){
      toast.error("User not Created patient profile")
      return redirect("/createpatient")
  }

    toast.error(error?.response?.data?.msg)
    return null
  }
}

const Appointments = () => {
  const {appointments, numOfPages, page} = useLoaderData()

  if (appointments.length === 0){
    return (
      <div className="bg-gray-100">
        <h1 className='text-xl'>No Appointments to display</h1>
      </div>
    )
  }
  
  return (
    <div className="bg-gray-100">
    <div className="max-w-7xl mx-auto py-6 sm:px-6 lg:px-8">
      {/* Filter Section */}
      {/* <FilterSection isFilterApplied={isFilterApplied} myFilter={myFilter} handleOpenFilter={handleOpenFilter} handleSearch={handleSearch}/> */}

      <Appointment appointments={appointments}/>

      {/* Pagination */}
      <Pagination numOfPages={numOfPages} page={page} />
      
    </div>
  </div>
  )
}

export default Appointments