import React, { Component } from 'react';
import {compose} from "recompose";
import {Link, withRouter} from "react-router-dom";
import {withBackaccessContext} from "../BackEnd";
import "bootstrap/dist/css/bootstrap.min.css";
import './listFellowUsers.css';
import Modal from "react-bootstrap/Modal";
import Button from "react-bootstrap/Button";
import Col from "react-bootstrap/Col";
import {Form} from "react-bootstrap";
import SendIcon from '@material-ui/icons/Send';
import AccountCircleIcon from '@material-ui/icons/AccountCircle';

const user = JSON.parse(localStorage.getItem("authUser"));

class ListFellowUsers extends Component {
    constructor(props) {
        super(props);
        this.state = {
            uri:"send_proposal",
            uri2:"user_userpreferences",
            token: '',
            filteredUser: this.props.filteredUser,
            age: '',
            openModal: false,
            proposition: '',
            avatar_path: ''
        };
        this.state.token = user.token;
        this.state.currentUser = user.id;;
        this.state.avatar_path = user.avatar_path
    }

    calculate_age(date) {
        const dateTime = new Date(date);
        const diff_ms = Date.now() - dateTime.getTime();
        const age_dt = new Date(diff_ms);
        return Math.abs(age_dt.getUTCFullYear() - 1970)
    }

    openModal = event => {
        this.setState({openModal: true});
    };
    handleClose = event => {
        this.setState({openModal: false});
    };
    onChange = event => {
        this.setState({ [event.target.name]: event.target.value });
    };
    sendProposition = event => {
        const {proposition, token} = this.state;
        this.props.backaccess.sendTripProposition({firstname: this.props.filteredUser.user.firstname, message: proposition, email: this.props.filteredUser.user.email, token : token}, this.state.uri)
            .then(() => {
                this.setState({openModal: false});
            }).catch(error => {
            this.setState({ error });
        });
        event.preventDefault();
    };

    render() {
        let url = "/user-alike/"+this.props.filteredUser.user_id;
        let img= "";
        if(this.props.filteredUser.user.avatar_path){
            img = <img src={this.props.filteredUser.user.avatar_path} className="round" height="120px" alt="user"/>
        }else{
            img = <img src="/img/Default.png" className="round" height="120px" alt="user"/>
        }
        return (
            <React.Fragment>
                <div className="col-md-auto col-xs-auto">
                    <div className="card-container">
                        {img}
                        <h3>{this.props.filteredUser.user.firstname} {this.props.filteredUser.user.lastname}</h3>
                        <h6>{this.calculate_age(this.props.filteredUser.user.date_of_birth)} ans, {this.props.filteredUser.user.city}</h6>
                        <h6>{this.props.filteredUser.user.sexe}</h6>

                        <div className="buttons pb-3 mt-3">
                            <Link to={url}>
                                <Button variant="light" className="mr-2">Mon profil <AccountCircleIcon/></Button>
                                {/* <button className="primary mr-2">
                                    <strong>Profil</strong>
                                </button> */}
                            </Link>
                            {/* <button className="primary ml-1" onClick={this.openModal}>
                                <strong>Tripper</strong>
                            </button> */}
                            <Button variant="light" onClick={this.openModal} className="ml-1">Un voyage ? <SendIcon/></Button>
                        </div>
                    </div>
                </div>
                <Modal show={this.state.openModal} onHide={this.handleClose} animation={false} aria-labelledby="contained-modal-title-vcenter" centered>
                    <Modal.Header closeButton>
                        <Modal.Title>Proposition de voyage</Modal.Title>
                    </Modal.Header>
                    <Modal.Body>
                        <Form.Group as={Col} controlId="formGridEmail">
                            <Form.Control as="textarea" rows="3" name="proposition" onChange={this.onChange} placeholder="En visitant votre profil et en voyant vos traits de personnalités, ..." />
                        </Form.Group>
                    </Modal.Body>
                    <Modal.Footer>
                        <Button variant="outline-danger" onClick={this.handleClose}>
                            Annuler
                        </Button>
                        <Button variant="success" onClick={this.sendProposition}>
                            Proposer un voyage
                        </Button>
                    </Modal.Footer>
                </Modal>
            </React.Fragment>
        )
    }
}

export default  compose(
    withRouter,
    withBackaccessContext,
)(ListFellowUsers);

