import React from 'react';
import { Router, Route, Switch } from 'react-router-dom';
import ReactNotification from 'react-notifications-component';
import history from "../../history";

import SignUpPage from '../Sign/SignUp';
import { SignInForm } from '../Sign/SignIn';
import Home from '../Home';
import Account from '../Account';
import Navigation from '../Navigation';
import {Createtrip, Updatetrip} from '../Trip';
import TripList from '../Trip/TripList';
import Footer from '../Footer';
import * as ROUTES from '../../const/routes';
import { withAuthentification } from '../Session';
import TripReviews from '../TripReviews';
import PasswordForgetPage from '../Password/PasswordForget/index';
import { ContactPage } from '../OtherPages/ContactPage';
import paswchange from '../Password/PasswordForget/paswchange';
//import Notification from '../Notification'; /*NOT USED CLEAR WARNING*/
import MyTripList from '../Trip/MyTripList';
import tripCard from '../Trip/tripCard';
import Message from '../Message';
import Notfound from '../OtherPages/404';
import ListTripLiked from '../Trip/ListTripLiked';
import ListTripParticipated from '../Trip/ListTripParticipated';
import ListUserByTrip from '../Trip/ListUserByTrip';
import Concept from '../OtherPages/Concept';
import CGU from '../OtherPages/CGU';
import PolitqueConf from '../OtherPages/PolitiqueConf';
import Wizard from '../Preferences/Preferences';
import ProfilSummary from "../Account/ProfileSummary";
import ProfileLikeSummary from "../Account/ProfileLikeSummary";


const App = () => (
	<Router history={history}>
		<div>
			<ReactNotification />
			<Navigation />

			
			<div>
			<Switch>
				<Route exact path={ROUTES.HOME} component={Home} />
				<Route path={ROUTES.SIGN_UP} component={SignUpPage} />
				<Route path={ROUTES.SIGN_IN} component={SignInForm} />
				<Route path={ROUTES.TRIPS} component={TripList} />
				<Route path={ROUTES.TRIP} component={tripCard} />
				<Route path={ROUTES.USERTRIPS} component={MyTripList} />
				<Route path={ROUTES.USERLIKES} component={ListTripLiked} />
				<Route path={ROUTES.USERPARTICIPATION} component={ListTripParticipated} />
				<Route path={ROUTES.LISTPARTICIPANTS} component={ListUserByTrip} />
				<Route path={ROUTES.CREATETRIP} component={Createtrip} />
				<Route path={ROUTES.TRIPSREVIEWS} component={TripReviews} />
				<Route path={ROUTES.PASSWORD_FORGET} component={PasswordForgetPage} />
				<Route path={ROUTES.CONTACT} component={ContactPage} />
				<Route exact path={ROUTES.ACCOUNT} component={Account} />
				<Route path={ROUTES.RESETPASSWORD} component={paswchange} />
				<Route path={ROUTES.CONCEPT} component={Concept} />
				<Route path={ROUTES.UPDATETRIP} component={Updatetrip} />
				<Route path={ROUTES.MESSAGERIE} component={Message} />
				<Route path={ROUTES.PREFERENCES} component={Wizard} />
				<Route path={ROUTES.CGU} component={CGU} />
				<Route path={ROUTES.POLITQUE} component={PolitqueConf} />
				<Route path={ROUTES.PROFILSUMMARY} component={ProfilSummary} />
				<Route path={ROUTES.PROFILELIKESUMMARY} component={ProfileLikeSummary}/>
				<Route component={Notfound} />
			</Switch>
			</div>
			<Footer/>
		</div>
	</Router>
);


export default withAuthentification(App);
