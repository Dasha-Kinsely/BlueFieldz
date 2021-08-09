import { applyMiddleware, createStore } from "redux";
import createSagaMiddleware from "redux-saga"
import {composeWithDevTools} from "redux-devtools-extension"
import {persistStore, persistReducer } from "redux-persist"
import storage from "redux-persist/lib/storage"
import {rootSaga} from "../sagas/index"
import rootReducer from "../reducers/index"


export default function configStore() {
    const sagaMiddleware = createSagaMiddleware()
    const middlewares = [sagaMiddleware]
    let initialState = {}
    if ('authedUser' in localStorage){
        try{
            initialState['authedUser'] = localStorage['authedUser']
        } catch(e){
            initialState['authedUser'] = null
        }
    }
    const persistConfig = {
        key: 'root',
        storage,
    }
    const store = createStore(
        persistReducer(persistConfig, rootReducer),
        rootReducer,
        initialState,
        composeWithDevTools(applyMiddleware(...middlewares))
    )
    const persistor = persistStore(store)
    sagaMiddleware.run(rootSaga)
    return {store, persistor}
}