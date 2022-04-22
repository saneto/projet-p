import React, {Component} from 'react';
import { withRouter } from 'react-router-dom';
import { compose } from 'recompose';
import { withBackaccessContext } from '../BackEnd';
import {Row, Container, Button, Col} from 'react-bootstrap';

class HomeContent extends Component{

    constructor(props){
        super(props);
        this.state = {
            trips :[],
            tripsHome: [],
            uri: 'trips'
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
        .getData(this.state.uri)
        .then((data) => {
            this.setState({ trips: data.response});
            //console.log(this.state.trips);
            const tripHome = [];
            if(this.state.trips.length > 0)
                for(let i = 0; i < this.state.trips.length; i++){
                    tripHome.push(this.state.trips[i]);
                }
            this.setState({ tripsHome: tripHome})
        }).catch(error => {
            console.log("not ok");
            this.setState({ error });
        });
    };

    render() {
        //console.log(this.state.tripsHome)
        return (
            <React.Fragment>
                <main>
                    <div className="container-fluid">
                        <h2 className="text-center mb-4">Les derniers voyages créés par les Tripiziens !</h2>
                        <div className="trip-teasing">

                            <div className="containerTripHome justify-content-center">
                               {(this.state.tripsHome && this.state.tripsHome.length > 0) ? this.state.tripsHome.map(trip => (
                                    <div className="box" key={trip.id}>
                                        <div className="imgBox" key={trip.id}>
                                            {trip.base64 &&
                                                <img src={trip.base64} alt="Trip custom"></img>
                                            }

                                            {!trip.base64 &&
                                                <img src="img/card1.jpg" alt="Trip default"></img>
                                            }
                                        </div>
                                        <a href={`trip/${trip.id}`}>
                                            <div className="details">
                                                <div className="content">
                                                    <h2>{this.filterText(trip.title, 42)}</h2>
                                                    <p>{this.filterText(trip.description, 96)}</p>
                                                    <div className="infoTrip mt-4">
                                                        <li className="badge badge-pill badge-success mr-2">{trip.trip_type}</li>
                                                        <li className="badge badge-pill badge-success mr-2">{trip.nb_traveler} voyageurs</li>
                                                        <li className="badge badge-pill badge-success">{trip.nb_days} jours</li>
                                                    </div>
                                                </div>
                                            </div>
                                        </a>
                                    </div>
                                )): ""}
                            </div>
                                        
                            <div className="row d-flex justify-content-center align-items-center pb-5">
                                <Button href="/trips" className="btn border rounded standardhigh">Afficher les voyages</Button>
                            </div>
                        </div>
                    </div>
                    <div className="container-fluid">
                        <div className="row d-flex justify-content-center align-items-center separator"><h2 className="px-5 text-center">"Ce qui importe, ce n'est pas le voyage, c'est celui avec qui on voyage"</h2></div>
                    </div>
                    <Container>
                            <Row>
                                <Col xs={0} lg={4}>
                                    <Row className="p-2"><img className="block" src="img/block1.jpg" alt="block1"/></Row>
                                    <Row className="p-2"><img className="block" src="img/block2.jpg" alt="block1"/></Row>
                                </Col>
                                <Col xs={12} lg={8} className="mt-2 mb-2">
                                    <Row className="block-box block-text-green">
                                        <Col>
                                        <h2 className="text-center py-5">「 LFTRIP 」</h2>
                                        <p className="text-justify px-5 font-weight-light">Il s'agit avant tout de <strong className="brand-strong-green">votre communauté</strong> avec laquelle vous écrivez vos propres <strong className="brand-strong-red">histoires</strong> et <strong className="brand-strong-blue">voyages</strong>.</p>
                                        <p className="text-justify px-5 font-weight-light"><strong className="brand-strong-blue">Envie</strong> d'aller quelque part mais avec personne pour vous accompagner ?</p>
                                        <p className="text-justify px-5 font-weight-light">Vous voulez <strong className="brand-strong-red">sortir</strong> de votre routine mais vous ne savez pas quoi faire ?</p>
                                        <p className="text-justify px-5 font-weight-light"><strong className="brand-strong-green">「 LFTRIP 」</strong>c'est la communauté qui permet de relier les personnes avec les mêmes envies et les mêmes passions.</p>
                                        <p className="text-center px-5 font-weight-light"><strong>Rejoignez-nous !</strong></p>
                                        </Col>
                                    </Row>  
                                </Col>
                            </Row>
                            <Row>
                                <Col xs={0} lg={4}>
                                    <Row className="px-2 pb-2"><img className="block" src="img/block3.jpg" alt="block1"/></Row>
                                </Col>
                                <Col xs={12} lg={4} className="mb-2">
                                    <Row className="block-box block-text-green">
                                        <Col>
                                        <h3 className="text-center py-3">「 Concept 」</h3>
                                        <p className="text-justify px-2 font-weight-light">Renseigner votre profil, matcher, voyager !</p>
                                        <p className="text-justify px-2 font-weight-light">C'est simple, après votre inscription, laissez-vous guider et intégrer la communauté !</p>
                                        </Col>
                                    </Row>
                                </Col>
                                
                                <Col xs={12} lg={4} className="mb-2">
                                <Row className="block-box block-text-green">
                                        <Col>
                                        <h3 className="text-center py-3">「 Fonctionnement 」</h3>
                                        <p className="text-justify px-2 font-weight-light">Une fois votre profil complété, proposez votre trip idéal ou bien cherchez celui qui vous correspond le mieux !</p>
                                        <p className="text-justify px-2 font-weight-light">Entrez en relation pour vérifier si cela correspond à vos attentes et voyager !</p>
                                        </Col>
                                    </Row>
                                </Col>
                            </Row>
                    </Container>
                </main>
            </React.Fragment>
        );
    }
}

export default compose(
	withRouter,
	withBackaccessContext,
)(HomeContent);