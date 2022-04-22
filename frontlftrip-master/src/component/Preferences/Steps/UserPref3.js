import React, {Component} from 'react';
import {Row, Container, Button, Col, Form} from 'react-bootstrap';
import { withRouter } from 'react-router-dom';
import { compose } from 'recompose';
import { withBackaccessContext } from '../../BackEnd';
import Radio from '@material-ui/core/Radio';
import RadioGroup from '@material-ui/core/RadioGroup';
import FormControlLabel from '@material-ui/core/FormControlLabel';
import FormControl from '@material-ui/core/FormControl';
import FormLabel from '@material-ui/core/FormLabel';


class UserPref3 extends Component{
    constructor(props){
        super(props);
        this.state = {
            food:"",
        }
    }
    onChange = event => {
        this.setState({ [event.target.name]: event.target.value });
    };

    render()
    {
        const { food } = this.state.food;
        return(
            <Container fluid >
                <div className="d-flex padding-top-2 padding-bottom-10 justify-content-center align-items-center">
                    <Row className="justify-content-md-center">
                            <Col>
                                <Row  className="justify-content-md-center">
                                    <Col  md="auto">
                                    <div className="section-title line-style no-margin">
                                                <h3 className="title">Régime alimentaire</h3>
                                                <FormControl component="fieldset">
                                                    <RadioGroup row aria-label="position" name="position" value={food} onChange={this.onChange}>
                                                        <FormControlLabel value="Halal" control={<Radio color="default" />} label="Halal"/>
                                                        <FormControlLabel value="Vegan" control={<Radio color="default" />} label="Vegan" />
                                                        <FormControlLabel value="Végétarien" control={<Radio color="default" />} label="Végétarien" />
                                                    </RadioGroup>
                                            </FormControl>
                                        </div>
                                    </Col>
                                </Row>
                            </Col>
                        </Row>
                </div>
            </Container>
        )
    }
}

export default  compose(
    withRouter,
    withBackaccessContext,
)(UserPref3);



