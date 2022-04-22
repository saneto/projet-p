import React, {Component} from 'react';
import Button from 'react-bootstrap/Button';
import { withRouter } from 'react-router-dom';
import { compose } from 'recompose';
import { withBackaccessContext } from '../BackEnd';
import * as ROUTES from '../../const/routes';
import MaterialTable from 'material-table';

const styles = {
div:{
    display: 'flex',
    flexFlow: 'wrap',
    //justifyContent: 'center',
    alignItems: 'center',
},
btnCr: {
    margin: '2vh',
},
};
class TripsTable extends Component{
    constructor(props){
        super(props);
        this.state = {
            trips :[],
            uri: 'trips',
            columns: [
                { field: 'url', title: 'Titre', render: rowData => <a href={`trip/${rowData.id}`}>{rowData.title}</a>},
                { title: 'Pays', field: 'country' },
                { title: 'Description', field: 'description' },
                { title: 'Type de logement', field: 'lodging' },
                { title: 'Nombre de jours', field: 'nb_days'},
                { title: 'Programme', field: 'program' },
                ],
        }
    }
    componentDidMount() {
        this.props.backaccess
        .getData(this.state.uri)
        .then((data) => {
            this.setState({ trips: data.response});
            //console.log(this.state.trips);
        }).catch(error => {
            console.log("not ok");
            this.setState({ error });
        });
    };

    render()
    {
        return(
            <div>
            <div className="title-top">
            <div className="titleContainerTrip pl-5 pr-5">
                <h1>Liste des voyages des Tripiziens</h1>
            </div>
        </div>
        <div className="pl-5" style={styles.btnCr}>            
                <Button href= {ROUTES.CREATETRIP} variant="primary">
                    Créer une nouvelle aventure
                </Button>
        </div>
        <div className="box">
        <MaterialTable
            title="Tous les voyages sur LFTrip"
            columns={this.state.columns}
            data={this.state.trips} />
        </div>
        </div>
        )
    }
}

export default  compose(
    withRouter,
    withBackaccessContext,
)(TripsTable);
