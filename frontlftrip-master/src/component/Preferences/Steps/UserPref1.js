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



class UserPref1 extends Component{
    constructor(props){
        super(props);
        this.state = {
            ShyOutgoing: "",
            CalmOrSpark:"",
            SundayMikeHornAdventurer:"",
            AthleteSportKezako:"",
            CalmOrSpark:"",
            FatMorningReadyLeave7AM:"",
            scootered_disorganized:"",
            trekking_beach_towel :"",
            
        }
    }
    onChange = event => {
        this.setState({ [event.target.name]: event.target.value });
    };

    render()
    {
        const { CalmOrSpark, ShyOutgoing, SundayMikeHornAdventurer, 
            AthleteSportKezako, FatMorningReadyLeave7AM, scootered_disorganized, 
            trekking_beach_towel } = this.state;

        return(
            <Container fluid >
                <div className="d-flex padding-top-2 padding-bottom-10 justify-content-center align-items-center">
                    <Row className="justify-content-md-center">
                            <Col>
                                <Row  className="justify-content-md-center">
                                    <Col  md="auto">
                                        <div className="section-title line-style no-margin">
                                                <h3 className="title">Personanalité</h3>
                                                <FormControl component="fieldset">
                                                    <RadioGroup row aria-label="position" name="position" value={ShyOutgoing} onChange={this.onChange}>
                                                        <FormControlLabel value="Timide" control={<Radio color="default" />} label="Timide" labelPlacement="start"/>
                                                        <FormControlLabel value="Extraverti" control={<Radio color="default" />} label="Extraverti" />
                                                    </RadioGroup>
                                                    <RadioGroup row aria-label="position" name="position" value={SundayMikeHornAdventurer} onChange={this.onChange}>
                                                    <FormControlLabel value="Aventurier comme indiana" control={<Radio color="default" />} label="Aventurier comme indiana" labelPlacement="start"/>
                                                    <FormControlLabel value="Aventurier du dimanche" control={<Radio color="default" />} label="Aventurier du dimanche" />
                                                </RadioGroup>
                                                <RadioGroup row aria-label="position" name="position" value={AthleteSportKezako} onChange={this.onChange}>
                                                    <FormControlLabel value="Le sport ? Ca se mange ?" control={<Radio color="default" />} label="Le sport ? Ca se mange ?" labelPlacement="start"/>
                                                    <FormControlLabel value="Le sport c'est votre dada" control={<Radio color="default" />} label="Le sport c'est votre dada" />
                                                </RadioGroup>
                                                <RadioGroup row aria-label="position" name="position" value={CalmOrSpark} onChange={this.onChange}>
                                                    <FormControlLabel value="Un vrai piment de Cayenne" control={<Radio color="default" />} label="Un vrai piment de Cayenne" labelPlacement="start"/>
                                                    <FormControlLabel value="Une petite perle de lait" control={<Radio color="default" />} label="Une petite perle de lait" />
                                                </RadioGroup>
                                            </FormControl>
                                            <h3 className="title">Rythme de vie</h3>
                                                <FormControl component="fieldset">
                                                    <RadioGroup row aria-label="position" name="position" value={FatMorningReadyLeave7AM} onChange={this.onChange}>
                                                        <FormControlLabel value="Une vrai marmotte" control={<Radio color="default" />} label="Une vrai marmotte" labelPlacement="start"/>
                                                        <FormControlLabel value="Frais et paré à partir à 7h12" control={<Radio color="default" />} label="Frais et paré à partir à 7h12" />
                                                    </RadioGroup>
                                                <RadioGroup row aria-label="position" name="position" value={scootered_disorganized} onChange={this.onChange}>
                                                    <FormControlLabel value="Organisé comme un scout" control={<Radio color="default" />} label="Orgnisé comme un scout" labelPlacement="start"/>
                                                    <FormControlLabel value="Désorganisé depuis que je suis né" control={<Radio color="default" />} label="Désorganisé depuis que je suis né" />
                                                </RadioGroup>
                                            </FormControl>
                                            <h3 className="title">Activités</h3>
                                                <FormControl component="fieldset">
                                                    <RadioGroup row aria-label="position" name="position" value={trekking_beach_towel} onChange={this.onChange}>
                                                        <FormControlLabel value="Aventurier de moyenne compétition" control={<Radio color="default" />} label="Aventurier de moyenne compétition" labelPlacement="start"/>
                                                        <FormControlLabel value="Serviette-Plage-Plage-Serviette" control={<Radio color="default" />} label="Serviette-Plage-Plage-Serviette" />
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
)(UserPref1);



