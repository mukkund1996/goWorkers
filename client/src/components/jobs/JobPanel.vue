<script setup lang="ts">
import { computed, ref } from "vue";
import { WEBSERVER_ROOT_URL } from "../../config/webserver";
const numWorkers = ref<string>("");

const updatedWorkers = computed(() =>
  numWorkers.value === "" ? "1" : numWorkers.value
);
const onShortJobSubmit = () => {
  const url = `http://${WEBSERVER_ROOT_URL}/short/${updatedWorkers.value}`;
  fetch(url, {
    method: "POST",
  });
  numWorkers.value = "1"
};

const onLongJobSubmit = () => {
  const url = `http://${WEBSERVER_ROOT_URL}/long/${updatedWorkers.value}`;
  fetch(url, {
    method: "POST",
  });
  numWorkers.value = "1"
};
</script>
<template>
  <div class="controls">
    <h2>Submit <span class="jobs-green">jobs</span></h2>
    <div class="input-container">
      <input
        type="number"
        name="worker count"
        id="workerCount"
        :value="numWorkers"
        @input="(event) => {
          const target = event.target as HTMLInputElement
          numWorkers = target.value
          }"
      />
      <button id="short" @click="onShortJobSubmit">Short</button>
      <button id="long" @click="onLongJobSubmit">Long</button>
    </div>
  </div>
</template>
<style>
.jobs-green {
  color: greenyellow;
  font-style: bold;
}

.controls {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
}
.input-container {
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: center;
  gap: 5%;
}
</style>
