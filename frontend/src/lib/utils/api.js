import axios from "axios"
import { PUBLIC_API_URL } from "$env/static/public"

export const api = axios.create({
    baseURL: PUBLIC_API_URL,
    withCredentials: true,
    // headers: {
    //     //"Access-Control-Allow-Origin": "*",
    //     "Access-Control-Allow-Credentials": true,
    //   },
})
