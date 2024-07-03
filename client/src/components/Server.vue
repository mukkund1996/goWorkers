<script lang="ts">
import { defineProps } from 'vue';
const props = defineProps<{isOnline: boolean}>()

</script>
<template>
<div :class="props.isOnline ? 'online': 'offline'">
    <i class="pi pi-cog worker"></i>
</div>
</template>
<style>
    .worker {
        font-size: 2.5rem;
    }

    .online {
        border: 1.5rem solid green;
        background-color: lightgreen;
    }
    .online .worker {
        color: green;
    }

    .offline {
        border: 1.5rem solid red;
        background-color: lightcoral;
    }
    .offline .worker {
        color: red
    }
</style>