import { api } from "$lib/utils/api.js"

export async function load({ cookies }) {

    const accessToken = cookies.get("accessToken")

    try {
        let booksRequest = await api.get("/livros/buscar", {
            headers: {
                Cookie: `accessToken=${accessToken}`
            }
        })

        return {
            books: booksRequest.data,
            error: null
        }
    } catch (error) {
        return {
            books: null,
            error: "Não há livros"
        }
    }


} 