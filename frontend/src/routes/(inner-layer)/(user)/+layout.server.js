import { redirect } from "@sveltejs/kit";
import { api } from "$lib/utils/api.js"

export async function load({ cookies }) {

    const accessToken = cookies.get("accessToken")

    try {
        let request = await api.post("/check", {}, {
            withCredentials: true,
            headers: {
                Cookie: `accessToken=${accessToken}`,
            }
        })
        
        console.log(request)
    } catch (error) {
        if (error?.code == 'ECONNREFUSED')
            redirect(303, "/error")

        else if (error.response.status == 403) 
            redirect(303, "/admin")
        
        else if (error.response.status == 401)
            redirect(303, "/login")

        else
            redirect(303, '/error')
    }

    return {}
}