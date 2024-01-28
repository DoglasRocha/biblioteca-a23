import { redirect } from "@sveltejs/kit"

export function load({ cookies }) {
    if (!cookies.get("logged"))
        return redirect(307, "/login")

    return {}
}