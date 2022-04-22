let API_ROUTE

if(process.env.NODE_ENV === 'development'){
    API_ROUTE = 'http://127.0.0.1:8080/api/v1'
}else{
    if(window.location.protocol !== 'https:'){
        API_ROUTE = 'http://lftrip.tk/api'
    }else{
        API_ROUTE = 'https://lftrip.tk/api'
    }
}


export default API_ROUTE

export const CONF = {
    uri : {
        "login" : 'login',
        "users" : 'users',
        "trips" : "trips",
        "trip" : "trip",
        "user_trips" : "user_trips",
        "comments" : "comments",
        "likes" : "likes",
        "chat" : "chat",
        "logout" : "logout",
        "elk" : "elk",
        "participation": "participation",
        "userlikes":"userlikes",
        "userapplies":"userapplies",
        "userpreferences":"userpreferences",
        "user_userpreferences":"user_userpreferences",
        "send_proposal": "send_proposal",
        "notification" : "notification",
        "elkM" : "elkMessage",
        "groupc" : "groupechat",
        "groupm" : "groupemessage",
        "message" : "messages"
    }
}
