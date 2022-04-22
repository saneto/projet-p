package models

import (
	"errors"
	"html"
	"log"
	"strings"
	"time"
	"fmt"
	"gitlab.com/esgi/lftrip/BackLFTrip/api/security"

	"github.com/badoux/checkmail"
	"github.com/jinzhu/gorm"
	uuid "github.com/twinj/uuid"
	"golang.org/x/crypto/bcrypt"
)

// User Struct
type User struct {
	UUID             uuid.UUID `gorm:"type:uuid;unique_index;" json:"uuid"`
	ID               uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Firstname        string    `valid:"required,alpha,length(2|255)" json:"firstname"`
	Lastname         string    `valid:"required,alpha,length(2|255)" json:"lastname"`
	Email            string    `gorm:"size:100;not null;unique" valid:"email" json:"email"`
	Password         string    `gorm:"size:100;not null;" json:"password"`
	Accesslevel      int       `valid:"range(0|1),numeric" json:"access_level"`
	Dateofbirth      string    `gorm:"size:250;null;" json:"dateofbirth"`
	Sexe             string    `gorm:"size:100;not null;" json:"sexe"`
	City             string    `gorm:"size:150;null;" json:"city"`
	IsGuide          string    `gorm:"size:100;null;" json:"is_guide"`
	CountryGuide     string    `gorm:"size:150;null;" json:"country_guide"`
	PhoneNumber      string    `gorm:"size:15;null" json:"phone_number"`
	DepartureAirport string    `gorm:"size:250;null;" json:"departure_airport"`
	Description      string    `gorm:"text;null;" json:"description"`
	AvatarPath       string    `gorm:"type:text;null" json:"avatar_path"`
	CreatedAt        time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt        time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	Chats            []*Chat   `gorm:"many2many:chat_users;association_foreignkey:id;foreignkey:id" json:"Chats,omitempty"`
}

// TableName : Gorm related
func (u *User) TableName() string {
	return "users"
}

// BeforeSave : call package security to hash the password
func (u *User) BeforeSave() error {
	hashedPassword, err := security.Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	u.UUID = uuid.NewV4()
	u.CreatedAt = time.Now()
	return nil
}

// BeforeUpdate is gorm hook that is triggered on every updated on user struct
func (u *User) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("UpdatedAt", time.Now())
	return nil
}

// Prepare : prepare  statements
func (u *User) Prepare() {
	u.UUID = uuid.NewV4()
	u.ID = 0
	u.Firstname = html.EscapeString(strings.TrimSpace(u.Firstname))
	u.Lastname = html.EscapeString(strings.TrimSpace(u.Lastname))
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	u.Sexe = html.EscapeString(strings.TrimSpace(u.Sexe))
	u.City = html.EscapeString(strings.TrimSpace(u.City))
	u.IsGuide = html.EscapeString(strings.TrimSpace(u.IsGuide))
	u.CountryGuide = html.EscapeString(strings.TrimSpace(u.CountryGuide))
	u.PhoneNumber = html.EscapeString(strings.TrimSpace(u.PhoneNumber))
	u.DepartureAirport = html.EscapeString(strings.TrimSpace(u.DepartureAirport))
	u.Description = html.UnescapeString(strings.TrimSpace(u.Description))
	u.Dateofbirth = html.EscapeString(strings.TrimSpace(u.Dateofbirth))
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}

// VerifyPassword : This method compare the password with the hash
func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

// Validate : function to check the data
func (u *User) Validate(action string) map[string]string {
	var errorMessages = make(map[string]string)
	var err error

	switch strings.ToLower(action) {
	case "update":
		if u.Email == "" {
			err = errors.New("Email requis")
			errorMessages["Required_email"] = err.Error()
		}
		if u.Email != "" {
			if err = checkmail.ValidateFormat(u.Email); err != nil {
				err = errors.New("Email invalide")
				errorMessages["Invalid_email"] = err.Error()
			}
		}

	case "login":
		if u.Password == "" {
			err = errors.New("Mot de passe requis")
			errorMessages["Required_password"] = err.Error()
		}
		if u.Email == "" {
			err = errors.New("Email requis")
			errorMessages["Required_email"] = err.Error()
		}
		if u.Email != "" {
			if err = checkmail.ValidateFormat(u.Email); err != nil {
				err = errors.New("Email invalide")
				errorMessages["Invalid_email"] = err.Error()
			}
		}
	case "forgotpassword":
		if u.Email == "" {
			err = errors.New("Email requis")
			errorMessages["Required_email"] = err.Error()
		}
		if u.Email != "" {
			if err = checkmail.ValidateFormat(u.Email); err != nil {
				err = errors.New("Email invalide")
				errorMessages["Invalid_email"] = err.Error()
			}
		}
	default:
		if u.Firstname == "" {
			err = errors.New("Prénom requis")
			errorMessages["Required_firstname"] = err.Error()
		}
		if u.Lastname == "" {
			err = errors.New("Nom requis")
			errorMessages["Required_lastname"] = err.Error()
		}
		if u.Password == "" {
			err = errors.New("Mot de passe requis")
			errorMessages["Required_password"] = err.Error()
		}
		if u.Password != "" && len(u.Password) < 6 {
			err = errors.New("Le mot de passe doit comporter au moins 6 caractères")
			errorMessages["Invalid_password"] = err.Error()
		}
		if u.Email == "" {
			err = errors.New("Email requis")
			errorMessages["Required_email"] = err.Error()

		}
		if u.Email != "" {
			if err = checkmail.ValidateFormat(u.Email); err != nil {
				err = errors.New("Email invalide")
				errorMessages["Invalid_email"] = err.Error()
			}
		}
	}
	return errorMessages

}

// AfterFind : values of return
func (u *User) AfterFind() (err error) {
	if err != nil {
		return err
	}
	/*if u.AvatarPath != "" {
		u.AvatarPath = os.Getenv("DO_SPACES_URL") + u.AvatarPath
	}*/
	//dont return the user password
	//u.Password = ""
	return nil
}

// SaveUser : Method Save User, triggered on every saved on user struct
func (u *User) SaveUser(db *gorm.DB) (*User, error) {
	var err error
	err = db.Debug().Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

// FindAllUsers : function to find all users
func (u *User) FindAllUsers(db *gorm.DB) (*[]User, error) {
	var err error
	users := []User{}
	err = db.Debug().Model(&User{}).Limit(100).Find(&users).Error
	if err != nil {
		return &[]User{}, err
	}
	return &users, err
}

// FindUserByID : function to find a user with an ID
func (u *User) FindUserByID(db *gorm.DB, uid uint64) (*User, error) {
	var err error
	err = db.Debug().Model(User{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &User{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &User{}, errors.New("User Not Found")
	}
	return u, err
}

// FindUserGuideByCountry : function to find a user guide with a country
func (u *User) FindUserGuideByCountry(db *gorm.DB, country string) (*[]User, error) {
	var err error
	users := []User{}
	err = db.Debug().Model(User{}).Where("is_guide = ? and country_guide = ?", "true", country).Find(&users).Error
	if err != nil {
		return &[]User{}, err
	}
	return &users, err
}

// UpdateAUser : update an user
func (u *User) UpdateAUser(db *gorm.DB, uid uint64) (*User, error) {
	// To hash the password
	err := u.BeforeSave()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(u.Dateofbirth)
	fmt.Println("dqsdqsd")
	fmt.Println(u.Sexe)
	db = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&User{}).UpdateColumns(
		map[string]interface{}{
			"firstname":         u.Firstname,
			"lastname":          u.Lastname,
			"email":             u.Email,
			"dateofbirth":       u.Dateofbirth,
			"sexe":              u.Sexe,
			"city":              u.City,
			"phone_number":      u.PhoneNumber,
			"is_guide":          u.IsGuide,
			"country_guide":     u.CountryGuide,
			"departure_airport": u.DepartureAirport,
			"description":       u.Description,
			"avatar_path":       u.AvatarPath,
			"update_at":         time.Now(),
		},
	)
	if db.Error != nil {
		return &User{}, db.Error
	}
	return u, nil
}

// UpdateAUserAvatar : update an avatar user
func (u *User) UpdateAUserAvatar(db *gorm.DB, uid uint64) (*User, error) {
	db = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&User{}).UpdateColumns(
		map[string]interface{}{
			"avatar_path": u.AvatarPath,
			"update_at":   time.Now(),
		},
	)
	if db.Error != nil {
		return &User{}, db.Error
	}
	return u, nil
}

// DeleteAUser : function to delete the user
func (u *User) DeleteAUser(db *gorm.DB, uid uint64) (int64, error) {
	db = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&User{}).Delete(&User{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

// UpdatePassword : funtion to update the password
func (u *User) UpdatePassword(db *gorm.DB) error {

	// To hash the password
	err := u.BeforeSave()
	if err != nil {
		log.Fatal(err)
	}

	db = db.Debug().Model(&User{}).Where("email = ?", u.Email).Take(&User{}).UpdateColumns(
		map[string]interface{}{
			"password":  u.Password,
			"update_at": time.Now(),
		},
	)
	if db.Error != nil {
		return db.Error
	}
	return nil
}

func (u *User) checkPassword(password string) (bool, error) {
	hashedPassword, err := security.Hash(password)
	if err != nil {
		return false, err
	}

	if u.Password != string(hashedPassword) {
		return false, nil
	}
	return true, nil
}
