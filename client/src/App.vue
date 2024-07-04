<script setup lang="ts">
import JobPanel from "./components/jobs/JobPanel.vue";
import Queue from "./components/jobs/Queue.vue";
import Results from "./components/jobs/Results.vue"
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
      <JobPanel></JobPanel>
      <div class="monitor-panel">
        <Queue></Queue>
        <Results></Results>
      </div>
    </div>
  </body>
</template>

<style scoped>
body {
  font-family: Comfortaa;
  background-color: #f4f4f4;
  padding: 1.2rem;
  border-radius: 2rem;
}
h1 {
  font-style: bold;
}

.header-container h1,
span {
  font-family: Comfortaa;
  text-align: center;
}

.body-container {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  width: 100%;
}

.go-header {
  color: greenyellow;
  font-style: bold;
}

.monitor-panel {
  display: flex;
  justify-content: center;
  align-items: center;
  flex-direction: row;
  gap: 10%;
  margin-top: 20px;
}
</style>
