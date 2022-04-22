import React from 'react';
import NavDropdown from 'react-bootstrap/NavDropdown';


import { withBackaccessContext } from '../../BackEnd';
const SignOutButton = ({ backaccess }) => (
	<NavDropdown.Item onClick={backaccess.onlogOut} >Se déconnecter</NavDropdown.Item>
);

export default withBackaccessContext(SignOutButton);