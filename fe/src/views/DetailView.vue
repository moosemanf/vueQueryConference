<template>
  <!-- <BaseInput v-model="data?.firstName" label="First name" />
    <BaseInput v-model="data?.lastName" label="Last name" /> -->
  <section class="flex flex-col gap-y-4">
    <div v-if="isLoading">chill man...</div>
    <div v-if="isError">Houston, we have a problem: {{ error?.message }}</div>
    <label class="w-28">
      ID:
      <span class="ml-16 pl-2">
        {{ data?.id }}
      </span>
    </label>
    <BaseInput
      v-for="field in data?.fields"
      :key="field.name"
      :value="field.value"
      :field="field"
      :options="nations"
      :class="{ 'w-60': field.type === 'select' }"
      @blur="changeHandler"
    />
  </section>
</template>

<script setup lang="ts">
import { computed, h, toRef } from 'vue'
import { NInput, NSelect, useMessage, type DataTableColumns } from 'naive-ui'
import { useRoute } from 'vue-router'
import {
  useGetCustomer,
  type Entity,
  type Field,
  type Selectable,
  customersOptions,
} from '@/api/customers'
import { useGetNations, nationOptions } from '@/api/nations'
import { useQuery, useQueryClient } from '@tanstack/vue-query'
import BaseInput from '@/components/BaseInput.vue'
import { useDirtyStore } from '@/stores/dirty'

const route = useRoute()
const dirtyStore = useDirtyStore()
const queryClient = useQueryClient()

const props = defineProps<{
  id: string
}>()

// const queryKey = computed(() => ['customer', props.id])
const {
  isLoading,
  isPending,
  data,
  isError,
  error,
} = useGetCustomer(
  toRef(props, 'id')
  // enabled: computed(() => !!props.id)
)

const { isLoading: isLoadingNations, data: nations } = useGetNations({
  isSelectable: true,
})

function changeHandler(field: Field) {
  const entities = queryClient.getQueryData(customersOptions().queryKey)

  if (!entities) return
  const oldEntity = entities.find((e) => String(e.id) === props.id)
  const fields =
    data?.value?.fields?.map((f) => {
      if (f.name === field.name) return field
      return f
    }) ?? []
  const firstName = fields.find((f) => f.name === 'firstName')?.value
  const lastName = fields.find((f) => f.name === 'lastName')?.value
  const entity = { ...oldEntity, firstName, lastName }
  dirtyStore.setDirtyState(true)

  const newE = entities.map((e) => {
    if (String(e.id) === props.id) return entity
    return e
  }) as Entity[]
  queryClient.setQueryData(customersOptions().queryKey, newE)

}
</script>
