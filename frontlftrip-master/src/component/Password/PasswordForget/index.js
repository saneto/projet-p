import React, { Component } from 'react';
import { Link } from 'react-router-dom';
import Container from 'react-bootstrap/Container';
import Form from 'react-bootstrap/Form';
import Button from 'react-bootstrap/Button';
//import { SignUpLink } from '../../Sign/SignUp'; /*NOT USED *
import { withBackaccessContext } from '../../BackEnd';
import * as ROUTES from '../../../const/routes';

const PasswordForgetPage = () => (
	<section id="user-profile">
		<div className="container">
			<div className="row">
					<PasswordForgetForm />
			</div>
		</div>
	</section>
);

const INITIAL_STATE = {
	email: '',
	error: null,
};

class PasswordForgetFormBase extends Component {
	constructor(props) {
		super(props);

		this.state = { ...INITIAL_STATE };
	}

	onSubmit = event => {
		const { email } = this.state;

		this.props.backaccess
			.doPasswordReset({email : email})
			.then(() => {
				this.setState({ ...INITIAL_STATE });
			})
			.catch(error => {
				this.setState({ error });
			});

		event.preventDefault();
	};

	onChange = event => {
		this.setState({ [event.target.name]: event.target.value });
	};

	render() {
		const { email, error } = this.state;
		const isInvalid = email === '';
		return (
			<Container fluid className="view-fixed">
				<div className="d-flex padding-top-10 padding-bottom-10 justify-content-center align-items-center">
					<Form onSubmit={this.onSubmit}>
					<h1>Mot de passe oublié	?</h1>
					<p>Pour réinitialiser ton mot de passe, entre ton adresse email.</p>
						<Form.Group className="marginb1">
							<Form.Label>Email</Form.Label>
							<Form.Control required type="email" name="email" value={email} onChange={this.onChange} className="form-control form-sign" id="email" aria-describedby="emailHelp" placeholder="Votre email"/>
						</Form.Group>
						<div className="text-center">
                            <Button variant="primary" type="submit" disabled={isInvalid}>Valider</Button>
                        </div>
						{error && <p>{error.message}</p>}
					</Form>
				</div>
			</Container>
		);
	}
}

const PasswordForgetLink = () => (
	<p>
		<Link to={ROUTES.PASSWORD_FORGET}>Mot de passe oublié ?</Link>
	</p> 
);

export default PasswordForgetPage;

const PasswordForgetForm = withBackaccessContext(PasswordForgetFormBase);

export { PasswordForgetForm, PasswordForgetLink };