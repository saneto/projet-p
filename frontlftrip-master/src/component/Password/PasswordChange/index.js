import React, { Component } from 'react';

import {withAuthentification} from './withAuthentification';

const INITIAL_STATE = {
	passwordOne: '',
	passwordTwo: '',
	oldpassword: '',
	error: null,
};

class PasswordChangeForm extends Component {
	
	constructor(props) {
		super(props);

		this.state = { ...INITIAL_STATE };
	}

  	onSubmit = event => {
		const { passwordOne } = this.state;

		this.props.firebase
			.doPasswordUpdate(passwordOne)
			.then(() => {
				this.setState({ ...INITIAL_STATE });
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
		const {oldpassword, passwordOne, passwordTwo, error } = this.state;



		return (
			<div>
				<h3 className="title">Change Your Password</h3>
				<form onSubmit={this.onSubmit}  className="grey-box">
					<span className="text">
						There are many variations of passages of Lorem Ipsum available, but the majority have suffered alteration inble. If you are of going.
					</span>
					<input value={oldpassword} onChange={this.onChange} type="password" className="form-control" id="old-password" name="old-password" placeholder="Old password"/>
					<input value={passwordOne} onChange={this.onChange} type="password" className="form-control" id="passwordOne" name="passwordOne" placeholder="New password"/>
					<input value={passwordTwo} onChange={this.onChange} type="password" className="form-control" id="passwordTwo" name="passwordTwo" placeholder="Confirm new password"/>
						
					<button className="btn btn-defaul" type="submit">
						Reset My Password
					</button>

					{error && <p>{error.message}</p>}
					
				</form>
			</div>
		);
	}
}

export default withAuthentification(PasswordChangeForm);