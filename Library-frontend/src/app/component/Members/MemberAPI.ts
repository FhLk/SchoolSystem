const URL_BASE = process.env.APP_BASE_URL

export interface memberData{
    id : string
    firstName : string
    middleName : string
    lastName : string
    class : string
    exp : Date
}

export const getAllMember = async () =>{
    try {
        const response = await fetch(`http://localhost:8080/api/member`)
        const json = await response.json()
        return json
    } catch (error) {
        console.error(error);
    }
}

export const getMemberByID = async ( id : string) =>{
    try {
        const response = await fetch(`http://localhost:8080/api/member/${id}`)
        const json = await response.json()
        return json
    } catch (error) {
        console.error(error);
    }
}

export const createNewMember = async (newMember : memberData) =>{
    try {
        const response = await fetch(`http://localhost:8080/api/member`,{
            method : "POST",
            headers : {
                "Content-Type": "application/json",
            },
            body : JSON.stringify(newMember)
        })
        const json = await response.json()
        return json
    } catch (error) {
        console.error(error);
    }
}

export const updateAllDataMember = async (newMember : memberData) =>{
    try {
        const response = await fetch(`http://localhost:8080/api/member/${newMember.id}`,{
            method : "PUT",
            headers : {
                "Content-Type": "application/json",
            },
            body : JSON.stringify(newMember)
        })
        const json = await response.json()
        return json
    } catch (error) {
        console.error(error);
    }
}

export const updateSomeDataMember = async (data : string, id :string) =>{
    try {
        const response = await fetch(`http://localhost:8080/api/member/${id}`,{
            method : "PATCH",
            headers : {
                "Content-Type": "application/json",
            },
            body : JSON.stringify(data)
        })
        const json = await response.json()
        return json
    } catch (error) {
        
    }
}

export const deleteMember = async (id : string) =>{
    try {
        await fetch(`http://localhost:8080/api/member/${id}`,{
            method : "DELETE",
        })
    } catch (error) {
        console.error(error);
    }
}