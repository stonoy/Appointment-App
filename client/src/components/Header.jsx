import React from 'react'
import { Link, useNavigate } from 'react-router-dom'
import {useDispatch, useSelector} from 'react-redux'
import { logout } from '../feature/user/userSlice'



const Header = () => {
  const {token, user} = useSelector((state) => state.user)
  const navigate = useNavigate()
  const dispatch = useDispatch()


  const handleLogout = () => {
    dispatch(logout())

    navigate("/login")
  }
 
  return (
    <header className="bg-gray-500 text-white p-2">
      <div className="container align-element mx-auto flex items-center">

        <div className='ml-auto flex'>
          <h2 className='mr-4'>{user.name}</h2>
          <div>
          {
            token ? 
            <button onClick={handleLogout} className="text-gray-300 hover:text-white mr-4">Logout</button>
            :
            <Link to="/login" className="text-gray-300 hover:text-white mr-4">Login/Register</Link>
          }
        
          </div>
        </div>
      </div>
    </header>
  )
}

export default Header