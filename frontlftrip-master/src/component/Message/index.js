import React from 'react';

import ChatRoom from './ChatRoom';
import UserCard from './UserCard';
import { Container, Row, Col } from 'react-bootstrap';
import { withBackaccessContext } from '../BackEnd';
//import Button from 'react-bootstrap/Button';
import { withAuthorization } from '../Session';
import { w3cwebsocket as W3CWebSocket } from "websocket";
import './index.css';
import { compose } from 'recompose';

class Message extends React.Component {
	constructor(props) {
		super(props);
		this.state = {
			discussions:[],
			currentUser : 0,
			m : [],
			listUser : [],
			existingDis : [],
			target : [], 
			senderid:"",
			sende:""
		};
		this.initSocket = this.initSocket.bind(this);
	}
	componentDidMount() {
		this.initdata();
		this.initSocket();
	}

	initdata = () =>{
		this.props.backaccess.getDataWitId(this.props.authUser.id, 'elkM').then( ret=>{
				let chats =(ret!== undefined )?ret.chat : [];
				if(ret!== undefined && Object.keys(chats).length !== 0 ){
					this.initChat(chats,ret.users )
				}else if(ret!== undefined){
					this.setState({ 
						listUser : ret.users
					});
					this.getChatFromSQL()
				}else{
					this.props.backaccess
						.getDataWithToken("users")
						.then((data) => {
							this.setState({ listUser :data.response  });
						})
						.catch(error => {
							this.setState({ error });
						});
				}

			}
		)
	}
	initChat = (chats, users) =>{
		let discussions = [];let existingDis=[];
		for(const key in chats){
			let discussion = {};
			discussion.id = chats[key][0].conversation_id;
			if(chats[key][0].group){
				discussion.groups= true;
				discussion.name = chats[key][0].conversation_name;
				discussion.users= [];
				chats[key][0].conversation_participant_ids.forEach(element=> {
					//console.log()
					if(element!==this.props.authUser.id){
						users.forEach(user=> {
							if(user.id===element){
								existingDis.push(user.id);
								discussion.users.push(user);
							}
						});
					}
				});
			}else{
				chats[key][0].conversation_participant_ids.forEach(element=> {
					//console.log()
					if(element!==this.props.authUser.id){
						users.forEach(user=> {
							if(user.id===element){
								existingDis.push(user.id)
								discussion.users = user
							}
						});
					}
				})
				discussion.groups= false
			}

			discussion.messages = Object.values(chats[key]).reverse()
			discussion.updated_at = chats[key][0].created_at
			if(discussion.users===undefined){
				discussion.users = this.props.authUser
			}
			discussions.push(discussion)
		}
		this.onchangeDataOnGroup(discussions[0])
		this.setState({ discussions: discussions , 
			m  :discussions[0] , 
			currentUser:discussions[0].id,
			target: discussions[0].users,
			existingDis: existingDis, 
			listUser : users
		});
	}

	getChatFromSQL = () =>{
		this.props.backaccess
			.getDataWithToken("chat")
			.then((data) => {
					if(data.response){
						let existingDis=[];
						let value = "";
						data.response.forEach((values, i)  => {
							let usr ;
							values.users.forEach(element=> {
							if(element.email !==this.props.authUser.email)
									{existingDis.push(element.id)
									usr = element;
									if(i ===0){
										value= element
									}}
							});
							values.users = usr;
						});
						//let m = (data.response[0].messages!==undefined)? data.response[0].messages : []
						this.setState({ discussions: data.response, 
										m  :data.response[0].messages , 
										currentUser:data.response[0].id,
										target: value,
										existingDis: existingDis });

					}
				})
				.catch(error => {
					this.setState({ error });
				});
	}
	initSocket () {
		this.ws = new W3CWebSocket("ws://lftrip.tk/api/messages?apiKey="+this.props.authUser.email);
		this.ws.onmessage = (msg) => {
			//console.log(msg)
			var obj = JSON.parse(msg.data);
			if(obj["type"]=== "Message" || obj["type"]=== "GroupMessage"){
				this.state.discussions.forEach((element, i) => {
					if(element.id === parseInt(obj.target_id)){
						let discussion = this.state.discussions[i].messages
						if(discussion===undefined){
							discussion=[]
						}
						discussion.push({
							"body": obj.message,
							"created_at": new Date().getTime(),
							"sender_id": obj.target_name
						})
					}
					this.setState({ discussions:this.state.discussions });
				});
			}else if(obj["type"]=== "GroupChat"){
				if(this.state.discussions.length === 0)
				{	let dis = {};
					dis.id=obj.target_id;
					dis.name=obj.target_name; 
					dis.groups = true;
					dis.updated_at = obj.created_at;
					let par  = obj.target_email.split(",")
					let usr = []
					this.state.listUser.forEach(element=>{
						if(element.id!==this.props.authUser.id && par.includes(String(element.id))){
							usr.push(element)
						}
					})
					dis.users = usr
					this.setState({
						currentUser : ( this.state.discussions.length === 0) ?  dis.id  : this.state.currentUser,
						discussions: [...this.state.discussions, dis],
						target:  usr,
						m : dis
					})
				}else{
					let dis=this.state.discussions;
					let notExist = false;
					this.state.discussions.forEach((discussion, id)=>{
						if(discussion.id === parseInt(obj.target_id) ){
							notExist = true
							let par  = obj.target_email.split(",")
							let usr = []
							this.state.listUser.forEach(element=>{
								if(element.id!==this.props.authUser.id && par.includes(String(element.id))){
									usr.push(element)
								}
							})
							dis[id].users=usr
						}
					})
					if(notExist){
						this.setState({
							discussions:dis
						})
					}else{
						dis = {};
						dis.id=obj.target_id;
						dis.name=obj.target_name; 
						dis.groups = true;
						dis.updated_at = obj.created_at;
						let par  = obj.target_email.split(",")
						let usr = []
						this.state.listUser.forEach(element=>{
							if(element.id!==this.props.authUser.id && par.includes(String(element.id))){
								usr.push(element)
							}
						})
						dis.users = usr
						this.setState({
							currentUser : ( this.state.discussions.length === 0) ?  dis.id  : this.state.currentUser,
							discussions: [...this.state.discussions, dis],
							target:  usr,
							m : dis
						})
					}
				}

			}else if(obj["type"]=== undefined){
					let usr ;
					obj.users.forEach(element=> {
						if(element.email !==this.props.authUser.email){
							usr = element;}
					});
					obj.users = usr;
					obj.messages =[];
					this.setState({
						currentUser : ( this.state.discussions.length === 0) ?  obj.id  : this.state.currentUser,
						discussions: [...this.state.discussions, obj],
						existingDis: [...this.state.existingDis, obj.users.id],
						target: obj.users 
					})
					if(this.state.discussions.length === 1){
						this.setState({
							m : obj.messages
						})
					}
				}
		}
	}
	
	onSendMessage  = (message)=> {
		//console.log("édqsdsssssaaa&&&&&qsdds")
		this.ws.send(
			JSON.stringify({
				type : "solo",
				createBy: String(this.props.authUser.id),
				targetUsername:this.state.target.email,
				discussionId:  String(this.state.currentUser),
				targetName:this.state.target.lastname,
				targetID: String(this.state.target.id),
				sender_name: this.props.authUser.firstname,
				content: ( message),
				
			})
		);
		let d = this.state.discussions;
		d.forEach((element, i) => {
			if(element.id ===this.state.currentUser){
				if(d[i].messages===undefined){
					d[i].messages=[]
				}
				d[i].messages.push({
					"body": message,
					"created_at": new Date().getTime(),
					"sender_id": this.props.authUser.id,
					"sender_name": this.props.authUser.firstname
				})
			}})
		this.setState({ discussions:d });
	}

	onSendMessageGroup  = (message)=> {
		//console.log("édqsdqsdds")
		if(this.state.senderid==="" || this.state.sende===""){
			this.onchangeDataOnGroup(this.state.m)
		}
		this.props.backaccess
			.postData({
				createBy: String(this.props.authUser.id),
				targetUsername: this.state.sende,
				discussionId:  String(this.state.currentUser),
				targetName:this.state.m.name,
				targetID: this.state.senderid+","+this.props.authUser.id ,
				content:  message,
				senderName: this.props.authUser.firstname
			},"groupm")
			.then((data) => {
				let d = this.state.discussions;
				d.forEach((element, i) => {
					if(element.id ===this.state.currentUser){
						if(d[i].messages===undefined){
							d[i].messages=[]
						}
						d[i].messages.push({
							"body": message,
							"created_at": new Date().getTime(),
							"sender_id": this.props.authUser.id,
							"sender_name": this.props.authUser.firstname
						})
					}})
				this.setState({ discussions:d });
			})
			.catch(error => {
				this.setState({ error });
			});
	}

    onChangeChat = value =>{
		this.setState({ m : value.messages, currentUser: value.discussionId, target: value.targeruser });	    
		this.onchangeDataOnGroup(value.messages)
	}

	onchangeDataOnGroup = value =>{
		if(value.groups === true){
			let e =""
			let azee = ""
			value.users.forEach((user, k)=>{
				if(k < value.users.length-1){
                    e+=user.email+",";
					azee+=String(user.id)+",";
                }else{
                    e+=user.email;
					azee+=String(user.id);
                }
			})
			this.setState({ senderid : azee  });
			this.setState({ sende : String(e) });			
		} else if(value.groups === false){
			this.setState({ senderid : "", 
			sendE: "" });
		}
	}
	
	creationChat = event =>{
		this.props.backaccess
			.postData({email : event.target.value },"chat")
			.then((data) => {
				if(data.data.response){
					let usr ;
					data.data.response.users.forEach(element=> {
						if(element.email !==this.props.authUser.email){
							usr = element;
						}	
					});
					data.data.response.users = usr;
					data.data.response.messages =[]
					let discussion = data.data.response
					discussion.messages=[]
					this.setState({
						currentUser : ( this.state.discussions.length === 0) ?  data.data.response.id  : this.state.currentUser,
						discussions: [...this.state.discussions, discussion],
						existingDis: [...this.state.existingDis, data.data.response.users.id],
						target:  data.data.response.users,
					})
					if(this.state.discussions.length === 1){
						this.setState({
							m : data.data.response.messages
						})
					}
				}
			})
			.catch(error => {
				this.setState({ error });
			});
		
	 }

	 retrieveMoreMsg =() =>{
		this.props.backaccess
		.postData({
			id: String(this.state.m.id),
			lastIndex: String(this.state.m.messages.length)
		},"message")
		.then((data) => { 
			let d = this.state.m;
			let tab = Object.values(data.data.response).reverse().concat(d.messages)
			d.messages= tab;
			//console.log(this.state)
			let ds = this.state.discussions
			ds.forEach((element, k) => {
				if(element.id===d.id){
					ds[k]=d
				}
			})
			this.setState({
				m: d,
				discussions: ds
			})
		})
	}
  	render () {
		const {discussions} =this.state;
		return (
			<Container>
				<br/>
				<Row>
					<Col xs="4">
						{/*listUser.map((user, i) =>

								((this.props.authUser.email!==user.email && this.state.existingDis.indexOf(user.id) ===-1)? <div key={i}>
									<Row>
										{user.lastname} {user.firstname}  {user.email} 
										<Button value ={user.email} className="icon fa fa-cog" onClick={this.creationChat} variant="primary"> création d'une discussions</Button>
									</Row>
								</div> : "" )
								
							
						)*/}
					</Col>
				</Row>
				<Row>
					<div id="frame">
						<div id="sidepanel">
							<div id="profile">
								<div className="wrap">
									{(this.props.authUser.avatar_path!=="") ? ( <img src={this.props.authUser.avatar_path}   className="online"  alt="" /> ):( <img id="profile-img"  className="online"  src="/img/Default.png" alt="" />)}
									<p>{this.props.authUser.lastname}  {this.props.authUser.firstname}</p>
								</div>
							</div>
							<div id="contacts">
								<ul>
									{discussions.map((discussion, i) => (
										(discussion.id!==this.state.currentUser)?(<li key={i} className='contact'>
											<UserCard key={discussion.id} message={(discussion.messages!==undefined && discussion.messages.length > 0)?discussion.messages.slice(-1)[0].body:""}  discussion = {discussion} discussionId={discussion.id} onChangeChat={this.onChangeChat} currentUser={this.state.currentUser} />
										</li>):
										(<li key={i} className="contact active">
											<UserCard key={discussion.id} message={(discussion.messages!==undefined && discussion.messages.length > 0)?discussion.messages.slice(-1)[0].body:""} discussion = {discussion} discussionId={discussion.id} onChangeChat={this.onChangeChat} currentUser={this.state.currentUser} />
										</li>)
									))}
								</ul>
							</div>
						</div>
						<ChatRoom retrieveMoreMsg={this.retrieveMoreMsg} send={this.onSendMessage} groupSend={this.onSendMessageGroup} message={this.state.m} authUser={this.props.authUser}  target={this.state.target} />
					</div>
				</Row>
				<br/>
			</Container>
		);
 	}
}

const condition = authUser => !!authUser;
export default  compose( withAuthorization(condition), withBackaccessContext  )(Message);
