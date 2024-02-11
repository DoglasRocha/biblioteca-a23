import axios from "axios"

export const api = axios.create({
    baseURL: "http://localhost:8080/api/",
    withCredentials: true,
    // headers: {
    //     //"Access-Control-Allow-Origin": "*",
    //     "Access-Control-Allow-Credentials": true,
    //   },
})
