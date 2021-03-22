import React from 'react'
import {Box, Link as ChakraLink} from '@chakra-ui/react'
import {Link} from 'react-router-dom'
import {useAuth0} from '@auth0/auth0-react'

const Home: React.FC = () => {
    const { isAuthenticated } = useAuth0()

    return (
        <Box p="1rem">
            {isAuthenticated && <ChakraLink as={Link} bg="teal.500" padding="0.5rem" borderRadius="0.5rem" color="white" to="/movies">See Movies</ChakraLink>}
        </Box>
    )
}

export default Home