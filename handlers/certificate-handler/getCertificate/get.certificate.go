package handlerGetCertificate

import (
	"fmt"
	"github.com/MgHtinLynn/final-year-project-mcc/controllers/certificate/getCertificate"
	util "github.com/MgHtinLynn/final-year-project-mcc/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type handler struct {
	service getCertificate.Service
}

func NewHandlerGetCertificate(service getCertificate.Service) *handler {
	return &handler{service: service}
}

func (h *handler) GetCertificateHandler(ctx *gin.Context) {
	var input getCertificate.InputGetCertificate

	input.ID, _ = strconv.Atoi(ctx.Param("id"))

	fmt.Println("ID ", input.ID)

	getCertificateById, errResultStudent := h.service.GetCertificateService(&input)

	switch errResultStudent {

	case "RESULT_Certificate_NOT_FOUND_404":
		util.APIResponse(ctx, "Certificate data is not exist or deleted", http.StatusNotFound, http.MethodGet, nil)
		return

	default:
		util.APIResponse(ctx, "Certificate data successfully", http.StatusOK, http.MethodGet, getCertificateById)
	}
}
