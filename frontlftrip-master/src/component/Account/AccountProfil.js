import React from 'react';
import Card from 'react-bootstrap/Card';
import Container from 'react-bootstrap/Container';
import Row from 'react-bootstrap/Row';
import Col from 'react-bootstrap/Col';
import Form from 'react-bootstrap/Form';
import Button from 'react-bootstrap/Button';
import { Typeahead } from 'react-bootstrap-typeahead';
import {compose} from "recompose";
import {withRouter} from "react-router-dom";
import {withBackaccessContext} from "../BackEnd";
//import FileBase64 from 'react-file-base64';
import "bootstrap/dist/css/bootstrap.min.css";

import options from '../Trip/data';

const AccountProfil = () => (
    <AccountProfilComposed />
);

const INITIAL_STATE = {
	user_id : '',
	email : '',
	token : '',
	firstname: '',
	lastname: '',
	sexe : '',
	city : '',
	date_of_birth : '',
	phone_number : '',
	airport : '',
	description : '',
	is_guide : '',
	country_guide : '',
	avatar_path:"",
};

class AccountProfilPanel extends React.Component{

	constructor(props) {
        super(props);

        this.state = { ...INITIAL_STATE };
	}

	backGoDateToBootstrapDate = (getdate) => {
		var date = new Date(getdate);
		var year = date.getFullYear();
		var month = date.getMonth()+1;
		var day = date.getDate();

		if (day < 10) {
		day = '0' + day;
		}
		if (month < 10) {
		month = '0' + month;
		}
		
		var formatedDate = year+'-' + month + '-'+day;
		return formatedDate;
	}

	frontBootstrapDateToBackGoDate = (getdate) => {
		var date = new Date(getdate);
		var formatedDate = date.toISOString();
		return formatedDate;
	}

	componentDidMount() {

        if (localStorage.getItem(('authUser'))) {
			const data = JSON.parse(localStorage.getItem('authUser'));
			var date_of_birth = this.backGoDateToBootstrapDate(data.date_of_birth);
			//console.log(data);
            this.setState(
				{
					token : data.token, 
					email: data.email,
					user_id: data.id,
					firstname : data.firstname,
					lastname : data.lastname,
					sexe : data.sexe,
					city : data.city,
					date_of_birth : date_of_birth,
					airport : data.departure_airport,
					phone_number : data.phone_number,
					description : data.description, 
					is_guide : data.is_guide,
					country_guide : data.country_guide,
					avatar_path: data.avatar_path,
				}
			)
        }
    }

	onChange = event => {
        this.setState({ [event.target.name]: event.target.value });
	};

	onChangeCountryGuide = event => {
        this.setState({ country_guide: event[0].label});
    };
    /*getFiles = event =>{
		this.setState({ avatar_path: event[0].base64 })
	}*/

	handleClick = event => {
        
        if (this.state.is_guide ==='false'){
			this.setState({ is_guide: 'true' });
		}
		else if (this.state.is_guide ==='true'){
			this.setState({ is_guide: 'false' });
		}
		else if (this.state.is_guide ===''){
			this.setState({ is_guide: 'true' });
		}
		//console.log(this.state.is_guide);
    };
	
	changeProfilInformation= event => {
		const formatedDate = this.frontBootstrapDateToBackGoDate(this.state.date_of_birth);
		const { user_id, email, token, firstname, lastname, sexe, city ,phone_number ,airport, description, is_guide , country_guide, avatar_path} = this.state;
    	//console.log(this.state)
        this.props.backaccess
            .updateUserAccountInformation(
				{
					id : user_id,
					email: email,
					firstname: firstname,
					lastname:lastname,
					sexe : sexe,
					city : city,
					date_of_birth : formatedDate,
					phone_number : phone_number,
					airport : airport,
					description : description,
					token : token,
					is_guide : is_guide,
					country_guide : country_guide,
					avatar_path: avatar_path
				}
			)

            .then(() => {
				// met à jour le localstorage
				var authUser = JSON.parse(localStorage.authUser);
				authUser.city = city;
				authUser.date_of_birth = formatedDate;
				authUser.phone_number = phone_number;
				authUser.departure_airport = airport;
				authUser.description = description;
				authUser.is_guide = is_guide;
				authUser.country_guide = country_guide;
				authUser.avatar_path = avatar_path;
				localStorage.setItem('authUser',  JSON.stringify(authUser));
				//met à jour le state
				var date_of_birth = this.backGoDateToBootstrapDate(formatedDate);
				this.setState({city : city , date_of_birth : date_of_birth ,phone_number : phone_number ,airport : airport, description : description, is_guide : is_guide, country_guide: country_guide, avatar_path:avatar_path});
            }).catch(error => {
            console.log("not ok");
            this.setState({ error });
        });
        event.preventDefault();
    };


	render(){
		/*let img= "";
		if(this.state.avatar_path){
			img = <img src={this.state.avatar_path} className="rounded-circle avatar" height="100px" width="100px" alt="avatar"/>
		}else{
			img = <img src="/img/Default.png" className="rounded-circle avatar" value={this.state.avatar_path} height="100px" width="100px" alt="avatar"/>
		}*/
        return (
            <React.Fragment>
				
				<blockquote className="blockquote text-center">
					<p className="mb-0 lead">Compléter votre profil afin qu'on puisse savoir quel type de voyageur vous êtes !</p>
					<footer className="blockquote-footer">Modifier les informations de votre profil visible à la communauté ici</footer>
				</blockquote>
                <Card>
					<Card.Body>

						<Container>
							<Row>

								{/* <Col lg="3">
									{img}
									<FileBase64 multiple={true} onDone={this.getFiles} />
								</Col> */}
								<Col lg="12" className="d-flex justify-content-start align-items-start">
									<Container>
										<Form>
											<Row><Form.Label>Ville</Form.Label></Row>
											<Row><Form.Control type="text" name="city" placeholder="Ville" defaultValue={this.state.city} onChange={this.onChange}/></Row>
											<Row><Form.Label>Date de naissance</Form.Label></Row>
											<Row><Form.Control type="date" name="date_of_birth" defaultValue={this.state.date_of_birth} onChange={this.onChange}/></Row>
                                            <Row><Form.Label>Téléphone</Form.Label></Row>
											<Row><Form.Control type="tel" name="phone_number" defaultValue={this.state.phone_number} onChange={this.onChange}/></Row>
                                            <Row><Form.Label>Aeroport</Form.Label></Row>
											<Row><Form.Control type="text" name="airport" defaultValue={this.state.airport} onChange={this.onChange}/></Row>
                                            <Row><Form.Label>Description</Form.Label></Row>
											<Row><Form.Control type="textarea" name="description" defaultValue={this.state.description} onChange={this.onChange}/></Row>
											<Row className="mt-3">
												{this.state.is_guide === "true" && <Form.Check type="checkbox" checked={true} id="guide" onClick={this.handleClick}/>}
												{this.state.is_guide !== "true" && <Form.Check type="checkbox" id="guide" onClick={this.handleClick}/>}
												<Form.Label>Être Guide</Form.Label>
											</Row>
											{this.state.is_guide === "true" && 
											<Row>
												<Typeahead
                                                    {...this.state}
                                                    name="country"
                                                    id="basic-example"
                                                    onChange={this.onChangeCountryGuide}
													options={options}
													placeholder={this.state.country_guide}
                                                    value={this.state.country_guide}
                                                />
											</Row>
											}
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
									<Button variant="success" onClick={this.changeProfilInformation}>Enregistrer</Button>
								</Col>
							</Row>
						</Container>
					</Card.Footer>
				</Card>
				
            </React.Fragment>

        );
    }
}

const AccountProfilComposed = compose(
    withRouter,
    withBackaccessContext,
)(AccountProfilPanel);

export default AccountProfil;
