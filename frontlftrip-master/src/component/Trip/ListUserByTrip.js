import React, {Component} from 'react';
import {TaggedContentCard} from 'react-ui-cards';

import { withRouter } from 'react-router-dom';
import { compose } from 'recompose';
import { withBackaccessContext } from '../BackEnd';
//import * as ROUTES from '../../const/routes'; /*NOT USED*/

const styles = {
div:{
    display: 'flex',
    flexFlow: 'wrap',
    //justifyContent: 'center',
    alignItems: 'center',
}
};
//const user = JSON.parse(localStorage.getItem("authUser")); /* NOT USED*/
class ListUserByTrip extends Component{
    constructor(props){
        super(props);
        this.state = {
            users :[],
            uri: 'participation',
            id: this.props.match.params.id,
        }
    }
    componentDidMount() {
        this.props.backaccess
        .getDataWitId(this.state.id,this.state.uri)
        .then((data) => {
            this.setState({ users: data.response});
            //console.log(this.state.users);
        }).catch(error => {
            console.log("not ok");
            this.setState({ error });
        });
    };
    render()
    {
        return(
            <div style={styles.div}>
            {this.state.users.map(user => 
                (
            <div key={user.user.id}>
            {user.user.avatar_path && 
                <TaggedContentCard
                href= {`user-alike/${user.user.id}`}
                thumbnail={user.user.avatar_path}    
                title={user.user.firstname +' '+ user.user.lastname}
                description={user.user.description}
                tags={[
                    `${user.user.city}`,
                    `${user.user.departure_airport}`,
                    `${user.user.sexe}`
                ]}
                />
            }
            {!user.user.avatar_path && 
                <TaggedContentCard
                href= {`user-alike/${user.user.id}`}
                thumbnail="/img/Default.png"   
                title={user.user.firstname +' '+ user.user.lastname}
                description={user.user.description}
                tags={[
                    `${user.user.city}`,
                    `${user.user.departure_airport}`,
                    `${user.user.sexe}`
                ]}
                />
            }

        </div>
            ))}
        </div>
        )
    }
}

export default  compose(
    withRouter,
    withBackaccessContext,
)(ListUserByTrip);