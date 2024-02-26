import { api } from "$lib/utils/api.js"

export async function load({cookies}) {

    let accessToken = cookies.get("accessToken")
    try {
        let request = await api.get("/emprestimos/ativo", {
            headers: {
                Cookie: `accessToken=${accessToken}`
            }
        })

        return {
            loan: request.data,
            error: null
        }
    }
    catch (error) {
        return {
            loan: null,
            error: error.response.data
        }
    }
}