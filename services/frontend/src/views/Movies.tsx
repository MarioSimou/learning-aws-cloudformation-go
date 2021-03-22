import React from 'react'
import {Box} from '@chakra-ui/react'
import {useAuth0} from '@auth0/auth0-react'
import { useHistory } from 'react-router-dom'

const Movies: React.FC = () => {
    const {isAuthenticated} = useAuth0()
    const history = useHistory()

    React.useEffect(() => {
        if(!isAuthenticated){
            history.push('/')
        }
    },[isAuthenticated, history])


    return (
        <Box>
            movies
        </Box>
    )
}

export default Movies