import React from 'react';

class SubmitMessage extends React.Component {
	constructor(props) {
		super(props);
		this.state = {value: ''};
		this.handleChange = this.handleChange.bind(this);
        this.handleOnKeyDown = this.handleOnKeyDown.bind(this);
        this.onSendMessage = this.onSendMessage.bind(this);
	}

	handleChange(event) {
		this.setState({value: event.target.value});
	}

	handleOnKeyDown(event) {
		if (event.keyCode === 13 && this.state.value !== "") {
			this.props.onSubmit(this.state.value);
			this.setState({ value: '' })  ;
		}
    }
    
    onSendMessage(event) {
		if(this.state.value!==""){
			this.props.onSubmit(this.state.value);
			this.setState({ value: '' });
		}
	}

	render() {
		return (
			<div className="message-input">
				<div className="wrap">
					<input type="text" placeholder="Write your message..." 
						value={this.state.value}
						onChange={this.handleChange}
						onKeyDown={this.handleOnKeyDown} />
					<button onClick={this.onSendMessage} className="submit"><i className="fa fa-paper-plane" aria-hidden="true"></i></button>
				</div>
            </div>
		);
	}
}

export default  SubmitMessage;