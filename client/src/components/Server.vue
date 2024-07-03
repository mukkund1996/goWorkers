<script lang="ts">
import { defineProps } from 'vue';
const props = defineProps<{
  online: boolean;
}>();

</script>
<template>
<div :class="props.online ? 'online': 'offline'">
    <i class="pi pi-cog worker"></i>
</div>
</template>
<style>
    .worker {
        font-size: 2.5rem;
    }

    .online {
        border: 2px solid green;
        background-color: lightgreen;
    }
    .online .worker {
        color: green;
    }

    .offline {
        border: 2px solid red;
        background-color: lightgray;
    }
    .offline .worker {
        color: red
    }
</style>