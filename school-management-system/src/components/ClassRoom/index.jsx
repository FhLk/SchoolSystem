import React, { useState } from 'react'
import StudentList from './StudentList'
import { Link } from 'react-router-dom'

function index() {
    const [presentCount, setPresentCount] = useState(0)
    const [classNameLevel, setclassNameLevel] = useState("1/1")

    const handleCheckList = (data) => {
        setPresentCount(data.filter((student) => student.isPresent).length)
    }

    return (
        <div className='p-5'>
            <Link to="/">
            <button class="bg-white hover:bg-gray-100 text-gray-800 font-semibold py-2 px-4 border border-gray-400 rounded shadow">
                Back
            </button>
            </Link>
            <h1 className='text-3xl'>Class Room {classNameLevel}</h1>
            <StudentList handleOnCheck={handleCheckList} />
            <div>
                Total Present: {presentCount}
            </div>
        </div>
    )
}

export default index
