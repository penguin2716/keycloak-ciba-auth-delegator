const API_BASE = import.meta.env.VITE_API_BASE_URL || "";

export interface Settings {
  keycloak: {
    base_url: string;
    realm: string;
  };
}

export const settingsApi = {
  async get(): Promise<Settings> {
    const resp = await fetch(`${API_BASE}/api/settings`);
    if (!resp.ok) {
      throw new Error("Failed to fetch settings");
    }
    return resp.json();
  },

  async update(settings: Settings): Promise<Settings> {
    const resp = await fetch(`${API_BASE}/api/settings`, {
      method: "PUT",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(settings),
    });
    if (!resp.ok) {
      throw new Error("Failed to update settings");
    }
    return resp.json();
  },
};
