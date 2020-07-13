import React from 'react'
import { BrowserRouter, Route } from 'react-router-dom'
import SignIn from '~/components/pages/auth/SignIn'
import SignUp from '~/components/pages/auth/SignUp'

const App = (): JSX.Element => {
    return (
        <BrowserRouter>
            <Route path="/" exact component={SignIn} />
            <Route path="/accounts/signup" component={SignUp} />
        </BrowserRouter>
    )
}

export default App
