<script setup lang="ts">
</script>

<template>
  <h1>Welcome to modern state management (in UI)</h1>
    <h2>There are three kinds of state</h2>
    <ul class="list-disc ml-8">
      <li><strong>Local client</strong> state:</li>
      <ul class="list-disc ml-8">
        <li>values that are scoped directly to a single component or its descendants</li>
        <li>right where you need it</li>
      </ul>
      <li><strong>Global client</strong> state:</li>
      <ul class="list-disc ml-8">
        <li>for example: darkmode, user (-settings),... </li>
        <li>think Vuex, Pinia, Redux</li>
        <li>we use that way too often (think modals)</li>
      </ul>
      <li><strong>Server</strong> state:</li>
      <ul class="list-disc ml-8">
        <li>values that are fetched from a server via an API and cached on the client</li>
        <li>Is persisted remotely in a location you do not control or own</li>
        <li>Requires asynchronous APIs for fetching and updating</li>
        <li>Implies shared ownership and can be changed by other people without your knowledge</li>
        <li>Can potentially become "out of date" in your applications if you're not careful</li>
      </ul>
    </ul>
    <h2 class="font-bold text-xl ml-6 mt-4">Challenges for server state</h2>
    <ul class="list-disc ml-4">
<li><strong>Caching...</strong> (possibly the hardest thing to do in programming)</li>
<li><strong>Deduping</strong> multiple requests for the same data into a single request</li>
<li>Updating <strong>"out of date"</strong> data in the background</li>
<li>Knowing when data is "out of date"</li>
<li>Reflecting updates to data <strong>as quickly as possible</strong></li>
<li>Performance optimizations like pagination and <strong>lazy loading data</strong></li>
<li>Managing memory and <strong>garbage collection of server state</strong></li>
<li>Memoizing query results with structural sharing</li>
    </ul>
</template>

<style>
li {
  @apply ml-12;
}
</style>