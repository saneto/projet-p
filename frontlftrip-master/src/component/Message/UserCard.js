import React,{Component} from "react";

class UserCard extends Component
{
    constructor(props)
    {
        super(props)        
        this.state={
            targetUser:this.props.discussion.users
        }
    }

    changeDis = ()=>{
       this.props.onChangeChat({
            messages : this.props.discussion,
            discussionId : this.props.discussionId,
            targeruser : this.state.targetUser
        })
    }

    render()
    {
        const discussion = this.state.targetUser;
        return(
            (discussion!==undefined && this.props.discussion.groups)?    
                <div value={this.props.discussion.name} onClick={this.changeDis}  className="wrap">
                    <div className="meta">
                        <p className="name"> <img id="profile-img"  className="online"  src="/img/group.png" alt="" /> {this.props.discussion.name}</p>
                        <p className="preview">{this.props.message}</p>
                    </div>
                </div>:
                 <div value={discussion.username} onClick={this.changeDis}  className="wrap">
                    <span className="contact-status online"></span>
                    {(discussion.src!=="") ? (<img src="/img/Default.png"alt="" />):( <img src={discussion.src} alt="" />)}
                    <div className="meta">
                        <p className="name">{discussion.lastname}  {discussion.firstname}</p>
                        <p className="preview">{this.props.message}</p>
                    </div>
                </div>
        );
    }
}

export default UserCard;  