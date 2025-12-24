import { ref } from "vue";
import { settingsApi, type Settings } from "@/api/settings";

export function useSettings() {
  const settings = ref<Settings>({
    keycloak: {
      base_url: "",
      realm: "",
    },
  });
  const loading = ref(false);
  const error = ref<string | null>(null);

  const loadSettings = async () => {
    loading.value = true;
    error.value = null;
    try {
      settings.value = await settingsApi.get();
    } catch (err) {
      error.value = "Failed to fetch settings";
    } finally {
      loading.value = false;
    }
  };

  const saveSettings = async () => {
    loading.value = true;
    error.value = null;
    try {
      const updated = await settingsApi.update(settings.value);
      settings.value = updated;
      return true;
    } catch (err) {
      error.value = "Failed to save settings";
      return false;
    } finally {
      loading.value = false;
    }
  };

  return {
    settings,
    loading,
    error,
    loadSettings,
    saveSettings,
  };
}
