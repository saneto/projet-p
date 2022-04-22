import React from 'react';
import Card from 'react-bootstrap/Card';
import Container from 'react-bootstrap/Container';
import Row from 'react-bootstrap/Row';
import Col from 'react-bootstrap/Col';
import Form from 'react-bootstrap/Form';
import Button from 'react-bootstrap/Button';
import {compose} from "recompose";
import {withRouter} from "react-router-dom";
import {withBackaccessContext} from "../BackEnd";
import Modal from "react-bootstrap/Modal";
import FileBase64 from 'react-file-base64';
import "bootstrap/dist/css/bootstrap.min.css";
import { withAuthorization } from '../Session';

const AccountInformation = () => (
    <AccountInformationComposed />
);

const INITIAL_STATE = {
	user_id : '',
	token : '',
	email: '',
	firstname: '',
	lastname: '',
	city : '',
	sexe : '',
	date_of_birth : '',
	phone_number : '',
	airport : '',
	description : '',
	is_guide : '',
    current_password: '',
    new_password: '',
	retype_password: '',
	avatar_path:"",
	openDeleteModal: false,
};

class AccountInformationPanel extends React.Component{

	constructor(props) {
        super(props);

		this.state = { ...INITIAL_STATE };

	}

	componentDidMount() {

        if (localStorage.getItem(('authUser'))) {
			const data = JSON.parse(localStorage.getItem('authUser'));
            this.setState(
				{
					token : data.token,
					email: data.email,
					user_id: data.id, 
					firstname : data.firstname,
					lastname : data.lastname,
					sexe : data.sexe,
					city : data.city,
					date_of_birth : data.date_of_birth,
					airport : data.departure_airport,
					phone_number : data.phone_number,
					description : data.description,
					is_guide : data.is_guide,
					avatar_path:data.avatar_path,
				}
			)

		}
    }

    openModal = event => {
		this.setState({openDeleteModal: true});
	};
	handleClose = event => {
		this.setState({openDeleteModal: false});
	};
	getFiles = event =>{
		this.setState({ avatar_path: event[0].base64 })
	}

	changePassword = event => {
        const { email, current_password, new_password, retype_password, token } = this.state;
        //console.log(token)
        this.props.backaccess
            .updateUserPassword({email : email, current_password : current_password, new_password : new_password, retype_password : retype_password, token : token})
            .then(() => {
				this.quitPasswordMode();
            }).catch(error => {
            this.setState({ ...INITIAL_STATE });
            console.log("not ok");
            this.setState({ error });
        });
        event.preventDefault();
	};

	changeAccountInformation= event => {

        const { user_id, email, firstname, lastname, sexe, token, city ,phone_number ,airport, description, date_of_birth, avatar_path, is_guide } = this.state;
        //console.log(token)
        this.props.backaccess
            .updateUserAccountInformation({id : user_id, email : email, firstname: firstname, lastname: lastname, token : token, sexe: sexe, city : city , date_of_birth : date_of_birth ,phone_number : phone_number ,airport : airport, description : description, avatar_path:avatar_path, is_guide : is_guide})
            .then(() => {
				// met à jour le localstorage
				var authUser = JSON.parse(localStorage.authUser);
				authUser.email = email;
				authUser.firstname = firstname;
				authUser.lastname = lastname;
				authUser.avatar_path = avatar_path;
				localStorage.setItem('authUser',  JSON.stringify(authUser));
				//met à jour le state
				this.setState({email: email, firstname:firstname, lastname:lastname, avatar_path:avatar_path});
				this.quitInformationMode();
            }).catch(error => {
            console.log("not ok");
            this.setState({ error });
        });
        event.preventDefault();
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

	onChange = event => {
        this.setState({ [event.target.name]: event.target.value });
    };

	// change DOM struct to get input field to edit user information
    editInformationMode = () => {
        this.switchInputable("edit-mode", "read-only-mode")
	}

	// change DOM struct to get input field to edit user information
    editPasswordMode = () => {
        this.switchInputable("password-mode", "edit-mode")
	}

	// change DOM struct to quit editor mode (need to reinitialize state value)
    quitInformationMode = () => {
		this.switchInputable("read-only-mode", "edit-mode")
	}

	// change DOM struct to quit password mode (need to reinitialize state value)
    quitPasswordMode = () => {
		this.switchInputable("read-only-mode", "password-mode")
	}

	// param : on is class name to display, off is class name to hidechange Dom to display and hide
    // change DOM struct to hide and display different inputs
    switchInputable = (on, off) => {
        var toHide = document.getElementsByClassName(off);
        for (var i = 0; i < toHide.length; i ++) {
            toHide[i].style.display = 'none';
        }
        var toShow = document.getElementsByClassName(on);
        for (var j = 0; j < toShow.length; j ++) {
            toShow[j].style.display = 'block';
        }
    }

	render(){
		let img= "";
		if(this.state.avatar_path){
			img = <img src={this.state.avatar_path} className="rounded-circle avatar" height="100px" width="100px" alt="avatar"/>
		}else{
			img = <img src="/img/Default.png" className="rounded-circle avatar" value={this.state.avatar_path} height="100px" width="100px" alt="avatar"/>
		}
        return (
            <React.Fragment>
				<blockquote className="blockquote text-center">
					<p className="mb-0 lead">Besoin de modifier des informations liées à votre compte ?</p>
					<footer className="blockquote-footer">Modifier votre email, nom, prénom ou/et votre mot de passe ci-dessous</footer>
				</blockquote>
                <Card className="pt-3 mb-3">
					<Card.Body>

						<Container className="read-only-mode">
							<Row>
								<Col lg="3">
									{img}
								</Col>
								<Col lg="7" className="d-flex justify-content-start align-items-start">
									<Container>
										<Row><b>{this.state.firstname} {this.state.lastname}</b></Row><br/>
										<Row>{this.state.email}</Row>
									</Container>
								</Col>
								<Col lg="2" className="d-flex justify-content-center align-items-center justify-content-md-end align-items-md-start">
									<Button className="icon fa fa-cog" onClick={this.editInformationMode} variant="primary">Modifier</Button>
								</Col>
							</Row>
						</Container>

						<Container className="edit-mode">
							<Row>
								<Col lg="3">
									{img}
									<FileBase64 multiple={true} onDone={this.getFiles} />
								</Col>
								<Col lg="9" className="d-flex justify-content-start align-items-start">
									<Container>
										<Form>
											<Row><Form.Label>Email</Form.Label></Row>
											<Row><Form.Control type="email" name="email" placeholder="name@example.com" defaultValue={this.state.email} onChange={this.onChange}/></Row>
											<Row><Form.Label>Nom</Form.Label></Row>
											<Row><Form.Control type="text" name="lastname" placeholder="Nom" defaultValue={this.state.lastname} onChange={this.onChange}/></Row>
											<Row><Form.Label>Prénom</Form.Label></Row>
											<Row><Form.Control type="text" name="firstname" placeholder="Prénom" defaultValue={this.state.firstname} onChange={this.onChange}/></Row><br/>
											<Row><Button variant="link" onClick={this.editPasswordMode}>Vous souhaitez modifier votre mot de passe ?</Button></Row>
										</Form>
									</Container>
								</Col>
							</Row>
						</Container>

						<Container className="password-mode">
							<Row>
								<Col lg="3"><img src="https://mdbootstrap.com/img/Photos/Avatars/avatar-8.jpg" className="rounded-circle avatar" height="100px" width="100px" alt="avatar"/></Col>
								<Col lg="9" className="d-flex justify-content-start align-items-start">
									<Container>
										<Form>
											<Row><Form.Label>Email</Form.Label></Row>
											<Row><Form.Control type="email" name="email" onChange={this.onChange} placeholder="name@example.com"/></Row>
											<Row>Mot de passe actuel</Row>
											<Row><Form.Control type="password" name="current_password" onChange={this.onChange}/></Row>
											<Row>Nouveau mot de passe</Row>
											<Row><Form.Control type="password" name="new_password" onChange={this.onChange}/></Row>
											<Row>Confirmer le nouveau mot de passe</Row>
											<Row><Form.Control type="password" name="retype_password" onChange={this.onChange}/></Row>
										</Form>
									</Container>
								</Col>
							</Row>
						</Container>

					</Card.Body>
					<Card.Footer className="edit-mode">
						<Container>
							<Row>
								<Col xs={12} md={4} className="d-flex justify-content-center align-items-center justify-content-md-start align-items-md-start">
									<Button variant="outline-danger" onClick={this.openModal}>Supprimer le compte </Button>
								</Col>
								<Col xs={12} md={8} className="d-flex justify-content-center align-items-center justify-content-md-end align-items-md-end pt-3 pt-sm-0">
									<Button variant="secondary" className="mr-3" onClick={this.quitInformationMode}>Annuler</Button>
									<Button variant="success" onClick={this.changeAccountInformation}>Enregistrer</Button>
								</Col>
							</Row>
						</Container>
					</Card.Footer>
					<Card.Footer className="password-mode">
						<Container>
							<Row>
								<Col className="d-flex justify-content-end align-items-end">
									<Button variant="secondary" className="mr-3" onClick={this.quitPasswordMode}>Annuler</Button>
									<Button variant="success" onClick={this.changePassword}>Enregistrer</Button>
								</Col>
							</Row>
						</Container>
					</Card.Footer>
				</Card>

				<Card className="read-only-mode">
					<Card.Footer>
						<Container>
							<Row>
								<Col className="d-flex justify-content-start align-items-center">
									La suppression de votre compte sera irréversible
								</Col>
								<Col className="d-flex justify-content-end align-items-center">
									<Button variant="outline-danger" onClick={this.openModal}>Supprimer le compte </Button>
								</Col>
							</Row>
						</Container>
					</Card.Footer>
				</Card>

				<Modal show={this.state.openDeleteModal} onHide={this.handleClose} animation={false}>
					<Modal.Header closeButton>
						<Modal.Title>Demande de confirmation</Modal.Title>
					</Modal.Header>
					<Modal.Body>Attention vous êtes sur le point de supprimez définitivement votre compte et toutes les informations relatives à celui-ci.</Modal.Body>
					<Modal.Footer>
						<Button variant="secondary" onClick={this.handleClose}>
							Annuler
						</Button>
						<Button variant="danger" onClick={this.deleteAccount}>
							Confirmer
						</Button>
					</Modal.Footer>
				</Modal>
            </React.Fragment>
        );
    }
}

const condition = authUser => !!authUser;

const AccountInformationComposed = compose(
	withAuthorization(condition),
    withRouter,
    withBackaccessContext,
)(AccountInformationPanel);

export default AccountInformation;
