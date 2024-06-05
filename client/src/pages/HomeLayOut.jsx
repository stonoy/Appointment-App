import React from 'react'
import { Outlet, useNavigation } from 'react-router-dom'
import { Loader } from '../components'

const HomeLayOut = () => {
  const navigation = useNavigation()
  const isLoading = navigation.state === 'loading'
  return (
    <>
      {
        isLoading ?
        <Loader/> 
        :
        <section className='align-element py-10'><Outlet/></section>
      }
    </>
  )
}

export default HomeLayOut