import React from 'react';

const Appointment = ({ appointments }) => {
  return (
    <div className="grid gap-6 sm:grid-cols-2 lg:grid-cols-3">
      {appointments.map((appointment) => (
        <div key={appointment.id} className="bg-white shadow-md rounded-lg overflow-hidden">
          <div className="p-4">
            <h3 className="text-xl font-bold mb-2">{appointment.treatment}</h3>
            <p className="text-gray-700 mb-1"><strong>Status:</strong> {appointment.status}</p>
            <p className="text-gray-700 mb-1"><strong>Doctor:</strong> {appointment.doctor_name}</p>
            <p className="text-gray-700 mb-1"><strong>Specialty:</strong> {appointment.specialty}</p>
            <p className="text-gray-700 mb-1"><strong>Location:</strong> {appointment.location}</p>
            <p className="text-gray-700 mb-1"><strong>Timing:</strong> {new Date(appointment.timing).toLocaleString()}</p>
            <p className="text-gray-700 mb-1"><strong>Duration:</strong> {appointment.duration} minutes</p>
          </div>
        </div>
      ))}
    </div>
  );
};

export default Appointment;
