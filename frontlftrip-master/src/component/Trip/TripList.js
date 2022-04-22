import React, {Component} from 'react';
import AppBar from '@material-ui/core/AppBar';
import Button from 'react-bootstrap/Button';
import {Container, Row} from 'react-bootstrap';
import TripSearchForm from '../TripSearch/TripSearchForm';
import { withRouter } from 'react-router-dom';
import { compose } from 'recompose';
import { withBackaccessContext } from '../BackEnd';
import TripsTable from './TripsTable';
import * as ROUTES from '../../const/routes';
import MapLeaflet from "./Map";
import history from '../../history';
import Fab from '@material-ui/core/Fab';
import AddIcon from '@material-ui/icons/Add';
import FormatListBulletedIcon from '@material-ui/icons/FormatListBulleted';
import ExploreIcon from '@material-ui/icons/Explore';
import AppsIcon from '@material-ui/icons/Apps';
import Tabs from '@material-ui/core/Tabs';
import Tab from '@material-ui/core/Tab';
import PropTypes from 'prop-types';


const styles = {
div:{
    display: 'flex',
    flexFlow: 'wrap',
    //justifyContent: 'center',
    alignItems: 'center',
},
margin: {
    position: 'fixed',
    bottom: '1%',
    right: '1%',
    zIndex: '9999',
},
root: {
    flexGrow: 1,
    width: '100%',
},
btnCr: {
    margin: '2vh',
},

};

function TabPanel(props) {
    const { children, value, index, ...other } = props;
    
    return (
        <div
        role="tabpanel"
        hidden={value !== index}
        id={`simple-tabpanel-${index}`}
        aria-labelledby={`simple-tab-${index}`}
        {...other}
        >
        {value === index && (
            children
        )}
        </div>
    );
    }
    
    TabPanel.propTypes = {
    children: PropTypes.node,
    index: PropTypes.any.isRequired,
    value: PropTypes.any.isRequired,
    };
    
    function a11yProps(index) {
    return {
        id: `simple-tab-${index}`,
        'aria-controls': `simple-tabpanel-${index}`,
    };
    }
class Triplist extends Component{
    constructor(props){
        super(props);
        this.state = {
            trips :[],
            uri: 'trips',
            value: 0,
        }
    }

    filterText = (text, nbChar) => {
        var filteredText = text;
		if (text.length > nbChar){
            filteredText = text.substring(0, nbChar) + "...";
        }
        return filteredText;
	}

    componentDidMount() {
        this.props.backaccess
        .getData(this.state.uri)
        .then((data) => {
            this.setState({ trips: data.response});
        }).catch(error => {
            console.log("not ok");
            this.setState({ error });
        });
    };

    setTrip = value => {
        this.setState({ trips: value});
    }

    handleChange = (event, newValue) => {
        this.setState({value: newValue});
    };
    
    render()
    {
        //console.log(this.state.trips)
        return(
            <div style={styles.root}>
                <AppBar position="static" color="default">
                <Tabs
                    value={this.state.value}
                    onChange={this.handleChange}
                    variant="fullWidth"
                    indicatorColor="secondary"
                    textColor="secondary"
                    aria-label="icon label tabs example"
                >
                    <Tab icon={<ExploreIcon />} label="Voyages à la carte" {...a11yProps(0)} />
                    <Tab icon={<FormatListBulletedIcon />} label="Voyages en liste" {...a11yProps(1)} />
                    <Tab icon={<AppsIcon/>} label="Et même en card !" {...a11yProps(2)} />
                </Tabs>
                </AppBar>
                <TabPanel value={this.state.value} index={0}>
                    <div className="title-top">
                        <div className="titleContainerTrip">
                            <h1 className="pl-5">Notre carte</h1>
                        </div>
                    </div>
                    <div className="pl-5" style={styles.btnCr}>            
                        <Button href= {ROUTES.CREATETRIP} variant="primary">
                            Créer une nouvelle aventure
                        </Button>
                    </div>
                    <MapLeaflet/>
                    <br />    
                </TabPanel>

                <TabPanel value={this.state.value} index={1}>
                    <TripsTable />
                </TabPanel>

                <TabPanel value={this.state.value} index={2}>
                <TripSearchForm setTrip={this.setTrip} backaccess={this.props.backaccess}/>
                    <Container>
                        <Row className="pt-5 pl-3">
                            <Button href= {ROUTES.CREATETRIP} variant="primary">
                                Créer une nouvelle aventure
                            </Button>
                        </Row>     
                    </Container>
                    <Container>
                    <div className="containerTrip">
                        {this.state.trips.map(trip => (
                            <div className="box" key={trip.id}>
                                <div className="imgBox" key={trip.id}>
                                    {trip.base64 &&
                                        <img src={trip.base64} alt="Trip custom"></img>
                                    }
                                    {!trip.base64 &&
                                        <img src="img/card1.jpg" alt="Trip default"></img>
                                    }                            
                                </div>
                                <a href={`trip/${trip.id}`}>
                                    <div className="details">
                                        <div className="content">
                                            <h2>{this.filterText(trip.title,42)}</h2>
                                            <p>{this.filterText(trip.description,96)}</p>
                                            <div className="infoTrip mt-4">
                                                <li className="badge badge-pill badge-success mr-2">{trip.trip_type}</li>
                                                <li className="badge badge-pill badge-success mr-2">{trip.nb_traveler} voyageurs</li>
                                                <li className="badge badge-pill badge-success">{trip.nb_days} jours</li>
                                            </div>
                                        </div>
                                    </div>
                                </a>
                            </div>                    
                        ))}
                    </div>
                    </Container>
                </TabPanel>

                <Fab onClick= {() =>(history.push(ROUTES.CREATETRIP))} color="primary" aria-label="add" style={styles.margin}>
                        <AddIcon />
                </Fab>
            </div>
            
            
        )
    }
}

export default compose(
	withRouter,
	withBackaccessContext,
)(Triplist);