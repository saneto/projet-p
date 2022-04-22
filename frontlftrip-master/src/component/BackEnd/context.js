import React from 'react';

const BackaccessContext  = React.createContext (null);

export const withBackaccessContext = Component  => props => (
        <BackaccessContext.Consumer>
            {backaccess => <Component {...props} backaccess = {backaccess} />}
        </BackaccessContext.Consumer>
);

export default BackaccessContext;