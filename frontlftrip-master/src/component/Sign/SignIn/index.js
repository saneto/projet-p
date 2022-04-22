import React, { Component } from 'react';
import { withRouter } from 'react-router-dom';
import { compose } from 'recompose';
import Container from 'react-bootstrap/Container';
import Form from 'react-bootstrap/Form';
import Button from 'react-bootstrap/Button';

import { SignUpLink } from '../SignUp';
import { withBackaccessContext } from '../../BackEnd';
import { PasswordForgetLink } from '../../Password/PasswordForget';

	const SignInPage = () => (
			<div >
				<SignInForm />
			</div>
	);

const INITIAL_STATE = {
	email: '',
	password: '',
	error: null,
};

class SignInFormBase extends Component {
	constructor(props) {
		super(props);

		this.state = { ...INITIAL_STATE };
	}

	onSubmit = event => {
		const { email, password } = this.state;
		this.props.backaccess
			.dosignWithEmailAndPassword({email : email, password : password})
			.then(() => {
			}).catch(error => {
				this.setState({ ...INITIAL_STATE });
				console.log("not ok");
				this.setState({ error });
			});
		event.preventDefault();
	};

	onChange = event => {
		this.setState({ [event.target.name]: event.target.value });
	};

	render() {
		const { email, password, error } = this.state;
		//const isInvalid = password === '' || email === ''; /* CLEAR WARNING */
		return (
			<Container fluid className="view-fixed slide1 covervh">
					<div className="d-flex padding-top-10 padding-bottom-10 justify-content-center align-items-center">
						<Form className="form" onSubmit={this.onSubmit}>
						<h1>Se connecter</h1>
							<Form.Group className="marginb1">
								<Form.Label>Email</Form.Label>
								<Form.Control required type="email" name="email" value={email} onChange={this.onChange} className="form-control form-sign" id="email" aria-describedby="emailHelp" placeholder="Votre email"/>
							</Form.Group>
							<Form.Group className="marginb1">
								<Form.Label>Mot de passe</Form.Label>
								<Form.Control required type="password" minLength= "6" name="password" value={password} onChange={this.onChange} className="form-control form-sign" id="password" placeholder="Votre mot de passe"/>
							</Form.Group>
							<div className="text-center">
								<Button variant="primary" type="submit">Se connecter</Button>
							</div>
							<br></br>
							<SignUpLink />
							<PasswordForgetLink />
							{error && <p>{error.message}</p>}
						</Form>
					</div>
			</Container>
		);
	}
}

const SignInForm = compose(
	withRouter,
	withBackaccessContext,
)(SignInFormBase);

export default SignInPage;

export { SignInForm };