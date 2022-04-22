import React from 'react';
import Button from 'react-bootstrap/Button';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faHandPeace, faSmile, faHeart } from '@fortawesome/free-solid-svg-icons'


const Concept = () => (
    <React.Fragment>
        <div className="title-top">
            <div className="titleContainerTrip container">
                <h1>Concept LFTrip</h1>
            </div>
        </div>

        <div className="container padding-top-5 padding-bottom-10 justify-content-center align-items-center">
            <h3>Une approche innovante du voyage de groupe</h3>
            <p>LFTRIP, c'est avant tout une nouvelle approche du voyage de groupe, avec la possibilité d'échanger avec ses futurs
            compagnons de voyage, de consulter leurs profils et ainsi de mieux les connaître avant de réserver sa place. 
            Fini les voyages où l'on ne sait pas avec qui on part avant d'arriver sur place.</p>
            <hr/>
            <h3>Des voyageurs aux mêmes attentes et motivations à voyager que toi.</h3>
            <p>Pars en petits groupes de voyageurs
            composés par affinités. Un mélange d'amitiés, de découvertes et d'expériences où l'humain prime.
            Que tu sois philosophe, boute-en-train ou explora  teur, pars avec des voyageurs comme toi. 
            LFTrip, c'est un état d'esprit, pas une question d'âge, de sexe ou de situation.</p>
            <Button variant="outline-dark" href="/register" className="btn">Nous rejoindre</Button>
            <hr/>
            <h2 className="mar-b-50 main-heading">Nos valeurs</h2>
            <p className="mar-b-50">A travers le partage de valeurs communes, 
            les expériences avec LFTRip permettent de vivre des moments de bonheur et des aventures extraordinaires.</p>
            <h3><FontAwesomeIcon icon={faHandPeace} /> Cohésion</h3>
            <p>Elle est centrale, c’est le coeur de LFTRIP. 
            Nous souhaitons replacer l’humain au centre du voyage car on a tendance à s’égarer dans la sphère du virtuel
            et de l’individualisme. Nous revenons aux basiques pour se rencontrer sur des valeurs partagées, 
            pour apprécier l’instant présent et vivre des expériences humaines intenses. 
            A plus long terme, la volonté de créer de vraies amitiés est notre finalité.</p>
            <hr/>
            <h3><FontAwesomeIcon icon={faSmile} />  Respect</h3>
            <p>La communauté LFTrip : les Tripiziens, est à l’image du monde qui nous entoure : unique, dynamique et variée. 
            Le respect est ce qui nous unit et nous permet de nous intégrer au sein d’un groupe. 
            Cette valeur est essentielle pour que chacun puisse se sentir partout accepté.</p>
            <hr/>
            <h3><FontAwesomeIcon icon={faHeart} /> Sincérité</h3>
            <p>Notre communauté repose sur la confiance et la sincérité. Nous attendons de nos voyageurs de trouver un juste
            équilibre entre attentes partagées, authenticité et échanges honnêtes.</p>
        </div>

    </React.Fragment>
);


export default Concept;