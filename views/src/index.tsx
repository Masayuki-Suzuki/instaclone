import 'ress/dist/ress.min.css'
import React from 'react'
import ReactDOM from 'react-dom'
import { applyMiddleware, compose, createStore } from 'redux'
import { Provider } from 'react-redux'
import { library } from '@fortawesome/fontawesome-svg-core'
import { faCheckCircle, faTimesCircle } from '@fortawesome/free-regular-svg-icons'
import { faGoogle } from '@fortawesome/free-brands-svg-icons'
import thunk from 'redux-thunk'
import reducers from './reducers'
import App from './components/App'

/* Use redux devtools */
interface ExWindow extends Window {
    __REDUX_DEVTOOLS_EXTENSION_COMPOSE__?: typeof compose
}
// eslint-disable-next-line
declare var window: ExWindow
const composeReduxDevToolsEnhancers =
    (typeof window === 'object' && window.__REDUX_DEVTOOLS_EXTENSION_COMPOSE__) || compose

const store = createStore(reducers, composeReduxDevToolsEnhancers(applyMiddleware(thunk)))

library.add(faGoogle, faCheckCircle, faTimesCircle)

ReactDOM.render(
    <Provider store={store}>
        <App />
    </Provider>,
    document.getElementById('root')
)
