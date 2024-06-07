import React from 'react'
import { Outlet, useNavigation } from 'react-router-dom'
import { Header, Loader, Navbar } from '../components'

const HomeLayOut = () => {
  const navigation = useNavigation()
  const isLoading = navigation.state === 'loading'
  return (
    <>
      <Header/>
      <Navbar/>
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