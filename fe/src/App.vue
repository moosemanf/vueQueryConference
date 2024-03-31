<script setup lang="ts">
import { RouterLink, RouterView } from 'vue-router'
import HelloWorld from './components/HelloWorld.vue'
</script>

<template>
  <header>
    <!-- <img alt="Vue logo" class="logo" src="@/assets/logo.svg" width="125" height="125" /> -->

    <div class="wrapper">
      <span class="text-red-400">oida</span>
      <nav class="bg-red-400">
        <RouterLink to="/" class="text-bold">Home</RouterLink>
        <RouterLink to="/master">Master / Client</RouterLink>
        <RouterLink to="/dropdown">Dropdown</RouterLink>
      </nav>
    </div>
  </header>

  <RouterView />
</template>
