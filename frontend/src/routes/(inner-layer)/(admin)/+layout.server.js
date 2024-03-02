import { redirect } from "@sveltejs/kit";
import { api } from "$lib/utils/api.js";

export const config = {
    runtime: 'edge'
}

export async function load({ cookies }) {

    const accessToken = cookies.get("accessToken")

    try {
        await api.post("admin/check", {}, {
            withCredentials: true,
            headers: {
                Cookie: `accessToken=${accessToken}`,
            }
        })
    } catch (error) {
        if (error?.code == 'ECONNREFUSED')
            redirect(303, "/error")

        else if (error.response.status == 403) 
            redirect(303, "/")
        
        else if (error.response.status == 401)
            redirect(303, "/admin/login")

        else if (error.response.status == 406)
            redirect(303, "/error")
    }

    return {}
}