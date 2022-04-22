import React from 'react';
import ReactContactForm from 'react-mail-form';

const ContactPage = () => (
    <div className="fluid view-fixed">
        <div className="d-flex padding-top-10 justify-content-center align-items-center">
            <h5>Contactez nous :</h5>
        </div>
        <div className="d-flex padding-top-5 padding-bottom-10 justify-content-center align-items-center">
        <ReactContactForm buttonText="Envoyer" className="contact" to="contact-lftrip@gmail.com" titlePlaceholder="Sujet" contentsPlaceholder="Message et informations de contact"/>
        </div>
    </div>
    
);

export {ContactPage};