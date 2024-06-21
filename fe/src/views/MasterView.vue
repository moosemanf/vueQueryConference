<template>
  <h1>Master with Details and Crud action</h1>
  <div class="space-x-4 mb-8">
    <n-button type="primary"> New </n-button>
    <n-button :disabled="!store.isDirty"> Save </n-button>
    <!-- <n-button type="error" :disabled="!id"> Delete </n-button> -->
    <n-button type="error" ghost :disabled="!id" @click="onReset">
      Reset
    </n-button>
  </div>

  <div v-if="isPending">master is loading...</div>
  <div v-if="isError">showing my errors: {{ error }}</div>

  <div v-else class="flex gap-x-12">
    <n-data-table
      class="w-[400px]"
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
import { computed, h, ref } from 'vue'
import {
  NDataTable,
  NButton,
  useMessage,
  type DataTableColumns,
} from 'naive-ui'
import { useRoute, useRouter } from 'vue-router'
import { useDirtyStore } from '@/stores/dirty'
import { useMutation, useQuery, useQueryClient } from '@tanstack/vue-query'
import {
  getCustomers,
  type Entity,
  useGetAllCustomers,
  deleteCustomerQuery,
} from '@/api/customers'

const props = defineProps<{
  id?: string
}>()

const router = useRouter()
const route = useRoute()
const store = useDirtyStore()
const queryClient = useQueryClient()
// type Entity = { id: string }

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

const { isError, data, error, isPending } = useGetAllCustomers(
  computed(() => !store.isDirty)
)

// const data: Song[] = [
//   { id: 3, title: 'Wonderwall', length: '4:18' },
//   { id: 4, title: "Don't Look Back in Anger", length: '4:48' },
//   { id: 12, title: 'Champagne Supernova', length: '7:27' }
// ]

const message = useMessage()
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

function clickedHandler(id: number) {
  // message.warning('Yuhuu!!')

  router.push({ name: 'Detail', params: { id } })
}

function onDelete(id: number) {
  console.log('onDelete', id)
  mutate(id)
  // handleOptimisticDelete(id)
}
const { mutate } = useDeleteCustomer()

function useDeleteCustomer() {
  return useMutation({
    mutationFn: (id: number) => deleteCustomerQuery(id),
    onMutate: async (id) => {
      console.log('onMutate', id)
      // Cancel current queries for the todos list
      await queryClient.cancelQueries({ queryKey: ['customers'] })

      queryClient.setQueryData<Entity[]>(['customers'], (old) => {
          return old?.filter((entity) => entity.id !== id) ?? []
        })

      // Return context 
      return { id }
    },
    onError: (error, variables, context) => {
      console.log('onError', 'variables', variables)
      console.log('onError', 'error', error)
      console.log('onError', 'context', context)
      // An error happened!
      

      queryClient.setQueryData<Entity[]>(['customers'], (old: any) => {
        console.log('onError_inMap', old)
          return (old?.length ?? 0) > 0
            ? [...old, variables]
            : [variables]
        })
      if (context) {
        console.log(`rolling back optimistic update with id ${context.id}`)
      }
    },
    onSuccess: (_, id, __) => {
  message.success(`Deleted: ${id}`)

    },
    // onSettled: (data, error, variables, context) => {
    //   console.log('onSettled', 'data', data)
    //   console.log('onSettled', 'error', error)
    //   console.log('onSettled', 'variables', variables)
    //   console.log('onSettled', 'context', context)
    //   // Error or success... doesn't matter!
    // },
  })
}

// function handleOptimisticDelete(id: number) {
//   queryClient.
// }

function changeHandler(input: string) {
  store.setDirtyState(true)
}

function onReset() {
  console.log('onReset')
  queryClient.invalidateQueries({ queryKey: ['customers', props.id] })
  store.setDirtyState(false)
}
</script>

<style>
.selected-row {
  @apply !bg-red-300;
  background-color: blue !important;
}
</style>
