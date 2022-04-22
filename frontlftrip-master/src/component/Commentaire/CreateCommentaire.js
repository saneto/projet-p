import React, { Component } from 'react';
import { compose } from 'recompose';
import { withBackaccessContext } from '../BackEnd';
import {withRouter} from 'react-router-dom';
import TextField from '@material-ui/core/TextField';
import Container from 'react-bootstrap/Container';
import Row from 'react-bootstrap/Row';
import Col from 'react-bootstrap/Col';
import Card from 'react-bootstrap/Card';
import SendIcon from '@material-ui/icons/Send';

const INITIAL_STATE = {
    uri:'comments',
    iduser :'',
    error: null,
    idtrip:'',
    token:'',
    body:'',
};

const user = JSON.parse(localStorage.getItem("authUser"));

class CreateComment extends Component {

    constructor(props) {
        super(props);
        this.state = { ...INITIAL_STATE };
        this.state.iduser = user.id;
        this.state.token = user.token;
        this.state.idtrip =  this.props.match.params.id;
    }
    onSubmit = event => {
        const { body } = this.state;
        //console.log(this.props.match.params.id)
        const id = this.props.match.params.id;
        this.props.backaccess
            .postDataWithId({data: body, id: id, token: this.state.token}, this.state.uri)
            .then(response => {
            //console.log(response.data.response.body)
                this.setState({ ...INITIAL_STATE });
                this.props.onComment(response.data.response);
            })
            .catch(error => {
                this.setState({ error });
            });

            event.preventDefault();
    };

    onChange = event => {
        this.setState({ [event.target.name]: event.target.value });
    };

    render() {
        return (
            <React.Fragment>
                <Container fluid>
                    <Row>
                        <Col>
                        <Card className="card-outline-secondary mb-4">
                        <Card.Header as="h5" className="trip-card-detail-header m-2">Ecrire un commentaire</Card.Header>
                        <Card.Body>
                            <TextField required fullWidth name="body" id="body" label="Mon commentaire" multiline rows="3" variant="outlined" value={this.state.body} onChange={this.onChange}/>
                        </Card.Body>
                        <Card.Footer className="d-flex justify-content-end align-items-center">
                            <button href="#" onClick={this.onSubmit} className="btn green-button">Commenter <SendIcon/></button>
                        </Card.Footer>
                            
                        </Card>
                        </Col>
                    </Row>
                </Container>

                
            </React.Fragment>
        
        );
    }
}

const CommentForm = compose(
    withRouter,
    withBackaccessContext,
)(CreateComment);

export default CommentForm;
