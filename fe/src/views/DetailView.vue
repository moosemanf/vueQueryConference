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
      v-for="field in data?.fields ?? []"
      :key="field.name"
      :field
      @blur="changeHandler"
      />
      <!-- :class="{ 'w-60': field.type === 'select' }" -->
  </section>
</template>

<script setup lang="ts">
import { computed, toRef } from 'vue'
import { useQueryClient } from '@tanstack/vue-query'
import { onBeforeRouteLeave, onBeforeRouteUpdate } from 'vue-router'
import { useDirtyStore } from '@/stores/dirty'
import {
  useGetCustomer,
  type Entity,
  type Field,
  customersOptions,
} from '@/api/customers'
import BaseInput from '@/components/BaseInput.vue'

const dirtyStore = useDirtyStore()
const queryClient = useQueryClient()

const props = defineProps<{
  id: string
}>()

// const queryKey = computed(() => ['customer', props.id])
const { isLoading, data, isError, error } = useGetCustomer(
  toRef(props, 'id'),
  computed(() => !dirtyStore?.isDirty)
)


function changeHandler(field: Field) {
  handleChangeAndOptimisticallyUpdateFE(field)
}

function handleChangeAndOptimisticallyUpdateFE(field: Field) {
  const { queryKey } = customersOptions(toRef(false))
  // get Masterlist
  const entities = queryClient.getQueryData(queryKey)
  const oldEntity = entities?.find((e) => String(e.id) === props.id)

  if (!entities || !oldEntity) return

  const newEntity = { ...oldEntity, [field.name]: field.value }
  dirtyStore.setDirtyState(true)

  const newEntities = entities.map<Entity>((old) => {
    return String(old.id) === props.id ? newEntity : old
  })
  queryClient.setQueryData(queryKey, newEntities)
}

onBeforeRouteLeave(() => routeGuard())

onBeforeRouteUpdate(() => routeGuard())

function routeGuard() {
  if (dirtyStore.isDirty) {
    const answer = window.confirm(
      'You have unsaved changed, do you really want to leave?'
    )
    // cancel the navigation and stay on the same page
    if (!answer) return false
    dirtyStore.setDirtyState(false)
    queryClient.resetQueries({ queryKey: [ 'customers', props.id], exact: true })
    queryClient.invalidateQueries({ queryKey: [ 'customers'], exact: true })
  }
}
</script>
