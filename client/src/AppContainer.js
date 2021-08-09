import {useEffect, React} from "react"
import {Switch, Route, useLocation, Redirect} from "react-router-dom"
import { useDispatch, useSelector } from "react-redux"

import Test from "./pages/Test"

const AppContainer = () => {
    const location = useLocation()
    const dispatch = useDispatch()
    return (
            <div id="app-container">
                <Switch location={location} key={location.pathname}>
                    <Route component={Test} exact path="/test"/>
                </Switch>
            </div>
    )
}

export default AppContainer