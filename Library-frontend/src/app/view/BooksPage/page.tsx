import ListBook from '@/app/component/Books/ListBook'
import React from 'react'

function BooksPage() {
  return (
    <div className='min-h-screen flex justify-between'>
      <div className='bg-white w-[200px] my-[17px] rounded-xl ml-[30px] mr-[40px]'>
        <h1>wow 1</h1>
      </div>
      <div className='w-[1240px] my-[17px] rounded-xl flex flex-col'>
        <div className='flex justify-end'>
          <input className='border-2 border-black w-1/3 rounded-lg h-[45px] px-5' placeholder='Search'/>
        </div>
        <div className='bg-white h-full mt-[22px] rounded-xl p-[50px]'>
          <h1>Books</h1>
          <button>Add</button>
        </div>
      </div>
      <div className='bg-white w-[340px] my-[17px] rounded-xl ml-[40px] mr-[30px]'>
        <h1>wow 3</h1>
      </div>
    </div>
  )
}

export default BooksPage
