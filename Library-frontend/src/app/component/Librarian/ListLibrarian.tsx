"use client"

import React, { FC, useState } from 'react'
import { accountData, librarianData } from './LibrarianAPI'

const ListLibrarian : FC =()=> {
    const [librarians , setLibrarians] = useState<librarianData[]>([])
    const [accounts, setAccounts] = useState<accountData[]>([])
  return (
    <div>
      
    </div>
  )
}

export default ListLibrarian
