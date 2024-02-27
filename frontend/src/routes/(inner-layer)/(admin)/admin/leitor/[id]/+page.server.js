import { api } from "$lib/utils/api.js"

export async function load({ cookies, params }) {
    const accessToken = cookies.get("accessToken")

    try {
        let request = await api.get(`/admin/leitor/${params.id}`, {
            headers: {
                Cookie: `accessToken=${accessToken}`
            }
        })

        return {
            readerData: request.data,
            error: null
        }
    } catch (error) {
        return {
            readerData: null,
            error: error.response.data
        }
    }
}