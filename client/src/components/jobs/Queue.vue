<script setup lang="ts">
import { onMounted, ref } from "vue";
import { WEBSERVER_ROOT_URL } from "../../config/webserver";

const queue = ref<string[]>([]);

onMounted(() => {
  const ws = new WebSocket(`ws://${WEBSERVER_ROOT_URL}/queue`);
  ws.onmessage = (event) => {
    queue.value = JSON.parse(event.data) as string[];
  };
});
</script>
<template>
  <div class="queue-container">
    <h2>Queue ({{ queue.length }})</h2>
    <ul>
      <li class="list-items" v-for="job in queue">
        {{ job }}
      </li>
    </ul>
  </div>
</template>
<style>
.queue-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 5%;
  height: 300px;
  width: 400px;
  border: solid 2px lightseagreen;
  overflow-y: scroll;
}

.list-items {
  font-size: small;
  font-style: normal;
}
</style>
