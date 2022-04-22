import React from 'react';
import {withRouter} from "react-router-dom";
import {withBackaccessContext} from "../BackEnd";
import {compose} from "recompose";
import Container from 'react-bootstrap/Container';
//import ListFellowUsers from "./ListFellowUsers";
//import Comment from "../Commentaire/Comment";
import Row from "react-bootstrap/Row";
import Col from "react-bootstrap/Col";
import EditIcon from "@material-ui/icons/Edit";
import Button from "react-bootstrap/Button";
import Modal from "react-bootstrap/Modal";
import {Form} from "react-bootstrap";
import SendIcon from '@material-ui/icons/Send';
import { withAuthorization } from '../Session';

const data = JSON.parse(localStorage.getItem('authUser'));

const INITIAL_STATE = {
    uri:"userpreferences",
    href:'',
    uri2:"user_userpreferences",
    uri3:"send_proposal",
    user_id : '',
    email: '',
    firstname: '',
    lastname: '',
    city : '',
    sexe : '',
    date_of_birth : '',
    phone_number : '',
    airport : '',
    description : '',
    age : '',
    type_trip: '',
    avatar_path: '',
    filteredUsers: [],
    ShyOutgoing: "",
    CalmOrSpark:"",
    SundayMikeHornAdventurer:"",
    AthleteSportKezako:"",
    FatMorningReadyLeave7AM:"",
    scootered_disorganized:"",
    trekking_beach_towel :"",
    LuxuryCharmingLodging : "",
    food:"",
    userPref: [],
    proposition: '',
    token: '',
    openModal: false
};


class ProfileLikeSummary extends React.Component{

    constructor(props) {
        super(props);

        this.state = { ...INITIAL_STATE };
        this.state.userId = this.props.match.params.id;
        this.state.token = data.token;
    }

    componentDidMount() {
        this.setState({userPref: []});
        this.setState({userPref: this.state.userPref});
        this.props.backaccess.getDataWitId(this.state.userId, this.state.uri2).then((data) => {
            this.setState({type_trip: data.response[0].traveler_type, email: data.response[0].user.email, avatar_path: data.response[0].user.avatar_path, firstname : data.response[0].user.firstname, lastname : data.response[0].user.lastname, sexe : data.response[0].user.sexe, city : data.response[0].user.city, date_of_birth : data.response[0].user.date_of_birth, airport : data.response[0].user.departure_airport, phone_number : data.response[0].user.phone_number, description : data.response[0].user.description});
            this.setState({age: this.calculate_age(data.response[0].user.date_of_birth)});
            this.setState({
                ShyOutgoing: data.response[0].shy_outgoing,
                CalmOrSpark: data.response[0].calm_or_spark,
                SundayMikeHornAdventurer: data.response[0].sunday_mike_horn_adventurer,
                AthleteSportKezako: data.response[0].athlete_sport_kezako,
                FatMorningReadyLeave7AM: data.response[0].fat_morning_ready_leave_7_am,
                scootered_disorganized: data.response[0].scootered_disorganized,
                trekking_beach_towel : data.response[0].trekking_beach_towel,
                LuxuryCharmingLodging : data.response[0].luxury_charming_lodging,
                food: data.response[0].food});
            // populate user pref array with preference
            this.state.userPref.push(this.state.ShyOutgoing, this.state.CalmOrSpark, this.state.SundayMikeHornAdventurer, this.state.AthleteSportKezako, this.state.FatMorningReadyLeave7AM, this.state.scootered_disorganized,
                this.state.trekking_beach_towel, this.state.LuxuryCharmingLodging, this.state.food);
            this.setState({userPref: this.state.userPref});

            // remove empty value from array
            
            let pref = this.state.userPref.filter(pref => pref !== undefined);
            this.setState({userPref: pref});
        });
    }

    calculate_age(date) {
        const dateTime = new Date(date);
        const diff_ms = Date.now() - dateTime.getTime();
        const age_dt = new Date(diff_ms);
        return Math.abs(age_dt.getUTCFullYear() - 1970)
    }

    openModal = event => {
        this.setState({openModal: true});
    };
    handleClose = event => {
        this.setState({openModal: false});
    };
    onChange = event => {
        this.setState({ [event.target.name]: event.target.value });
    };
    sendProposition = event => {
        const {proposition, token} = this.state;
        this.props.backaccess.sendTripProposition({firstname: this.state.firstname, message: proposition, email: this.state.email, token : token,}, this.state.uri3)
            .then(() => {
                this.setState({openModal: false});
            }).catch(error => {
            this.setState({ error });
        });
        event.preventDefault();
    };

    render(){
        let img= "";
        if(this.state.avatar_path){
            img = <img src={this.state.avatar_path} className="rounded-circle" height="150px" alt="avatar"/>
        }else{
            img = <img src="/img/Default.png" className="rounded-circle" value={this.state.avatar_path} height="150px" alt="avatar"/>
        }
        const {userPref} = this.state;
        return (
            <Container fluid>
                <Row className="d-flex justify-content-center align-items-center pt-5 pb-5 profil-summary-background">
                    <Col>
                        <div className="text-center">
                            {img}
                            <h1>{this.state.firstname} {this.state.lastname}</h1>
                            <p>{this.state.age} ans, {this.state.sexe}</p>
                            <p>{this.state.type_trip}</p>
                            <br/>
                            <Button variant="light" onClick={this.openModal}>Me contacter <SendIcon/></Button>
                        </div>
                    </Col>
                </Row>
                <Row className="d-flex justify-content-end align-items-end">
                    <Col xs={12} md={4} className="mt-3 mar-b-6">
                        <p className="intro">{this.state.description}</p>
                    </Col>
                    <Col xs={12} md={3}>
                        <div className="card-container reverse-bloc sub-block sub-block-like">
                            <strong>Traits de personnalité</strong>
                            <EditIcon/>
                            <ul className="text-interest mar-b-10">
                                {userPref.map((pref, index) =>
                                    <li key={index}>{pref}</li>
                                )}
                            </ul>
                        </div>
                    </Col>
                </Row>
                <Modal show={this.state.openModal} onHide={this.handleClose} animation={false} aria-labelledby="contained-modal-title-vcenter" centered>
                    <Modal.Header closeButton>
                        <Modal.Title>Formulaire de contact</Modal.Title>
                    </Modal.Header>
                    <Modal.Body>
                        <Form.Group as={Col} controlId="propositionTrip">
                            <Form.Label className="label-title">Message</Form.Label>
                            <Form.Control as="textarea" rows="3" name="proposition" onChange={this.onChange} placeholder="En visitant votre profil et en voyant vos traits de personnalités, ..." />
                        </Form.Group>
                    </Modal.Body>
                    <Modal.Footer>
                        <Button variant="outline-danger" onClick={this.handleClose}>
                            Annuler
                        </Button>
                        <Button variant="success" onClick={this.sendProposition}>
                            Envoyer
                        </Button>
                    </Modal.Footer>
                </Modal>
            </Container>
        );
    }
}

const condition = authUser => !!authUser;

export default compose(
    withAuthorization(condition),
    withRouter,
    withBackaccessContext,
)(ProfileLikeSummary);

