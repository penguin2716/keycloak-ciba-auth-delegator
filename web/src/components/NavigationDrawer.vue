<template>
  <v-navigation-drawer v-model="app.drawer">
    <v-list nav>
      <v-list-item title="Delegations" prepend-icon="mdi-account-question" to="/delegations"></v-list-item>
      <v-list-item title="Settings" prepend-icon="mdi-cog" to="/settings"></v-list-item>
    </v-list>
  </v-navigation-drawer>
</template>

<script lang="ts" setup>
import { useAppStore } from "@/stores/app";
const app = useAppStore();
</script>
