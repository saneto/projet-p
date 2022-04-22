import React, {Component} from 'react';
import {Row, Container, Button, Col, Form} from 'react-bootstrap';
import { withRouter } from 'react-router-dom';
import { compose } from 'recompose';
import { withBackaccessContext } from '../BackEnd';
import FileBase64 from 'react-file-base64';
import {Typeahead} from "react-bootstrap-typeahead";
import options from "./data";
import Card from 'react-bootstrap/Card';
import latlong from './latlong';
import {store} from "react-notifications-component";
import { withAuthorization } from '../Session';
//import  { post } from 'axios';

class Updatetrip extends Component{
    constructor(props){
        super(props);
        this.state = {
            uri : "trips",
            country: "",
            title: "",
            nb_days : 0,
            description: "",
            start_date: "",
            end_date: "",
            middle_age: 0,
            nb_traveler: 0,
            budget: 0,
            lodging : "",
            program: "",
            base64:"",
            id :  this.props.match.params.id,
            is_active: true,
            trip_type : "",
            latitude:0,
            longitude:0,
        }
    }

    componentDidMount() {
        this.props.backaccess
        .getDataWitId(this.state.id,this.state.uri)
        .then(data => {
            this.setState(data.response);
            let date =new Date(this.state.start_date)
            let dateString = new Date(date.getTime() - (date.getTimezoneOffset() * 60000 ))
                    .toISOString()
                    .split("T")[0];
            this.setState({start_date : dateString});
            let date2 =new Date(this.state.end_date)

            let dateString2 = new Date(date2.getTime() - (date2.getTimezoneOffset() * 60000 ))
                    .toISOString()
                    .split("T")[0];
            this.setState({end_date : dateString2});

            //this.setState({ img: this.state.trip.base64});
           /// console.log(this.state.trip);
            //console.log(this.state);
        }).catch(error => {
            console.log("not ok");
            console.log(error);
            this.setState({ error });
        });
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
                    message: "L'âge moyen doit supèrieur à 16 ou inférieur à 100",
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
                    //console.log(this.state.latitude);
                    //console.log(JSON.stringify(this.state.latitude));
                    //console.log(JSON.parse(this.state.latitude));
                    let data = {
                        "country": this.state.country,
                        "title": this.state.title,
                        "nb_days" : nb_Days,
                        "description":this.state.description,
                        "start_date": this.state.start_date+"T00:00:00.00Z",
                        "end_date": this.state.end_date+"T00:00:00.00Z",
                        "middle_age": this.state.middle_age,
                        "nb_traveler": parseInt(this.state.nb_traveler) ,
                        "budget":  parseInt(this.state.budget),
                        "lodging" : this.state.lodging,
                        "Program": this.state.program,
                        'base64' : this.state.base64,
                        "is_active": this.state.is_active,
                        "trip_type" : this.state.trip_type,
                        "latitude": this.state.latitude,
                        "longitude": this.state.longitude
                    }
                    this.props.backaccess.updateData(data, this.state.uri, this.state.id).then( ret =>{
                            //console.log(ret)
                        }
                    ).catch( error =>{

                    })
                }
            }
        } else{
            store.addNotification({
                title: 'Erreur sur les dates',
                message: 'Date de début doit être supérieur à la date de fin',
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

    onChangeCountry = event => {
        if(event[0] !== undefined){
            this.setState({ country: event[0].label});
            
            latlong.forEach((latlongv) => {
                if(event[0].alpha2 === latlongv.alpha2){
                    //console.log(latlongv);
                    this.setState({latitude: latlongv.latitude, longitude: latlongv.longitude});
                }
            })}

    };

    onChange = event => {
		this.setState({ [event.target.name]: event.target.value });
    };

    getFiles = event =>{
        //console.log(event[0])
        this.setState({ base64: event[0].base64 })
    }

    render()
    {
        const { country,  title,
                description, start_date, end_date,
                middle_age, nb_traveler, budget,
                program , lodging} = this.state;
        let img= "";
        if(this.state.base64){
            img = <img className="card-img-top img-fluid" src={this.state.base64} alt=""/>
        }else{
            img =  <img className="card-img-top img-fluid" src="/img/card1.jpg" value={this.state.base64} alt=""/>
        }
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
                                                    <h3 className="title">Modifier le Trip qui vous ressemble</h3>
                                                </div>
                                                
                                            </Col>
                                        </Row>
                                        <Card.Body className="create-trip-card">
                                        <Row>
                                            <Form.Group as={Col} xs={12} md={8} controlId="formGridEmail">
                                                <Form.Label className="label-title"><strong className="chevron">{">"}</strong> Un titre accrocheur pour votre trip !</Form.Label>
                                                <Form.Control type="text" name="title" onChange={this.onChange} value ={title} placeholder="Ex : Un roadtrip en Italie !" required/>
                                            </Form.Group>
                                            <Form.Group as={Col} xs={12} md={4} controlId="formGridEmail">
                                                <Form.Label className="label-title"><strong className="chevron">{">"}</strong> Destination</Form.Label>
                                                <Typeahead
                                                    {...this.state}
                                                    name="country"
                                                    id="basic-example"
                                                    required
                                                    onChange={this.onChangeCountry}
                                                    options={options}
                                                    placeholder={country}
                                                    value={country}
                                                    disabled
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
                                                <Form.Control type="text"  name="lodging" onChange={this.onChange} value ={lodging} placeholder="Airbnb, hotel, camping" />
                                            </Form.Group>
                                            <Form.Group as={Col} xs={12} md={4} controlId="formGridEmail">
                                                <Form.Label className="label-title"><strong className="chevron">{">"}</strong> Date de début</Form.Label>
                                                <Form.Control type="date" name="start_date" onChange={this.onChange} value ={start_date}/>
                                            </Form.Group>

                                            <Form.Group as={Col} xs={12} md={4} controlId="formGridEmail">
                                                <Form.Label className="label-title"><strong className="chevron">{">"}</strong> Date de fin</Form.Label>
                                                <Form.Control type="date"  name="end_date" onChange={this.onChange}  value ={end_date}/>
                                            </Form.Group>
                                        
                                        </Row>
                                        <Row>
                                            
                                            <Form.Group as={Col} xs={12} md={4} controlId="formGridEmail">
                                                <Form.Label className="label-title"><strong className="chevron">{">"}</strong> Taille du groupe</Form.Label>
                                                <Form.Control name="nb_traveler" onChange={this.onChange} value ={nb_traveler} placeholder="Nombre de voyageurs au total voulu" required/>
                                            </Form.Group>

                                            <Form.Group as={Col} xs={12} md={4} controlId="formGridEmail">
                                                <Form.Label className="label-title"><strong className="chevron">{">"}</strong> Âge moyen</Form.Label>
                                                <Form.Control type="number"  name="middle_age" onChange={this.onChange} value ={middle_age}/>
                                            </Form.Group>
                                            <Form.Group as={Col} xs={12} md={4} controlId="formGridEmail">
                                                <Form.Label className="label-title"><strong className="chevron">{">"}</strong> Budget</Form.Label>
                                                <Form.Control type="number"  name="budget" onChange={this.onChange} value ={budget}/>
                                            </Form.Group>  

                                        </Row>
                                        <Row>
                                                <Form.Group as={Col} controlId="formGridEmail">
                                                    <Form.Label className="label-title"><strong className="chevron">{">"}</strong> Description</Form.Label>
                                                    <Form.Control as="textarea" rows="3" name="description" onChange={this.onChange} value ={description} required placeholder="Une description du voyage ainsi qu'une petite présentation du groupe de voyageurs et de ce que vous recherchez dans un trip !" />
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
                                                Mettre à jour le voyage
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
)(Updatetrip);



