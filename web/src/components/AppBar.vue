<template>
<v-app-bar>
  <v-app-bar-nav-icon variant="text" @click.stop="app.drawer = !app.drawer"></v-app-bar-nav-icon>
  <v-app-bar-title>Authentication Channel Provider for Keycloak</v-app-bar-title>
</v-app-bar>
</template>

<script lang="ts" setup>
import { useAppStore } from "@/stores/app";
const app = useAppStore();
</script>
