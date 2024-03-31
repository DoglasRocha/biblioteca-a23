import axios from "axios"
import { env } from "$env/dynamic/public"

export const api = axios.create({
    baseURL: env.PUBLIC_LOGIN_URL,
    withCredentials: true,
    headers: {
        "Access-Control-Allow-Origin": "*",
        //"Access-Control-Allow-Credentials": true,
      },
})
