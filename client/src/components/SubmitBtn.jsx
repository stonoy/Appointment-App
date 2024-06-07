import React from 'react'
import { FaSpinner } from 'react-icons/fa'

const SubmitBtn = ({text, type, isLoading}) => {
  return (
    <div>
              <button
                type={type}
                disabled={isLoading}
                className="w-full flex justify-center py-2 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
              >
                {/* Loading spinner */}
        {isLoading && (
          <FaSpinner className="animate-spin mr-2" /> // Add animate-spin class for animation
        )}
        {/* Button text */}
                {text}
              </button>
            </div>
  )
}

export default SubmitBtn