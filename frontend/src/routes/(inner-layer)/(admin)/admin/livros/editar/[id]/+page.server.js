import { api } from "$lib/utils/api.js"

export async function load({ cookies, params }) {
    const accessToken = cookies.get("accessToken")

    try {
        let bookRequest = await api.get(`/admin/livros/buscar/${params.id}`, {
            headers: {
                Cookie: `accessToken=${accessToken}`
            }
        })

        return {
            book: bookRequest.data,
            error: null
        }
    } catch (error) {
        return {
            book: null,
            error: "Não há livro com esse identificador"
        }
    }
}