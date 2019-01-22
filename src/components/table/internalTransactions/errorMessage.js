import React from 'react';
import Jumbotron from 'react-bootstrap/lib/Jumbotron';

const errorMessage = (props) => {
    return (       
         <Jumbotron>
            <h1>Internal Transaction Status</h1>
            <p>
                Internal Transactions seem to be empty for this request.
            </p>            
        </Jumbotron>
    )
}

export default errorMessage;