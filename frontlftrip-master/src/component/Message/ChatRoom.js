import React from 'react';

import SubmitMessage from './SubmitMessage.js';

class ChatRoom extends React.Component {
	constructor(props) {
		super(props);
		this.state = { };
		this.sendMessage = this.sendMessage.bind(this);
	}

  	sendMessage (message) {
		if(this.props.message !==[]){
			//console.log(this.message)
			if(this.props.message.groups){
				this.props.groupSend(message)
			}else{
				this.props.send(message)
			}}

  	}

  	componentDidUpdate() {
	  if(this.newData!==undefined && this.newData!==null){
		this.newData.scrollIntoView({ behavior: "smooth" })}
	}
	handleScroll = (e) => {
		if(e.target.scrollTop === 0){
			//console.log("ototot")
			this.props.retrieveMoreMsg()
		}
	}
  render () {
	const messages = (this.props.message!==undefined && this.props.message.messages!==undefined)?  this.props.message.messages : [] ;
	//console.log(messages);
	let barProfil = "";
	if(this.props.message !== undefined){
		if(!this.props.message.groups){
			if(this.props.target.avatar_path!==""){
				barProfil=<div><img id="profile-img"  className="online"  src="/img/Default.png" alt="" /><p>{this.props.target.lastname} {this.props.target.firstname}</p></div>
			}else{
				barProfil=<div><img src={this.props.target.avatar_path} alt="" /><p>{this.props.target.lastname} {this.props.target.firstname}</p></div>
			}
		}else{
			barProfil=<div><img id="profile-img"  className="online"  src="/img/group.png" alt="" /><p>{this.props.message.name}</p></div>
		}
	}

    return (
		<div className="content">
			<div className="contact-profile">
				{barProfil}
			</div>
			<div className="messages" onScroll={this.handleScroll}>
				<ul >
					{messages.map((message, i) => (
						(message.body!=="")?
							((message.sender_id===this.props.authUser.id)?
							<li key={i} className="replies" ref={(ref) => this.newData = ref}>
								<img className="online"  src="/img/Default.png" alt="" />
								<p>{message.body}<br/>
								<span className="msg_time_send_replies"> par {message.sender_name}</span></p>
							</li> :
							<li key={i} className="sent" ref={(ref) => this.newData = ref}>
								<img className="online"  src="/img/Default.png" alt="" />
								<p>{message.body}<br/>
								<span className="msg_time_send_sent"> par {message.sender_name}</span></p>
							</li>): ''
					))}
				</ul>
			</div>
			<SubmitMessage
				placeholder="speak your mind"
				onSubmit={this.sendMessage}  />
		</div>
        
    )
  }
}

export default ChatRoom;