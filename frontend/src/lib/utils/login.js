import axios from "axios"
import { PUBLIC_LOGIN_URL } from "$env/dynamic/public"

export const api = axios.create({
    baseURL: PUBLIC_LOGIN_URL,
    withCredentials: true,
    // headers: {
    //     //"Access-Control-Allow-Origin": "*",
    //     "Access-Control-Allow-Credentials": true,
    //   },
})
