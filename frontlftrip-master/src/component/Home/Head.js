import React from 'react';
import Button from 'react-bootstrap/Button';
import Carousel from 'react-bootstrap/Carousel'

class Head extends React.Component{

    render() {
        return (
            <React.Fragment>
                <div className="cover mb-5"> 
                    <Carousel interval="8000" pause="false">
                        <Carousel.Item>
                            <div className="view slide1">

                                <div className="mask-home d-flex justify-content-end align-items-center">

                                    <div className="text-right white-text mx-5">
                                        <h1 className="mb-4">
                                            <strong>A la recherche
                                            <div className="green-text highlight"> d'une nouvelle aventure ?</div></strong>
                                        </h1>

                                        <p className="mb-4 d-none d-md-block">
                                            <strong>Participe aux nouvelles histoires de notre communauté</strong>
                                        </p>

                                        <Button variant="outline-light" href="/trips" className="green-hover">Je cherche mon voyage idéal</Button>
                                    </div>

                                </div>

                            </div>
                        </Carousel.Item>
                        <Carousel.Item>
                            <div className="view slide2">

                                <div className="mask-home d-flex justify-content-end align-items-center">

                                    <div className="text-right white-text mx-5">
                                        <h1 className="mb-4">
                                            <strong>Organise ton futur voyage avec
                                            <div className="green-text highlight">des voyageurs qui te ressemblent</div></strong>
                                        </h1>

                                        <p className="mb-4 d-none d-md-block">
                                            <strong>Crée ta propre aventure avec nous</strong>
                                        </p>

                                        <Button variant="outline-light" href="/createtrip" className="btn border border-light rounded white-text green-hover">Je crée mon voyage</Button>
                                    </div>
                                </div>

                            </div>
                        </Carousel.Item>
                        <Carousel.Item>
                            <div className="view slide3">

                                <div className="mask-home d-flex justify-content-end align-items-center">

                                    <div className="text-right white-text mx-5">
                                        <h1 className="mb-4">
                                            <strong>Profitez de l'instant présent
                                            <div className="green-text highlight">Avec le meilleur des guides</div></strong>
                                        </h1>

                                        <p className="mb-4 d-none d-md-block">
                                            <strong>Une personne locale pour vous montrer les perles rares</strong>
                                        </p>

                                        <Button variant="outline-light" href="/trips" className="btn border border-light rounded white-text green-hover">A la recherche d'un guide</Button>
                                    </div>

                                </div>

                            </div>
                        </Carousel.Item>
                    </Carousel>
                </div>
            </React.Fragment>
        );
      }
}


export default Head;