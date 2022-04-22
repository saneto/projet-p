import React from 'react';

const InternalError1 = () => (
    <div>
    <h1>500</h1>
    <h3>Internal server error.</h3>
    <h5>Le serveur a rencontré un problème inattendu qui ne lui a pas permis de terminer la demande.</h5>
    <button href="/" className="btn">revenir au site</button>

    </div>
);


export default InternalError1;