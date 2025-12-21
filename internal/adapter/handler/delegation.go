package handler

import (
	"delegator/internal/application/usecase"
	"delegator/internal/domain/model"
	"delegator/internal/domain/repository"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type DelegationHandler struct {
	usecase *usecase.DelegationUsecase
}

func NewDelegationHandler(uc *usecase.DelegationUsecase) *DelegationHandler {
	return &DelegationHandler{
		usecase: uc,
	}
}

type CreateDelegationBody struct {
	AcrValues       string `json:"acr_values"`
	BindingMessage  string `json:"binding_message"`
	ConsentRequired bool   `json:"consent_required"`
	LoginHint       string `json:"login_hint"`
	Scope           string `json:"scope"`
}

type DelegationJSON struct {
	ID              string    `json:"id"`
	Status          string    `json:"status"`
	AcrValues       string    `json:"acr_values"`
	BindingMessage  string    `json:"binding_message"`
	ConsentRequired bool      `json:"consent_required"`
	LoginHint       string    `json:"login_hint"`
	Scope           string    `json:"scope"`
	AuthToken       string    `json:"auth_token"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

func toDelegationJSON(m *model.Delegation) *DelegationJSON {
	return &DelegationJSON{
		ID:              m.ID,
		Status:          string(m.Status),
		AcrValues:       m.AcrValues,
		BindingMessage:  m.BindingMessage,
		ConsentRequired: m.ConsentRequired,
		LoginHint:       m.LoginHint,
		Scope:           m.Scope,
		AuthToken:       m.AuthToken,
		CreatedAt:       m.CreatedAt,
		UpdatedAt:       m.UpdatedAt,
	}
}

func (h *DelegationHandler) List(c echo.Context) error {
	// limitクエリパラメータの準備
	qLimit := c.QueryParam("limit")
	if len(qLimit) == 0 {
		qLimit = "20"
	}
	limit, err := strconv.ParseInt(qLimit, 10, 32)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	// offsetクエリパラメータの準備
	qOffset := c.QueryParam("offset")
	if len(qOffset) == 0 {
		qOffset = "0"
	}
	offset, err := strconv.ParseInt(qOffset, 10, 32)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	// 一括取得
	delegations, err := h.usecase.List(c.Request().Context(), int(limit), int(offset))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	// JSON用の構造に変換して返却
	response := []*DelegationJSON{}
	for _, delegation := range delegations {
		response = append(response, toDelegationJSON(delegation))
	}
	return c.JSON(http.StatusOK, response)
}

func (h *DelegationHandler) Create(c echo.Context) error {
	// リソース作成のためのパラメータを取得
	body := &CreateDelegationBody{}
	err := c.Bind(body)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	// Authorizationヘッダからトークンを抽出
	token := c.Request().Header.Get("Authorization")
	if len(token) <= len("Bearer ") {
		return echo.NewHTTPError(http.StatusBadRequest, errors.New("Authorization header must be set with bearer token"))
	}
	token = token[len("Bearer "):]

	// リソースを作成
	created, err := h.usecase.Create(c.Request().Context(), &usecase.CreateDelegationInput{
		AcrValues:       body.AcrValues,
		BindingMessage:  body.BindingMessage,
		ConsentRequired: body.ConsentRequired,
		LoginHint:       body.LoginHint,
		Scope:           body.Scope,
		AuthToken:       token,
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	// 作成されたリソースを返却
	return c.JSON(http.StatusCreated, toDelegationJSON(created))
}

func (h *DelegationHandler) GetById(c echo.Context) error {
	// 指定されたIDのリソースを取得
	delegation, err := h.usecase.GetById(c.Request().Context(), c.Param("delegationId"))
	if err == repository.ErrRecordNotFound {
		// 見つからなければ 404 Not Found を返却
		return echo.NewHTTPError(http.StatusNotFound, err)
	} else if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	// 見つかったリソースを返却
	return c.JSON(http.StatusOK, toDelegationJSON(delegation))
}

func (h *DelegationHandler) ApproveById(c echo.Context) error {
	// 指定したIDのリソースをApproveする
	delegation, err := h.usecase.ApproveById(c.Request().Context(), c.Param("delegationId"))
	if err == repository.ErrRecordNotFound {
		return echo.NewHTTPError(http.StatusNotFound, err)
	} else if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, toDelegationJSON(delegation))
}

func (h *DelegationHandler) CancelById(c echo.Context) error {
	// 指定したIDのリソースをCancelする
	delegation, err := h.usecase.CancelById(c.Request().Context(), c.Param("delegationId"))
	if err == repository.ErrRecordNotFound {
		return echo.NewHTTPError(http.StatusNotFound, err)
	} else if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, toDelegationJSON(delegation))
}

func (h *DelegationHandler) UnauthorizeById(c echo.Context) error {
	// 指定したIDのリソースをApproveする
	delegation, err := h.usecase.UnauthorizeById(c.Request().Context(), c.Param("delegationId"))
	if err == repository.ErrRecordNotFound {
		return echo.NewHTTPError(http.StatusNotFound, err)
	} else if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, toDelegationJSON(delegation))
}

func (h *DelegationHandler) DeleteById(c echo.Context) error {
	// 指定されたIDのリソースを削除
	err := h.usecase.DeleteById(c.Request().Context(), c.Param("delegationId"))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return c.NoContent(http.StatusNoContent)
}
