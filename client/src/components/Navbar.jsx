import React, { useState } from 'react'
import { links } from '../utils'
import { RxHamburgerMenu } from "react-icons/rx"
import { useSelector } from 'react-redux'
import Navlink from './Navlink'

const Navbar = () => {
    const [showDropDown, setShowDropDown] = useState(false)

    // get the user to check that user is online
    const {token, user:{name,role}} = useSelector(state => state.user)

  return (
    <header className="bg-gray-800 text-white p-4">
      <div className="container align-element mx-auto flex justify-between items-center">
        <div className="text-xl font-bold">Appointments</div>
        <nav className="hidden flex-grow text-center sm:block">
          <ul className="flex justify-center space-x-4">
            {links.map(link => {
              
              if (!token && link.name === "Appointments"){return null}
              if (role !== "doctor" && link.name === "Availability"){return null}
              if (role !== "admin" && link.name === "Admin"){return null}
              return <Navlink key={link.id} {...link} bigScreen={true}/>
            })}
          </ul>
        </nav>
        <div className='sm:hidden'>
        <RxHamburgerMenu onClick={() => setShowDropDown(!showDropDown)} className="fa-solid text-2xl fa-bars text-md"></RxHamburgerMenu>
        </div>
      </div>
      {/* Dropdown */}
        {showDropDown && 
            (
                <div className=" absolute top-20 right-4 bg-gray-800 z-10">
                <ul className="flex flex-col justify-center items-center py-2">
                  {links.map(link => {
                    if (!token && link.name === "Appointments"){return null}
                    if (role !== "doctor" && link.name === "Availability"){return null}
                    if (role !== "admin" && link.name === "Admin"){return null}
                    return <Navlink key={link.id} {...link} bigScreen={false}/>
                  })}
                </ul>
              </div>
            )
        }
    </header>
  )
}

export default Navbar