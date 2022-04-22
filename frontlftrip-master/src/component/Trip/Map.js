import React, { Component, Fragment } from 'react';
import {Button, Card} from 'react-bootstrap';
import { withRouter } from 'react-router-dom';
import { compose } from 'recompose';
import { withBackaccessContext } from '../BackEnd';
import { Map, TileLayer, Marker, Popup } from 'react-leaflet';
import L from 'leaflet';

delete L.Icon.Default.prototype._getIconUrl;

L.Icon.Default.mergeOptions({
    iconRetinaUrl: require('leaflet/dist/images/marker-icon-2x.png'),
    iconUrl: require('leaflet/dist/images/marker-icon.png'),
    shadowUrl: require('leaflet/dist/images/marker-shadow.png')
});


const MyPopupMarker = ({ id, latitude, longitude, country, title, base64 }) => (
  <Marker position={[latitude, longitude]}>
    <Popup> 
    <Card style={{ width: '18rem' }}>
        {base64 &&
            <Card.Img src={base64} alt="Trip custom"/>
        }

        {!base64 &&
            <Card.Img src="img/card1.jpg" alt="Trip default"/>
        }   
        <Card.Body>
          <Card.Title>{title}</Card.Title>
          <Card.Text>
          {country} 
          </Card.Text>
          <Button variant="outline-dark" href={`trip/${id}`}>Aller voir les détails du voyage</Button>
        </Card.Body>
      </Card>
    </Popup>
  </Marker>
)

const MyMarkersList = ({markers})  => {
  const items = markers.map(({ uuid, ...props }) => (
    <MyPopupMarker key={uuid} {...props} />
  ))
  return <Fragment>{items}</Fragment>
}


class MapLeaflet extends Component{
  state = {
    trips:[],
    uri: 'trips',
  }

  componentDidMount() {
    this.props.backaccess
    .getData(this.state.uri)
    .then((data) => {
        this.setState({ trips: data.response});
        //console.log(this.state.trips);
    }).catch(error => {
        console.log("not ok");
        this.setState({ error });
    });
};

  render() {
    return (
      <Map center={[46.227638, 2.213749]} zoom={4} style={{ width: '100%', height: '500px'}}>
        <TileLayer
          attribution='&amp;copy <a href="http://osm.org/copyright">OpenStreetMap</a> contributors'
          url="https://{s}.tile.openstreetmap.fr/osmfr/{z}/{x}/{y}.png"
        />
        <MyMarkersList markers={this.state.trips} />
      </Map>
    )
  }
}

export default  compose(
  withRouter,
  withBackaccessContext,
)(MapLeaflet);