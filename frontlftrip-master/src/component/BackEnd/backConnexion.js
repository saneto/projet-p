import  API_ROUTE, {CONF} from '../../const/params';
import axios from "axios";
import {store} from "react-notifications-component";
import history from "../../history";

const headConfig = {
    headers: { 
        Accept: 'application/json',
        'Content-Type': 'application/json',
        Authorization: `Bearer ${ localStorage.getItem("token")}` }
};

export default class BackConnexion{
    constructor(){
        this.authListener = {}
    }

    doCreateUserWithEmailAndPassword = (data, uri) =>{

        return new Promise((resolve, reject) => { resolve(axios.post(
        `${API_ROUTE}/users`,
            {
                "email": data.email,
                "password" : data.password,
                "firstname" : data.firstname,
                "lastname" : data.lastname,
                "is_guide": "false"
            }).then(response=> {
                if(response.status === 201){
                    store.addNotification({
                        title: 'Inscription réussie',
                        message: 'Bienvenue sur LFTrip',
                        type: 'success',                         // 'default', 'success', 'info', 'warning'
                        container: 'top-right',                // where to position the notifications
                        animationIn: ["animated", "fadeIn"],     // animate.css classes that's applied
                        animationOut: ["animated", "fadeOut"],   // animate.css classes that's applied
                        dismiss: {
                            duration: 5000,
                            onScreen: true
                        }
                    })
                    //console.log(response.data)
                    localStorage.setItem('authUser',  JSON.stringify(response.data.response));
                    localStorage.setItem('token',  response.data.response.token);
                    history.push("/preferences");  
                    window.location.reload(false);
                }
            }).catch(error =>{
                const listerr = error.response.data.error;
                for (let [key, value] of Object.entries(listerr)) {
                    store.addNotification({
                        title: `Erreur: ${key}`,
                        message: `${value}`,
                        type: 'danger',                         // 'default', 'success', 'info', 'warning'
                        container: 'top-right',                // where to position the notifications
                        animationIn: ["animated", "fadeIn"],     // animate.css classes that's applied
                        animationOut: ["animated", "fadeOut"],   // animate.css classes that's applied
                        dismiss: {
                            duration: 5000,
                            onScreen: true
                        }
                    })
                };
            })) });
    }

    dosignWithEmailAndPassword = (data, uri) =>{

        return new Promise((resolve, reject) => { resolve(axios.post(
        `${API_ROUTE}/login`,
            {
                "email": data.email,
                "password" : data.password
            }).then(response=> {
                if(response.status === 200){
                    store.addNotification({
                        title: 'Connexion réussie',
                        message: 'Bienvenue sur LFTrip',
                        type: 'success',                         // 'default', 'success', 'info', 'warning'
                        container: 'top-right',                // where to position the notifications
                        animationIn: ["animated", "fadeIn"],     // animate.css classes that's applied
                        animationOut: ["animated", "fadeOut"],   // animate.css classes that's applied
                        dismiss: {
                            duration: 1500
                        }
                    })

                    localStorage.setItem('authUser',  JSON.stringify(response.data.response));
                    localStorage.setItem('token',  response.data.response.token);
                    history.push("/");
                    window.location.reload(false);
                }else{
                    reject( console.log(response.data));
                }
            }).catch(error =>{
                const listerr = error.response.data.error;
                for (let [key, value] of Object.entries(listerr)) {
                    store.addNotification({
                        title: `Erreur: ${key}`,
                        message: `${value}`,
                        type: 'danger',                         // 'default', 'success', 'info', 'warning'
                        container: 'top-right',                // where to position the notifications
                        animationIn: ["animated", "fadeIn"],     // animate.css classes that's applied
                        animationOut: ["animated", "fadeOut"],   // animate.css classes that's applied
                        dismiss: {
                            duration: 5000,
                            onScreen: true
                        }
                    })
                };
            })) });
    }

    doPasswordReset = (data, uri) =>{

        return new Promise((resolve, reject) => { resolve(axios.post(
        `${API_ROUTE}/password/forgot`,
            {
                "email": data.email,
            }).then(response=> {
                if(response.status === 200){
                    //console.log(response.data);
                    store.addNotification({
                        title: 'Email envoyé avec succès',
                        message: 'Veuillez vérifier votre boite de passagerie et de spam',
                        type: 'success',                         // 'default', 'success', 'info', 'warning'
                        container: 'top-right',                // where to position the notifications
                        animationIn: ["animated", "fadeIn"],     // animate.css classes that's applied
                        animationOut: ["animated", "fadeOut"],   // animate.css classes that's applied
                        dismiss: {
                            duration: 5000,
                            onScreen: true
                        }
                    })
                    data = {
                        data : response.data
                    }
                    localStorage.setItem('authUser',  JSON.stringify(data));
                    history.push("/");
                }else{
                    ///reject( console.log(response.data));
                }
            }).catch(error =>{
                const listerr = error.response.data.error;
                for (let [key, value] of Object.entries(listerr)) {
                    store.addNotification({
                        title: `Erreur: ${key}`,
                        message: `${value}`,
                        type: 'danger',                         // 'default', 'success', 'info', 'warning'
                        container: 'top-right',                // where to position the notifications
                        animationIn: ["animated", "fadeIn"],     // animate.css classes that's applied
                        animationOut: ["animated", "fadeOut"],   // animate.css classes that's applied
                        dismiss: {
                            duration: 5000,
                            onScreen: true
                        }
                    })
                };
            })) });
    }

    onlogOut = () =>{
        store.addNotification({
            message: 'Déconnexion réussie',
            type: 'success',                         // 'default', 'success', 'info', 'warning'
            container: 'top-right',                // where to position the notifications
            animationIn: ["animated", "fadeIn"],     // animate.css classes that's applied
            animationOut: ["animated", "fadeOut"],   // animate.css classes that's applied
            slidingExit: {
                duration: 5000,
                timingFunction: 'ease-out',
                delay: 0
              }
        })
        localStorage.removeItem('authUser');
        localStorage.removeItem('token');
        history.push("/");
        window.location.reload(false);
        history.push("/");
    }

    
    getDataWithToken= (uri) =>{
        return new Promise((resolve, reject) => { 
            resolve(
            axios.get(
                `${API_ROUTE}/${CONF["uri"][uri]}`
                ,headConfig).then(response => {
                if(response.status === 200){
                    return response.data;
                }else{
                    reject( response.status);
                }
            }).catch(error =>{
                reject( error);
            }))});
    }

    postData= (data, uri) =>{

        return new Promise((resolve, reject) => {

            axios.post(
                `${API_ROUTE}/${CONF["uri"][uri]}`,
            data
            , headConfig).then(response => {
                resolve(response)
            }).catch(error =>{
                const listerr = error.response.data.error;
                for (let [key, value] of Object.entries(listerr)) {
                    store.addNotification({
                        title: `Erreur: ${key}`,
                        message: `${value}`,
                        type: 'danger',                         // 'default', 'success', 'info', 'warning'
                        container: 'top-right',                // where to position the notifications
                        animationIn: ["animated", "fadeIn"],     // animate.css classes that's applied
                        animationOut: ["animated", "fadeOut"],   // animate.css classes that's applied
                        dismiss: {
                            duration: 3000,
                            onScreen: true
                        }
                    })
                }
            })});
    }

    postDataWithId= (data, uri) =>{
       /* let headConfig = {
            headers: {
                Accept: 'application/json',
                'Content-Type': 'application/json',
                'authorization': 'Bearer ' + data.token,
            }
        }        */
        return new Promise((resolve, reject) => {

            axios.post(
                `${API_ROUTE}/${CONF["uri"][uri]}/${data.id}`,
            {
                body: data.data,
            },headConfig).then(response => {
                
                    store.addNotification({
                        title: 'Commentaire envoyé',
                        message: 'Commentaire envoyé',
                        type: 'success',                         // 'default', 'success', 'info', 'warning'
                        container: 'bottom-center',                // where to position the notifications
                        animationIn: ["animated", "fadeIn"],     // animate.css classes that's applied
                        animationOut: ["animated", "fadeOut"],   // animate.css classes that's applied
                        dismiss: {
                            duration: 3000,
                            onScreen: true
                        }
                    });
                
                resolve(response)
            }).catch(error =>{
                reject( error);
            })});
    }

    editCommentary= (data, uri) =>{
        return new Promise((resolve, reject) => {

            axios.put(
                `${API_ROUTE}/${CONF["uri"][uri]}/${data.id}`,
                {
                    body: data.data,
                },headConfig).then(response => {
                if(response.status === 200){
                    store.addNotification({
                        title: 'Commentaire mis à jour',
                        message: 'Commentaire mis à jour',
                        type: 'success',                         // 'default', 'success', 'info', 'warning'
                        container: 'top-right',                // where to position the notifications
                        animationIn: ["animated", "fadeIn"],     // animate.css classes that's applied
                        animationOut: ["animated", "fadeOut"],   // animate.css classes that's applied
                        dismiss: {
                            duration: 5000,
                            onScreen: true
                        }
                    });
                }

                resolve(response)
            }).catch(error =>{
                reject( error);
            })});
    }

    getDataWitId = (id, uri) =>{
        return new Promise(
            (resolve, reject) => {
                resolve(
                    axios.get(
                        `${API_ROUTE}/${CONF["uri"][uri]}/${id}`,
                        {
                            id : id
                        }).then(response=> {
                            //console.log(response)
                            if(response.status === 200){
                                //console.log(response.data)
                                return response.data;
                            }else{
                                //console.log(response.status)
                                reject( response.data);
                            }
                        }).catch(error =>{
                            //console.log(error)
                            reject( error);
                        })
                )
            });
    }

    getDataWitParam = (data, uri) =>{        
        return new Promise(
            (resolve, reject) => {
                resolve(
                    axios.post(
                        `${API_ROUTE}/${CONF["uri"][uri]}`,  data).then(response=> {
                            //console.log(response)
                            if(response.status === 200){
                                //console.log(response.data)
                                return response.data;
                            }else{
                                //console.log(response.status)
                                reject( response.data);
                            }
                        }).catch(error =>{
                            console.log(error)
                            reject( error);
                        })
                )
            });
    }


    getData = ( uri) =>{        
        return new Promise((resolve, reject) => { 
            resolve(axios.get(
                `${API_ROUTE}/${CONF["uri"][uri]}`
            ).then(response=> {

                if(response.status === 200){
                    return response.data;
                }else{
                    reject( response.data);
                }
            }).catch(error =>{
                reject( error);
            }))
        });
    }

    updateData = (data, uri,id) =>{        
        return new Promise((resolve, reject) => { 
            axios.put(
                `${API_ROUTE}/${CONF["uri"][uri]}/${id}`,
                data,
                headConfig
            )
                .then(response=> {
                    if(response.status === 200){
                        store.addNotification({
                            title: 'Mise à jour',
                            message: 'Vos modifications ont été prise en compte',
                            type: 'success',                         // 'default', 'success', 'info', 'warning'
                            container: 'top-right',                // where to position the notifications
                            animationIn: ["animated", "fadeIn"],     // animate.css classes that's applied
                            animationOut: ["animated", "fadeOut"],   // animate.css classes that's applied
                            dismiss: {
                                duration: 3000,
                                onScreen: true
                            }
                        })
                        history.goBack();
                    }
                }).catch(error =>{
                const listerr = error.response.data.error;
                for (let [key, value] of Object.entries(listerr)) {
                    store.addNotification({
                        title: `Erreur: ${key}`,
                        message: `${value}`,
                        type: 'danger',                         // 'default', 'success', 'info', 'warning'
                        container: 'top-right',                // where to position the notifications
                        animationIn: ["animated", "fadeIn"],     // animate.css classes that's applied
                        animationOut: ["animated", "fadeOut"],   // animate.css classes that's applied
                        dismiss: {
                            duration: 5000,
                            onScreen: true
                        }
                    })
                }
            })
        });
    }

    updateUserPreferenceData = (data, uri,id) =>{        
        return new Promise((resolve, reject) => { 
            axios.put(
                `${API_ROUTE}/${CONF["uri"][uri]}/${id}`,
                data,
                headConfig
            )
                .then(response=> {
                    if(response.status === 200){
                        store.addNotification({
                            title: 'Mise à jour',
                            message: 'vos modifications ont été prises en compte',
                            type: 'success',                         // 'default', 'success', 'info', 'warning'
                            container: 'top-right',                // where to position the notifications
                            animationIn: ["animated", "fadeIn"],     // animate.css classes that's applied
                            animationOut: ["animated", "fadeOut"],   // animate.css classes that's applied
                            dismiss: {
                                duration: 5000,
                                onScreen: true
                            }
                        })
                        history.push("/account");
                    }
                }).catch(error =>{
                const listerr = error.response.data.error;
                for (let [key, value] of Object.entries(listerr)) {
                    store.addNotification({
                        title: `Erreur: ${key}`,
                        message: `${value}`,
                        type: 'danger',                         // 'default', 'success', 'info', 'warning'
                        container: 'top-right',                // where to position the notifications
                        animationIn: ["animated", "fadeIn"],     // animate.css classes that's applied
                        animationOut: ["animated", "fadeOut"],   // animate.css classes that's applied
                        dismiss: {
                            duration: 5000,
                            onScreen: true
                        }
                    })
                }
            })
        });
    }

    updateUserPassword = (data) =>{
        //console.log(data.token)
/*         let headConfig = {
            headers: {
                Accept: 'application/json',
                'Content-Type': 'application/json',
                'authorization': 'Bearer ' + data.token,
            }
        } */
        return new Promise((resolve, reject) => { resolve(axios.post(
            `${API_ROUTE}/password/userreset`,
            {
                    'email' : data.email,
                    'current_password' : data.current_password,
                    'new_password' : data.new_password,
                    'retype_password' : data.retype_password,
            }, headConfig
            )
            .then(response=> {
            if(response.status === 200){
                store.addNotification({
                    title: 'Mot de passe mis à jour',
                    message: 'Mot de passe mis à jour',
                    type: 'success',                         // 'default', 'success', 'info', 'warning'
                    container: 'top-right',                // where to position the notifications
                    animationIn: ["animated", "fadeIn"],     // animate.css classes that's applied
                    animationOut: ["animated", "fadeOut"],   // animate.css classes that's applied
                    dismiss: {
                        duration: 5000,
                        onScreen: true
                    }
                })
                localStorage.removeItem('authUser');
                history.push("/login");
            }
        }).catch(error =>{
                const listerr = error.response.data.error;
                for (let [key, value] of Object.entries(listerr)) {
                    store.addNotification({
                        title: `Erreur: ${key}`,
                        message: `${value}`,
                        type: 'danger',                         // 'default', 'success', 'info', 'warning'
                        container: 'top-right',                // where to position the notifications
                        animationIn: ["animated", "fadeIn"],     // animate.css classes that's applied
                        animationOut: ["animated", "fadeOut"],   // animate.css classes that's applied
                        dismiss: {
                            duration: 5000,
                            onScreen: true
                        }
                    })
                }
            })) });
    };

    deleteUserAccount = (data) =>{
        //console.log(data.token)
/*         let headConfig = {
            headers: {
                Accept: 'application/json',
                'Content-Type': 'application/json',
                'authorization': 'Bearer ' + data.token,
            }
        } */
        return new Promise((resolve, reject) => { resolve(axios.delete(
            `${API_ROUTE}/users/${data.id}`, headConfig
        )
            .then(response=> {
                if(response.status === 200){
                    store.addNotification({
                        title: 'Compte supprimé',
                        message: 'Votre compte a bien été supprimé',
                        type: 'success',                         // 'default', 'success', 'info', 'warning'
                        container: 'top-right',                // where to position the notifications
                        animationIn: ["animated", "fadeIn"],     // animate.css classes that's applied
                        animationOut: ["animated", "fadeOut"],   // animate.css classes that's applied
                        dismiss: {
                            duration: 5000,
                            onScreen: true
                        }
                    })
                    localStorage.removeItem('authUser');
                    history.push("/login");
                }
            }).catch(error =>{
                const listerr = error.response.data.error;
                for (let [key, value] of Object.entries(listerr)) {
                    store.addNotification({
                        title: `Erreur: ${key}`,
                        message: `${value}`,
                        type: 'danger',                         // 'default', 'success', 'info', 'warning'
                        container: 'top-right',                // where to position the notifications
                        animationIn: ["animated", "fadeIn"],     // animate.css classes that's applied
                        animationOut: ["animated", "fadeOut"],   // animate.css classes that's applied
                        dismiss: {
                            duration: 5000,
                            onScreen: true
                        }
                    })
                }
            })) });
    };

    deleteUserComment = (data) =>{
/*         let headConfig = {
            headers: {
                Accept: 'application/json',
                'Content-Type': 'application/json',
                'authorization': 'Bearer ' + data.token,
            }
        }; */
        return new Promise((resolve, reject) => { resolve(axios.delete(
            `${API_ROUTE}/comments/${data.id}`, headConfig
        )
            .then(response=> {
                if(response.status === 200){
                    store.addNotification({
                        title: 'Commentaire supprimé',
                        message: 'ce commentaire a été supprimé',
                        type: 'danger',                         // 'default', 'success', 'info', 'warning'
                        container: 'bottom-center',                // where to position the notifications
                        animationIn: ["animated", "fadeIn"],     // animate.css classes that's applied
                        animationOut: ["animated", "fadeOut"],   // animate.css classes that's applied
                        dismiss: {
                            duration: 5000,
                            onScreen: true
                        }
                    })
                }
            }).catch(error =>{
                const listerr = error.response.data.error;
                for (let [key, value] of Object.entries(listerr)) {
                    store.addNotification({
                        title: `Erreur: ${key}`,
                        message: `${value}`,
                        type: 'danger',                         // 'default', 'success', 'info', 'warning'
                        container: 'top-right',                // where to position the notifications
                        animationIn: ["animated", "fadeIn"],     // animate.css classes that's applied
                        animationOut: ["animated", "fadeOut"],   // animate.css classes that's applied
                        dismiss: {
                            duration: 5000,
                            onScreen: true
                        }
                    })
                }
            })) });
    };

    deleteUserTrip = (data, uri) =>{
        return new Promise((resolve, reject) => { resolve(axios.delete(
            `${API_ROUTE}/${CONF["uri"][uri]}/${data.id}`, headConfig
        )
            .then(response=> {
                if(response.status === 200){
                    store.addNotification({
                        title: 'Voyage supprimé',
                        message: 'ce voyage a été supprimé',
                        type: 'success',                         // 'default', 'success', 'info', 'warning'
                        container: 'top-right',                // where to position the notifications
                        animationIn: ["animated", "fadeIn"],     // animate.css classes that's applied
                        animationOut: ["animated", "fadeOut"],   // animate.css classes that's applied
                        dismiss: {
                            duration: 5000,
                            onScreen: true
                        }
                    });
                    history.push("/trips");
                }
            }).catch(error =>{
                const listerr = error.response.data.error;
                for (let [key, value] of Object.entries(listerr)) {
                    store.addNotification({
                        title: `Erreur: ${key}`,
                        message: `${value}`,
                        type: 'danger',                         // 'default', 'success', 'info', 'warning'
                        container: 'top-right',                // where to position the notifications
                        animationIn: ["animated", "fadeIn"],     // animate.css classes that's applied
                        animationOut: ["animated", "fadeOut"],   // animate.css classes that's applied
                        dismiss: {
                            duration: 5000,
                            onScreen: true
                        }
                    })
                }
            })) });
    };

    deleteData= (data, uri) =>{
        return new Promise((resolve, reject) => { resolve(axios.delete(
            `${API_ROUTE}/${CONF["uri"][uri]}`,
            {
                data
            }).then(response=> {

                if(response.status === 200){
                    return response.status;
                }else{
                    reject( response.status);
                }
            }).catch(error =>{
                reject( console.log(error));
            }))
        });
    };

    resetPassword= (data) =>{
        //console.log(data.token)
        let headConfig = {
            headers: {
                Accept: 'application/json',
                'Content-Type': 'application/json',
            }
        }
        return new Promise((resolve, reject) => { 
            resolve(axios.post(
            `${API_ROUTE}/password/reset`,
            {
                    'token' : data.token,
                    'new_password' : data.new_password,
                    'retype_password' : data.retype_password,
            }, headConfig
            )
            .then(response=> {
            if(response.status === 200){
                store.addNotification({
                    title: 'Mot de passe mis à jour',
                    message: 'Mot de passe mis à jour',
                    type: 'success',                         // 'default', 'success', 'info', 'warning'
                    container: 'top-right',                // where to position the notifications
                    animationIn: ["animated", "fadeIn"],     // animate.css classes that's applied
                    animationOut: ["animated", "fadeOut"],   // animate.css classes that's applied
                    dismiss: {
                        duration: 5000,
                        onScreen: true
                    }
                })
                localStorage.removeItem('authUser');
                history.push("/login");
            }
        }).catch(error =>{
                const listerr = error.response.data.error;
                for (let [key, value] of Object.entries(listerr)) {
                    store.addNotification({
                        title: `Erreur: ${key}`,
                        message: `${value}`,
                        type: 'danger',                         // 'default', 'success', 'info', 'warning'
                        container: 'top-right',                // where to position the notifications
                        animationIn: ["animated", "fadeIn"],     // animate.css classes that's applied
                        animationOut: ["animated", "fadeOut"],   // animate.css classes that's applied
                        dismiss: {
                            duration: 5000,
                            onScreen: true
                        }
                    })
                }
            })) 
        });
    }

    postLikeWithId= (data, uri) =>{
/*         let headConfig = {
            headers: {
                Accept: 'application/json',
                'Content-Type': 'application/json',
                'authorization': 'Bearer ' + data.token,
            }
        }     */    
        return new Promise((resolve, reject) => { 
                resolve(
                    axios.post(
                        `${API_ROUTE}/${CONF["uri"][uri]}/${data.id}`,
                    {
                    },headConfig).then(response => {

                        if(response.status === 201){
                            //console.log(response)
                        return response.data;
                        }else{
                            reject( response.status);
                        }
                    }).catch(error =>{
                        reject( error);
                    })
                )
            }
        );
    }

deleteLikeWithId= (data, uri) =>{
/*     let headConfig = {
        headers: {
            Accept: 'application/json',
            'Content-Type': 'application/json',
            'authorization': 'Bearer ' + data.token,
        }
    } */
    return new Promise((resolve, reject) => {
        resolve(
        axios.delete(
            `${API_ROUTE}/${CONF["uri"][uri]}/${data.id}`,headConfig)
            .then(response => {
            if(response.status === 200){
                return response.data;
            }else{
                reject( response.status);
            }
        }).catch(error =>{
            reject( error);
        }))});
}
    //CALL
    updateUserAccountInformation = (data) =>{
        //console.log(data.token)
/*         let headConfig = {
            headers: {
                Accept: 'application/json',
                'Content-Type': 'application/json',
                'authorization': 'Bearer ' + data.token,
            }
        } */
        return new Promise((resolve, reject) => { resolve(axios.put(
            `${API_ROUTE}/users/${data.id}`,
            {
                'email' : data.email,
                'firstname' : data.firstname,
                'lastname' : data.lastname,
                'city' : data.city,
                'date_of_birth' : data.date_of_birth,
                'phone_number' : data.phone_number,
                'departure_airport' : data.airport,
                'description' : data.description,
                'is_guide' : data.is_guide,
                'country_guide' : data.country_guide,
                'avatar_path' : data.avatar_path
            }, headConfig
        )
        .then(response=> {
            if(response.status === 200){
                store.addNotification({
                    title: 'Informations de compte modifiés',
                    message: 'Les changements ont bien été pris en compte',
                    type: 'success',                         // 'default', 'success', 'info', 'warning'
                    container: 'bottom-center',                // where to position the notifications
                    animationIn: ["animated", "fadeIn"],     // animate.css classes that's applied
                    animationOut: ["animated", "fadeOut"],   // animate.css classes that's applied
                    dismiss: {
                        duration: 5000,
                        onScreen: true
                    }
                })
            }
        }).catch(error =>{
            const listerr = error.response.data.error;
            for (let [key, value] of Object.entries(listerr)) {
                store.addNotification({
                    title: `Erreur: ${key}`,
                    message: `${value}`,
                    type: 'danger',                         // 'default', 'success', 'info', 'warning'
                    container: 'top-right',                // where to position the notifications
                    animationIn: ["animated", "fadeIn"],     // animate.css classes that's applied
                    animationOut: ["animated", "fadeOut"],   // animate.css classes that's applied
                    dismiss: {
                        duration: 5000,
                        onScreen: true
                    }
                })
            }
        })) });
    }

    //CALL
    getGuideListForTrip = (trip) =>{
        return new Promise((resolve, reject) => { resolve(axios.get(
            `${API_ROUTE}/users_guides/${trip.country}`
        )
        .then(response=> {
            return response;
        }).catch(error =>{
            const listerr = error.response.data.error;
            for (let [key, value] of Object.entries(listerr)) {
                store.addNotification({
                    title: `Erreur: ${key}`,
                    message: `${value}`,
                    type: 'danger',                         // 'default', 'success', 'info', 'warning'
                    container: 'top-right',                // where to position the notifications
                    animationIn: ["animated", "fadeIn"],     // animate.css classes that's applied
                    animationOut: ["animated", "fadeOut"],   // animate.css classes that's applied
                    dismiss: {
                        duration: 5000,
                        onScreen: true
                    }
                })
            }
        })) });
    }

    // send email to user to talk about a trip
    sendTripProposition= (data, uri) =>{
    //console.log(data);
        return new Promise((resolve, reject) => {

            axios.post(
                `${API_ROUTE}/${CONF["uri"][uri]}`,
                {
                    body: data.data,
                },headConfig).then(response => {
                if(response.status === 201){
                    store.addNotification({
                        title: 'Proposition de trip envoyé',
                        message: 'Votre proposition a bien été envoyé',
                        type: 'success',                         // 'default', 'success', 'info', 'warning'
                        container: 'top-right',                // where to position the notifications
                        animationIn: ["animated", "fadeIn"],     // animate.css classes that's applied
                        animationOut: ["animated", "fadeOut"],   // animate.css classes that's applied
                        dismiss: {
                            duration: 5000,
                            onScreen: true
                        }
                    });
                }

                resolve(response)
            }).catch(error =>{
                const listerr = error.response.data.error;
                for (let [key, value] of Object.entries(listerr)) {
                    store.addNotification({
                        title: `Erreur: ${key}`,
                        message: `${value}`,
                        type: 'danger',                         // 'default', 'success', 'info', 'warning'
                        container: 'top-right',                // where to position the notifications
                        animationIn: ["animated", "fadeIn"],     // animate.css classes that's applied
                        animationOut: ["animated", "fadeOut"],   // animate.css classes that's applied
                        dismiss: {
                            duration: 5000,
                            onScreen: true
                        }
                    })
                }
            })});
    }

}
