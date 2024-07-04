<script setup lang="ts">
import { onMounted, ref } from "vue";
import { WEBSERVER_ROOT_URL } from "../../config/webserver";
import { WorkerStatus } from "./Workers.types";
import Server from "./Server.vue";

const workerStatus = ref<WorkerStatus[]>();

onMounted(() => {
  const ws = new WebSocket(`ws://${WEBSERVER_ROOT_URL}/workerStatus`);
  ws.onmessage = (event) => {
    workerStatus.value = JSON.parse(event.data) as WorkerStatus[]
  }
});
</script>
<template>
  <div class="server" v-for="worker in workerStatus">
    <Server :busy="worker.Busy"></Server>
    <h4>Worker - {{ worker.Id }}</h4>
  </div>
</template>
<style>

.server {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: 3%;
}


</style>
