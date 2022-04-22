import React, { Component } from 'react';
import { compose } from 'recompose';
import { withBackaccessContext } from '../../BackEnd';
import {withRouter} from 'react-router-dom';
import Button from '@material-ui/core/Button';
import CssBaseline from '@material-ui/core/CssBaseline';
import TextField from '@material-ui/core/TextField';
import Grid from '@material-ui/core/Grid';
import Box from '@material-ui/core/Box';
import Typography from '@material-ui/core/Typography';
import Container from '@material-ui/core/Container';

const PasswordChangePage = () => (
    <section id="user-pwd-update">
        <PasswordChangeForm />
    </section>
);

const INITIAL_STATE = {
    new_password: '',
    retype_password: '',
    error: null,
    token:  '',
};

class PasswordChangeFormBase extends Component {

    constructor(props) {
        super(props);
        this.state = { ...INITIAL_STATE };
        this.state.token =  this.props.match.params.token;
    }
    onSubmit = event => {
        const { new_password } = this.state;
        const { retype_password } = this.state;
        const {token} = this.state;

        this.props.backaccess
            .resetPassword({token : token, new_password : new_password, retype_password : retype_password})
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
        return (
<Container className="" component="main" maxWidth="xs">
    <CssBaseline />

    <div className="">

        <Typography component="h1" variant="h5">
            Reset password
        </Typography>

        <form className="" onSubmit={this.onSubmit}>
            <Grid container spacing={2}>
                <Grid item xs={12}>
                    <TextField
                        variant="outlined"
                        required
                        fullWidth
                        id="new_password"
                        label="Nouveau mot de passe"
                        name="new_password"
                        autoComplete="new_password"
                        type="password"
                        value={this.new_password}
                        onChange={this.onChange}
                    />
                </Grid>
                <Grid item xs={12}>
                    <TextField
                        variant="outlined"
                        required
                        fullWidth
                        id="retype_password"
                        label="Vérification du nouveau mot de passe"
                        name="retype_password"
                        autoComplete="retype_password"
                        type="password"
                        value={this.retype_password}
                        onChange={this.onChange}
                    />
                </Grid>
            </Grid>
            <Button
                fullWidth
                type="submit"
                variant="contained"
                color="primary"
                className=""
            >
                Modifier mon mot de passe
            </Button>
        </form>
    </div>
    <Box mt={5}>
    </Box>
</Container>
/*             <div>
                <h3 className="title">Change Your Password</h3>
                <form onSubmit={this.onSubmit}  className="grey-box">
                    <span className="text">
                        There are many variations of passages of Lorem Ipsum available, but the majority have suffered alteration inble. If you are of going.
                    </span>
                    <input value={passwordOne} onChange={this.onChange} type="password" className="form-control" id="passwordOne" name="passwordOne" placeholder="New password"/>
                    <input value={passwordTwo} onChange={this.onChange} type="password" className="form-control" id="passwordTwo" name="passwordTwo" placeholder="Confirm new password"/>
						
                    <button className="btn btn-defaul" type="submit">
                        Reset My Password
                    </button>

                    {error && <p>{error.message}</p>}
					
                </form>
            </div> */
        );
    }
}

const PasswordChangeForm = compose(
    withRouter,
    withBackaccessContext,
)(PasswordChangeFormBase);

export default PasswordChangePage;
