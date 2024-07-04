<script setup lang="ts">
const props = defineProps<{
  busy: boolean;
}>();
</script>
<template>
  <div :class="props.busy ? 'busy' : 'free'">
    <i class="pi pi-cog worker"></i>
  </div>
</template>
<style>
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
