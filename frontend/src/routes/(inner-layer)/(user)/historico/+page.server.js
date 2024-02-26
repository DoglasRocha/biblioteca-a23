import { api } from "$lib/utils/api.js"

export async function load({ cookies }) {
    const accessToken = cookies.get("accessToken")
    const headers = {
        Cookie: `accessToken=${accessToken}`
    }

    try {
        let requests = await api.get(`/solicitacoes`, {
            headers: headers
        })

        let loans = await api.get("/emprestimos", {
            headers: headers
        })

        return {
            requests: requests.data,
            loans: loans.data,
            error: null
        }
    } catch (error) {
        return {
            requests: [],
            loans: [],
            error: error.response.data
        }
    }
}