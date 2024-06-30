<script setup lang="ts">
import { RouterLink, RouterView } from 'vue-router'
import { NMessageProvider } from 'naive-ui'
import { VueQueryDevtools } from '@tanstack/vue-query-devtools'
import Sidebar from './SideBar.vue'
</script>

<template>
  <VueQueryDevtools />
  <header class="h-20 bg-slate-300 flex items-center">
    <!-- <img alt="Vue logo" class="logo" src="@/assets/logo.svg" width="125" height="125" /> -->
    <span class="font-bold px-10">Vue-Query</span>

    <div class="wrapper">
      <nav class="space-x-6">
        <RouterLink to="/" class="text-bold">Home</RouterLink>
        <RouterLink to="/master">Master&thinsp;/&thinsp;Detail</RouterLink>
        <RouterLink to="/dropdown">Dropdown</RouterLink>
        <RouterLink to="/sse">Server side events</RouterLink>
      </nav>
    </div>
  </header>
  <div class="flex">
    <Sidebar />
    <main class="my-12 mx-20">
      <n-message-provider>
        <RouterView />
      </n-message-provider>
    </main>
  </div>
</template>
