import { redirect } from "@sveltejs/kit";

export function load({ cookies }) {
    const accessToken = cookies.get("accessToken")

    if (!accessToken)
        redirect(303, "/login")

    return {}
}