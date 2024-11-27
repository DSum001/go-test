package unit

import (
	//"fmt"
	"example.com/sa-67-example/entity"
	"testing"
	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestPhoneNumber(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run(`phone_number check 10 digit`, func(t *testing.T) {
		member := entity.Member{
			PhoneNumber: "123456789",
			Password: "asdfghjk",
			Url:  "https://github.com/sut67/unit-test-se67/blob/main/test/account_test.go",
		}

		ok, err := govalidator.ValidateStruct(member)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("Phone Number length is not 10 digits."))

	})
}

func TestUrl(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run(`Url is required`, func(t *testing.T) {
		member := entity.Member{
			PhoneNumber: "1234567890",
			Password: "asdfghjk",
			Url:  "test",
			//Url:  "https://github.com/sut67/unit-test-se67/blob/main/test/account_test.go",
		}

		ok, err := govalidator.ValidateStruct(member)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("Url is invalid."))

	})
}

func TestAll(t *testing.T)  {
	g := NewGomegaWithT(t)

	t.Run(`all is valid`, func(t *testing.T) {
		user := entity.Member{
			PhoneNumber: "1234567890",
			Password: "asdfghjk",
			Url:  "https://github.com/sut67/unit-test-se67/blob/main/test/account_test.go",
		}

		ok, err := govalidator.ValidateStruct(user)

		g.Expect(ok).To(BeTrue())
		g.Expect(err).To(BeNil())

	})
}