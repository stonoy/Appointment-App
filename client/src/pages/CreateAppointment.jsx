import { redirect } from "react-router-dom"
import { customFetch } from "../utils"
import { toast } from "react-toastify"


export const action = (store) => async({ params}) => {
    const {id} = params
    const token = store.getState().user.token

    try {
        const resp = await customFetch.post(`/createappointment`,{id}, {
            headers : {
                'Authorization': `Bearer ${token}`
            }
        })
        toast.success(resp?.data?.msg)
    } catch (error) {
        const errMsg = error?.response?.data?.msg || "Error in creating appointment"
        if (error?.response?.status == '401' || error?.response?.status == '403'){
            toast.error("Login to continue")
            return redirect("/")
        }

        toast.error(errMsg)
        
    }
    return redirect("/")
}