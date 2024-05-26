import React from 'react'
import { Link } from 'react-router-dom'

function Classlist() {
  return (
    <div className='flex justify-around py-3'>
      <Link to="/class">
        <button className="btn text-3xl h-20">M. 1/1</button>
      </Link>
      <Link to="/class">
        <button className="btn text-3xl h-20">M. 1/2</button>
      </Link>
      <Link to="/class">
        <button className="btn text-3xl h-20">M. 1/3</button>
      </Link>
    </div>
  )
}

export default Classlist
