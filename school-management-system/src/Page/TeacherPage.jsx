import React from "react";

function TeacherPage() {
  return (
    <div className="infor-div">
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
      <div>
        <h3>Class Room List</h3>
        <ul>
          <li>M. 1/1 <button>Check List</button></li>
          <li>M. 1/2 <button>Check List</button></li>
          <li>M. 1/3 <button>Check List</button></li>
        </ul>
      </div>
    </div>
  );
}

export default TeacherPage;
