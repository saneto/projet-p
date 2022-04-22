import React from 'react';
import Container from 'react-bootstrap/Container';
import Navbar from 'react-bootstrap/Navbar';
import Nav from 'react-bootstrap/Nav';
import NavDropdown from 'react-bootstrap/NavDropdown';
import Button from 'react-bootstrap/Button';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faTwitter, faFacebook, faInstagram } from '@fortawesome/free-brands-svg-icons';
import { AuthUserContext } from '../Session';
import * as ROUTES from '../../const/routes';
import SignOutButton from '../Sign/SignOut';
import Notification from '../Notification';
const Navigation = () => (
	<AuthUserContext.Consumer>
		{authUser => authUser ? (
				<NavigationAuth />
			) : (
				<NavigationNonAuth />
			)
		}
	</AuthUserContext.Consumer>
	);


	const NavigationNonAuth = () =>{
		return (
			<div >
				<Navbar sticky="top" collapseOnSelect expand="lg" variant="dark">
					<Container>
					<Navbar.Brand href="/"><img className ="logo" src="/img/logo.png" alt="logo"/></Navbar.Brand>
					<Navbar.Toggle aria-controls="responsive-navbar-nav" />
					<Navbar.Collapse className="px-5" id="responsive-navbar-nav">
	
						<Nav className="mr-auto">
							<Nav.Link href="/">Accueil</Nav.Link>
							<Nav.Link href="/trips">Voyages</Nav.Link>
							{/* <Nav.Link href="/tripsreviews">Témoignages</Nav.Link> */}
							<Nav.Link href="/concept">Concept & Valeur</Nav.Link>
							
						</Nav>
	
						<Nav>
							<Nav.Item>
								<a href="https://www.facebook.com/lftrip/" className="nav-link" target="_blank" rel="noopener noreferrer">
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
								<Button href="/register" variant="dark" className="nav-link border border-light rounded black-hover btnnav">Inscription</Button>
							</Nav.Item>
							<Nav.Item>
								<Button href="/login" variant="dark" className="nav-link border border-light rounded black-hover btnnav">Connexion</Button>
							</Nav.Item>
						</Nav>
					</Navbar.Collapse>
					</Container>
				</Navbar>
		</div>
		
	)};

const NavigationAuth  = () => {
	let user = JSON.parse(localStorage.getItem("authUser"));
	let username = user.firstname + ' ' + user.lastname;
	return (
		<div >
			<Navbar sticky="top" collapseOnSelect expand="lg" variant="dark">
				<Container>
				<Navbar.Brand href="/"><img className ="logo" src="/img/logo.png" alt="logo"/></Navbar.Brand>
				<Navbar.Toggle aria-controls="responsive-navbar-nav" />
				<Navbar.Collapse id="responsive-navbar-nav">
				
					<Nav className="mr-auto">
						<Nav.Link href="/">Accueil</Nav.Link>
						<Nav.Link href="/trips">Voyages</Nav.Link>
						<Nav.Link href="/concept">Concept & Valeur</Nav.Link>
						<Nav.Link href={ROUTES.MESSAGERIE}>Discussions</Nav.Link>


					</Nav>

					<Nav>
						<Nav.Item>
							<a href="https://www.facebook.com/lftrip/" className="nav-link" target="_blank" rel="noopener noreferrer">
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
						
						<Notification />
						<NavDropdown title={username} id="basic-nav-dropdown">
							<NavDropdown.Item href="/account">Mon compte</NavDropdown.Item>
							<NavDropdown.Item href="/profilsummary">Mon profil</NavDropdown.Item>
							<NavDropdown.Item href="/usertrips">Mes voyages créés</NavDropdown.Item>
							<NavDropdown.Item href="/tripsliked">Mes favoris</NavDropdown.Item>
							<NavDropdown.Item href="/tripsparticipated">Voyages réservés</NavDropdown.Item>
							<NavDropdown.Item href={ROUTES.MESSAGERIE}>Mes discussions</NavDropdown.Item>

							<NavDropdown.Divider />
							<SignOutButton />
							
						</NavDropdown>
					</Nav>
				</Navbar.Collapse>
				</Container>
			</Navbar>


    	</div>

)
};


export default Navigation;
