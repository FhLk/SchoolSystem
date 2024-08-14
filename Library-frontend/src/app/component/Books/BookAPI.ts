

const URL_BASE = process.env.APP_BASE_URL

export const getAllBooks = async () => {
    try {
        const response = await fetch("http://localhost:8080/book",{
            method :"GET",
            headers: {
                'Content-Type': 'application/json',
            }
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