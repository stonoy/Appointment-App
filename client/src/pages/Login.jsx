import React from 'react'
import { Form, Link,redirect, useNavigation } from 'react-router-dom';
import { FormInput, SubmitBtn } from '../components';
import { customFetch } from '../utils';

import { toast } from 'react-toastify';
import { setUser } from '../feature/user/userSlice';
 

export const action = (store) => async ({request}) => {
  const formData = await request.formData()
  const data = Object.fromEntries(formData)
  // console.log(data)

  

  try {
    const resp = await customFetch.post('/login', data)
    // console.log(resp.data)
    store.dispatch(setUser(resp.data))
    toast.success("user logged in successfully")
    return redirect('/')
  } catch (error) {
    
    const errorMsg = error?.response?.data?.msg || 'Error in Logging user'

    toast.error(errorMsg)

    return null
  }
  
}


const Login = () => {
  const navigation = useNavigation()
  const isLoading = navigation.state == "submitting"
  return (
    <div className="min-h-screen align-element bg-gray-100 flex flex-col justify-center py-12 sm:px-6 lg:px-8">
      <div className="sm:mx-auto sm:w-full sm:max-w-md">
        <h2 className="mt-6 text-center text-3xl font-extrabold text-gray-900">Sign in to your account</h2>
      </div>

      <div className="mt-8 sm:mx-auto sm:w-full sm:max-w-md">
        <div className="bg-white py-8 px-4 shadow sm:rounded-lg sm:px-10">
          <Form method='post' className="space-y-6" >
          <FormInput 
            type="text"
            label="Email"
            name="email"
            defaultValue=""
          />

<FormInput 
            type="password"
            label="Password"
            name="password"
            defaultValue=""
          />

          <SubmitBtn
            type="submit"
            text="Sign In"
            isLoading={isLoading}
          />
          </Form>
          <p className="mt-2 text-sm text-center text-gray-600">Not a member? <Link to="/register">Register</Link></p>
          <p className="mt-2 hidden text-sm text-center text-red-600">the error</p>

          {/* Link to Products Page */}
          <p className="mt-2 text-sm text-center text-gray-600">
              <Link to="/" className="text-blue-500">Browse Doctor Availability</Link>
            </p>
        </div>
      </div>
    </div>
  );
  
}

export default Login