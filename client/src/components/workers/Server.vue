<script setup lang="ts">
const props = defineProps<{
  busy: boolean;
  name: string;
}>();
</script>
<template>
  <div :class="`${props.busy ? 'busy' : 'free'} worker-container`">
    <i class="pi pi-cog worker"></i>
    <span class="worker-label">{{ props.name }}</span>
  </div>
</template>
<style>
.worker-container {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  padding: 1.2rem 1.2rem;
  gap: 1.5rem;
}

.worker-label {
  width: 100%;
}

.worker {
  font-size: 4.5rem;
}

.free {
  border: 2px solid green;
  background-color: lightgreen;
}
.free .worker {
  color: green;
}

.busy {
  border: 2px solid red;
  background-color: lightgray;
}
.busy .worker {
  color: red;
}
</style>
