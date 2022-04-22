import React from 'react';
import Card from 'react-bootstrap/Card';
import Container from 'react-bootstrap/Container';
import Row from 'react-bootstrap/Row';
import Col from 'react-bootstrap/Col';
import Badge from 'react-bootstrap/Badge'

class TripInformations extends React.Component{

    constructor(props){
        super(props);

        this.state = {
            
        }
    }

    render(){
        return (
            <React.Fragment>
                
                <Card className="mb-5">
                    <Card.Header as="h5" className="trip-card-detail-header m-2">Informations générales sur le voyage</Card.Header>
                    <Card.Body>
                        <Card.Title as="h4" className="pb-3">
                            <b>{this.props.trip.title}</b>
                        </Card.Title>
                        <Container className="pb-3">
                            <Row>
                                <Col className="d-flex justify-content-center align-items-center"><Badge variant="info p-2">{this.props.trip.country}</Badge>{' '}</Col>
                                <Col className="d-flex justify-content-center align-items-center"><Badge variant="info p-2">{this.props.trip.nb_traveler} personnes</Badge>{' '}</Col>
                                <Col className="d-flex justify-content-center align-items-center"><Badge variant="info p-2">{this.props.trip.nb_days} jours</Badge>{' '}</Col>
                                <Col className="d-flex justify-content-center align-items-center"><Badge variant="info p-2">{this.props.trip.budget}€ à prévoir</Badge>{' '}</Col>
                            </Row>
                        </Container>
                        <h5 className="mt-3"><u>Description :</u></h5>
                        {this.props.trip.description}
                        <h5 className="mt-3"><u>Programme :</u></h5>
                        {this.props.trip.program}
                        <h5 className="mt-3"><u>Hébergement :</u></h5>
                        {this.props.trip.lodging}
                        <p className="mt-3"><strong><u>Age moyen du groupe recherché :</u></strong> {this.props.trip.middle_age} ans</p>
     
                        <h5 className="mt-4 mb-4"><u>A propos de l'organisateur :</u></h5>  
                        {this.props.author &&
                            <React.Fragment>
                                <Container>
                                    <Row>
                                        <Col xs={12} md={1}>
                                            {this.props.author.avatar_path &&
                                            <img src={this.props.author.avatar_path} className="rounded-circle avatar" height="50px" width="50px" alt="avatar"/>
                                            }
                                            {!this.props.author.avatar_path &&
                                            <img src="/img/Default.png" className="rounded-circle avatar" height="50px" width="50px" alt="avatar"/>
                                            }
                                        </Col>
                                        <Col className="d-flex justify-content-start align-items-center">
                                            <a href={`/user-alike/${this.props.author.id}`}>
                                                <span className="bold pr-2">
                                                    {this.props.author.firstname &&
                                                        " " + this.props.author.firstname + " "
                                                    }
                                                    {this.props.author.lastname &&
                                                        this.props.author.lastname + " "
                                                    }
                                                </span>
                                            </a>
                                            <span className="italic pr-2">{this.props.author.city && 
                                                " (originaire de " + this.props.author.city + ")"
                                            }</span>
                                            <span>{this.props.author.is_guide &&  this.props.author.is_guide === "true" &&
                                               <span className="grey-text italic bold">Guide</span>
                                            }</span>
                                        </Col>
                                    </Row>
                                    <Row className="pt-3">
                                        <b><u>Ma p'tite description</u></b>   
                                        <span>
                                            {this.props.author.description &&
                                                " : " + this.props.author.description
                                            }
                                        </span>
                                        
                                    </Row>
                                </Container>
                                                     
                            </React.Fragment>  
                        }   
                    </Card.Body>
                </Card>
            </React.Fragment>
        );
    }
}

export default TripInformations;