<template>
  <section class="flex flex-col gap-y-4">
    <div v-if="isLoading">chill guys...</div>
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
import { computed, h, toRef, type Ref } from 'vue'
import { NInput, NSelect, useMessage, type DataTableColumns } from 'naive-ui'
import { useRoute } from 'vue-router'
import {
  useGetCustomer,
  type Entity,
  type Field,
  type Selectable,
  customersOptions,
  deleteCustomerQuery,
} from '@/api/customers'
import { useGetNations, nationOptions } from '@/api/nations'
import { useMutation, useQuery, useQueryClient } from '@tanstack/vue-query'
import BaseInput from '@/components/BaseInput.vue'
import { useDirtyStore } from '@/stores/dirty'

const route = useRoute()
const dirtyStore = useDirtyStore()
const queryClient = useQueryClient()

const props = defineProps<{
  id: string
}>()

// const queryKey = computed(() => ['customer', props.id])
const { isLoading, data, isError, error } = useGetCustomer(
  toRef(props, 'id'),
  computed(() => !dirtyStore.isDirty)
)

const { isLoading: isLoadingNations, data: nations } = useGetNations({
  isSelectable: true,
})

function changeHandler(field: Field) {
  handleChangeAndOptimisticallyUpdateFE(field)
}

function handleChangeAndOptimisticallyUpdateFE(field: Field) {
  const { queryKey } = customersOptions()
  // get Masterlist
  const entities = queryClient.getQueryData(queryKey)
  const oldEntity = entities?.find((e) => String(e.id) === props.id)

  if (!entities || !oldEntity) return

  // prepare for optimistic update
  const fields = updateFieldList(data, field)
  const firstName = findFieldValue(fields, 'firstName')
  const lastName = findFieldValue(fields, 'lastName')

  const newEntity = { ...oldEntity, firstName, lastName }
  // const newEntity = { ...oldEntity, fields }
  dirtyStore.setDirtyState(true)

  const newEntities = entities.map((old) => {
    return String(old.id) === props.id ? newEntity : old
  }) as Entity[]
  queryClient.setQueryData(queryKey, newEntities)
}

function updateFieldList(data: Ref<Entity> | Ref<undefined>, field: Field) {
  return (
    data?.value?.fields?.map((oldField) => {
      return oldField.name === field.name ? field : oldField
    }) ?? []
  )
}

function findFieldValue(fields: Field[], name: string) {
  return fields.find((f) => f.name === name)?.value ?? ''
}
</script>
