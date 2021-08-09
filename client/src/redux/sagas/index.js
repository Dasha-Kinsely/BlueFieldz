import { all, call } from "@redux-saga/core/effects";
import {authSagas} from "./auth/auth.sagas";

export function* rootSaga() {
    yield all([
        call(authSagas)
    ])
}