import React, {Component} from 'react';
import Button from 'react-bootstrap/Button';
import { withRouter } from 'react-router-dom';
import { compose } from 'recompose';
import { withBackaccessContext } from '../BackEnd';
import { withAuthorization } from '../Session';
import Container from 'react-bootstrap/Container';
/* NOT USED
const styles = {
div:{
    display: 'flex',
    flexFlow: 'wrap',
    //justifyContent: 'center',
    alignItems: 'center',
}
};
*/
const user = JSON.parse(localStorage.getItem("authUser"));
class ListTripParticipated extends Component{
    constructor(props){
        super(props);
        this.state = {
            participations :[],
            uri: 'userapplies',
            id: user.id
        }
    }

    filterText = (text, nbChar) => {
        var filteredText = text;
		if (text.length > nbChar){
            filteredText = text.substring(0, nbChar) + "...";
        }
        return filteredText;
	}

    componentDidMount() {
        this.props.backaccess
        .getDataWitId(this.state.id,this.state.uri)
        .then((data) => {
            this.setState({ participations: data.response});
            //console.log(this.state.participations.trip);
        }).catch(error => {
            console.log("not ok");
            this.setState({ error });
        });
    };
    render()
    {
        return(
            <React.Fragment>
                <div className="title-top">
                    <div className="titleContainerTrip">
                        <h1>Voyages réservés</h1>
                    </div>
                </div>

                {!this.state.participations.length > 0 &&
                    <section className="no-trip-display">
                        <div className="no-trip-display-container">
                            <div>
                                <p>Aucune réservation pour le moment.</p>
                                <p>
                                    <Button variant="outline-dark" href="/trips" className="btn border border-dark rounded standardhigh red-hover btn-no-trip">Parcours nos voyages</Button>
                                </p>
                            </div>
                        </div>
                    </section>
                }
                <Container>
                    <div className="containerTrip">
                        {this.state.participations.map(participation => (
                            <div className="box" key={participation.trip_id}>
                                <div className="imgBox" key={participation.trip_id}>
                                    {participation.trip.base64 &&
                                        <img src={participation.trip.base64} alt="Trip custom"></img>
                                    }

                                    {!participation.trip.base64 &&
                                        <img src="img/card1.jpg" alt="Trip default"></img>
                                    }                            
                                </div>
                                <a href={`trip/${participation.trip_id}`}>
                                    <div className="details">
                                        <div className="content">
                                            <h2>{this.filterText(participation.trip.title,42)}</h2>
                                            <p>{this.filterText(participation.trip.description,96)}</p>
                                            <div className="infoTrip mt-4">
                                                <li className="badge badge-pill badge-success mr-2">{participation.trip.trip_type}</li>
                                                <li className="badge badge-pill badge-success mr-2">{participation.trip.nb_traveler} voyageurs</li>
                                                <li className="badge badge-pill badge-success">{participation.trip.nb_days} jours</li>
                                            </div>
                                        </div>
                                    </div>
                                </a>
                            </div>                    
                        ))}
                    </div>
                </Container>
            </React.Fragment>
        )
    }
}

const condition = authUser => !!authUser;

export default  compose(
    withAuthorization(condition),
    withRouter,
    withBackaccessContext,
)(ListTripParticipated);