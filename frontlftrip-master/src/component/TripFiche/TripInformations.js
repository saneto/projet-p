import React from 'react';
import Card from 'react-bootstrap/Card';
import Button from 'react-bootstrap/Button';

class TripInformations extends React.Component{

    render(){
        return (
            <React.Fragment>
                <h2 id="list">Liste des trips de la communauté !  </h2>
                <Card>
                    <Card.Header as="h5">Featured</Card.Header>
                    <Card.Body>
                        <Card.Title>Special title treatment</Card.Title>
                        <Card.Text>
                        With supporting text below as a natural lead-in to additional content.
                        </Card.Text>
                        <Button variant="primary">Go somewhere</Button>
                    </Card.Body>
                </Card>
                
            </React.Fragment>
        );
    }
}

export default TripInformations;