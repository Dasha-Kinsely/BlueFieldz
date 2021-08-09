import { applyMiddleware, createStore } from "redux";
import logger from "redux-logger"
import createSagaMiddleware from "redux-saga"
import {composeWithDevTools} from "redux-devtools-extension"
import {persistStore, persistReducer } from "redux-persist"
import storage from "redux-persist/lib/storage"
import {rootSaga} from "../sagas/index"
import rootReducer from "../reducers/index"


export default function configStore() {
    const sagaMiddleware = createSagaMiddleware()
    // Dev only
    const middlewares = [sagaMiddleware, logger]
    let initialState = {}
    //-------------------------------------------
    const persistConfig = {
        key: 'root',
        storage,
        whitelist: []
    }
    const store = createStore(
        persistReducer(persistConfig, rootReducer),
        initialState,
        composeWithDevTools(applyMiddleware(...middlewares))
    )
    const persistor = persistStore(store)
    sagaMiddleware.run(rootSaga)
    return {store, persistor}
}