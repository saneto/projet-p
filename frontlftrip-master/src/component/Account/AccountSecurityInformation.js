import React from 'react';
import Typography from '@material-ui/core/Typography';
import Container from '@material-ui/core/Container';
import {compose} from "recompose";
import {withRouter} from "react-router-dom";
import {withBackaccessContext} from "../BackEnd";
import Row from 'react-bootstrap/Row';
import Col from 'react-bootstrap/Col';
import Form from 'react-bootstrap/Form';
import Card from 'react-bootstrap/Card';
import Button from 'react-bootstrap/Button';

const AccountSecurityInformationPage = () => (
    <section id="user-pwd-update">
        <AccountSecurityInformationForm />
    </section>
);


const INITIAL_STATE = {
    user_id : '',
    token : '',
    email: '',
    current_password: '',
    new_password: '',
    retype_password: '',
};


class AccountSecurityInformationBase extends React.Component{
    constructor(props) {
        super(props);

        this.state = { ...INITIAL_STATE };
    }

    componentDidMount() {
        
        if (localStorage.getItem(('authUser'))) {
            const data = JSON.parse(localStorage.getItem('authUser'));
            this.setState({token : data.token, email : data.email, user_id : data.id})
        }
    }


    onSubmit = event => {
        const { email, current_password, new_password, retype_password, token } = this.state;
        //console.log(token)
        this.props.backaccess
            .updateUserPassword({email : email, current_password : current_password, new_password : new_password, retype_password : retype_password, token : token})
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

    deleteAccount = event => {
        const {user_id, token} = this.state;
        this.props.backaccess.deleteUserAccount({id : user_id, token : token})
            .then(() => {

            }).catch(error => {
                this.setState({ ...INITIAL_STATE });
                console.log("not ok");
                this.setState({ error });
            });
        event.preventDefault();
    };

    render(){
        return (
            <Container className="" component="main" maxWidth="xs">
                <div className="">
                    <Typography component="h1" variant="h5">
                        Changer mon mot de passe
                    </Typography>
                </div>
                <Card className="pt-3">
                <Card.Body>
                <Container>
                    <Row>
                        <Col lg="9" className="d-flex justify-content-start align-items-start">
                            <Container>
                                <Form>
                                    <Row><Form.Label>Email</Form.Label></Row>
                                    <Row><Form.Control type="email" value={this.email} name="email" onChange={this.onChange} placeholder="name@example.com" /></Row>
                                    <Row>Mot de passe actuel</Row>
                                    <Row><Form.Control type="password" value={this.current_password} name="current_password" onChange={this.onChange}/></Row>
                                    <Row>Nouveau mot de passe</Row>
                                    <Row><Form.Control type="password" value={this.new_password} name="new_password" onChange={this.onChange}/></Row>
                                    <Row>Confirmer le mot de passe</Row>
                                    <Row><Form.Control type="password" value={this.retype_password} name="retype_password" onChange={this.onChange}/></Row>
                                </Form>
                            </Container>
                        </Col>

                    </Row>
                </Container>
                </Card.Body>
                <Card.Footer>
                <Container>
                    <Row>
                        <Col className="d-flex justify-content-end align-items-end">
                            {/* <Button variant="secondary" onClick={this.quitPasswordMode}>Annuler</Button> */}
                            <Button variant="success" onClick={this.onSubmit}>Modifier mon mot de passe</Button>
                        </Col>
                    </Row>
                </Container>
                </Card.Footer>
            </Card>
            <Button size="small" type="submit" variant="danger" onClick={this.deleteAccount}>
                    Supprimer mon compte
            </Button>
            </Container>
        );
    }
}
const AccountSecurityInformationForm = compose(
    withRouter,
    withBackaccessContext,
)(AccountSecurityInformationBase);

export default AccountSecurityInformationPage;
