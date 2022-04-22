import React from 'react';
import { compose } from 'recompose';
import { withBackaccessContext } from '../../BackEnd';
import { withRouter  } from 'react-router-dom';
import Card from 'react-bootstrap/Card';
import Button from 'react-bootstrap/Button';
import Row from 'react-bootstrap/Row';
import Col from 'react-bootstrap/Col';
import SendIcon from '@material-ui/icons/Send';
import AccountCircleIcon from '@material-ui/icons/AccountCircle';

class TripGuide extends React.Component{

    constructor(props){
        super(props);
        this.state = {
        }
    }    
    
    createChat = event =>{
        this.props.createChat(event)
     }

    render(){
        return (
            <React.Fragment>
                <Card className="mb-5">
                    <Card.Header as="h5" className="trip-card-detail-header">Les guides disponibles pour la destination</Card.Header>
                    <Card.Body>
                        <Card.Body className="pb-2">
                            {this.props.guide.length === 0 &&
                                "Il n'y a pas encore de guide au sein de la communauté pour ce pays =/"
                            }
                            {this.props.guide.length > 0 &&
                                
                                this.props.guide.map(user =>
                                    <Row key={user.id} className="guide-card mb-4 py-2">
                                        <Col xs={12} md={2} className="d-flex justify-content-center align-items-center">
                                            {user.avatar_path &&
                                            <img src={user.avatar_path} className="rounded-circle avatar" height="80px" width="80px" alt="avatar"/>
                                            }
                                            {!user.avatar_path &&
                                            <img src="/img/Default.png" className="rounded-circle avatar" height="80px" width="80px" alt="avatar"/>
                                            }
                                        </Col>
                                        <Col xs={12} md={7} className="d-flex justify-content-center justify-content-md-start align-items-center">
                                            {user.firstname} {user.lastname}
                                        </Col>
                                    
                                            
                                        <Col xs={12} md={3} className="d-flex justify-content-center justify-content-md-end align-items-center">
                                            <Button value={user.email} onClick={this.createChat} variant="light">Contacter <SendIcon/></Button>
                                            <Button variant="light" href={`/user-alike/${user.id}`}>Voir le profil <AccountCircleIcon/></Button>
                                        </Col>
                                    </Row>
                                )                              
                            }
                            
                        </Card.Body>
                        
                     
                         
                    </Card.Body>
                </Card>
            </React.Fragment>
        );
    }
}
export default compose(
    withRouter,
    withBackaccessContext,
)(TripGuide);