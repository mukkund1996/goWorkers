<script setup lang="ts">
import Workers from "./components/workers/Workers.vue";
</script>

<template>
  <header>
    <nav class="header-container">
      <h1><span class="go-header">go</span>WORKERS</h1>
    </nav>
  </header>
  <body>
    <div class="body-container">
      <Workers></Workers>
    </div>
  </body>
</template>

<style scoped>
body {
  font-family: Comfortaa;
}
h1 {
  font-style: bold;
}

.header-container {
}

.header-container h1,
span {
  font-family: Comfortaa;
  text-align: center;
}

.body-container {
  display: flex;
  justify-content: center;
  align-items: center;
  width: 100%;
}

.go-header {
  color: greenyellow;
  font-style: bold;
}
</style>
