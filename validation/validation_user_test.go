package validation_test

import (
	"testing"
	. "github.com/onsi/gomega"
	"github.com/peraponglim1125/Pre5/entity"
)

func TestValidateUser (t *testing.T){
	g := NewGomegaWithT(t)

	user := entity.User{
		Username: "Phirapong",
        Password: "password123",
        Email: "Phirapong@example.com",   
        Age: 20,  
	}
	ok, err := entity.ValidationUser(&user)
	g.Expect(ok).To(BeTrue())
	g.Expect(err).To(BeNil())
}

func TestInValidateName (t *testing.T){
	g := NewGomegaWithT(t)

	user := entity.User{
		Username: "",
        Password: "password123",
        Email: "Phirapong@example.com",   
        Age: 20,  
	}
	ok, err := entity.ValidationUser(&user)
	g.Expect(ok).To(BeFalse())
	g.Expect(err.Error()).To(Equal("Username is required"))
}

func TestInValidatePassword (t *testing.T){
	g := NewGomegaWithT(t)

	user := entity.User{
		Username: "phirapong",
        Password: "123",
        Email: "Phirapong@example.com",   
        Age: 20,  
	}
	ok, err := entity.ValidationUser(&user)
	g.Expect(ok).To(BeFalse())
	g.Expect(err.Error()).To(Equal("Password must be at least 8 characters"))
}
func TestInValidateEmail (t *testing.T){
	g := NewGomegaWithT(t)

	user := entity.User{
		Username: "phirapong",
        Password: "password123",
        Email: "Phirapongcom",   
        Age: 20,  
	}
	ok, err := entity.ValidationUser(&user)
	g.Expect(ok).To(BeFalse())
	g.Expect(err.Error()).To(Equal("Invalid email format"))
}
func TestInValidateAge (t *testing.T){
	g := NewGomegaWithT(t)

	user := entity.User{
		Username: "phirapong",
        Password: "password123",
        Email: "Phirapong@example.com",   
        Age: 1,  
	}
	ok, err := entity.ValidationUser(&user)
	g.Expect(ok).To(BeFalse())
	g.Expect(err.Error()).To(Equal("Age must be at least 13"))
}


