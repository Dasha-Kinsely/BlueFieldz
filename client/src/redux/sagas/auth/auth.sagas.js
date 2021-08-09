import {all, call, takeLatest} from "redux-saga/effects"

import {authActionTypes, signInFailed, signinSuccess, signoutFailed, signoutSuccess, emailSignupSuccess, emailSignupFailed, checkUserSession} from "../../actions/auth.actions"

/*export function* checkIfAuthenticated() {
    try{
        const
        yield 
    } catch(e) {
        yield
    }
}

export function* checkUserSessionSaga() {
    try{
        yield takeLatest(authActionTypes.CHECK_USER_SESSION, )
    } catch(e) {
        yield
    }
}

export function* () {
    try{
        yield
    } catch(e) {
        yield
    }
}

export function* () {
    try{
        yield
    } catch(e) {
        yield
    }
}

export function* () {
    try{
        yield
    } catch(e) {
        yield
    }
}

export function* () {
    try{
        yield
    } catch(e) {
        yield
    }
}

export function* () {
    try{
        yield
    } catch(e) {
        yield
    }
}

export function* () {
    try{
        yield
    } catch(e) {
        yield
    }
}

export function* () {
    try{
        yield
    } catch(e) {
        yield
    }
}

export function* () {
    try{
        yield
    } catch(e) {
        yield
    }
}

export function* () {
    try{
        yield
    } catch(e) {
        yield
    }
}

export function* () {
    try{
        yield
    } catch(e) {
        yield
    }
}*/

export function* checkUserSessionSaga() {
    try {
        yield null
    } catch (e) {
        yield null
    }
}

export function* placeholder() {
    yield takeLatest(authActionTypes.CHECK_USER_SESSION, checkUserSessionSaga)
}

export function* authSagas() {
    yield all([
        call(placeholder),
    ]);
}