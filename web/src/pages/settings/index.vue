<template>
  <v-container fluid class="mt-2 px-8">
    <v-row>
      <v-col>
        <div class="text-h5">Settings</div>
      </v-col>
    </v-row>

    <v-row>
      <v-col>
        <div>Keycloak</div>
      </v-col>
    </v-row>
    <v-row>
      <v-col>
        <v-text-field
          v-model="settings.keycloak.base_url"
          label="Keycloak Base URL"
          placeholder="https://keycloak.example.com"
          hide-details
        ></v-text-field>
      </v-col>
    </v-row>
    <v-row>
      <v-col>
        <v-text-field
          v-model="settings.keycloak.realm"
          label="Keycloak Realm Name"
          hide-details
        ></v-text-field>
      </v-col>
    </v-row>

    <v-row>
      <v-col>
        <v-btn color="primary" @click="handleSave" :disabled="loading">
          save
        </v-btn>
      </v-col>
    </v-row>
  </v-container>
</template>

<script lang="ts" setup>
import { type Ref, ref, onMounted } from "vue";
import { useSettings } from "@/composables/useSettings";

const { settings, loading, loadSettings, saveSettings } = useSettings();

const handleSave = async () => {
  await saveSettings();
};

onMounted(() => {
  loadSettings();
});
</script>
