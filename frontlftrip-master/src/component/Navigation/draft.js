/*
function Navigation () {

	

    return(
<div >
    <Navbar sticky="top" collapseOnSelect expand="lg" variant="dark">
        <Container>
        <Navbar.Brand href="/"><img className ="logo" src="/img/logo.png" alt="logo"/></Navbar.Brand>
        <Navbar.Toggle aria-controls="responsive-navbar-nav" />
        <Navbar.Collapse id="responsive-navbar-nav">

            <Nav className="mr-auto">
                <Nav.Link href="/">Home</Nav.Link>
                <Nav.Link href="/trips">Voyages</Nav.Link>
                <Nav.Link href="/usertrips">My Trips</Nav.Link>
                <Nav.Link href="/tripsreviews">Reviews</Nav.Link>
				
            </Nav>

            <Nav>
                <Nav.Item>
                    <a href="https://www.facebook.com/" className="nav-link" target="_blank" rel="noopener noreferrer">
                        <FontAwesomeIcon icon={faFacebook} />
                    </a>
                </Nav.Item>
                <Nav.Item>
                    <a href="https://twitter.com/" className="nav-link" target="_blank" rel="noopener noreferrer">
                        <FontAwesomeIcon icon={faTwitter} />
                    </a>
                </Nav.Item>
                <Nav.Item>
                    <a href="https://instagram.com/" className="nav-link" target="_blank" rel="noopener noreferrer">
                        <FontAwesomeIcon icon={faInstagram} />
                    </a>
                </Nav.Item>
                <Nav.Item>
                    <Button href="/register" variant="dark" className="nav-link border border-light rounded black-hover btnnav">Register</Button>
                </Nav.Item>
                <Nav.Item onClick={() => setModalShow(true)}>
                    <Button variant="dark" className="nav-link border border-light rounded black-hover btnnav">Login</Button>
                </Nav.Item>
                <NavDropdown title="Elisa Delawara" id="basic-nav-dropdown">
                    <NavDropdown.Item href="/account">Paramètres</NavDropdown.Item>
                    <NavDropdown.Divider />
                    <NavDropdown.Item href="/logout">Logout</NavDropdown.Item>
                </NavDropdown>
            </Nav>
        </Navbar.Collapse>
        </Container>
    </Navbar>
    
    <MyVerticallyCenteredModal
        show={modalShow}
        onHide={() => setModalShow(false)}
    />
    </div>
    );

}
*/