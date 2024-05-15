<template>
  <h1>Master with Details and Crud action</h1>
  <div class="space-x-4 mb-8">
    <n-button type="primary"> New </n-button>
    <n-button :disabled="!store.isDirty"> Save </n-button>
    <n-button type="error" :disabled="!id"> Delete </n-button>
  </div>

  <div class="flex gap-x-12">
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
import { useQuery } from '@tanstack/vue-query'
import { getCustomers, type Entity, useGetAllCustomers } from '@/api/customers'

const route = useRoute()
const router = useRouter()
const props = defineProps<{
  id?: string
}>()

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
            onClick: () => deleteEntity(row.id),
          },
          { default: () => 'Delete' }
        )
      },
    },
  ]
}

const { isPending, isError, isFetching, data, error, refetch } =
  useGetAllCustomers()

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

function deleteEntity(id: number) {
  message.error(`Soon to be deleted: ${id}`)
}

const store = useDirtyStore()

function changeHandler(input: string) {
  store.setDirtyState(true)
}
</script>

<style>
.selected-row {
  @apply !bg-red-300;
  background-color: blue !important;
}
</style>
