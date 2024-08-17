const URL_BASE = process.env.APP_BASE_URL

export interface bookData {
    id : string
    img : string
    title : string
    category : string
    author : string
    publisher : string
    year : Number
    status : Number
  }

export const getAllBooks = async () => {
    try {
        const response = await fetch("http://localhost:8080/api/book");
        const jsonData = await response.json();
        return jsonData
    }
    catch (e) {
        console.error(e);
    }
}

export const getBookByID = async (id : string) =>{
    try {
        const response = await fetch(`http://localhost:8080/api/book/${id}`);
        const jsonData = await response.json();
        return jsonData
    }
    catch (e) {
        console.error(e);
    }
}

export const createNewBook = async (newBook : bookData) =>{
    try {
        const response = await fetch(`http://localhost:8080/api/book/`,{
            method : "POST",
            headers:{
                "Content-Type": "application/json",
            },
            body: JSON.stringify(newBook)
        });
        const jsonData = await response.json();
        return jsonData
    }
    catch (e) {
        console.error(e);
    }
}

export const updateBookAllData = async (newBook : bookData) => {
    try {
        const response = await fetch(`http://localhost:8080/api/book/${newBook.id}`,{
            method : "PUT",
            headers:{
                "Content-Type": "application/json",
            },
            body: JSON.stringify(newBook)
        });
        const jsonData = await response.json();
        return jsonData
    }
    catch (e) {
        console.error(e);
    }
}

export const updateBookSomeData = async (data : string, id :string) => {
    try {
        const response = await fetch(`http://localhost:8080/api/book/${id}`,{
            method : "PATCH",
            headers:{
                "Content-Type": "application/json",
            },
            body: JSON.stringify(data)
        });
        const jsonData = await response.json();
        return jsonData
    }
    catch (e) {
        console.error(e);
    }
}

export const deleteBook = async (id : string) =>{
    try {
        const response = await fetch(`http://localhost:8080/api/book/${id}`,{
            method : "DETELE",
            headers:{
                "Content-Type": "application/json",
            },
        });
        const jsonData = await response.json();
        return jsonData
    }
    catch (e) {
        console.error(e);
    }
}