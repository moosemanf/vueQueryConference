<template>
  <div class="bg-slate-100 basis-1/5 px-8 py-8" >

    <h2 class="ml-0">Top customers:</h2>
    
    <div v-if="isPending">master is loading...</div>
    <div v-else-if="isError">showing my errors: {{ error }}</div>
    
    <div v-else class="flex flex-col gap-x-4">
      <div v-for="customer in data" :key="customer.id">
        {{ customer.firstName }}
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { useDirtyStore } from '@/stores/dirty'
import { useGetAllCustomers } from '@/api/customers'

const store = useDirtyStore()

const { isError, data, error, isPending } = useGetAllCustomers(
  computed(() => !store.isDirty)
)
</script>
