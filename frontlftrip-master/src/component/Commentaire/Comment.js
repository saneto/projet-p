import React, {Component} from 'react';
import Dropdown from "react-bootstrap/Dropdown";
import { withRouter  } from 'react-router-dom';
import { compose } from 'recompose';
import { withBackaccessContext } from '../BackEnd';
import Container from "react-bootstrap/Container";
import Form from "react-bootstrap/Form";
import Row from "react-bootstrap/Row";
import Button from "react-bootstrap/Button";
import Col from "react-bootstrap/Col";
import MoreHorizRoundedIcon from '@material-ui/icons/MoreHorizRounded';
import Card from 'react-bootstrap/Card';
import IconButton from '@material-ui/core/IconButton';
import DeleteIcon from '@material-ui/icons/Delete';
import CreateIcon from '@material-ui/icons/Create';
import EditIcon from '@material-ui/icons/Edit';
import CancelIcon from '@material-ui/icons/Cancel';

const user = JSON.parse(localStorage.getItem("authUser"));

const CustomToggle = React.forwardRef(({ children, onClick }, ref) => (
    <span
      
      ref={ref}
      onClick={e => {
        e.preventDefault();
        onClick(e);
      }}
    >
      {
          <IconButton>
          <MoreHorizRoundedIcon/>
          </IconButton>
      /* Render custom icon here */}
      
    </span>
  ));

class Comment extends Component{
    constructor(props) {
        super(props);
        //console.log(this.props);
        this.state = {
            uri:'comments',
            token: '',
            comment: this.props.comment,
            currentUser: '',
            editCom: false,
            body: ''
        };
        this.state.token = user.token;
        this.state.currentUser = user.id;;
    }

    editCom = event => {
        this.setState({editCom: true});
        event.preventDefault();
    };
    onSubmit = (commentId) => event => {
        const { body } = this.state;
        this.props.backaccess
            .editCommentary({data: body, id: commentId, token: this.state.token}, this.state.uri)
            .then(response => {
                this.setState({editCom: false});
                this.setState({comment: response.data.response});
            })
            .catch(error => {
                this.setState({ error });
            });

        event.preventDefault();
    };

    onChange = event => {
        this.setState({ [event.target.name]: event.target.value });
    };

    render()
    {
        const {comment, currentUser, editCom} = this.state;
        return(
            <React.Fragment>
            <Card.Body key={comment.id} className="pb-1">
                {editCom && <Col className="d-flex justify-content-start align-items-start">
                    <Container>
                        <Form>
                            <Row>
                                <Col lg={10}><Form.Control name="body" type="text" as="textarea" rows="2" defaultValue={comment.body} onChange={this.onChange}/></Col>
                                <Col lg={2} className="d-flex justify-content-center align-items-center">
                                    <Row><Button className="button-link-black" variant="link" onClick={this.onSubmit(this.props.id)}><EditIcon/> Modifier</Button></Row>
                                    <Row><Button className="button-link-red" variant="link" onClick={() => this.setState({editCom: false})}><CancelIcon/> Annuler</Button></Row>
                                </Col>
                            </Row>
                            
                        </Form>
                    </Container>
                </Col>}
                {!editCom && 
                <Row>
                    <Col xs={10}>{comment.body}</Col>
                    <Col xs={2} className="d-flex justify-content-end align-items-center">
                    {user && comment.user_id === currentUser &&
                        <Dropdown>
                            <Dropdown.Toggle as={CustomToggle} id="dropdown-custom-components">   
                            </Dropdown.Toggle>
                            <Dropdown.Menu>
                                <Dropdown.Item onClick={this.editCom}><CreateIcon/>Modifier</Dropdown.Item>
                                <Dropdown.Item onClick={this.props.deleteCom(this.props.id)}><DeleteIcon/>Supprimer</Dropdown.Item>    
                            </Dropdown.Menu>
                        </Dropdown>                
                    }
                    </Col>
                </Row>
                }
                <small className="text-muted">Posté par {comment.user.firstname} {comment.user.lastname} le {comment.created_at.substring(0, 10)} {" "}
                    {comment.user.avatar_path &&
                        <img src={user.avatar_path} className="rounded-circle avatar" height="20px" width="20px" alt="avatar"/>
                        }
                        {!comment.user.avatar_path &&
                        <img src="/img/Default.png" className="rounded-circle avatar" height="20px" width="20px" alt="avatar"/>
                    }</small>
                <hr  className="mb-0"/>
            </Card.Body>
            </React.Fragment>
        )
    }
}

export default  compose(
    withRouter,
    withBackaccessContext,
)(Comment);
