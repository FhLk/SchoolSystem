import { bookData } from "./ListBook";


const URL_BASE = process.env.APP_BASE_URL

export const getAllBooks = async () => {
    try {
        const response = await fetch("http://localhost:8080/api/book");
        if (!response.ok) {
            throw new Error(`Response status: ${response.status}`);
        }
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
        if (!response.ok) {
            throw new Error(`Response status: ${response.status}`);
        }
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
        if (!response.ok) {
            throw new Error(`Response status: ${response.status}`);
        }
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
        if (!response.ok) {
            throw new Error(`Response status: ${response.status}`);
        }
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
        if (!response.ok) {
            throw new Error(`Response status: ${response.status}`);
        }
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
        if (!response.ok) {
            throw new Error(`Response status: ${response.status}`);
        }
        const jsonData = await response.json();
        return jsonData
    }
    catch (e) {
        console.error(e);
    }
}