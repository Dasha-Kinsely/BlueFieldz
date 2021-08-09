export const authActionTypes = {
    CHECK_USER_SESSION: "CHECK_USER_SESSION",
    EMAIL_SIGNIN: "EMAIL_SIGNIN",
    GOOGLE_SIGNIN: "GOOGLE_SIGNIN",
    SIGNIN_SUCCESS: "SIGNIN_SUCCESS",
    SIGNIN_FAILED: "SIGNIN_FAILED",
    EMAIL_SIGNUP: "EMAIL_SIGNUP",
    EMAIL_SIGNUP_SUCCESS: "EMAIL_SIGNUP_SUCCESS",
    EMAIL_SIGNUP_FAILED: "EMAIL_SIGNUP_FAILED",
    EMAIL_VERIFICATION: "EMAIL_VERIFICATION",
    EMAIL_VERIFICATION_SUCCESS: "EMAIL_VERIFICATION_SUCCESS",
    SIGNOUT: "SIGNOUT",
    SIGNOUT_SUCCESS: "SIGNOUT_SUCCESS",
    SIGNOUT_FAILED: "SIGNOUT_FAILED",
    CHANGE_PASSWORD: "CHANGE_PASSWORD",
    CHANGE_PASSWORD_FAILED: "CHANGE_PASSWORD_FAILED"
}

export const checkUserSession = () => ({
    type: authActionTypes.CHECK_USER_SESSION
})

export const emailSignin = (p) => ({
    type: authActionTypes.EMAIL_SIGNIN,
    payload: p
})

export const googleSignin = () => ({
    type: authActionTypes.GOOGLE_SIGNIN
})

export const signinSuccess = (p) => ({
    type: authActionTypes.SIGNIN_SUCCESS,
    payload: p
})

export const signInFailed = (e) => ({
    type: authActionTypes.SIGNIN_FAILED,
    payload: e
})

export const emailSignup = (p) => ({
    type: authActionTypes.EMAIL_SIGNUP,
    payload: p
})

export const emailSignupSuccess = (p) => ({
    type: authActionTypes.EMAIL_SIGNUP_SUCCESS,
    payload: p
})

export const emailSignupFailed = (e) => ({
    type: authActionTypes.EMAIL_SIGNUP_FAILED,
    payload: e
})

export const emailVerification = (p) => ({
    type: authActionTypes.EMAIL_VERIFICATION,
    payload: p
})

export const emailVerificationSuccess = (p) => ({
    type: authActionTypes.EMAIL_VERIFICATION_SUCCESS,
    payload: p
})

export const signout = () => ({
    type: authActionTypes.SIGNOUT
})

export const signoutSuccess = () => ({
    type: authActionTypes.SIGNOUT_SUCCESS
})

export const signoutFailed = (e) => ({
    type: authActionTypes.SIGNOUT_FAILED,
    payload: e
})

export const changePassword = (p) => ({
    type: authActionTypes.CHANGE_PASSWORD,
    payload: p
})

export const changePasswordFailed = (e) => ({
    type: authActionTypes.CHANGE_PASSWORD_FAILED,
    payload: e
})