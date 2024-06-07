import React from 'react'
import { Doctors, Pagination } from '../components'
import { customFetch } from '../utils'
import { redirect, useLoaderData } from 'react-router-dom'
import { toast } from 'react-toastify'

export const loader = (store) => async() => {
  const token = store.getState().user.token

  try {
    const resp = await customFetch("/getalldoctors", {
      headers : {
        "Authorization" : `Bearer ${token}`
      }
    })
    return resp?.data
  } catch (error) {
    const errMsg = error?.response?.data?.msg || "Error in getting doctors list"
    if (error?.response?.status === 401 || error?.response?.status === 403){
        toast.error("Not Authorized")
        return redirect("/")
    }

    toast.error(errMsg)
    return null
  }
}

const AdminLanding = () => {
  const {doctors, numOfPages, page} = useLoaderData()

  if (doctors.length === 0){
    return (
      <div className="bg-gray-100">
        <h1 className='text-xl'>No Doctor profile created</h1>
      </div>
    )
  }
  
  return (
    <div className="bg-gray-100">
    <div className="max-w-7xl mx-auto py-6 sm:px-6 lg:px-8">
      {/* Filter Section */}
      {/* <FilterSection isFilterApplied={isFilterApplied} myFilter={myFilter} handleOpenFilter={handleOpenFilter} handleSearch={handleSearch}/> */}

      <Doctors doctors={doctors}/>

      {/* Pagination */}
      <Pagination numOfPages={numOfPages} page={page} />
      
    </div>
  </div>
  )
}

export default AdminLanding