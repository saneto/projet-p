import React from 'react';
import Container from 'react-bootstrap/Container';
import Row from 'react-bootstrap/Row';
import Col from 'react-bootstrap/Col';
import ListGroup from 'react-bootstrap/ListGroup';
import Tab from 'react-bootstrap/Tab';

import AccountInformation from './AccountInformation';
import AccountProfil from "./AccountProfil";
import PreferencesProfil from "./PreferencesProfil";

class AccountPage extends React.Component{

    render(){
        return (
            <React.Fragment>
                <Container fluid className="covervh">
                    <h2 className="text-center mb-4 pt-5">Mon Compte</h2>
                    <div className="trip-teasing">

                        <Tab.Container id="list-group-tabs-example" defaultActiveKey="#account">
                            <Row className="d-flex justify-content-start align-items-start pt-5 pb-5">

                                <Col className="padding-left-10" md="3">
                                    <ListGroup>
                                        <ListGroup.Item action href="#account">
                                            Mes Informations Personnelles
                                        </ListGroup.Item>
                                        <ListGroup.Item action href="#profil">Mon Profil</ListGroup.Item>
                                        <ListGroup.Item action href="#preferences">Mes Préférences</ListGroup.Item>
                                        {/* <ListGroup.Item action href="#link3" disabled>Abonnement</ListGroup.Item>
                                        <ListGroup.Item action href="#link4" disabled>Confidentialité & Sécurité</ListGroup.Item> */}
                                        <ListGroup.Item variant="secondary">Abonnement</ListGroup.Item>
                                        <ListGroup.Item variant="secondary">Confidentialité & Sécurité</ListGroup.Item>
                                    </ListGroup><br/>
                                    <span>Paramètres de l'application</span>
                                    <ListGroup>
                                        <ListGroup.Item action href="#link5">Notifications</ListGroup.Item>
                                        {/* <ListGroup.Item action href="#link6" disabled>Textes & Images</ListGroup.Item> */}
                                        <ListGroup.Item variant="secondary">Textes & Images</ListGroup.Item>
                                    </ListGroup>
                                </Col>

                                <Col md="9">
                                <Tab.Content>
                                    <Tab.Pane eventKey="#account">
                                        <AccountInformation/>
                                    </Tab.Pane>
                                    <Tab.Pane eventKey="#profil">
                                        <AccountProfil/>
                                    </Tab.Pane>
                                    <Tab.Pane eventKey="#preferences">
                                        <PreferencesProfil/>
                                    </Tab.Pane>
                                </Tab.Content>
                                </Col>

                            </Row>
                        </Tab.Container>

                    </div>
                </Container>
            </React.Fragment>
        );
    }
}

export default AccountPage;
