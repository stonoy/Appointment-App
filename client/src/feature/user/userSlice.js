import {createSlice} from '@reduxjs/toolkit'

const initialState = {
    token: "",
    user: {
        name : "Guest",
        role : "",
    }
}

const userSlice = createSlice({
    name: "user",
    initialState: JSON.parse(localStorage.getItem("user")) || initialState,
    reducers: {
        setUser : (state, {payload : {token, user : {role, name}}}) => {
            state.token = token
            state.user.name = name
            state.user.role = role
            localStorage.setItem("user", JSON.stringify(state))
        },
        logout : () => {
            localStorage.removeItem("user")
            return initialState
        }
    }
})

export const {setUser, logout} = userSlice.actions

export default userSlice.reducer