import React from 'react'
import { NavLink } from 'react-router-dom'

const Navlink = ({name, link, bigScreen}) => {
  
  return (
    <li><NavLink to={link} className={bigScreen ? "hover:text-gray-300" : "block text-white py-2 px-4 hover:bg-gray-700"}>{name}</NavLink></li>
  )
}

export default Navlink