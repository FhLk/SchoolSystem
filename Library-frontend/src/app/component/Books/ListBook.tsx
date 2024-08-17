"use client"

import React, { FC, useEffect, useState } from 'react'
import { bookData, getAllBooks, getBookByID } from './BookAPI'



const ListBook : FC = ()=> {
  const [books, setBooks] = useState<bookData[]>([])

  useEffect(()=>{
    const fetchBooks = async () => {
      // const books = await getBookByID("00");
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

