import { api } from "$lib/utils/api.js"

export async function load({ cookies }) {
    let accessToken = cookies.get("accessToken")

    try {
        let request = await api.get("/admin/admins", {
            headers: {
                Cookie: `accessToken=${accessToken}`
            }
        })

        return {
            admins: request.data,
            error: null
        }
    }
    catch (error) {
        return {
            admins: [],
            error: error.response.data
        }
    }
}