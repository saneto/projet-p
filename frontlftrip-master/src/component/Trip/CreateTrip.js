import React, {Component} from 'react';
import {Row, Container, Button, Col, Form} from 'react-bootstrap';
import Card from 'react-bootstrap/Card';
import { withRouter } from 'react-router-dom';
import { compose } from 'recompose';
import { withBackaccessContext } from '../BackEnd';
import FileBase64 from 'react-file-base64';
//import  { post } from 'axios';
import { Typeahead } from 'react-bootstrap-typeahead';
import 'react-bootstrap-typeahead/css/Typeahead.css';
import options from './data';
import latlong from './latlong';
import { withAuthorization } from '../Session';
import {store} from "react-notifications-component";
import history from "../../history";

class Createtrip extends Component{
    constructor(props){
        super(props);
        this.state = {
            uri : "trips",
            country: "",
            title: "",
            nb_days : 0,
            description: "",
            start_date: new Date().now,
            end_date: new Date().now,
            middle_age: 0,
            nb_traveler: 0,
            Budget: 0,
            image_program: "",
            program: "",
            base64:"",
            is_active: true,
            trip_type : "",
            latitude:0,
            longitude:0,
            lodging : ""
        }
    }

    checkDateValidation(startDate, endDate) {
        // check the dates
        if ((new Date(startDate) > new Date(endDate)) || (new Date(endDate) < new Date(startDate))) {
          return false; 
        } else {
          return true;
        }
    }

    onSubmit = () =>{
        let responseDate = this.checkDateValidation(this.state.start_date, this.state.end_date);
        if(responseDate){
            if(this.state.middle_age < 16 || this.state.middle_age > 100){
                store.addNotification({
                    title: "Erreur sur l'âge moyen",
                    message: "L'âge moyen doit être supérieur à 16 ou inférieur à 100",
                    type: 'danger',
                    container: 'top-right',                
                    animationIn: ["animated", "fadeIn"],
                    animationOut: ["animated", "fadeOut"],
                    dismiss: {
                        duration: 5000,
                        onScreen: true
                    }
                })
            } else {
                if(this.state.nb_traveler < 1 || this.state.Budget < 1){
                    store.addNotification({
                        title: "Erreur sur le nombre de voyageurs ou le budget",
                        message: "Le nombre de voyageurs et le budget doit être supérieur à 0 et ne pas être négative",
                        type: 'danger',
                        container: 'top-right',                
                        animationIn: ["animated", "fadeIn"],
                        animationOut: ["animated", "fadeOut"],
                        dismiss: {
                            duration: 5000,
                            onScreen: true
                        }
                    })
                } else {
                    var dateD = Date.parse(this.state.start_date+"T00:00:00.00Z")
                    var dateF= Date.parse(this.state.end_date+"T00:00:00.00Z")
                    var nb_Days = (dateF - dateD) / 86400000;
                    let data = {
                        "country": this.state.country,
                        "title": this.state.title,
                        "nb_days" : nb_Days,
                        "description":this.state.description,
                        "start_date": this.state.start_date+"T00:00:00.00Z",
                        "end_date": this.state.end_date+"T00:00:00.00Z",
                        "middle_age": parseInt(this.state.middle_age),
                        "nb_traveler": parseInt(this.state.nb_traveler) ,
                        "Budget":  parseInt(this.state.Budget),
                        "image_program": this.state.image_program,
                        "Program": this.state.program,
                        "base64" : this.state.base64,
                        "is_active": this.state.is_active,
                        "trip_type" : this.state.trip_type,
                        "lodging" : this.state.lodging,
                        "latitude": JSON.stringify(this.state.latitude),
                        "longitude": JSON.stringify(this.state.longitude)
                    }
                    this.props.backaccess.postData(data, this.state.uri).then( ret =>{
                        /* store.addNotification({
                            title: 'Création de voyage',
                            message: 'Votre voyage a été créer avec succès',
                            type: 'success',                         // 'default', 'success', 'info', 'warning'
                            container: 'top-right',                // where to position the notifications
                            animationIn: ["animated", "fadeIn"],     // animate.css classes that's applied
                            animationOut: ["animated", "fadeOut"],   // animate.css classes that's applied
                            dismiss: {
                                duration: 5000,
                                onScreen: true
                            }
                        }) */
                        history.push(`trip/${ret.data.response.id}`);
                        //console.log(ret)
                    }).catch( error =>{

                    })
                }
            }
        }else{
            store.addNotification({
                title: 'Erreur sur les dates',
                message: 'La date de début doit être inférieur à la date de fin',
                type: 'danger',
                container: 'top-right',                
                animationIn: ["animated", "fadeIn"],
                animationOut: ["animated", "fadeOut"],
                dismiss: {
                    duration: 5000,
                    onScreen: true
                }
            })
        }
    }

    onChange = event => {
		this.setState({ [event.target.name]: event.target.value });
    };
    handleClick = event => {
        var i;
        var checkboxArray = document.getElementsByClassName("checkbox");
        if (this.state.trip_type !==''){
            for(i=0; i<checkboxArray.length; i++) {
                // checkboxArray[i].firstChild.id checkbox qu'on vérifie
                // this.state.trip_type ce qui est actuellement coché
                // event.target.id ce sur quoi on appuie
                if (checkboxArray[i].firstChild.id === this.state.trip_type && checkboxArray[i].firstChild.checked === true){
                    checkboxArray[i].firstChild.checked = false;
                }
                if (checkboxArray[i].firstChild.id !== this.state.trip_type && checkboxArray[i].firstChildid === event.target.id){
                    checkboxArray[i].firstChild.checked = true;
                }
            }
        }

        this.setState({ trip_type: event.target.value });
        
    };
    onChangeCountry = event => {
        if(event[0] !== undefined){
        this.setState({ country: event[0].label});
        
        //console.log(this.state.latitude)
        latlong.forEach( latlongv => {
            if(event[0].alpha2 === latlongv.alpha2){
                this.setState({latitude: latlongv.latitude, longitude: latlongv.longitude});
            }
        })}

    };

    getFiles = event =>{
        //console.log(event[0])
        this.setState({ base64: event[0].base64 })
    }

    render()
    {
        const { country,  title,
                description, start_date, end_date,
                middle_age, nb_traveler, Budget,
                lodging,  program } = this.state;
        return(
            <Container fluid className="covervh create-trip-page">
                    <Row>
                        <Col md={4} className="p-0"><img className="create-banner" src="/img/createtripbanner.jpg" alt="img"/></Col>
                        <Col md={8}>
                            <Container>
                            <Row className="justify-content-md-center">
                            <Col>
                                
                                <Row className="justify-content-center align-items-center">

                                    <Card className="create-trip-card-border mt-4">
                                        <Row className="justify-content-md-center mt-4">
                                            <Col md="auto">
                                            
                                                <div className="section-title line-style no-margin create-trip-card-title pl-5 pr-5">
                                                    <h3 className="title">Créer un voyage qui vous ressemble</h3>
                                                </div>
                                                
                                            </Col>
                                        </Row>
                                        <Card.Body className="create-trip-card">
                                        <Row xs={12}>
                                            <Form.Group as={Col} xs={12} md={8} controlId="formGridEmail">
                                                <Form.Label className="label-title"><strong className="chevron">{">"}</strong> Un titre accrocheur pour votre voyage !</Form.Label>
                                                <Form.Control type="text" name="title" onChange={this.onChange} value ={title} placeholder="Ex : Un roadtrip en Italie !" required/>
                                            </Form.Group>
                                            <Form.Group as={Col} xs={12} md={4} controlId="formGridEmail">
                                                <Form.Label className="label-title"><strong className="chevron">{">"}</strong> Destination</Form.Label>
                                                <Typeahead
                                                    {...this.state}
                                                    name="country"
                                                    id="basic-example"
                                                    onChange={this.onChangeCountry}
                                                    options={options}
                                                    placeholder="Ex : Italie"
                                                    required
                                                    value={country}
                                                />
                                            </Form.Group>
                                        </Row>
                                        <Row>
                                        
                                            <Form.Group as={Col}>
                                                <Form.Label className="label-title"><strong className="chevron">{">"}</strong> Type de voyage</Form.Label>
                                                
                                                    <div className="mb-2">
                                                    <Form.Check
                                                        className="checkbox"
                                                        custom
                                                        inline
                                                        label="Voyage culturel"
                                                        type="checkbox"
                                                        id="culturel"
                                                        value="Culturel"
                                                        onClick={this.handleClick}
                                                    />
                                                    <Form.Check
                                                        className="checkbox"
                                                        custom
                                                        inline
                                                        label="Voyage détente"
                                                        type="checkbox"
                                                        id="detente"
                                                        value="Détente"
                                                        onClick={this.handleClick}
                                                    />
                                                    <Form.Check
                                                        className="checkbox"
                                                        custom
                                                        inline
                                                        label="Voyage circuit"
                                                        type="checkbox"
                                                        id="circuit"
                                                        value="Circuit"
                                                        onClick={this.handleClick}
                                                    />
                                                    <Form.Check
                                                        className="checkbox"
                                                        custom
                                                        inline
                                                        label="Voyage fêtard"
                                                        type="checkbox"
                                                        id="fetard"
                                                        value="Fetard"
                                                        onClick={this.handleClick}
                                                    />
                                                    <Form.Check
                                                        className="checkbox"
                                                        custom
                                                        inline
                                                        label="Voyage sportif"
                                                        type="checkbox"
                                                        id="sportif"
                                                        value="Sportif"
                                                        onClick={this.handleClick}
                                                    />
                                                    </div>
                                                
                                            </Form.Group>
                                        </Row>
                                        <Row>
                                        
                                            <Form.Group as={Col} xs={12} md={4} controlId="formGridEmail">
                                                <Form.Label className="label-title"><strong className="chevron">{">"}</strong> Type de logement</Form.Label>
                                                <Form.Control type="text" name="lodging" onChange={this.onChange} value ={lodging} placeholder="Airbnb, hotel, camping" />
                                            </Form.Group>
                                            <Form.Group as={Col} xs={12} md={4} controlId="formGridEmail">
                                                <Form.Label className="label-title"><strong className="chevron">{">"}</strong> Date de début</Form.Label>
                                                <Form.Control type="date" name="start_date" onChange={this.onChange} value ={start_date} placeholder="" required/>
                                            </Form.Group>

                                            <Form.Group as={Col} xs={12} md={4} controlId="formGridEmail">
                                                <Form.Label className="label-title"><strong className="chevron">{">"}</strong> Date de fin</Form.Label>
                                                <Form.Control type="date"  name="end_date" onChange={this.onChange}  value ={end_date} placeholder="" required/>
                                            </Form.Group>
                                        
                                        </Row>
                                        <Row>
                                            
                                            <Form.Group as={Col} xs={12} md={4} controlId="formGridEmail">
                                                <Form.Label className="label-title"><strong className="chevron">{">"}</strong> Taille du groupe</Form.Label>
                                                <Form.Control name="nb_traveler" onChange={this.onChange} value ={nb_traveler} placeholder="Nombre de voyageurs au total voulu" required/>
                                            </Form.Group>

                                            <Form.Group as={Col} xs={12} md={4} controlId="formGridEmail">
                                                <Form.Label className="label-title"><strong className="chevron">{">"}</strong> Âge moyen</Form.Label>
                                                <Form.Control type="number"  name="middle_age" onChange={this.onChange} value ={middle_age} placeholder="" required/>
                                            </Form.Group>
                                            <Form.Group as={Col} xs={12} md={4} controlId="formGridEmail">
                                                <Form.Label className="label-title"><strong className="chevron">{">"}</strong> Budget</Form.Label>
                                                <Form.Control type="number"  name="Budget" onChange={this.onChange} value ={Budget} placeholder="" required/>
                                            </Form.Group>  

                                        </Row>
                                        <Row>
                                                <Form.Group as={Col} controlId="formGridEmail">
                                                    <Form.Label className="label-title"><strong className="chevron">{">"}</strong> Description</Form.Label>
                                                    <Form.Control as="textarea" rows="3" name="description" onChange={this.onChange} value ={description} placeholder="Une description du voyage ainsi qu'une petite présentation du groupe de voyageurs et de ce que vous recherchez dans un trip !" required/>
                                                </Form.Group>
                                        </Row>
                                        <Row>
                                                
                                            <Form.Group as={Col} controlId="formGridEmail">
                                                <Form.Label className="label-title"><strong className="chevron">{">"}</strong> Programme</Form.Label>
                                                <Form.Control as="textarea" rows="3" name="program" onChange={this.onChange} value ={program} placeholder="Escalader les Alpes ? Passer la semaine à la plage ? Précisez les activités qui ont déjà été prévues !" />
                                            </Form.Group>
                                            

                                        </Row>
                                        <Row>
                                            <Form.Label as={Col} className="label-title"><strong className="chevron">{">"}</strong> Image</Form.Label>
                                        </Row>
                                        <Row className="ml-2">
                                        <FileBase64    multiple={ true }    onDone={ this.getFiles} />
                                        </Row>
                                        <Row>
                                        </Row>
                                            <br/>
                                        <Row className="justify-content-center align-items-center">
                                            <Button className="green-button" onClick={this.onSubmit}>
                                                Créer le voyage
                                            </Button>
                                        </Row>
                                        </Card.Body>
                                    </Card>
                                </Row>
                                
                            </Col>
                        </Row>
                            </Container>

                        </Col>
                    </Row>
                    
            </Container>
        )
    }
}

const condition = authUser => !!authUser;

export default  compose(
    withAuthorization(condition),
	withRouter,
	withBackaccessContext,
)(Createtrip);