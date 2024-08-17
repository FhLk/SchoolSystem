import { accountData } from "./AccountAPI"

const URL_BASE = process.env.APP_BASE_URL

export interface librarianData {
    id : string
    firstName : string
    middleName : string
    lastName : string
    email : string
    role : string
    account : accountData
}

export const getAllLibrarians = async ()=>{
    try{
        const response = await fetch(`http://localhost:8080/api/librarian`)
        const json = await response.json()
        return json
    }catch (e){
        console.error(e);
    }
}

export const getLibrarianByID = async (id : string) =>{
    try{
        const response = await fetch(`http://localhost:8080/api/librarian/${id}`)
        const json = await response.json()
        return json
    }catch (e){
        console.error(e);
    }
}

export const createNewLibrarian = async (newLibrarian : librarianData) =>{
    try{
        const response = await fetch(`http://localhost:8080/api/librarian`,{
            method : "POST",
            headers : {
                "Content-Type": "application/json",
            },
            body : JSON.stringify(newLibrarian)
        })
        const json = await response.json()
        return json
    }catch (e){
        console.error(e);
    }
}

export const updateAllDataLibrarian = async (newLibrarian : librarianData) =>{
    try{
        const response = await fetch(`http://localhost:8080/api/librarian/${newLibrarian.id}`,{
            method : "PUT",
            headers : {
                "Content-Type": "application/json",
            },
            body : JSON.stringify(newLibrarian)
        })
        const json = await response.json()
        return json
    }catch (e){
        console.error(e);
    }
}

export const updateSomeDataLibrarian = async (data : string, id : string)=>{
    try{
        const response = await fetch(`http://localhost:8080/api/librarian/${id}`,{
            method : "PATCH",
            headers : {
                "Content-Type": "application/json",
            },
            body : JSON.stringify(data)
        })
        const json = await response.json()
        return json
    }catch (e){
        console.error(e);
    }
}

export const deteleLibrarian = async (id : string)=>{
    try{
        await fetch(`http://localhost:8080/api/librarian/${id}`,{
            method : "DELETE",
        })
    }catch (e){
        console.error(e);
    }
}