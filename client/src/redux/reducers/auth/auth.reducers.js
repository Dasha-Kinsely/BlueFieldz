import { authActionTypes } from "../../actions/auth.actions";

const initialiState = {
    currentUser: null,
    error: null,
    loading: false
}

// only handle synced actions
const authReducer = (state=initialiState, action) => {
    switch(action.type) {
        case authActionTypes.EMAIL_SIGNUP:
            return {
                ...state,
                loading: true
            }
        case authActionTypes.SIGNIN_SUCCESS:
            return {
                ...state,
                currentUser: action.payload,
                loading: false,
                error: null
            }
        case authActionTypes.SIGNOUT_SUCCESS:
            return {
                ...state,
                currentUser: null,
                loading: false,
                error: null
            }
        case authActionTypes.SIGNOUT_FAILED:
            return {
                ...state,
                loading: false,
                error: action.apyload
            }
        default:
            return state
    }
}

export default authReducer