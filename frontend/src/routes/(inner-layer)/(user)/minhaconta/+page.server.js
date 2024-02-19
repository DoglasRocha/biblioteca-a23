import { api } from "$lib/utils/api.js"
import { redirect } from "@sveltejs/kit"

export async function load({cookies}) {

    let accessToken = cookies.get("accessToken")
    try {
        let request = await api.get("/minhaconta", {
            headers: {
                Cookie: `accessToken=${accessToken}`
            }
        })
        let user = request.data.User, reader = request.data
        return {
            userData: {
                name: user.name,
                surname: user.surname,
                email: user.email,
                birthday: reader.birthday.split("T")[0],
                phone_number: reader.phone_number,
                address: reader.address
            },
            error: null
        }
    }
    catch (error) {

        redirect(303, "/error")
    }

    
}