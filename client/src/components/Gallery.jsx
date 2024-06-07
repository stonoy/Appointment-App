import React from 'react';
import {Form, Link} from 'react-router-dom'

const Gallery = ({ avaliabilities }) => {
    console.log(avaliabilities)
  return (
    <div className="grid gap-6 sm:grid-cols-2 lg:grid-cols-3">
      {avaliabilities.map((availability) => (
        <div key={availability.id} className="bg-white shadow-md rounded-lg overflow-hidden">
          <div className="p-4">
            <h3 className="text-xl font-bold mb-2">{availability.treatment}</h3>
            <p className="text-gray-700 mb-1"><strong>Doctor:</strong> {availability.doctor_name}</p>
            <p className="text-gray-700 mb-1"><strong>Specialty:</strong> {availability.specialty}</p>
            <p className="text-gray-700 mb-1"><strong>Location:</strong> {availability.location}</p>
            <p className="text-gray-700 mb-1"><strong>Timing:</strong> {new Date(availability.timing).toLocaleString()}</p>
            <p className="text-gray-700 mb-1"><strong>Duration:</strong> {availability.duration} mins</p>
            <p className="text-gray-700 mb-1"><strong>Patients:</strong> {availability.current_patient}/{availability.max_patient}</p>
            <Form method='post' action={`/createappointment/${availability.id}`}>
            <button className="mt-4 inline-block bg-indigo-500 text-white px-4 py-2 rounded-lg hover:bg-indigo-600">
                  See
                </button>
            </Form>
          </div>
        </div>
      ))}
    </div>
  );
};

export default Gallery;
