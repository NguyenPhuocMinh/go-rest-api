package tests_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	commons "fast-food-api-client/commons"
	db "fast-food-api-client/core/database"
	controllersV1 "fast-food-api-client/src/controllers/v1"

	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.mongodb.org/mongo-driver/mongo"
)

const apiEndpoint = "/api/v1/auth"

var _ = Describe("[AuthController]", func() {
	Describe("[Test AuthRegister]", func() {
		Context("Register Success Without Error", func() {
			BeforeEach(func() {
				// mock database
				db.Count = func(schemaType string, filter interface{}) int {
					return 0
				}
				db.CreateOne = func(schemaType string, data interface{}) (*mongo.InsertOneResult, error) {
					InsertedID := &mongo.InsertOneResult{
						InsertedID: "123",
					}
					return InsertedID, nil
				}
			})

			It("Should be return status 200", func() {
				customerReqBody := map[string]interface{}{
					"firstName": "test",
					"lastName":  "test",
					"email":     "test@gmail.com",
					"password":  "123",
				}
				body, _ := json.Marshal(customerReqBody)
				req, err := http.NewRequest("POST", apiEndpoint+"/register", bytes.NewBuffer(body))
				w := httptest.NewRecorder()
				Expect(err).To(BeNil())

				c, _ := gin.CreateTestContext(w)
				c.Request = req

				controllersV1.RegisterAuth(c)

				var res commons.ResponseModel
				json.Unmarshal(w.Body.Bytes(), &res)

				// Assert data
				Expect(res.Error).To(BeNil())
				Expect(res.StatusCode).To(Equal(int(200)))
				Expect(res.Data).NotTo(BeNil())
			})
		})

		Context("Register Failed With Duplicated Data", func() {
			BeforeEach(func() {
				// mock database
				db.Count = func(schemaType string, filter interface{}) int {
					return 1
				}
			})

			It("Should be return status 409", func() {
				customerReqBody := map[string]interface{}{
					"firstName": "test",
					"lastName":  "test",
					"email":     "test@gmail.com",
					"password":  "123",
				}
				body, _ := json.Marshal(customerReqBody)
				req, err := http.NewRequest("POST", apiEndpoint+"/register", bytes.NewBuffer(body))
				w := httptest.NewRecorder()
				Expect(err).To(BeNil())

				c, _ := gin.CreateTestContext(w)
				c.Request = req

				controllersV1.RegisterAuth(c)

				var res commons.ResponseModel
				json.Unmarshal(w.Body.Bytes(), &res)

				// Assert data
				Expect(res.Error).NotTo(BeNil())
				Expect(res.StatusCode).To(Equal(int(409)))
				Expect(res.Data).To(BeNil())
			})
		})
	})
})
