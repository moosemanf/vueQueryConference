<template>
  <h1>Master with Details and Crud action</h1>
  <div class="space-x-4 mb-8">
    <n-button type="primary"> New </n-button>
    <n-button :disabled="!store.isDirty" @click="onUpdate"> Save </n-button>
    <n-button type="error" ghost :disabled="!id" @click="onReset">
      Reset
    </n-button>
  </div>

  <div v-if="isPending">master is loading...</div>
  <div v-else-if="isError">showing my errors: {{ error }}</div>

  <div v-else class="flex gap-x-12">
    <n-data-table
      class="w-[300px]"
      :columns="columns"
      :data="data"
      :bordered="true"
      :row-props="rowProps"
      :row-class-name="rowName"
    />
    <RouterView />
  </div>
</template>

<script setup lang="ts">
import { computed, h, toRef, type Ref } from 'vue'
import { NDataTable, NButton, useMessage } from 'naive-ui'
import { useRoute, useRouter } from 'vue-router'
import { useDirtyStore } from '@/stores/dirty'
import { useMutation, useQueryClient } from '@tanstack/vue-query'
import {
  type Entity,
  useGetAllCustomers,
  deleteCustomerQuery,
  updateCustomerQuery,
  useGetCustomer,
} from '@/api/customers'

const props = defineProps<{
  id?: string
}>()

const router = useRouter()
const route = useRoute()
const store = useDirtyStore()
const queryClient = useQueryClient()
const message = useMessage()

const { isError, data, error, isPending } = useGetAllCustomers(
  computed(() => !store.isDirty)
)

function clickedHandler(id: number) {
  router.push({ name: 'Detail', params: { id } })
}

/*********** delete */

const { mutate: deleteEntity } = useDeleteCustomer()

function onDelete(id: number) {
  const entityToDelete = data.value?.find((entity) => entity.id === id)
  if (!entityToDelete) {
    message.error(`Could not find ${id}!`)
    return
  }

  deleteEntity(entityToDelete)
}

function useDeleteCustomer() {
  const queryKey = ['customers']
  return useMutation({
    mutationFn: (entity: Entity) =>
      deleteCustomerQuery(entity.id),

    onMutate: async (entityToDelete) => {
      // Cancel current queries for the customers list
      await queryClient.cancelQueries({ queryKey })

      // Remove entity opportunistically from list
      queryClient.setQueryData<Entity[]>(queryKey, (old) => {
        return old?.filter((entity) => entity.id !== entityToDelete.id) ?? []
      })
    },

    onError: (error, entityNotDeleted) => {
      // An error happened!
      message.error(
        `NOT deleted: ${entityNotDeleted.firstName} (${nice(error)})`
      )

      // Add entity back to list
      queryClient.setQueryData<Entity[]>(queryKey, (old) => {
        if (!old || old.length === 0) return [entityNotDeleted]
        return [...old, entityNotDeleted].sort((a, b) => a.id - b.id)
      })
    },

    onSuccess: (_, entityDeleted) => {
      message.success(`Deleted: ${entityDeleted.firstName}`)
    },
    // onSettled: (data, error, variables, context) => {
    //   // Error or success... doesn't matter!
    // },
  })
}

/*********** update */

const { mutate: update } = useUpdateCustomer()

function onUpdate() {
  const id = Number(props.id ?? route.params.id)
  const entityToUpdate = data.value?.find((entity) => entity.id === id)
  if (!entityToUpdate) {
    message.error(`Could not find ${id}!`)
    return
  }

  update(entityToUpdate)
}

function useUpdateCustomer() {
  const queryKey = ['customers']
  return useMutation({
    mutationFn: (entity: Entity) => updateCustomerQuery(entity),

    onMutate: async (entity: Entity) => {
      // Cancel current queries for the customers list
      // await queryClient.cancelQueries({ queryKey })
      await queryClient.cancelQueries({ queryKey: [ ...queryKey, entity.id] })

      // List was updated via field change already!
    },

    onError: (error, entityUpdated) => {
      // An error happened!
      message.error(`NOT updated: ${entityUpdated.id} (${nice(error)})`)
      // Add entity back to list
      store.setDirtyState(false)
      queryClient.invalidateQueries({ queryKey })
      queryClient.resetQueries({ queryKey: [ ...queryKey, String(entityUpdated.id)] })
    },

    onSuccess: (_, entityUpdated) => {
      message.success(`Updated: ${entityUpdated.firstName}`)
      store.setDirtyState(false)
    },
    // onSettled: (data, error, variables, context) => {
    //   // Error or success... doesn't matter!
    // },
  })
}

function onReset() {
  console.log('onReset')
  queryClient.invalidateQueries({ queryKey: ['customers'] })
  queryClient.resetQueries({ queryKey: [ 'customers', props.id]})
  store.setDirtyState(false)
}

const nice = (error: any) => {
  return error.response?.data?.error
}

const createColumns = () => {
  return [
    {
      title: 'ID',
      key: 'id',
    },
    {
      title: 'First name',
      key: 'firstName',
    },
    {
      title: 'Last Name',
      key: 'lastName',
    },
    {
      title: 'Action',
      key: 'actions',
      render(row: Entity) {
        return h(
          NButton,
          {
            strong: true,
            tertiary: true,
            size: 'small',
            onClick: () => onDelete(row.id),
          },
          { default: () => 'Delete' }
        )
      },
    },
  ]
}
const columns = createColumns()

const rowProps = (item: Entity) => {
  return {
    style: 'cursor: pointer; whitespace-pre-wrap;',
    onClick: () => {
      clickedHandler(item.id)
      return {
        class: 'selected-row',
      }
    },
  }
}

const rowName = (row: Entity): string => {
  const id = props.id ?? route.params.id
  const selected = id && String(row.id) === String(id) ? 'selected-row' : ''
  return selected
}
</script>

<style>
.selected-row td {
  @apply !bg-red-300;
  background-color: skyblue !important;
}
</style>
