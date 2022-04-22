import React from 'react';
import Comment from './Comment';
import CommentForm from './CreateCommentaire';
import { withRouter  } from 'react-router-dom';
import { compose } from 'recompose';
import { withBackaccessContext } from '../BackEnd';
import { Card } from 'react-bootstrap';

const user = JSON.parse(localStorage.getItem("authUser"));

class  ListComments extends React.Component{
    constructor(props){
        super(props);
        this.state= {
            comments : [],
            uri: 'comments',
            id: '',
            token: '',
            currentUser: ''
        };
        this.state.token = user.token;
        this.state.currentUser = user.id;
        this.state.id =  this.props.match.params.id;
    }
    componentDidMount() {
        this.props.backaccess
        .getDataWitId(this.state.id,this.state.uri)
        .then((data) => {
            this.setState({comments: data.response});
        }).catch(error => {
            console.log("not ok");
            this.setState({ error });
        });
    };

    onComment = (commentText) =>{
        var joined = [commentText].concat(this.state.comments);
        this.setState({ comments: joined })
    };

    deleteCommentary = (commentId) => event => {
        const {token} = this.state;
        this.props.backaccess.deleteUserComment({id : commentId, token : token})
            .then(() => {
                let commentStep = this.state.comments.filter(comment => comment.id !== commentId);
                this.setState({comments: commentStep});
            }).catch(error => {
            this.setState({ error });
        });
        event.preventDefault();
    };

    render(){
        const {comments} = this.state;
        //console.log(comments);
        return(
            <React.Fragment>
                <CommentForm  onComment= {this.onComment}/>
                <div className="card card-outline-secondary my-4">
                    <Card.Header as="h5" className="trip-card-detail-header m-2">Les réactions de nos Tripiziens</Card.Header>
                    {comments.map(comment =>(
                        <Comment key={comment.id} comment={comment} deleteCom={this.deleteCommentary} id={comment.id}/>
                        )
                    )}
                </div>
            </React.Fragment>
        );
    }
}

export default  compose(
    withRouter,
    withBackaccessContext,
)(ListComments);
