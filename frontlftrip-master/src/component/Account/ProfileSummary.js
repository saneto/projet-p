import React from 'react';
import {withRouter} from "react-router-dom";
import {withBackaccessContext} from "../BackEnd";
import {compose} from "recompose";
import Container from 'react-bootstrap/Container';
import ListFellowUsers from "./ListFellowUsers";
import {Row} from "react-bootstrap";
import './listFellowUsers.css';
import Col from "react-bootstrap/Col";
import EditIcon from '@material-ui/icons/Edit';
import { withAuthorization } from '../Session';

let data= "";
if (localStorage.getItem(('authUser'))) {
    data = JSON.parse(localStorage.getItem('authUser'));
}

const INITIAL_STATE = {
    uri:"userpreferences",
    href:'',
    uri2:"user_userpreferences",
    user_id : data.id || '',
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
    avatar_path:'',
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
    userPref: []
};


class ProfileSummary extends React.Component{

    constructor(props) {
        super(props);

        this.state = { ...INITIAL_STATE };
    }

    componentDidMount() {
        //this.state.userPref = []; /* CLEAR WARNING */

        this.setState({userPref: []});
        this.setState({userPref: this.state.userPref});

        this.props.backaccess.getDataWitId(this.state.user_id, this.state.uri2).then((data) => {
            //console.log(data);
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
            //this.state.userPref = this.state.userPref.filter(pref => pref !== undefined);
            this.setState({userPref: this.state.userPref.filter(pref => pref !== undefined)});
            this.setState({userPref: this.state.userPref});
        });
        this.props.backaccess
            .getData(this.state.uri)
            .then((data) => {
                let filteredUser = data.response.filter(preference => ((Math.abs((this.calculate_age(preference.user.date_of_birth)) - (this.state.age)) <= 5) && (preference.user_id !== this.state.user_id)) || ((preference.user.city === this.state.city) && (preference.user_id !== this.state.user_id)) || ((preference.traveler_type === this.state.type_trip) && (preference.user_id !== this.state.user_id)));
                this.setState({filteredUsers : filteredUser});
                //console.log(filteredUser);
                //console.log(this.state.filteredUsers);
                })
            .catch(error => {
            this.setState({ error });
        });
    }

    calculate_age(date) {
        const dateTime = new Date(date);
        const diff_ms = Date.now() - dateTime.getTime();
        const age_dt = new Date(diff_ms);
        return Math.abs(age_dt.getUTCFullYear() - 1970)
    }

    render(){
        let img= "";
        if(this.state.avatar_path){
            img = <img src={this.state.avatar_path} className="rounded-circle" height="150px" alt="avatar"/>
        }else{
            img = <img src="/img/Default.png" className="rounded-circle" value={this.state.avatar_path} height="150px" alt="avatar"/>
        }
        const {filteredUsers, userPref} = this.state;
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
                        </div>
                    </Col>
                </Row>
                <Row className="d-flex justify-content-end align-items-end">
                    <Col xs={12} md={4} className="mt-3 mar-b-6">
                        <p className="intro">{this.state.description}</p>
                    </Col>
                    <Col xs={12} md={3}>
                        <div className="card-container reverse-bloc sub-block">
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

                <Row className="d-flex justify-content-center align-items-center pt-5 pb-5">
                    <Container>
                        <Row className="d-flex justify-content-center align-items-center pb-5">
                            <h5>Ces voyageurs pourraient te ressembler</h5>
                        </Row>
                        <Row className="d-flex justify-content-center align-items-center">
                            {filteredUsers.map(filteredUser =>(
                                    <ListFellowUsers key={filteredUser.id} filteredUser={filteredUser} userFilteredId={filteredUser.id}/>
                                )
                            )}
                        </Row>
                    </Container>
                </Row>
            </Container>
        );
    }
}

const condition = authUser => !!authUser;

export default compose(
    withAuthorization(condition),
    withRouter,
    withBackaccessContext,
)(ProfileSummary);

