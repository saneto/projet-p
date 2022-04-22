import React, { Component } from "react";
import withStyles from "@material-ui/styles/withStyles";
import { withRouter } from "react-router-dom";
import CssBaseline from "@material-ui/core/CssBaseline";
import Typography from "@material-ui/core/Typography";
import Grid from "@material-ui/core/Grid";
import Paper from "@material-ui/core/Paper";
import Button from "@material-ui/core/Button";
import Stepper from "@material-ui/core/Stepper";
import Step from "@material-ui/core/Step";
import StepLabel from "@material-ui/core/StepLabel";
import {Row, Container, Col} from 'react-bootstrap';
import FormControlLabel from "@material-ui/core/FormControlLabel";
import Radio from '@material-ui/core/Radio';
import RadioGroup from '@material-ui/core/RadioGroup';
import FormControl from "@material-ui/core/FormControl";
import { withBackaccessContext } from '../BackEnd';
import { compose } from 'recompose';
import AccountProfil from '../Account/AccountProfil';

const styles = theme => ({
root: {
    flexGrow: 1,
    overflow: "hidden",
    backgroundSize: "cover",
    backgroundPosition: "0 400px",
    marginTop: 10,
    padding: 20,
    paddingBottom: 200
},
grid: {
    margin: `0 2px`
},
smallContainer: {
    width: "60%"
},
bigContainer: {
    width: "80%"
},
stepContainer: {
    display: "flex",
    flexDirection: "column",
    alignItems: "center"
},
stepGrid: {
    width: "80%"
},
backButton: {
    marginRight: 1
},
outlinedButtom: {
    textTransform: "uppercase",
    margin: 1
},
stepper: {
    backgroundColor: "transparent",
},
paper: {
    padding: 3,
    textAlign: "left",
    color: "secondary"
},
topInfo: {
    display: "flex",
    justifyContent: "space-between",
    alignItems: "center",
    marginBottom: 42
},
formControl: {
    width: "100%"
},
selectEmpty: {
    marginTop: 2
},
borderColumn: {
    borderBottom: `1px solid grey["100"]`,
    paddingBottom: 24,
    marginBottom: 24
},
flexBar: {
    marginTop: 32,
    display: "flex",
    justifyContent: "center"
},
radio: {
    marginLeft: 3
},
});

const getSteps = () => {
return ["Informations Personnelles", "Type d'utilisateur", "Personnalité & Activités", "Hébergement", "Régime alimentaire", "Terminé"];
};

let data= "";
if (localStorage.getItem(('authUser'))) {
    data = JSON.parse(localStorage.getItem('authUser'));
}

class Wizard extends Component {
    constructor(props){
        super(props);
        this.state = {
            activeStep: 0,
            labelWidth: 0,
            uri:"userpreferences",
            uri2:"user_userpreferences",
            token:data.token,
            userpref:[],
            id:"",
            iduser:data.id,
            traveler_type:"",
            ShyOutgoing: "",
            CalmOrSpark:"",
            SundayMikeHornAdventurer:"",
            AthleteSportKezako:"",
            FatMorningReadyLeave7AM:"",
            scootered_disorganized:"",
            trekking_beach_towel :"",
            LuxuryCharmingLodging : "",
            food:"",
        };
    }

componentDidMount() {
    this.props.backaccess
    .getDataWitId(this.state.iduser,this.state.uri2)
    .then((data) => {
        this.setState({ userpref: data.response[0]});
        this.setState({ id: data.response[0].id});
        //console.log();
    }).catch(error => {
        console.log("not ok");
        this.setState({ error });
    });
}
handleNext = () => {
    this.setState(state => ({
    activeStep: state.activeStep + 1
    }));
};

handleBack = () => {
    this.setState(state => ({
    activeStep: state.activeStep - 1
    }));
};

handleReset = () => {
    this.setState({
    activeStep: 0
    });
};

handleChange = event => {
    this.setState({ [event.target.name]: event.target.value });
};

onChange = event => {
    this.setState({ [event.target.name]: event.target.value });
};



stepActions() {
    if (this.state.activeStep === 4) {
    return "Valider";
    }
    if (this.state.activeStep === 5) {
    return "Aller à la page d'accueil";
    }
    return "Suivant";
}

goToDashboard = event => {
    const queryString = this.props.location.search;
    this.props.history.push({
    pathname: "/",
    search: queryString
    });
};

onSubmit = () =>{
    let data = {
        "traveler_type":this.state.traveler_type,
        "shy_outgoing":this.state.ShyOutgoing,
        "calm_or_spark":this.state.CalmOrSpark,
        "sunday_mike_horn_adventurer":this.state.SundayMikeHornAdventurer,
        "athlete_sport_kezako":this.state.AthleteSportKezako,
        "luxury_charming_lodging":this.state.LuxuryCharmingLodging,
        "fat_morning":this.state.FatMorningReadyLeave7AM,
        "food":this.state.food,
        "scootered_disorganized":this.state.scootered_disorganized,
        "trekking_beach_towel":this.state.trekking_beach_towel,
    }
    //console.log(data);
    this.props.backaccess.updateUserPreferenceData(data, this.state.uri, this.state.id)
    .then( ret =>{
            //console.log(ret);
            this.setState(state => ({
                activeStep: state.activeStep + 1
                }))
        }
    ).catch( error =>{

    })
}

render() {
    const { classes } = this.props;
    const steps = getSteps();
    const { traveler_type } = this.state.traveler_type;
    const { activeStep } = this.state;
    const {CalmOrSpark} = this.state.CalmOrSpark;
    const {ShyOutgoing} = this.state.ShyOutgoing;
    const {SundayMikeHornAdventurer} = this.state.SundayMikeHornAdventurer;
    const {AthleteSportKezako} = this.state.AthleteSportKezako;
    const {FatMorningReadyLeave7AM} = this.state.FatMorningReadyLeave7AM;
    const {scootered_disorganized} = this.state.scootered_disorganized;
    const {trekking_beach_towel} = this.state.trekking_beach_towel;
    const {LuxuryCharmingLodging} = this.state.LuxuryCharmingLodging;
    const {food} = this.state.food;

    return (
    <React.Fragment>
        <CssBaseline />
        <div className={classes.root}>
        <Grid container justify="center">
            <Grid
            spacing={10}
            alignItems="center"
            justify="center"
            container
            className={classes.grid}
            >
            <Grid item xs={12}>
                <div className={classes.stepContainer}>
                <div className={classes.bigContainer}>
                    <Stepper
                    classes={{ root: classes.stepper }}
                    activeStep={activeStep}
                    alternativeLabel
                    >
                    {steps.map(label => {
                        return (
                        <Step key={label}>
                            <StepLabel>{label}</StepLabel>
                        </Step>
                        );
                    })}
                    </Stepper>
                </div>
                {activeStep === 0 && (
                    <div className={classes.bigContainer}>
                    <Paper className={classes.paper}>
                    <AccountProfil />
                    </Paper>
                    </div>
                )}
                {activeStep === 1 && (
                    <div className={classes.bigContainer}>
                        <Paper className={classes.paper}>
                            <Grid item container xs={12}>
                                <Container fluid >
                                    <div className="d-flex padding-top-2 padding-bottom-10 justify-content-center align-items-center">
                                        <Row className="justify-content-md-center">
                                            <Col>
                                                <Row>
                                                    <Col>
                                                        <div className="section-title line-style no-margin">
                                                            <blockquote className="blockquote text-center">
                                                                <p className="mt-3"><b>Quel type de voyageur te correspond le plus ?</b></p>
                                                            </blockquote>
                                                            <FormControl component="fieldset" className="d-flex">
                                                                <RadioGroup row aria-label="traveler_type" name="traveler_type" value={traveler_type} onChange={this.onChange} className="justify-content-around">
                                                                    <FormControlLabel value="Tripizien culturel" control={<Radio color="default" />} label="Tripizien culturel" labelPlacement="start"/>
                                                                    <FormControlLabel value="Tripizien des temps modernes" control={<Radio color="default" />} label="Tripizien des temps modernes" labelPlacement="bottom"/>
                                                                    <FormControlLabel value="Tripizien Accompli" control={<Radio color="default" />} label="Tripizien Accompli" labelPlacement="end"/>
                                                                </RadioGroup>
                                                            </FormControl>

                                                            <Row>
                                                                <Col lg="4" className="d-flex justify-content-between align-items-start mt-4">
                                                                    <p>Le <b>Tripizien culturel</b> aime assouvir sa soif de connaissances et apprécie être en totale immersion avec son environnement. 
                                                                    Il aime s'imprégner des endroits qu'il visite et d'en découvrir la culture, les gens et le paysage.</p>
                                                                </Col>
                                                                <Col lg="4" className="d-flex justify-content-between align-items-start mt-4">
                                                                    <p>Le <b>Tripizien des temps modernes</b> est sociable, toujours en mouvement et de bonne humeur. En voyage, il aime l’adrénaline, 
                                                                    les émotions fortes, et la nouveauté. Le voyage est pour lui une merveilleuse aventure qu’il vit à deux cent à l’heure. 
                                                                    </p>
                                                                </Col>
                                                                <Col lg="4" className="d-flex justify-content-between align-items-start mt-4">
                                                                    <p>Le <b>Tripizien accompli</b> est attentionné, altruiste et généreux, il est attentif aux autres. 
                                                                    En voyage, il est respectueux de l’environnement et il privilégie les voyages ayant une dimension de développement personnel et de bien-être.</p>
                                                                </Col>
                                                            </Row>
                                                        </div>
                                                    </Col>
                                                </Row>
                                            </Col>
                                        </Row>
                                    </div>
                                </Container>
                            </Grid>
                        </Paper>
                    </div>
                )}
                {activeStep === 2 && (
                    <div className={classes.bigContainer}>
                        <Paper className={classes.paper}>
                            <Grid item container xs={12}>  
                                <Container fluid >
                                    <div className="d-flex padding-top-2 padding-bottom-10 justify-content-center align-items-center">
                                        <Row className="justify-content-md-center">
                                            <Col>
                                                <Row>
                                                    <Col>
                                                        <div className="section-title line-style no-margin">
                                                            <blockquote className="blockquote text-center">
                                                                <p className="mt-3"><b>Personnalité</b></p>
                                                            </blockquote>

                                                            <FormControl component="fieldset">
                                                                <Row className="d-flex justify-content-center">
                                                                    <RadioGroup row aria-label="ShyOutgoing" name="ShyOutgoing" value={ShyOutgoing} onChange={this.onChange} className="radioWidth">
                                                                        <Col className="d-flex justify-content-end">
                                                                            <FormControlLabel value="Timide" control={<Radio color="default" />} label="Timide" labelPlacement="start"/>
                                                                        </Col>
                                                                        <Col className="d-flex justify-content-start">
                                                                            <FormControlLabel className={classes.radio} value="Extraverti" control={<Radio color="default" />} label="Extraverti" />
                                                                        </Col>
                                                                    </RadioGroup>
                                                                </Row>

                                                                <Row className="d-flex justify-content-center">
                                                                    <RadioGroup row aria-label="SundayMikeHornAdventurer" name="SundayMikeHornAdventurer" value={SundayMikeHornAdventurer} onChange={this.onChange} className="radioWidth">
                                                                        <Col className="d-flex justify-content-end">
                                                                            <FormControlLabel value="Aventurier comme Mike Horn" control={<Radio color="default" />} label="Aventurier comme Mike Horn" labelPlacement="start"/>
                                                                        </Col>
                                                                        <Col className="d-flex justify-content-start">
                                                                            <FormControlLabel className={classes.radio} value="Aventurier du dimanche" control={<Radio color="default" />} label="Aventurier du dimanche" />
                                                                        </Col>
                                                                    </RadioGroup>
                                                                </Row>

                                                                <Row className="d-flex justify-content-center">
                                                                    <RadioGroup row aria-label="AthleteSportKezako" name="AthleteSportKezako" value={AthleteSportKezako} onChange={this.onChange} className="radioWidth">
                                                                        <Col className="d-flex justify-content-end">
                                                                            <FormControlLabel value="Le sport ? Ca se mange ?" control={<Radio color="default" />} label="Le sport ? Ca se mange ?" labelPlacement="start"/>
                                                                        </Col>
                                                                        <Col className="d-flex justify-content-start">
                                                                            <FormControlLabel className={classes.radio} value="Le sport c'est la vie" control={<Radio color="default" />} label="Le sport c'est la vie" />
                                                                        </Col>
                                                                    </RadioGroup>
                                                                </Row>

                                                                <Row className="d-flex justify-content-center">
                                                                    <RadioGroup row aria-label="CalmOrSpark" name="CalmOrSpark" value={CalmOrSpark} onChange={this.onChange} className="radioWidth">
                                                                        <Col className="d-flex justify-content-end">
                                                                            <FormControlLabel value="Un vrai piment de Cayenne" control={<Radio color="default" />} label="Un vrai piment de Cayenne" labelPlacement="start"/>
                                                                        </Col>
                                                                        <Col className="d-flex justify-content-start">
                                                                            <FormControlLabel className={classes.radio} value="Une petite perle de lait" control={<Radio color="default" />} label="Une petite perle de lait" />
                                                                        </Col>
                                                                    </RadioGroup>
                                                                </Row>                        
                                                            </FormControl>

                                                            <div className="divider div-transparent"></div>

                                                            <blockquote className="blockquote text-center">
                                                                <p className="mt-2"><b>Rythme de vie</b></p>
                                                            </blockquote>

                                                            <FormControl component="fieldset">
                                                                <Row className="d-flex justify-content-center">
                                                                    <RadioGroup row aria-label="FatMorningReadyLeave7AM" name="FatMorningReadyLeave7AM" value={FatMorningReadyLeave7AM} onChange={this.onChange} className="radioWidth">
                                                                        <Col className="d-flex justify-content-end">
                                                                            <FormControlLabel value="Grasse mat’ c’est le kiff !" control={<Radio color="default" />} label="Grasse mat’ c’est le kiff !" labelPlacement="start"/>
                                                                        </Col>
                                                                        <Col className="d-flex justify-content-start">
                                                                            <FormControlLabel className={classes.radio} value="Frais et paré à partir à 7h12 !" control={<Radio color="default" />} label="Frais et paré à partir à 7h12 !" />
                                                                        </Col>
                                                                    </RadioGroup>
                                                                </Row>

                                                                <Row className="d-flex justify-content-center">
                                                                    <RadioGroup row aria-label="scootered_disorganized" name="scootered_disorganized" value={scootered_disorganized} onChange={this.onChange} className="radioWidth">
                                                                        <Col className="d-flex justify-content-end">
                                                                            <FormControlLabel value="Organisé comme un scout" control={<Radio color="default" />} label="Orgnisé comme un scout" labelPlacement="start"/>
                                                                        </Col>
                                                                        <Col className="d-flex justify-content-start">
                                                                            <FormControlLabel className={classes.radio} value="Désorganisé depuis que je suis né" control={<Radio color="default" />} label="Désorganisé depuis que je suis né" />
                                                                        </Col>
                                                                    </RadioGroup>
                                                                </Row>
                                                            </FormControl>

                                                            <div className="divider div-transparent"></div>

                                                            <blockquote className="blockquote text-center">
                                                                <p className="mt-2"><b>Activités</b></p>
                                                            </blockquote>

                                                            <FormControl component="fieldset">
                                                                <Row className="d-flex justify-content-center">
                                                                    <RadioGroup row aria-label="trekking_beach_towel" name="trekking_beach_towel" value={trekking_beach_towel} onChange={this.onChange} className="radioWidth">
                                                                        <Col className="d-flex justify-content-end">
                                                                            <FormControlLabel value="Partir à l’aventure !" control={<Radio color="default" />} label="Partir à l’aventure !" labelPlacement="start"/>
                                                                        </Col>
                                                                        <Col className="d-flex justify-content-start">
                                                                            <FormControlLabel className={classes.radio} value="Serviette, Plage et Bronzette" control={<Radio color="default" />} label="Serviette, Plage et Bronzette"/>
                                                                        </Col>
                                                                    </RadioGroup>
                                                                </Row>
                                                            </FormControl>
                                                        </div>
                                                    </Col>
                                                </Row>
                                            </Col>
                                        </Row>
                                    </div>
                                </Container>
                            </Grid>
                        </Paper>
                    </div>
                )}
                {activeStep === 3 && (
                    <div className={classes.bigContainer}>
                        <Paper className={classes.paper}>
                            <Container fluid >
                                <div className="d-flex padding-top-2 padding-bottom-10 justify-content-center align-items-center">
                                    <Row className="justify-content-md-center">
                                        <Col>
                                            <Row>
                                                <Col>
                                                    <div className="section-title line-style no-margin">
                                                        <blockquote className="blockquote text-center">
                                                            <p className="mt-3"><b>Hébergement</b></p>
                                                        </blockquote>

                                                        <FormControl component="fieldset">
                                                            <Row className="d-flex justify-content-center">
                                                                <RadioGroup row aria-label="LuxuryCharmingLodging" name="LuxuryCharmingLodging" value={LuxuryCharmingLodging} onChange={this.onChange} className="radioWidth">
                                                                    <Col className="d-flex justify-content-end">
                                                                        <FormControlLabel value="Plus y’a d’étoiles, plus je brille" control={<Radio color="default" />} label="Plus y’a d’étoiles, plus je brille" labelPlacement="start"/>
                                                                    </Col>
                                                                    <Col className="d-flex justify-content-start">
                                                                        <FormControlLabel className={classes.radio} value="Un petit lieu de charme" control={<Radio color="default" />} label="Un petit lieu de charme" />
                                                                    </Col>
                                                                </RadioGroup>
                                                            </Row>
                                                        </FormControl>

                                                    </div>
                                                </Col>
                                            </Row>
                                        </Col>
                                    </Row>
                                </div>
                            </Container>
                        </Paper>
                    </div>
                )}
                {activeStep === 4 && (
                    <div className={classes.bigContainer}>
                        <Paper className={classes.paper}>
                            <Container fluid >
                                <div className="d-flex padding-top-2 padding-bottom-10 justify-content-center align-items-center">
                                    <Row className="justify-content-md-center">
                                        <Col>
                                            <Row>
                                                <Col>
                                                    <div className="section-title line-style no-margin">
                                                        <blockquote className="blockquote text-center">
                                                            <p className="mt-3"><b>Régime alimentaire</b></p>
                                                        </blockquote>

                                                        <FormControl component="fieldset">
                                                            <Row className="d-flex justify-content-center">
                                                                <RadioGroup row aria-label="food" name="food" value={food} onChange={this.onChange}>
                                                                    <FormControlLabel value="Halal" control={<Radio color="default" />} label="Halal"/>
                                                                    <FormControlLabel value="Vegan" control={<Radio color="default" />} label="Vegan"/>
                                                                    <FormControlLabel value="Végétarien" control={<Radio color="default" />} label="Végétarien"/>
                                                                    <FormControlLabel value="Casher" control={<Radio color="default" />} label="Casher"/>
                                                                    <FormControlLabel value="Aucun régime" control={<Radio color="default" />} label="Aucun régime"/>
                                                                </RadioGroup>
                                                            </Row>
                                                        </FormControl>

                                                    </div>
                                                </Col>
                                            </Row>
                                        </Col>
                                    </Row>
                                </div>
                            </Container>
                        </Paper>
                    </div>
                )}
                {(activeStep === 5 || activeStep === 6) && (
                    <div className={classes.bigContainer}>
                    <Paper className={classes.paper}>
                        <Grid item container xs={12}>
                        <Grid item xs={12}>
                            <Typography variant="subtitle1" gutterBottom>
                            Congratulations{" "}
                            <span role="img" aria-label="conrats emoji">
                                🎉
                            </span>
                            </Typography>
                            <Typography variant="body1" gutterBottom>
                            Vos informations ont été mises à jour.
                            </Typography>
                            <Button href="/account" fullWidth variant="outlined">
                            Aller vers votre profil
                            </Button>
                        </Grid>
                        </Grid>
                    </Paper>
                    </div>
                )}
                <div className={classes.flexBar}>
                    {activeStep !== 5 && (
                    <Button
                        disabled={activeStep === 0}
                        onClick={this.handleBack}
                        className={classes.backButton}
                        size="large"
                    >
                        Précedent
                    </Button>
                    )}
                    {activeStep !== 5 && (
                    <Button
                    variant="contained"
                    color="primary"
                    onClick={
                        activeStep !== 4 ? this.handleNext : this.onSubmit
                    }
                    size="large"
                    >
                    {this.stepActions()}
                    </Button>
                    )}
                    {activeStep === 5 && (
                    <Button
                    variant="contained"
                    color="primary"
                    onClick={this.goToDashboard}
                    size="large"
                    >
                    {this.stepActions()}
                    </Button>
                    )}
                </div>
                </div>
            </Grid>
            </Grid>
        </Grid>
        </div>
    </React.Fragment>
    );
}
}

export default compose(
    withRouter,
    withBackaccessContext,
    withStyles(styles),
)(Wizard);