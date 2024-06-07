import React from 'react'
import { customFetch } from '../utils'
import { toast } from 'react-toastify'
import { useLoaderData } from 'react-router-dom'
import { Gallery, Pagination } from '../components'

export const loader = async () => {
  try {
    const resp = await customFetch('/getavailabilities')
    // console.log(resp.data)
    return resp.data
  } catch (error) {
    const errorMsg = error?.response?.data?.msg || 'Error in getting doctor availabilities'

      toast.error(errorMsg)

      return null
  }
}

const Landing = () => {
  const {avaliabilities, numOfPages, page} = useLoaderData()

  if (avaliabilities.length === 0){
    return (
      <div className="bg-gray-100">
        <h1 className='text-xl'>No Doctor is available now</h1>
      </div>
    )
  }

  return (
    <div className="bg-gray-100">
      <div className="max-w-7xl mx-auto py-6 sm:px-6 lg:px-8">
        {/* Filter Section */}
        {/* <FilterSection isFilterApplied={isFilterApplied} myFilter={myFilter} handleOpenFilter={handleOpenFilter} handleSearch={handleSearch}/> */}

        <Gallery avaliabilities={avaliabilities}/>

        {/* Pagination */}
        <Pagination numOfPages={numOfPages} page={page} />

        {/* Filter */}
        
      </div>
    </div>
  )
}

export default Landing