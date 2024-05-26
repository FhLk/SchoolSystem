import React, { useState } from 'react'
import Classlist from './Classlist'
import Rating from './Rating'

function index() {
  return (
    <div data-theme="light" className='p-10 px-32'>
      <div className='flex'>
        <img
          src="example1.jpg"
          alt=""
          style={{
            width: "300px",
            height: "300px",
            borderRadius: "10%",
            objectFit: "cover",
            objectPosition: "-20px 0",
          }}
        />
        <div className='px-32 text-5xl w-full'>
          <p>Mr. Paramat Pet-in</p>
          <p className='text-sm py-3'>Subject: Computor Programing</p>
          <Rating/>
          <h1 className='text-3xl'>Class List</h1>
          <Classlist />
        </div>
      </div>
    </div>
  )
}

export default index
