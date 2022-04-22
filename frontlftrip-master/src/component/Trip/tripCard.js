import React, {Component} from 'react';
import { withRouter } from 'react-router-dom';
import { compose } from 'recompose';
import { withBackaccessContext } from '../BackEnd';
import TripInformations from '../Trip/TripCard/TripInformations';
import TripGuide from '../Trip/TripCard/TripGuide'
import ListComments from '../Commentaire';
import TripParticipants from '../Trip/TripCard/TripParticipants';
import Container from 'react-bootstrap/Container';
import Row from 'react-bootstrap/Row';
import Col from 'react-bootstrap/Col';
import Tabs from 'react-bootstrap/Tabs';
import Tab from 'react-bootstrap/Tab';
import Button from "react-bootstrap/Button";
import { Card } from 'react-bootstrap';
import FavoriteIcon from '@material-ui/icons/Favorite';
import AddIcon from '@material-ui/icons/Add';
import IconButton from '@material-ui/core/IconButton';
import CancelIcon from '@material-ui/icons/Cancel';
import InfoIcon from '@material-ui/icons/Info';
import CommentIcon from '@material-ui/icons/Comment';
import BookmarksIcon from '@material-ui/icons/Bookmarks';
import PeopleIcon from '@material-ui/icons/People';
import { withAuthorization } from '../Session';

/* const styles = {
div:{
    display: 'flex',
    flexFlow: 'wrap',
    justifyContent: 'center',
    alignItems: 'center',
}
};*/
const user = ( localStorage.getItem("authUser")) ?  JSON.parse(localStorage.getItem("authUser"))  : [];
const token = ( localStorage.getItem("token")) ?  localStorage.getItem("token")  : '';


function ControlledTabs(data) {
    const [key, setKey] = React.useState('home');
    return (
      <Tabs
        id="controlled-tab-example"
        activeKey={key}
        onSelect={(k) => setKey(k)}
      >
        <Tab eventKey="home" title={<span><InfoIcon/>{" Détails"}</span>}>
            <TripInformations trip={data.trip} author={data.trip.author}/>
        </Tab>
        <Tab eventKey="profile" title={<span><CommentIcon/>{" Commentaires"}</span>}>
            <ListComments/>
        </Tab>
        <Tab eventKey="guide" title={<span><BookmarksIcon/>{" Guides"}</span>}>
            <TripGuide createChat={data.createChat} guide={data.guide} />
        </Tab>
        <Tab eventKey="participant" title={<span><PeopleIcon/>{" Participants"}</span>}>
            <TripParticipants createGroupeChat={data.createGroupeChat} createChat={data.createChat} participants={data.participants}/>
        </Tab>
      </Tabs>
    );
  }

class TripDetails extends Component{
    constructor(props){
        super(props);
        this.state = {
            trip :[],
            uri: 'trips',
            id: '',
            uri2:'likes',
            uri3:'participation',
            participation:[],
            likes: [],
            updated: null,
            participated: null,
            token: '',
            iduser:'',
            likeid: 0,
            img : "",
            debutDateForUser : "",
            endDateForUser : "",
            likedCount : "",
            participatedCount : "",
            userGuide : [],
        }
        this.state.id =  this.props.match.params.id;
        this.state.iduser = user.id;
        this.state.token = token;
    }
    componentDidMount() {
        // get trip details
        this.props.backaccess
        .getDataWitId(this.state.id,this.state.uri)
        .then((data) => {
            this.setState({ trip: data.response});
            this.props.backaccess
            .getGuideListForTrip(this.state.trip)
            .then((res) => {
                this.setState( {userGuide: res.data.response})
            }).catch(error => {
                console.log(error);
                
            });
        }).catch(error => {
            console.log("not ok");
            this.setState({ error });
        });
        // get likes
        this.props.backaccess
        .getDataWitId(this.state.id,this.state.uri2)
        .then((data) => {
            this.setState({ likes: data.response, img: this.state.trip.base64 });
            this.setState({ likedCount: this.state.likes.length });
            for(let i = 0 ; i < this.state.likes.length; i++){
                if(this.state.likes[i].user_id === this.state.iduser){
                    this.setState({ updated: true});
                }else{
                    this.setState({ updated: false});
                }
            }
        }).catch(error => {
            this.setState({ error });
        });
        // get participation
        this.props.backaccess
        .getDataWitId(this.state.id,this.state.uri3)
        .then((data) => {
            this.setState({ participation: data.response});
            this.setState({ participatedCount: this.state.participation.length });
            for(let i = 0 ; i < this.state.participation.length; i++){
                if(this.state.participation[i].user_id === this.state.iduser){
                    this.setState({ participated: true});
                }else{
                    this.setState({ participated: false});
                }
            }
        }).catch(error => {
            this.setState({ error });
        });
    };

    actualisePariticpants = () => {
        this.props.backaccess
        .getDataWitId(this.state.id,this.state.uri3)
        .then((data) => {
            this.setState({ participation: data.response});
            this.setState({ participatedCount: this.state.participation.length });
            for(let i = 0 ; i < this.state.participation.length; i++){
                if(this.state.participation[i].user_id === this.state.iduser){
                    this.setState({ participated: true});
                }else{
                    this.setState({ participated: false});
                }
            }
        }).catch(error => {
            this.setState({ error });
        });
    };

    updateLikes = () => {
        if(!this.state.updated) {
            this.props.backaccess
            .postLikeWithId({iduser: this.state.iduser, id: this.state.id, token: this.state.token}, this.state.uri2)
            .then((data) => {
                this.setState({ likes: data.response, likedCount: this.state.likedCount +1});
                //console.log(this.state.likes.id);
            }).catch(error => {
                console.log("not ok");
                this.setState({ error });
            });
            this.setState((prevState, props) => {
                return {
                //likes: prevState.likes + 1,
                updated: true
                };
            });
        } else {
            this.props.backaccess
            .deleteLikeWithId({iduser: this.state.iduser, id: this.state.id, token: this.state.token}, this.state.uri2)
            .then((data) => {
                this.setState({ likes: data.response, likedCount: this.state.likedCount -1});
                //console.log(this.state.likes);
            }).catch(error => {
                console.log("not ok");
                this.setState({ error });
            });
            this.setState((prevState, props) => {
                return {
                //likes: prevState.likes - 1,
                updated: false
                };
            });

            }
    }

    updateParticiper = () => {
        if(!this.state.participated) {
            this.props.backaccess
            .postLikeWithId({iduser: this.state.iduser, id: this.state.id, token: this.state.token}, this.state.uri3)
            .then((data) => {
                this.actualisePariticpants();
            }).catch(error => {
                console.log("not ok");
                this.setState({ error });
            });
            this.setState((prevState, props) => {
                return {
                participated: true
                };
            });
        } else {
            this.props.backaccess
            .deleteLikeWithId({iduser: this.state.iduser, id: this.state.id, token: this.state.token}, this.state.uri3)
            .then((data) => {
                this.actualisePariticpants();
            }).catch(error => {
                console.log(error);
                this.setState({ error });
            });
            this.setState((prevState, props) => {
                return {
                participated: false
                };
            });

            }
    }
    
    deleteTrip = (tripId) => event => {
        this.props.backaccess.deleteUserTrip({id : tripId}, this.state.uri)
            .then(() => {
                //console.log('ok delete');
            }).catch(error => {
            this.setState({ error });
        });
        event.preventDefault();
    };

    getBackGoDateToUserFriendlyDate = (getdate) => {
		var date = new Date(getdate);
		var year = date.getFullYear();
		var month = date.getMonth()+1;
		var day = date.getDate();
        var textMonth = ''
		if (day < 10) {
		day = '0' + day;
		}
		if (month < 10) {
		month = '0' + month;
        }
        
        switch (month){
            case '01':
                textMonth = "Janvier";
                break;
            case '02':
                textMonth = "Février";
                break;
            case '03':
                textMonth = "Mars";
                break;
            case '04':
                textMonth = "Avril";
                break;
            case '05':
                textMonth = "Mai";
                break;
            case '06':
                textMonth = "Juin";
                break;
            case '07':
                textMonth = "Juillet";
                break;
            case '08':
                textMonth = "Aout";
                break;
            case '09':
                textMonth = "Septembre";
                break;
            case 10:
                textMonth = "Octobre";
                break;
            case 11:
                textMonth = "Novembre";
                break;
            case 12:
                textMonth = "Décembre";
                break;
            default:
                break;
        }
        
		var formatedDate = day+' ' + textMonth + ' '+year;
		return formatedDate;
	}

    createGroupeChat = (data, email) =>{
		this.props.backaccess
            .postData({
                tripId: String(this.state.trip.id),
                tripName:this.state.trip.title,
                name: user.lastname,
                particiapantID: data,
                emails:email
            },"groupc")
			.then((data) => {
                //console.log(this.props)
				if(data.data.response){
                    this.props.history.push({
                        pathname: '/message',
                        state: { id: data.data.response.id }
                        })
				}
			})
			.catch(error => {
				this.setState({ error });
			});
    }

    creationChat = event =>{
        //console.log( event.target.value)
		this.props.backaccess
			.postData({email : event.target.value },"chat")
			.then((data) => {
                //console.log(data)
				if(data.data.response){
                    this.props.history.push({
                        pathname: '/message',
                        state: { id: data.data.response.id }
                      })
				}
			})
			.catch(error => {
				this.setState({ error });
			});
		
	 }


    render()
    {
        let img = "";
        //console.log(this.state.img)
        if(this.state.trip.base64){
            img = <img className="card-img-top img-fluid" src={this.state.trip.base64} alt=""/>
        }else{
            img =  <img className="card-img-top img-fluid" src="/img/card1.jpg" value={this.state.img} alt=""/>
        }

        let url = "/updatetrip/"+this.state.id;
        //let url2 = "/listeparticipants/"+this.state.id; /* NOT USED - CLEAR WARNING */
        //const {likes} = this.state; /* NOT USED - CLEAR WARNING */

        return(
            <React.Fragment>
                
                <Container>
                    <Row className="mt-5">
                        <Container><h3>{this.state.trip.title}</h3></Container>
                        <Container className="margin-top-negative">
                            <Row>
                                <Col>
                                    <IconButton  color="secondary" aria-label="liked">
                                        <FavoriteIcon />
                                    </IconButton>
                                    {this.state.likedCount} personnes aiment le voyage
                                </Col>
                                <Col className="d-flex justify-content-end align-items-center">
                                    {this.state.participated &&
                                        <IconButton color="secondary" aria-label="participate">
                                            <CancelIcon />
                                        </IconButton>
                                    }
                                    {this.state.participated &&
                                        <Button variant="link" onClick={this.updateParticiper} className="participate-link">
                                            Ce voyage ne m'intéresse plus
                                            {/* Je ne participe plus à ce Trip */}
                                        </Button>
                                    }
                                    {!this.state.participated && 
                                        <IconButton color="primary" aria-label="participate">
                                            <AddIcon />
                                        </IconButton>
                                    }
                                    {!this.state.participated && 
                                        <Button variant="link" onClick={this.updateParticiper} className="participate-link">
                                            Le voyage m'intéresse !
                                            {/* Je participe à ce Trip !*/}
                                        </Button>
                                    }
                                </Col>
                            </Row>
                        
                        </Container>
                    </Row>
                </Container>
                <Container fluid className="body-page mt-3 pt-3">
                    <Container>
                        <Row>
                            <Col md={8}>
                                <Card>
                                    {img}
                                </Card>
                            </Col>
                            <Col md={4}>
                                <Card>
                                    <Card.Header className="card-trip-overview">{this.state.participatedCount} tripiziens ont postulé au voyage</Card.Header>
                                    <Card.Body key={this.state.trip.id}>
                                        <h5><u>Destination</u></h5>
                                        <p className="card-text"><b>{this.state.trip.country}</b></p>
                                        <h5><u>Date de départ</u></h5>
                                        <p className="card-text"><b>{this.getBackGoDateToUserFriendlyDate(this.state.trip.start_date)}</b></p>
                                        <h5><u>Date de retour</u></h5>
                                        <p className="card-text"><b>{this.getBackGoDateToUserFriendlyDate(this.state.trip.end_date)}</b></p>
                                        {this.state.participation.length > 0 && this.state.participation.length < 4 &&
                                            <div>
                                                <h5><u>Les participants</u></h5>
                                                {this.state.participation.map(participant =>(
                                                    <span key={participant.id}>
                                                    {participant.user.avatar_path &&
                                                    <img src={participant.user.avatar_path} className="rounded-circle avatar" height="30px" width="30px" alt="avatar"/>
                                                    }
                                                    {!participant.user.avatar_path &&
                                                    <img src="/img/Default.png" className="rounded-circle avatar" height="30px" width="30px" alt="avatar"/>
                                                    }                                        
                                                    </span>
                                                ))}
                                                {/* <div>
                                                    <Link to ={url2}>voir la liste des participants</Link>
                                                </div>  */}                                           
                                            </div>
                                        }
                                        {this.state.participation.length > 0 && this.state.participation.length > 3 &&
                                            <div>
                                                {this.state.participation.slice(0,3).map(participant =>(
                                                    <span key={participant.id}>
                                                    {participant.user.avatar_path &&
                                                    <img src={participant.user.avatar_path} className="rounded-circle avatar" height="30px" width="30px" alt="avatar"/>
                                                    }
                                                    {!participant.user.avatar_path &&
                                                    <img src="/img/Default.png" className="rounded-circle avatar" height="30px" width="30px" alt="avatar"/>
                                                    }                                        
                                                    </span>
                                                ))}
                                                <span>
                                                    {" et "}{this.state.participation.length-3} 
                                                    {this.state.participation.length-3 === 1 && " autre personne"}
                                                    {this.state.participation.length-3 !== 1 && " autres personnes"}
                                                </span>
                                                {/* <div>
                                                    <Link to ={url2}>Voir la liste des participants</Link>
                                                </div> */}
                                            </div>
                                        }
                                    </Card.Body>
                                </Card>
                            </Col>
                        </Row>
                        <Row>
                            <Col xs={12} md={2}>
                                {this.state.updated &&
                                    <span>
                                        <IconButton onClick={this.updateLikes} color="secondary" aria-label="add to favorites">
                                            <FavoriteIcon />
                                        </IconButton>
                                        {"Je n'aime pas"}
                                    </span>
                                }
                                {!this.state.updated &&
                                    <span>
                                        <IconButton onClick={this.updateLikes} aria-label="add to favorites">
                                            <FavoriteIcon />
                                        </IconButton>
                                        {"J'aime"}
                                    </span>
                                }
                            </Col>
                            <Col xs={12} md={6} className="d-flex justify-content-end align-items-center">
                                {user && this.state.trip.author_id === this.state.iduser &&
                                    <div>
                                        <Button href={url} className="blue-button mr-1">Modifier</Button>
                                        <Button variant="danger" onClick={this.deleteTrip(this.state.trip.id)}>Supprimer le voyage</Button>
                                </div>
                                }
                            </Col>
                        </Row>
                        <Row className="mt-3">
                            <Col>
                                <ControlledTabs createGroupeChat={this.createGroupeChat} createChat={this.creationChat} trip={this.state.trip} guide={this.state.userGuide} participants={this.state.participation}/>
                            </Col>
                        </Row>
                    </Container>
                </Container>
            </React.Fragment>
        )
    }
}

const condition = authUser => !!authUser;

export default compose(
    withAuthorization(condition),
    withRouter,
    withBackaccessContext,
)(TripDetails);
