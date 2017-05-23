package concourse_test

import (
	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/onsi/gomega/ghttp"
)

var _ = Describe("InvalidateCache", func() {
	Context("when ATC request succeeds", func() {
		BeforeEach(func() {
			expectedURL := "/api/v1/teams/some-team/pipelines/mypipeline/resources/myresource"
			atcServer.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("DELETE", expectedURL),
					ghttp.RespondWithJSONEncoded(http.StatusOK, nil),
				),
			)
		})

		It("sends invalidate cache request to ATC", func() {
			found, err := team.InvalidateCache("mypipeline", "myresource")
			Expect(err).NotTo(HaveOccurred())
			Expect(found).To(BeTrue())

			Expect(atcServer.ReceivedRequests()).To(HaveLen(1))
		})
	})

	Context("when pipeline or resource does not exist", func() {
		BeforeEach(func() {
			expectedURL := "/api/v1/teams/some-team/pipelines/mypipeline/resources/myresource"
			atcServer.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("DELETE", expectedURL),
					ghttp.RespondWithJSONEncoded(http.StatusNotFound, nil),
				),
			)
		})

		It("returns a ResourceNotFoundError", func() {
			found, err := team.InvalidateCache("mypipeline", "myresource")
			Expect(err).NotTo(HaveOccurred())
			Expect(found).To(BeFalse())
		})
	})

	Context("when ATC responds with an error", func() {
		BeforeEach(func() {
			expectedURL := "/api/v1/teams/some-team/pipelines/mypipeline/resources/myresource"

			atcServer.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("DELETE", expectedURL),
					ghttp.RespondWithJSONEncoded(http.StatusInternalServerError, nil),
				),
			)
		})

		It("returns an error", func() {
			found, err := team.InvalidateCache("mypipeline", "myresource")
			Expect(err).To(HaveOccurred())
			Expect(found).To(BeFalse())
		})
	})
})
