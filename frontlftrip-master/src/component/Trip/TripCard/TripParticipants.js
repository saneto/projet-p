import React  from 'react';

import Card from 'react-bootstrap/Card';
import Button from 'react-bootstrap/Button';
import Row from 'react-bootstrap/Row';
import Col from 'react-bootstrap/Col';

import SendIcon from '@material-ui/icons/Send';
import PersonIcon from '@material-ui/icons/Person';

class TripParticipants extends React.Component{

    constructor(props){
        super(props);
        this.state = {
        }
    }
    createChat = event =>{
        this.props.createChat(event)
     }

    createGroupeChat = () =>{
        if(this.props.participants.length > 0){
            let lUserId=""
            let lemail=""
            this.props.participants.forEach((participant, i) => {
                if(i < this.props.participants.length-1){
                    lUserId+= participant.id.toString()+","
                    lemail+= participant.email+","
                }else{
                    lUserId+= participant.id.toString()
                    lemail+= participant.email
                }

            })
            this.props.createGroupeChat(lUserId, lemail)
        }
    }

    render(){
        return (
            <React.Fragment>
                <Card className="mb-5">
                    <Card.Header as="h5" className="trip-card-detail-header">
                        <Col xs={12} md={7}className="d-flex justify-content-start align-items-center">
                            Ils y participent
                        </Col>
                           { (this.props.participants.length > 0)&&<Col xs={12} md={3}className="d-flex justify-content-end align-items-center">
                                <Button onClick={this.createGroupeChat} variant="light">Créer un groupe <SendIcon/></Button>
                            </Col>}
                    </Card.Header>
                    <Card.Body>
                        <Card.Body className="pb-2">
                            {this.props.participants.length === 0 &&
                                <div>
                                <p>Personne ne participe encore à ce trip =/</p>
                                <p>Soit le premier à y participer !</p>
                                </div>
                            }
                            {this.props.participants.length > 0 &&
                                
                                this.props.participants.map(participant =>
                                    <Row key={participant.id} className="guide-card mb-4 py-2">
                                        <Col xs={12} md={2} className="d-flex justify-content-center align-items-center">
                                            <a href={"../user-alike/" + participant.user.id}>
                                                {participant.user.avatar_path &&
                                                <img src={participant.user.avatar_path} className="rounded-circle avatar" height="80px" width="80px" alt="avatar"/>
                                                }
                                                {!participant.user.avatar_path &&
                                                <img src="/img/Default.png" className="rounded-circle avatar" height="80px" width="80px" alt="avatar"/>
                                                }
                                            </a>
                                        </Col>
                                        <Col xs={12} md={7} className="d-flex justify-content-start align-items-center">
                                            {participant.user.firstname} {participant.user.lastname}
                                        </Col>
                                        <Col xs={12} md={3} className="d-flex justify-content-end align-items-center">
                                            <Button value={participant.user.email} onClick={this.createChat} variant="light">Contacter <SendIcon/></Button>
                                            <Button href={"../user-alike/" + participant.user.id} variant="light"><PersonIcon/> Voir le profil</Button>
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
export default TripParticipants;