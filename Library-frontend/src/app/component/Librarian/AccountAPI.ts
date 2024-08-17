const URL_BASE = process.env.APP_BASE_URL

export interface accountData {
    account_id : string
    username : string
    password : string
}

export const getAllAccount = async () =>{
    try {
        const response = await fetch(`http://localhost:8080/api/account`)
        const json = await response.json()
        return json
    } catch (error) {
        console.error(error);
    }
}

export const getAccountByID = async (id : string) =>{
    try {
        const response = await fetch(`http://localhost:8080/api/account/${id}`)
        const json = await response.json()
        return json
    } catch (error) {
        console.error(error);
    }
}

export const createNewAccount = async (newAccount : accountData) => {
    try {
        const response = await fetch(`http://localhost:8080/api/account`,{
            method : "POST",
            headers : {
                "Content-Type": "application/json",
            },
            body : JSON.stringify(newAccount)
        })
        const json = await response.json()
        return json
    } catch (error) {
        console.error(error);
    }
}

export const updateAllDataAccount = async (newAccount : accountData) => {
    try {
        const response = await fetch(`http://localhost:8080/api/account/${newAccount.account_id}`,{
            method : "PUT",
            headers : {
                "Content-Type": "application/json",
            },
            body : JSON.stringify(newAccount)
        })
        const json = await response.json()
        return json
    } catch (error) {
        console.error(error);
    }
}

export const updateSomeDataAccount = async (data : string, id : string) => {
    try {
        const response = await fetch(`http://localhost:8080/api/account/${id}`,{
            method : "PATCH",
            headers : {
                "Content-Type": "application/json",
            },
            body : JSON.stringify(data)
        })
        const json = await response.json()
        return json
    } catch (error) {
        console.error(error);
    }
}

export const deleteAccount = async (id : string) => {
    try {
        await fetch(`http://localhost:8080/api/account/${id}`,{
            method : "DELETE",
        })
    } catch (error) {
        console.error(error);
    }
}
