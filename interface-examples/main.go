package main

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
)

type AccountNotifier interface {
	NotifyAccountCreated(context.Context, Account) error
}

type EmailAccountNotifier struct {
}

func (e EmailAccountNotifier) NotifyAccountCreated(ctx context.Context, account Account) error {
	slog.Info("email notification :- new account created", "username", account.Username, "email", account.Email)
	return nil
}

type DiscordAccountNotifier struct {
}

func (d DiscordAccountNotifier) NotifyAccountCreated(ctx context.Context, account Account) error {
	slog.Info("discord notification :- new account created", "username", account.Username, "email", account.Email)
	return nil
}

type Account struct {
	Username string
	Email    string
}

type AccountHandler struct {
	AccountNotifier AccountNotifier
}

func (h *AccountHandler) handleCreateAccount(w http.ResponseWriter, r *http.Request) {
	var account Account
	if err := json.NewDecoder(r.Body).Decode(&account); err != nil {
		slog.Error(" failed to decode request body: %v", err)
		return
	}
	if err := h.AccountNotifier.NotifyAccountCreated(r.Context(), account); err != nil {
		slog.Error("failed to notify account created: %v", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(account)
	if err != nil {
		return
	}
}

func main() {
	mux := http.NewServeMux()
	AccountHandler := &AccountHandler{
		AccountNotifier: DiscordAccountNotifier{}, // change to EmailAccountNotifier{}
	}
	mux.HandleFunc("POST /account", AccountHandler.handleCreateAccount)

	err := http.ListenAndServe(":3000", mux)
	if err != nil {
		return
	}
}
