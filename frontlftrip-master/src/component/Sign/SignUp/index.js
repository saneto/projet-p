import React, { Component } from 'react';
import { Link, withRouter } from 'react-router-dom';
import { compose } from 'recompose';
import { withBackaccessContext } from '../../BackEnd';
import * as ROUTES from '../../../const/routes';

import Container from 'react-bootstrap/Container';
import Form from 'react-bootstrap/Form';
import Button from 'react-bootstrap/Button';

const SignUpPage = () => (
    <section id="user-profile">  
        <SignUpForm />
    </section>
);

const INITIAL_STATE = {
    firstname: '',
    lastname: '',
    email: '',
    password: '',
    passwordTwo: '',
    isAdmin: false,
    formErrors: {firstname: '', lastname: '', email: '', password: '', passwordTwo: ''},
    firstnameValid: false,
    nameValid: false,
    emailValid: false,
    passwordValid: false,
    passwordTwoValid: false,
    formValid: false,
    error: null
};

class SignUpFormBase extends Component {
    constructor(props) {
        super(props);

        this.state = { ...INITIAL_STATE };
    }

    onSubmit = event => {
        const { firstname, lastname, email, password } = this.state;
        this.props.backaccess
            .doCreateUserWithEmailAndPassword({firstname: firstname, lastname: lastname, email : email, password : password})
            .then(() => {
            })
            .catch(error => {
                //this.setState({ ...INITIAL_STATE });
                console.log("not ok");
                this.setState({ error });
            });

        event.preventDefault();
};

    onChange = event => {
        this.setState({ [event.target.name]: event.target.value });
    };

    onChangeCheckbox = event => {
        this.setState({ [event.target.name]: event.target.checked });
    };

    render() {
        const { firstname, email, password, error, lastname, passwordTwo } = this.state;
        const disable = firstname === "" || email === "" || password === "" || passwordTwo === "" || password !== passwordTwo;
        return (
            <Container fluid className="view-fixed slide1">

                <div className="d-flex padding-top-10 padding-bottom-10 justify-content-center align-items-center">
                    <Form className="form" onSubmit={this.onSubmit}>
                    <h1>S'inscrire</h1>
                        <Form.Group className="marginb1">
                            <Form.Label>Prénom</Form.Label>
                            <Form.Control required type="text" minLength= "2" name="firstname" value={firstname} onChange={this.onChange} className="form-control form-sign" id="firstname" aria-describedby="emailHelp" placeholder="Votre prénom" />
                        </Form.Group>
                        <Form.Group className="marginb1">
                            <Form.Label>Nom</Form.Label>
                            <Form.Control required type="text" minLength= "2" name="lastname" value={lastname} onChange={this.onChange} className="form-control form-sign" id="lastname" aria-describedby="emailHelp" placeholder="Votre nom"/>
                        </Form.Group>
                        <Form.Group className="marginb1">
                            <Form.Label>Email</Form.Label>
                            <Form.Control required type="email" name="email" value={email} onChange={this.onChange} className="form-control form-sign" id="email" aria-describedby="emailHelp" placeholder="Votre email"/>
                        </Form.Group>
                        <Form.Group className="marginb1">
                            <Form.Label>Mot de passe</Form.Label>
                            <Form.Control required type="password" minLength= "6" name="password" value={password} onChange={this.onChange} className="form-control form-sign" id="password" placeholder="Votre mot de passe"/>
                        </Form.Group>
                        <Form.Group className="marginb1">
                            <Form.Label>Confirmer votre mot de passe</Form.Label>
                            <Form.Control required type="password" minLength= "6" name="passwordTwo" value={passwordTwo} onChange={this.onChange} className="form-control form-sign" id="password2" placeholder="Confirmer votre mot de passe"/>
                        </Form.Group>
                        <Form.Group className="marginb1">
                            <Form.Control as="textarea" readOnly className='form-sign term-condition'
                            defaultValue="En vous inscrivant, vous acceptez les conditions générales d'utilisation"
                            />
                        </Form.Group>
                        <Form.Group>
                        <Form.Check inline required type="checkbox" label="J'accepte les conditions générales d'utilisation" id="exampleCheck1"/>
                        </Form.Group>
                        
                        <div className="text-center">
                            <Button variant="primary" type="submit" disabled={disable}>S'enregistrer</Button>
                        </div>
                        <br></br>
                        <SignInLink/>
                        {error && <p>{error.message}</p>}
                    </Form>
                </div>
            </Container>
        );
    }
}

    const SignUpLink = () => (
        <p>
            Pas encore de compte ? <Link to={ROUTES.SIGN_UP}>Inscris-toi</Link>
        </p>
    );
    const SignInLink = () => (
        <p>
            Déjà inscrit ? <Link to={ROUTES.SIGN_IN}>Se connecter</Link>
        </p>
    );
    

const SignUpForm = compose(
    withRouter,
    withBackaccessContext,
)(SignUpFormBase);

export default SignUpPage;

export { SignUpForm, SignUpLink, SignInLink };