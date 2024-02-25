import { api } from "$lib/utils/api.js"

export async function load({ cookies }) {

    const accessToken = cookies.get("accessToken")

    try {
        let booksRequest = await api.get("/admin/emprestimos/ativos", {
            headers: {
                Cookie: `accessToken=${accessToken}`
            }
        })

        return {
            loans: booksRequest.data,
            error: null
        }
    } catch (error) {
        return {
            loans: [],
            error: error.response.data
        }
    }


}
