import { api } from "$lib/utils/api.js"

export async function load({ cookies }) {
    const accessToken = cookies.get("accessToken")

    try {
        let requests = await api.get(`/solicitacoes`, {
            headers: {
                Cookie: `accessToken=${accessToken}`
            }
        })

        return {
            requests: requests.data,
            error: null
        }
    } catch (error) {
        return {
            requests: [],
            error: error.response.data
        }
    }
}