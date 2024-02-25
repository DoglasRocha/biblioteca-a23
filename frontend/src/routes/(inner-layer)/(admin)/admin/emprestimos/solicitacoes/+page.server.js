import { api } from "$lib/utils/api.js"

export async function load({ cookies }) {

    const accessToken = cookies.get("accessToken")

    try {
        let booksRequest = await api.get("/admin/emprestimos/solicitacoes", {
            headers: {
                Cookie: `accessToken=${accessToken}`
            }
        })

        return {
            requests: booksRequest.data,
            error: null
        }
    } catch (error) {
        return {
            requests: null,
            error: "Não há livros"
        }
    }


}