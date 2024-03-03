import { redirect } from "@sveltejs/kit";
import { api } from "$lib/utils/api.js";

export async function load({ cookies }) {

    const accessToken = cookies.get("accessToken")
    console.log(accessToken)

    try {
        let request = await api.post("admin/check", {}, {
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
            redirect(303, "/")
        
        else if (error.response.status == 401)
            redirect(303, "/admin/login")

        else if (error.response.status == 406)
            redirect(303, "/error")

        else
            redirect(303, '/error')
    }

    return {}
}