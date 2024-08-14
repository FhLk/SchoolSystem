"use client"

import React, { FC, useEffect, useState } from 'react'
import { getAllBooks } from './BookAPI'

interface bookData {
  id : string
  img : string
  title : string
  category : string
  author : string
  publisher : string
  year : Number
  status : Number
}

const ListBook : FC = ()=> {
  const [books, setBooks] = useState<bookData[]>([])

  useEffect(()=>{
    const fetchBooks = async () => {
      const books = await getAllBooks();
      setBooks(books);
   }
   fetchBooks()
  },[])
  return (
    <div>
      
    </div>
  )
}

export default ListBook

