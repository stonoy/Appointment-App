import {createSlice} from '@reduxjs/toolkit'

const initialState = {
    token: "",
    userName: "Guest"
}

const userSlice = createSlice({
    name: "user",
    initialState: localStorage.getItem("user") || initialState,
    reducers: {
        setUser : (state, action) => {
            console.log(action)
        },
    }
})

export const {setUser} = userSlice.actions

export default userSlice.reducer