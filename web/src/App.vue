<template>
  <v-app>
    <router-view />
  </v-app>
</template>

<script lang="ts" setup>
import { onMounted } from "vue";
import { useDisplay } from "vuetify";
import { useAppStore } from "@/stores/app";

const app = useAppStore();
const { mobile } = useDisplay();

onMounted(async () => {
  app.drawer = !mobile.value;
})
</script>
