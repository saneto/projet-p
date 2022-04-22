import React from 'react';
import {Button} from 'react-bootstrap';
import options from '../Trip/data';
import { Typeahead } from 'react-bootstrap-typeahead';
class TripSearchForm extends React.Component{
    constructor(props){
        super(props);
        this.state = {
            uri : "elk",
            country: "",
            nb_days : "",
            start_date: new Date().now,
            end_date: new Date().now,
            nb_traveler: "",
            budget: "",
            data : {}
        }
    }

    onSubmit = () =>{
        this.props.backaccess.getDataWitParam(this.state.data, this.state.uri).then( ret =>{
                this.props.setTrip(Object.values(ret.response))
            }
        ).catch( error =>{

        })
    }

    onChange = event => {
        let tab = this.state.data;
        if(event.target.value===""){
            delete this.state.data[event.target.name]
        }else{
            
            tab[event.target.name]= String(event.target.value)
        }
        
        this.setState({ [event.target.name]: event.target.value,
            data : tab });
    };

    onChangeDate = event => {
        let tab = this.state.data;
        let d  =  event.target.value //this.format(event.target.value)
        if(event.target.value===""){
            delete this.state.data[event.target.name]
        }else{
            tab[event.target.name]= String(d)
        }
        
        this.setState({ [event.target.name]: event.target.value,
            data : tab });
    };

    format(inputDate) {
        var date = new Date(inputDate);
        if (!isNaN(date.getTime())) {
            // Months use 0 index.
            let j = ""
            if(date.getDate()>= 10){
                j = ""+date.getDate()
            }else{
                j = "0"+ date.getDate()
            }
            let m = ""
            if((date.getMonth() + 1)>= 10){
                m =  ""+(date.getMonth() + 1)
            }else{
                m = "0"+ (date.getMonth() + 1)
            }
            return   j + '-' + m + '-' + date.getFullYear();
        }
    }

    onChangeCountry = event => {
        if(event.length===0){
            delete this.state.data["country"]
        }else{
            let d = this.state.data
            d["country"]= String(event[0].label)
            this.setState({ country: event[0].label,
                data : d });
        }
        
    };


    render()
    {
        const { country,   nb_days,
                start_date, end_date,
                nb_traveler, budget} = this.state;
        return(
        <React.Fragment>
            <div className="view-fixed slide1">

                <div className="mask">

                    <div className="row padding-top-20 text-center white-text mx-5 justify-content-center align-items-center">
                            
                        <div className="col-md-12">
                            <h1 className="mb-4">
                                <strong>A la recherche
                                <div className="green-text highlight"> d'une nouvelle aventure ?</div></strong>
                            </h1>
                        </div>
                        <form className="form">
                            <div className="row">
                                <div className="col-md-4">
                                    <div className="form-group">
                                        <label className="justify-content-start">Pays</label>
                                        <Typeahead {...this.state} id="country" name="country" onChange={this.onChangeCountry} options={options} placeholder="Séléctionner un pays..." value={country} />
                                    </div>
                                </div>
                                <div className="col-md-4">
                                    <div className="form-group">
                                        <label className="justify-content-start">Nombre de voyageurs</label>
                                        <input type="number" name="nb_traveler" onChange={this.onChange} value ={nb_traveler} className="form-control" id="nb_traveler" placeholder="Nombre de voyageurs"/>
                                    </div>
                                </div>
                                <div className="col-md-4">
                                    <div className="form-group">
                                        <label className="justify-content-start">Budget</label>
                                        <input type="number" name="budget" onChange={this.onChange} value ={budget} className="form-control" id="budget" placeholder="Budget"/>
                                    </div>
                                </div>
                            </div>
                            <div className="row">
                                <div className="col-md-4">
                                    <div className="form-group">
                                        <label className="justify-content-start">Nombre de jours</label>
                                        <input className="form-control" type="number" name="nb_days" onChange={this.onChange} value ={nb_days} placeholder="Nombre de jours"/>
                                    </div>
                                </div>
                                <div className="col-md-4">
                                    <div className="form-group">
                                        <label className="justify-content-start">Date de départ</label>
                                        <input type="date"  name="end_date" onChange={this.onChangeDate}  value ={end_date} className="form-control" id="departureDate" placeholder="Date de départ"/>
                                    </div>
                                </div>
                                <div className="col-md-4">
                                    <div className="form-group">
                                        <label className="justify-content-start">Date d'arrivée</label>
                                        <input type="date" name="start_date" onChange={this.onChangeDate} value ={start_date} className="form-control" id="arrivalDate" placeholder="Date d'arrivée"/>
                                    </div>
                                </div>
                            </div>

                            <div className="row justify-content-center mt-3">
                                <div className="col-md-2">
                                    <div className="text-center">
                                        <Button variant="primary" onClick={this.onSubmit}>
                                                Rechercher
                                        </Button>
                                    </div>
                                </div>
                            </div>
                                
                        </form>

                    </div>

                </div>

            </div>
        </React.Fragment>

        )
    }
}

export default TripSearchForm;