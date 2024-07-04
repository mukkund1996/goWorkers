<script setup lang="ts">
import Server from "./components/workers/Server.vue";
</script>

<template>
  <header>
    <nav class="header-container">
      <h1><span class="go-header">go</span>WORKERS</h1>
    </nav>
  </header>
  <body>
    <div class="body-container">
      <h1>More stuff</h1>
      <Server :online="true"></Server>
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
