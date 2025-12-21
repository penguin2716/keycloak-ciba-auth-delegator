package external

import (
	"bytes"
	"context"
	"delegator/internal/domain/model"
	"delegator/internal/domain/repository"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type KeycloakNotifier struct {
	client *http.Client
	repo   repository.SettingsRepository
}

func NewKeycloakNotifier(client *http.Client, repo repository.SettingsRepository) *KeycloakNotifier {
	return &KeycloakNotifier{
		client: client,
		repo:   repo,
	}
}

type AuthorizationResult struct {
	Status string `json:"status"`
}

func (n *KeycloakNotifier) CallbackURL(ctx context.Context) (string, error) {
	// DBから現在の設定を取得
	settings, err := n.repo.Load(ctx)
	if err != nil {
		return "", err
	}
	// 通知先URLを生成
	return fmt.Sprintf("%s/realms/%s/protocol/openid-connect/ext/ciba/auth/callback",
		settings.Keycloak.BaseURL, settings.Keycloak.Realm), nil
}

func (n *KeycloakNotifier) sendAuthorizationResult(ctx context.Context, m *model.Delegation) error {
	// 送信するペイロードを作成
	payload, err := json.Marshal(&AuthorizationResult{
		Status: string(m.Status),
	})
	if err != nil {
		return err
	}

	// HTTPリクエストの作成
	callbackURL, err := n.CallbackURL(ctx)
	if err != nil {
		return err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, callbackURL, bytes.NewReader(payload))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", m.AuthToken))

	// HTTPリクエストを送信
	resp, err := n.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Keycloak returned unexpected status: %d", resp.StatusCode)
	}

	// 結果を表示
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(string(respBody))

	return nil
}

func (n *KeycloakNotifier) DelegationApproved(ctx context.Context, m *model.Delegation) error {
	return n.sendAuthorizationResult(ctx, m)
}

func (n *KeycloakNotifier) DelegationCancelled(ctx context.Context, m *model.Delegation) error {
	return n.sendAuthorizationResult(ctx, m)
}

func (n *KeycloakNotifier) DelegationUnauthorized(ctx context.Context, m *model.Delegation) error {
	return n.sendAuthorizationResult(ctx, m)
}
