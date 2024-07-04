<script setup lang="ts">
import { onMounted, ref } from "vue";
import { WEBSERVER_ROOT_URL } from "../../config/webserver";

const results = ref<Record<string, number[]>>({});

onMounted(() => {
  const ws = new WebSocket(`ws://${WEBSERVER_ROOT_URL}/results`);
  ws.onmessage = (event) => {
    results.value = JSON.parse(event.data) as Record<string, number[]>;
  };
});
</script>
<template>
  <div class="results-container">
    <h2>Results</h2>
    <ul>
      <li class="list-items" v-for="job in Object.keys(results)">
        {{ job }}: {{ results[job].length }}
      </li>
    </ul>
  </div>
</template>
<style>
.results-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  gap: 5%;
  height: 100%;
}

.list-items {
  font-size: small;
  font-style: normal;
}
</style>
