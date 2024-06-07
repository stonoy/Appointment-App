import axios from 'axios'

const productUrl = 'http://localhost:8080/api/v1'

export const customFetch = axios.create({
    baseURL: productUrl
})

export const links = [
    {id: 1, name: "Home", link: "/"},
    {id: 2, name: "About", link: "/about"},
    {id: 3, name: "Appointments", link: "/appointments"},
    {id: 4, name: "Availability", link: "/availability"},
    {id: 5, name: "Admin", link: "/admin"},
]