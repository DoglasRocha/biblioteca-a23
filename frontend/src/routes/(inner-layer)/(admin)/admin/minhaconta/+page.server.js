import { api } from "$lib/utils/api.js"
import { redirect } from "@sveltejs/kit"

export async function load({cookies}) {

    let accessToken = cookies.get("accessToken")
    try {
        let request = await api.get("/admin/minhaconta", {
            headers: {
                Cookie: `accessToken=${accessToken}`
            }
        })
        let user = request.data.User
        return {
            userData: {
                name: user.name,
                surname: user.surname,
                email: user.email,
            },
            error: null
        }
    }
    catch (error) {

        redirect(303, "/admin/error")
    }

    
}