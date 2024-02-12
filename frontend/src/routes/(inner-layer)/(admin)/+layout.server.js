import { redirect } from "@sveltejs/kit";
import { api } from "$lib/utils/api.js"

export async function load({ cookies }) {

    const accessToken = cookies.get("accessToken")

    console.log(accessToken)
    try {
        const request = await api.post("admin/check", {}, {
            withCredentials: true,
            headers: {
                Cookie: `accessToken=${cookies.get("accessToken")}`,
            }
        })
        console.log(request)

    } catch (error) {
        //console.log(error)
    }

    if (!accessToken)
        redirect(303, "/login")

    return {}
}