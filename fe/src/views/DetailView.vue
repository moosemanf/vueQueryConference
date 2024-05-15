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
    <!-- <BaseInput :value="data?.fields." label="First Name" @blur="changeHandler" /> -->
    <!-- <label class="flex items-center">
      <span class="w-28"> First Name </span>
      <input
        :value="data?.firstName"
        @blur="(e) => changeHandler(e, 'firstName')"
      />
    </label>
    <label class="flex items-center">
      <span class="w-28"> Last Name </span>
      <input
        :value="data?.lastName"
        @blur="(e) => changeHandler(e, 'lastName')"
      />
    </label>
    <label class="flex items-center">
      <span class="w-28"> Note </span>
      <input
        :value="data?.note"
        placeholder="Tell me something"
        @blur="(e) => changeHandler(e, 'note')"
      />
    </label>
    <label class="flex items-center">
      <span class="w-28"> Nation </span>

      <n-select
        :value="data?.nationId ?? ''"
        :options="nations"
        class="w-60"
        @update:value="selectHandler"
      />
    </label> -->
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

const queryKey = computed(() => ['customer', props.id])
const {
  isLoading,
  isPending,
  data: customers,
  isError,
  error,
} = useGetCustomer(
  toRef(props, 'id')
  // enabled: computed(() => !!props.id)
)

const data = computed(() => {
  const fields =
    customers.value?.fields.map((field) => {
      // if (field.name === 'nationId') {
      //   return { ...field, options: () => useQuery(nationOptions(true)) }
      // }
      return field
    }) ?? []
  return { ...customers.value, fields }
})

const { isLoading: isLoadingNations, data: nations } = useGetNations({
  isSelectable: true,
})

function changeHandler(field: Field) {
  const entities = queryClient.getQueryData(customersOptions().queryKey)

  if (!entities) return
  const oldEntity = entities.find((e) => String(e.id) === props.id)
  const fields =
    data?.value.fields?.map((f) => {
      if (f.name === field.name) return field
      return f
    }) ?? []
  const firstName = fields.find((f) => f.name === 'firstName')?.value
  const lastName = fields.find((f) => f.name === 'lastName')?.value
  const entity = { ...oldEntity, firstName, lastName }
  dirtyStore.setDirtyState(true)

  // if (oldValue !== field.value) {
  // }

  // if (name in entities) {
  const newE = entities.map((e) => {
    if (String(e.id) === props.id) return entity
    return e
  }) as Entity[]
  const newEntities = [...entities, entity] as Entity[]
  console.log('newEntities:', newEntities)
  console.log('newEntities:', newE)
  queryClient.setQueryData(customersOptions().queryKey, newE)
  // }
  // const newFields = entity?.fields.map((field) => {
  //   if (
  //     isSingleSelectableField(changedField) &&
  //     !field.name.endsWith('Id') &&
  //     changedField.name.startsWith(field.name)
  //   ) {
  //     return { ...field, value: changedField.value.name }
  //   }
  //   return field.name === changedField.name ? changedField : field
  // })
}
</script>
