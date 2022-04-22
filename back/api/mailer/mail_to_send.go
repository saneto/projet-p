package mailer

import (
	"net/http"
//	"os"

	"github.com/matcornic/hermes/v2"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type sendMail struct{}

// SendMailer : Interface
type SendMailer interface {
	SendResetPassword(string, string, string, string) (*EmailResponse, error)
	SendWelcome(string, string, string) (*EmailResponse, error)
	SendCreateTrip(string, string, string) (*EmailResponse, error)
	SendParticipationUser(string, string, string, string) (*EmailResponse, error)
	SendParticipationUserCreatedTrip(string, string, string, string, string, string) (*EmailResponse, error)
	SendTravelProposal(string, string, string, string, string, string, string) (*EmailResponse, error)
}

// var SendMailer of type SendMail
var (
	SendMail SendMailer = &sendMail{} //this is useful when we start testing
)

// EmailResponse : Struct
type EmailResponse struct {
	Status   int
	RespBody string
}

func (s *sendMail) SendTravelProposal(ToUserProposal string, ToUserProposalEmail string, ToUserInvited string, ToUserInvitedEmail string, MessageProposal string, FromAdmin string, Sendgridkey string) (*EmailResponse, error) {
	h := hermes.Hermes{
		Product: hermes.Product{
			Name: "LFTrip",
		},
	}

	var intro, intro2 string
	intro = "Un Tripizien, " + ToUserProposal + " aimerait vous invitez à participer à un voyage."
	intro2 = "Échangez avec lui par email : " + ToUserProposalEmail + ", vous pourrez ensuite communiquer via d'autres réseaux."

	email := hermes.Email{
		Body: hermes.Body{
			Name: ToUserInvited,
			Intros: []string{
				intro,
				intro2,
				MessageProposal,
			},
			Signature: "On te souhaite de faire pleins de bonnes rencontres et de voyages !",
		},
	}
	emailBody, err := h.GenerateHTML(email)
	if err != nil {
		return nil, err
	}
	from := mail.NewEmail("LFTrip", FromAdmin)
	subject := "Un utilisateur vous propose de voyager avec vous."
	to := mail.NewEmail("Demande de voyage", ToUserInvitedEmail)
	message := mail.NewSingleEmail(from, subject, to, emailBody, emailBody)
	client := sendgrid.NewSendClient(Sendgridkey)
	_, err = client.Send(message)
	if err != nil {
		return nil, err
	}
	return &EmailResponse{
		Status:   http.StatusOK,
		RespBody: "Success, email send",
	}, nil
}

func (s *sendMail) SendParticipationUserCreatedTrip(ToUserCreatedTrip string, ToTrip string, ToUserParticipated string, ToUserParticipatedEmail string, FromAdmin string, Sendgridkey string) (*EmailResponse, error) {
	h := hermes.Hermes{
		Product: hermes.Product{
			Name: "LFTrip",
		},
	}

	var intro, intro2 string
	intro = "Un Tripizien, " + ToUserParticipated + ", est intéressé par un de vos voyage : " + ToTrip
	intro2 = "Échangez avec lui par email : " + ToUserParticipatedEmail + ", vous pourrez ensuite communiquer via d'autres réseaux."

	email := hermes.Email{
		Body: hermes.Body{
			Name: ToUserCreatedTrip,
			Intros: []string{
				intro,
				intro2,
			},
			Signature: "On te souhaite de faire pleins de bonnes rencontres et de voyages !",
		},
	}
	emailBody, err := h.GenerateHTML(email)
	if err != nil {
		return nil, err
	}
	from := mail.NewEmail("LFTrip", FromAdmin)
	subject := intro
	to := mail.NewEmail("Nouvelle Participation", ToUserCreatedTrip)
	message := mail.NewSingleEmail(from, subject, to, emailBody, emailBody)
	client := sendgrid.NewSendClient(Sendgridkey)
	_, err = client.Send(message)
	if err != nil {
		return nil, err
	}
	return &EmailResponse{
		Status:   http.StatusOK,
		RespBody: "Success, email send",
	}, nil
}

func (s *sendMail) SendParticipationUser(ToUser string, ToTrip string, FromAdmin string, Sendgridkey string) (*EmailResponse, error) {
	h := hermes.Hermes{
		Product: hermes.Product{
			Name: "LFTrip",
		},
	}

	var intro string
	intro = "Vous avez bien participé au voyage : " + ToTrip + " !"

	email := hermes.Email{
		Body: hermes.Body{
			Name: ToUser,
			Intros: []string{
				intro,
				"Les premiers échanges se font par email, l'auteur du voyage vous contactera prochainement, vous pourrez ensuite communiquer via d'autres réseaux.",
			},
			Actions: []hermes.Action{
				{
					Instructions: "D'autres voyages n'attendent que vous, n'hésitez pas à regarder les autres voyages créés par la communauté :)",
				},
			},
			Signature: "On te souhaite de faire pleins de bonnes rencontres et de voyages !",
		},
	}
	emailBody, err := h.GenerateHTML(email)
	if err != nil {
		return nil, err
	}
	from := mail.NewEmail("LFTrip", FromAdmin)
	subject := "Participation à un voyage"
	to := mail.NewEmail("Participation", ToUser)
	message := mail.NewSingleEmail(from, subject, to, emailBody, emailBody)
	client := sendgrid.NewSendClient(Sendgridkey)
	_, err = client.Send(message)
	if err != nil {
		return nil, err
	}
	return &EmailResponse{
		Status:   http.StatusOK,
		RespBody: "Success, email send",
	}, nil
}

func (s *sendMail) SendCreateTrip(ToUser string, FromAdmin string, Sendgridkey string) (*EmailResponse, error) {
	h := hermes.Hermes{
		Product: hermes.Product{
			Name: "LFTrip",
		},
	}

	var tripURL string
	tripURL = "https://lftrip.tk/trips"

	email := hermes.Email{
		Body: hermes.Body{
			Name: ToUser,
			Intros: []string{
				"Votre voyage a bien été créé ! LFTrip vous souhaite de trouver des voyageurs qui vous correspond !",
			},
			Actions: []hermes.Action{
				{
					Instructions: "D'autres voyages n'attendent que vous, n'hésitez pas à regarder les voyages créés par la communauté :)",
					Button: hermes.Button{
						Color: "#DC4D2F",
						Text:  "Découvre nos voyages",
						Link:  tripURL,
					},
				},
			},
			Signature: "On te souhaite de faire pleins de bonnes rencontres et de voyages !",
		},
	}
	emailBody, err := h.GenerateHTML(email)
	if err != nil {
		return nil, err
	}
	from := mail.NewEmail("LFTrip", FromAdmin)
	subject := "Voyage créé"
	to := mail.NewEmail("Voyage Créé", ToUser)
	message := mail.NewSingleEmail(from, subject, to, emailBody, emailBody)
	client := sendgrid.NewSendClient(Sendgridkey)
	_, err = client.Send(message)
	if err != nil {
		return nil, err
	}
	return &EmailResponse{
		Status:   http.StatusOK,
		RespBody: "Success, email send",
	}, nil
}

func (s *sendMail) SendWelcome(ToUser string, FromAdmin string, Sendgridkey string) (*EmailResponse, error) {
	h := hermes.Hermes{
		Product: hermes.Product{
			Name: "LFTrip",
		},
	}

	var tripURL string
	tripURL = "https://lftrip.tk/trips"

	email := hermes.Email{
		Body: hermes.Body{
			Name: ToUser,
			Intros: []string{
				"Bienvneue sur LFTrip ! Nous sommes heureux de te compter parmi notre belle communauté de voyageurs !",
			},
			Actions: []hermes.Action{
				{
					Instructions: "Clique sur le bouton ci-dessous pour découvrir les voyages créés par la communauté :)",
					Button: hermes.Button{
						Color: "#DC4D2F",
						Text:  "Découvre nos voyages",
						Link:  tripURL,
					},
				},
			},
			Signature: "On te souhaite de faire pleins de bonnes rencontres et de voyages !",
		},
	}
	emailBody, err := h.GenerateHTML(email)
	if err != nil {
		return nil, err
	}
	from := mail.NewEmail("LFTrip", FromAdmin)
	subject := "Bienvenue sur LFTrip"
	to := mail.NewEmail("Notre nouveau tripizien", ToUser)
	message := mail.NewSingleEmail(from, subject, to, emailBody, emailBody)
	client := sendgrid.NewSendClient(Sendgridkey)
	_, err = client.Send(message)
	if err != nil {
		return nil, err
	}
	return &EmailResponse{
		Status:   http.StatusOK,
		RespBody: "Success, Please click on the link provided in your email",
	}, nil
}

func (s *sendMail) SendResetPassword(ToUser string, FromAdmin string, Token string, Sendgridkey string) (*EmailResponse, error) {
	h := hermes.Hermes{
		Product: hermes.Product{
			Name: "LFTrip",
		},
	}
	var forgotURL string
	//if os.Getenv("APP_ENV") == "production" {
	forgotURL = "https://lftrip.tk/resetpassword/" + Token //this is the url of the local frontend app
	//}
	email := hermes.Email{
		Body: hermes.Body{
			Name: ToUser,
			Intros: []string{
				"Vous avez oublié votre mot de passe ? Aucune inquiétude !",
			},
			Actions: []hermes.Action{
				{
					Instructions: "Clique sur le bouton ci-dessous pour réinitialiser ton mot de passe",
					Button: hermes.Button{
						Color: "#DC4D2F",
						Text:  "Modifier mon mot de passe",
						Link:  forgotURL,
					},
				},
			},
			Outros: []string{
				"If you did not request a password reset, no further action is required on your part.",
			},
			Signature: "Thanks",
		},
	}
	emailBody, err := h.GenerateHTML(email)
	if err != nil {
		return nil, err
	}
	from := mail.NewEmail("LFTrip", FromAdmin)
	subject := "Reset Password"
	to := mail.NewEmail("Reset Password", ToUser)
	message := mail.NewSingleEmail(from, subject, to, emailBody, emailBody)
	client := sendgrid.NewSendClient(Sendgridkey)
	_, err = client.Send(message)
	if err != nil {
		return nil, err
	}
	return &EmailResponse{
		Status:   http.StatusOK,
		RespBody: "Success, Please click on the link provided in your email",
	}, nil
}
