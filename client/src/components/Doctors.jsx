import React from 'react';

const Doctors = ({ doctors }) => {
  return (
    <div className="overflow-x-auto">
      <table className="min-w-full bg-white">
        <thead>
          <tr>
            <th className="py-2 px-4 border-b text-left text-gray-700 font-bold">Name</th>
            <th className="py-2 px-4 border-b text-left text-gray-700 font-bold">Specialty</th>
            <th className="py-2 px-4 border-b text-left text-gray-700 font-bold">License Number</th>
          </tr>
        </thead>
        <tbody>
          {doctors.map((doctor) => (
            <tr key={doctor.id} className="bg-gray-100 hover:bg-gray-200">
              <td className="py-2 px-4 border-b">{doctor.name}</td>
              <td className="py-2 px-4 border-b">{doctor.specialty}</td>
              <td className="py-2 px-4 border-b">{doctor.licence_number}</td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
};

export default Doctors;
