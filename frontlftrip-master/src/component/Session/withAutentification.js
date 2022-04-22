import React from 'react';
import { withBackaccessContext } from '../BackEnd';
import AuthUserContext from './context';
const withAuthentification = Component => {
	class withAuthentification extends React.Component {
		constructor(props) {
			super(props);

			this.state = {
				authUser: JSON.parse(localStorage.getItem('authUser')),
			};
		}

		componentDidMount() {

			if (localStorage.getItem("authUser") === null) {
				this.setState({ authUser: null });
			}
		  }
		  componentWillUnmount() {
			if (localStorage.getItem("authUser") === null) {
				this.setState({ authUser: null });
			}
          }

		render() {
			return (
				<AuthUserContext.Provider value={this.state.authUser}>
					<Component {...this.props} authUser ={this.state.authUser}/>
				</AuthUserContext.Provider>
			);
		}
	}

	return withBackaccessContext(withAuthentification);
};

export default withAuthentification;
