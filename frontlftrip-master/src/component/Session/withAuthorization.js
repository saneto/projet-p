import React from 'react';
import AuthUserContext from './context';

import { withRouter } from 'react-router-dom';
import * as ROUTES from '../../const/routes';

const withAuthorization = condition => Component =>{
    class withAuthorization extends React.Component{
        constructor(props){
            super(props);
            this.state = {
                authUser : JSON.parse(localStorage.getItem('authUser'))
            }
        }

        componentDidMount(){
            if(!condition(this.state.authUser)){
                this.props.history.push(ROUTES.SIGN_IN)
            }
        }

        componentWillUnmount() {
            if(!condition(this.state.authUser)){
                this.props.history.push(ROUTES.SIGN_IN)
            }
          }

        render(){
            return(
                <AuthUserContext.Consumer>
                    {
                        authUser => condition(authUser) ? <Component {...this.props} authUser ={this.state.authUser}/> : null
                    }
                </AuthUserContext.Consumer>
            )
        }
    }
    return withRouter(withAuthorization);
};
export default withAuthorization;