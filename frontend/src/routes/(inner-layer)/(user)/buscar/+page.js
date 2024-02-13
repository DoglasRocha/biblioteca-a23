import { api } from "$lib/utils/api.js"

export async function load() {

    try {
        let booksRequest = await api.get("/livros/buscar")

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