package controllers

import "gitlab.com/esgi/lftrip/BackLFTrip/api/middlewares"

func (s *Server) initializeRoutes() {

	s.Melody.HandleMessage(s.pushMsg)
	s.Melody.HandleConnect(s.newConnection)

	v1 := s.Router.Group("/api/v1")
	{

		// Login Route
		v1.POST("/login", s.Login)

		// Reset password:
		v1.POST("/password/forgot", s.ForgotPassword)
		v1.POST("/password/reset", s.ResetPassword)
		v1.POST("/password/userreset", middlewares.TokenAuthMiddleware(), s.ResetUserPassword)

		//Users routes
		v1.POST("/users", s.CreateUser)
		v1.GET("/users", s.GetUsers)
		v1.GET("/users/:id", s.GetUser)
		v1.PUT("/users/:id", middlewares.TokenAuthMiddleware(), s.UpdateUser)
		v1.PUT("/avatar/users/:id", middlewares.TokenAuthMiddleware(), s.UpdateAvatar)
		v1.DELETE("/users/:id", middlewares.TokenAuthMiddleware(), s.DeleteUser)

		// Users Guide routes
		v1.GET("/users_guides/:country", s.GetUsersGuide)

		// Trip routes
		v1.POST("/trips", middlewares.TokenAuthMiddleware(), s.CreateTrip)
		v1.GET("/trips", s.GetTrips)
		v1.GET("/trips/:id", s.GetTrip)
		v1.PUT("/trips/:id", middlewares.TokenAuthMiddleware(), s.UpdateTrip)
		v1.DELETE("/trips/:id", middlewares.TokenAuthMiddleware(), s.DeleteTrip)
		v1.GET("/user_trips/:id", s.GetUserTrips)

		//Like routes
		v1.GET("/likes/:id", s.GetLikes)
		v1.POST("/likes/:id", middlewares.TokenAuthMiddleware(), s.LikeTrip)
		//v1.DELETE("/likes/:id", middlewares.TokenAuthMiddleware(), s.UnLikeTrip)
		v1.DELETE("/likes/:id", middlewares.TokenAuthMiddleware(), s.DeleteLike)

		v1.GET("/userlikes/:id", s.GetLikesByUserID)
		v1.GET("/likes", s.GetAllLikes)

		// Participation routes
		v1.GET("/participation/:id", s.GetParticipations)
		v1.POST("/participation/:id", middlewares.TokenAuthMiddleware(), s.ParticipationTrip)
		//v1.DELETE("/participation/:id", middlewares.TokenAuthMiddleware(), s.WithdrawParticipation)
		v1.DELETE("/participation/:id", middlewares.TokenAuthMiddleware(), s.DeleteParticipation)
		v1.GET("/userapplies/:id", s.GetParticipationByID)

		//Comment routes
		v1.POST("/comments/:id", middlewares.TokenAuthMiddleware(), s.CreateComment)
		v1.GET("/comments/:id", s.GetComments)
		v1.PUT("/comments/:id", middlewares.TokenAuthMiddleware(), s.UpdateComment)
		v1.DELETE("/comments/:id", middlewares.TokenAuthMiddleware(), s.DeleteComment)

		//ELK
		v1.POST("/searchtrips", s.searchTrips)
		v1.GET("/isertTrips", s.insertTestValue)
		v1.POST("/elk", s.searchELKTrip)

		//ws message
		v1.GET("/ws/messages", s.wsMessage)
		v1.GET("/messages/:id", s.getMessage)
		v1.POST("/messages", s.getELKConversation)
		v1.GET("/chat", middlewares.TokenAuthMiddleware(), s.getUserChat)
		v1.POST("/chat", middlewares.TokenAuthMiddleware(), s.createChat)

		//group part
		v1.POST("/groupechat", middlewares.TokenAuthMiddleware(), s.createGroupeChat)
		v1.POST("/groupemessage",middlewares.TokenAuthMiddleware(), s.appendGroupeMessage)


		//UserPreferences routes
		v1.POST("/userpreferences", middlewares.TokenAuthMiddleware(), s.CreateUserPreferences)
		v1.GET("/userpreferences", s.GetAllUserPreferences)
		v1.GET("/userpreferences/:id", s.GetUserPreference)
		v1.PUT("/userpreferences/:id", middlewares.TokenAuthMiddleware(), s.UpdateUserPreference)
		v1.DELETE("/userpreferences/:id", middlewares.TokenAuthMiddleware(), s.DeleteUserPreference)
		v1.GET("/user_userpreferences/:id", s.GetUserPreferenceByUser)

		// Send proposal trip
		v1.POST("/send_proposal", middlewares.TokenAuthMiddleware(), s.CreateProposal)
		v1.GET("/notification/:id", s.getELKNotification)

		//ws notification
		v1.GET("/ws/notification", s.wsNotification)
		v1.GET("/elkMessage/:id", s.getELKMessages)
	}
}
