import React, { useState } from 'react'
import Classlist from './Classlist'

function index() {
  return (
    <div>
      <h1>Teacher Information</h1>
      <div>
        <img
          src="example1.jpg"
          alt=""
          style={{
            width: "200px",
            height: "200px",
            borderRadius: "50%",
            objectFit: "cover",
            objectPosition: "-20px 0",
          }}
        />
      </div>
      <p>Mr. Paramat Pet-in</p>
      <p>Subject: Computor Programing</p>
      <Classlist/>
    </div>
  )
}

export default index
