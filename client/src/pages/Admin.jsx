import React from 'react'
import { NavLink, Outlet, useOutletContext } from 'react-router-dom';

const Admin = () => {
  return (
    <div className="container mx-auto p-4">
      <div className="flex pb-4 justify-between items-center">
        <h1 className="text-xl font-semibold">Admin</h1>
        <div className="flex space-x-4">
          <NavLink to="/admin/insights" className="text-blue-500 hover:text-blue-700">
            Insights
          </NavLink>
        </div>
      </div>
      <Outlet /> {/* The Outlet renders child components */}
    </div>
  );
}

export default Admin