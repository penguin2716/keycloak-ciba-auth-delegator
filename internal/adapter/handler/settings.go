package handler

import (
	"delegator/internal/application/usecase"
	"delegator/internal/domain/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

type SettingsHandler struct {
	usecase *usecase.SettingsUsecase
}

func NewSettingsHandler(uc *usecase.SettingsUsecase) *SettingsHandler {
	return &SettingsHandler{
		usecase: uc,
	}
}

type SettingsJSON struct {
	Keycloak struct {
		BaseURL string `json:"base_url"`
		Realm   string `json:"realm"`
	} `json:"keycloak"`
}

func (o *SettingsJSON) toDomain() *model.Settings {
	return &model.Settings{
		Keycloak: model.KeycloakSettings{
			BaseURL: o.Keycloak.BaseURL,
			Realm:   o.Keycloak.Realm,
		},
	}
}

func toSettingsJSON(m *model.Settings) *SettingsJSON {
	payload := &SettingsJSON{}
	payload.Keycloak.BaseURL = m.Keycloak.BaseURL
	payload.Keycloak.Realm = m.Keycloak.Realm
	return payload
}

func (h *SettingsHandler) Load(c echo.Context) error {
	// 現在の設定を取得
	m, err := h.usecase.Load(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	// 取得結果を返却
	return c.JSON(http.StatusOK, toSettingsJSON(m))
}

func (h *SettingsHandler) Save(c echo.Context) error {
	// リクエストボディを取得
	body := &SettingsJSON{}
	err := c.Bind(body)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	// 更新内容を保存
	updated, err := h.usecase.Save(c.Request().Context(), body.toDomain())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	// 更新結果を返却
	return c.JSON(http.StatusOK, toSettingsJSON(updated))
}
