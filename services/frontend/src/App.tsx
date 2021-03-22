import React from "react"
import {
  ChakraProvider,
  theme,
  CSSReset
} from "@chakra-ui/react"
import Home from './views/Home'
import SignIn from './views/SignIn'
import Movies from './views/Movies'
import {  BrowserRouter, Switch, Route } from 'react-router-dom'
import Navbar from './components/Navbar'
import { Auth0Provider } from '@auth0/auth0-react'

export const App = () => {
  const domain = process.env.REACT_APP_AUTH0_DOMAIN as string
  const clientID = process.env.REACT_APP_CLIENT_ID as string

  return (
    <ChakraProvider theme={theme}>
      <Auth0Provider domain={domain} clientId={clientID} redirectUri={window.location.origin}>
        <CSSReset/>
        <BrowserRouter>
          <Navbar/>
          <Switch>
            <Route path="/" exact component={Home}/>
            <Route path="/sign-in" exact component={SignIn}/>
            <Route path="/movies" exact component={Movies} />
          </Switch>
        </BrowserRouter>
      </Auth0Provider>
    </ChakraProvider>
  )
}
