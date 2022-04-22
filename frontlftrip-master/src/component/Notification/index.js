import React,{Component} from 'react';
import { w3cwebsocket as W3CWebSocket } from "websocket";
import { withAuthentification } from '../Session';
import {store} from "react-notifications-component";
import NavDropdown from 'react-bootstrap/NavDropdown';
import { withBackaccessContext } from '../BackEnd';
import { compose } from 'recompose';
class Notification extends Component {
    constructor(props) {
        super(props);
        this.state = { notifications: [] };
        this.initSocket = this.initSocket.bind(this);
      }
    
    initSocket () {
        this.ws = new W3CWebSocket("ws://lftrip.tk/api/wsnotification?apiKey="+this.props.authUser.email);
        this.ws.onmessage = (msg) => {
            let data = JSON.parse(msg.data)
            if(data.type !== undefined){
                if(data.type !== "Message" && data.type !== "GroupMessage"  )
                {
                    let notif = this.state.notifications
                    store.addNotification({
                        message: data.message,
                        type: 'success',                         // 'default', 'success', 'info', 'warning'
                        container: 'top-right',                // where to position the notifications
                        animationIn: ["animated", "fadeIn"],     // animate.css classes that's applied
                        animationOut: ["animated", "fadeOut"],   // animate.css classes that's applied
                        dismiss: {
                            duration: 1000
                        }
                    })
                    notif.push(data)
                    this.setState({
                        notifications : notif
                    })/**/
                }else if(!window.location.href.includes("message")){
                    store.addNotification({
                        message: "Vous avez un nouveau message",
                        type: 'default',                         // 'default', 'success', 'info', 'warning'
                        container: 'top-right',                // where to position the notifications
                        animationIn: ["animated", "fadeIn"],     // animate.css classes that's applied
                        animationOut: ["animated", "fadeOut"],   // animate.css classes that's applied
                        dismiss: {
                            duration: 1000
                        }
                    })
                }
            }
        }
    
    }
    
    componentDidMount() {
        this.initSocket();
        this.props.backaccess.getDataWitId(this.props.authUser.id,'notification').then(value=>{
            if(value !== undefined){
            this.setState({notifications : Object.values(value.response)})}
        })
	}
  
    render () 
    {
        return (
            <NavDropdown  title={
                <span><i className="fa fa-bell fa-fw"></i></span>
              } id="basic-nav-dropdown">
                  {
                      this.state.notifications.map( (notification, i)=>(
                        <div key={i}>
                            <NavDropdown.Item>{notification.message} par {notification.target_name} </NavDropdown.Item>
                            <NavDropdown.Divider />
                        </div>
                       ))
                  }
 
            </NavDropdown>
        );
    }
}
  
  export default  compose (withAuthentification, withBackaccessContext)(Notification);