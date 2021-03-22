import React from 'react'
import { Link } from 'react-router-dom'
import {HStack, Text, Link as ChakraLink} from '@chakra-ui/react'
import { useAuth0 } from '@auth0/auth0-react'

const Navbar: React.FC = () => {
    const { loginWithRedirect, isAuthenticated, logout } = useAuth0()
    const onClickLogout = () => logout({returnTo: window.location.origin})

    return (
        <HStack id="navbar" p="1rem 1.5rem" spacing="1.5rem" bg="teal.500" color="white">
            <ChakraLink as={Link} letterSpacing="widest" to="/">Home</ChakraLink>
            {!isAuthenticated && <ChakraLink as={Text} letterSpacing="widest" onClick={loginWithRedirect}>Login</ChakraLink>}
            {isAuthenticated && <ChakraLink as={Text} letterSpacing="widest" onClick={onClickLogout}>Logout</ChakraLink>}
        </HStack>
    )
}

export default Navbar