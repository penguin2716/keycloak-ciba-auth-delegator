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
        <v-btn color="primary" @click="save()">save</v-btn>
      </v-col>
    </v-row>
  </v-container>
</template>

<script lang="ts" setup>
import { type Ref, ref, onMounted } from "vue";

interface Settings {
  keycloak: {
    base_url: string;
    realm: string;
  };
}

const settings: Ref<Settings> = ref<Settings>({
  keycloak: {
    base_url: "",
    realm: "",
  },
});

const save = async () => {
  fetch((import.meta.env.VITE_API_BASE_URL || "") + "/api/settings", {
    method: "PUT",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(settings.value),
  })
    .then((resp) => resp.json())
    .then((resp) => {
      settings.value = resp;
    });
};

onMounted(() => {
  fetch((import.meta.env.VITE_API_BASE_URL || "") + "/api/settings")
    .then((resp) => resp.json())
    .then((resp) => {
      settings.value = resp;
    });
});
</script>
