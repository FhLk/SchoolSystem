import React, { useState } from 'react'

function StudentList({handleOnCheck}) {

  const [students, setStudents] = useState([
    {
      name: 'arun',
      gender: 'Male',
      isPresent: false
    },
    {
      name: 'rajesh',
      gender: 'Male',
      isPresent: false
    },
    {
      name: 'moorthy',
      gender: 'Male',
      isPresent: false
    },
    {
      name: 'raja',
      gender: 'Male',
      isPresent: false
    },
    {
      name: 'usha',
      gender: 'Female',
      isPresent: false
    },
  ])

  const handlePresent = (studentIndex) => {
    students[studentIndex].isPresent = true
  }

  const handleAbsent = (studentIndex) => {
    students[studentIndex].isPresent = false
  }

  function handleClick() {
    handleOnCheck(students);
  }

  return (
    <div>
      <ul>
        {students.map((student, i) => (
          <li key={i} className='flex justify-around items-center space-x-6 py-3'>
            {i + 1}. {student.name.toUpperCase()}
            <div className='flex items-center'>
              <input
                onChange={() => handlePresent(i)}
                id={`present-${i}`}
                type='radio'
                name={`student-${i}-presence`}
                className="w-4 h-4 text-blue-600 bg-gray-100 border-gray-300 focus:ring-blue-500 dark:focus:ring-blue-600 dark:ring-offset-gray-700 dark:focus:ring-offset-gray-700 focus:ring-2 dark:bg-gray-600 dark:border-gray-500" />
              <label
                htmlFor={`present-${i}`}
                className="w-full py-3 ms-2 text-sm font-medium">
                Present
              </label>
            </div>
            <div className='flex items-center'>
              <input
                onChange={() => handleAbsent(i)}
                id={`absent-${i}`}
                type='radio'
                name={`student-${i}-presence`}
                className="w-4 h-4 text-blue-600 bg-gray-100 border-gray-300 focus:ring-blue-500 dark:focus:ring-blue-600 dark:ring-offset-gray-700 dark:focus:ring-offset-gray-700 focus:ring-2 dark:bg-gray-600 dark:border-gray-500" />
              <label
                htmlFor={`absent-${i}`}
                className="w-full py-3 ms-2 text-sm font-medium">
                Absent
              </label>
            </div>
          </li>
        ))}
      </ul>
      <button onClick={handleClick} className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded">
        Check
      </button>
    </div>
  )
}

export default StudentList
