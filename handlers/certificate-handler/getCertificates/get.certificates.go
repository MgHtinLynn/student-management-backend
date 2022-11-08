package handlerGetCertificates

import (
	"github.com/MgHtinLynn/final-year-project-mcc/controllers/certificate/getCertificates"
	util "github.com/MgHtinLynn/final-year-project-mcc/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type handler struct {
	service getCertificates.Service
}

func NewHandlerGetCertificates(service getCertificates.Service) *handler {
	return &handler{service: service}
}

func (h *handler) GetCertificatesHandler(ctx *gin.Context) {
	getCertificateLists, count, errGetCertificates := h.service.GetCertificatesService(ctx)

	switch errGetCertificates {
	case "RESULT_EXAM_RESULT_NOT_FOUND_404":
		util.APIPaginationResponse(ctx, "Certificate data is not exists", http.StatusConflict, count, http.MethodPost, nil)

	default:
		util.APIPaginationResponse(ctx, "Results Certificate data successfully", http.StatusOK, count, http.MethodPost, getCertificateLists)
	}
}
