import { Provider } from "react-redux"
import { PersistGate} from "redux-persist/integration/react"
import {BrowserRouter} from "react-router-dom"
import configStore from "./redux/setup/index"
import AppContainer from "./AppContainer"


function App() {
  const {store, persistor} = configStore()
  return (
    <Provider store={store}>
      <BrowserRouter>
        <PersistGate loading={null} persistor={persistor}>
          <AppContainer/>
        </PersistGate>
      </BrowserRouter>
    </Provider>
  );
}

export default App;
