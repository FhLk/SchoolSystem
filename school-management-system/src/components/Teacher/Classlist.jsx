import React from 'react'
import { Link } from 'react-router-dom'

function Classlist() {
  return (
    <div>
      <h3>Class Room List</h3>
      <ul>
        <li>M. 1/1
          <Link to="/class">
            <button class="bg-white hover:bg-gray-100 text-gray-800 font-semibold py-2 px-4 border border-gray-400 rounded shadow">
              Attendance
            </button>
          </Link>
        </li>
        <li>M. 1/2
          <Link to="/class">
            <button class="bg-white hover:bg-gray-100 text-gray-800 font-semibold py-2 px-4 border border-gray-400 rounded shadow">
              Attendance
            </button>
          </Link>
        </li>
        <li>M. 1/3
          <Link to="/class">
            <button class="bg-white hover:bg-gray-100 text-gray-800 font-semibold py-2 px-4 border border-gray-400 rounded shadow">
              Attendance
            </button>
          </Link>
        </li>
      </ul>
    </div>
  )
}

export default Classlist
