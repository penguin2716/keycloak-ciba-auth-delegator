const API_BASE = import.meta.env.VITE_API_BASE_URL || "";

export interface Delegation {
  id: string;
  status: string;
  acr_values: string;
  binding_message: string;
  consent_required: boolean;
  login_hint: string;
  scope: string;
  auth_token: string;
  created_at: string;
  updated_at: string;
}

export const delegationApi = {
  async list(): Promise<Array<Delegation>> {
    const resp = await fetch(`${API_BASE}/api/delegations`);
    if (!resp.ok) {
      throw new Error("Failed to get delegations");
    }
    return resp.json();
  },

  async getById(id: string): Promise<Delegation> {
    const resp = await fetch(`${API_BASE}/api/delegations/${id}`);
    if (!resp.ok) {
      throw new Error("Failed to get delegation");
    }
    return resp.json();
  },

  async approveById(id: string): Promise<void> {
    const resp = await fetch(`${API_BASE}/api/delegations/${id}/approve`, {
      method: "PUT",
    });
    if (!resp.ok) {
      throw new Error("Failed to approve delegation");
    }
    return;
  },

  async cancelById(id: string): Promise<void> {
    const resp = await fetch(`${API_BASE}/api/delegations/${id}/cancel`, {
      method: "PUT",
    });
    if (!resp.ok) {
      throw new Error("Failed to cancel delegation");
    }
    return;
  },

  async unauthorizeById(id: string): Promise<void> {
    const resp = await fetch(`${API_BASE}/api/delegations/${id}/unauthorize`, {
      method: "PUT",
    });
    if (!resp.ok) {
      throw new Error("Failed to unauthorize delegation");
    }
    return;
  },

  async deleteById(id: string): Promise<void> {
    const resp = await fetch(`${API_BASE}/api/delegations/${id}`, {
      method: "DELETE",
    });
    if (!resp.ok) {
      throw new Error("Failed to delete delegation");
    }
    return;
  },
};
