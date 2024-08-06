"use client"
import React from 'react'

interface bookData {
  category : string,
  image : string | null,
  title : string,
  author : string,
  publisher: string,
  year : number,
  isbn : string | null,
  status: number
}

function ListBook() {
  const BOOKS : bookData[] = [
    { category: "คณิตศาสตร์",image : null,title: "แบบฝึกหัดคณิตศาสตร์ ป.6", author: "สถาบันส่งเสริมการสอนวิทยาศาสตร์และเทคโนโลยี", publisher: "สสวท.", year: 2024, isbn: null, status: 0 },
    { category: "คณิตศาสตร์",image : null,title: "แข่งขันคณิตศาสตร์ ม.ต้น", author: "ศ.ดร. นิวัฒน์ พาณิชยพงศ์", publisher: "ชญาน์บุ๊ค", year: 2024, isbn: null, status: 0 },
    { category: "วิทยาศาสตร์",image : null,title: "วิทยาศาสตร์รอบตัว", author: "กลุ่มนักเขียน", publisher: "สำนักพิมพ์อมรินทร์", year: 2023, isbn: "978-974-249-876-5",status: 0 },
    { category: "วิทยาศาสตร์",image : null,title: "ทดลองวิทยาศาสตร์ง่ายๆ", author: "อาจารย์ ดร. ณัฐพล  จันทร์เจริญ", publisher: "อมรินทร์คิดส์", year: 2021, isbn: "978-974-9759-87-6",status: 0 },
    { category: "ภาษาอังกฤษ",image : null,title: "Vocabulary Builder", author: "หมอแอน", publisher: "English For You", year: 2023, isbn: "978-616-7891-23-4",status: 0 },
    { category: "ภาษาอังกฤษ",image : null,title: "English Grammar in Use", author: "Raymond Murphy", publisher: "Cambridge University Press", year: 2021, isbn: "978-1-107-69308-9",status: 0 },
    { category: "ภาษาไทย",image : null,title: "ไวยากรณ์ภาษาไทย ม.3", author: "อาจารย์สุภัทร สุวรรณ", publisher: "พิงค์กีส์การศึกษา", year: 2022, isbn: "978-616-442-123-4",status: 1 },
    { category: "ภาษาไทย",image : null,title: "คำศัพท์ภาษาไทย", author: "กานต์พิชชา กุลวัฒนายานนท์", publisher: "านักเขียน", year: 2020, isbn: "978-616-9231-23-4",status: 0 },
    { category: "สังคมศึกษา",image : null,title: "ประวัติศาสตร์ไทย", author: "ชุมนุมนักประวัติศาสตร์ไทย", publisher: "คุรุสภา", year: 2020, isbn: "978-974-07-3589-7",status: 0 },
    { category: "สังคมศึกษา",image : null, title: "ภูมิศาสตร์โลก", author: "คณะกรรมการจัดทำหนังสือสังคมศึกษา", publisher: "สำนักพิมพ์ปทุมวัน", year: 2022, isbn: "978-974-267-123-4",status: 0 },
];


  return (
    <div>
      <ul className='grid grid-cols-3 gap-y-10'>
        {BOOKS.map((item,i)=>(
          <li key={i} className='flex flex-col items-center'>
            <img src='./image-placeholder.svg' width="100px" height="100px"/>
            <p>{item.title}</p>
            {/* <p>Author: {item.author}</p> */}
          </li>
        ))}
      </ul>
    </div>
  )
}

export default ListBook
