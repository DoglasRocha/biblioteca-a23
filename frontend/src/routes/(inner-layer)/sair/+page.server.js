import { redirect } from "@sveltejs/kit"

export function load({ cookies }) {
    cookies.delete("accessToken", {
        path: "/"
    })

    redirect(303, "/login")
}