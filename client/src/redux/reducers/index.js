import { combineReducers } from "redux";
import authReducer from "./auth/auth.reducers";

const rootReducer = combineReducers({
    authReducer
}
)
export default rootReducer