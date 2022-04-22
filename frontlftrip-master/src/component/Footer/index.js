import React from 'react';
import Container from 'react-bootstrap/Container';
import Row from 'react-bootstrap/Row';
import Col from 'react-bootstrap/Col';

const Footer = () => (
    <div className="footer">
        <Container fluid>
            <Row className="no-padding-margin">
                <Col xs={12} className="footer-text no-padding-margin">
                    <Container fluid>
                        <Row className="d-flex justify-content-center align-items-center py-5 m-0"><h3>Faites confiance à LFTRIP!</h3></Row>
                        <Row className="px-md-5 m-0">
                            <Col xs={12} md={6} className="px-md-5">
                                <p>LFTRIP a pour objectif de satisfaire ses utilisateurs et de les accompagner dans le bon déroulement de leurs voyages.</p>
                                <p>Du choix de la destination au choix de ses compagnons, LFTRIP vous assure une véritable expérience inédite !</p>
                            </Col>
                            <Col xs={12} md={6} className="px-5">
                                <Row>   
                                    <a href="/cgu">Conditions générales d'utilisation</a>
                                </Row>
                                <Row> 
                                    <a href="/politiquedeconf">Politique de confidentialité</a>
                                </Row>
                                <Row> 
                                    <a href="/concept">Notre concept</a>
                                </Row>
                                <Row> 
                                    <a href="/contact">Nous contacter</a>
                                </Row>
                                 
                            </Col>
                        </Row>
                        
                        <Row className="m-0">
                            <Col className="d-flex justify-content-center align-items-center pt-5 pb-5">LFTRIP 2020 - right reserved</Col>
                        </Row>
                    </Container>
                </Col>
            </Row>
        </Container>
    </div>
    
)

export default Footer;