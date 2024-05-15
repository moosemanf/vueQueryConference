<template>
  <label class="flex items-center">
    <span class="basis-1/3">{{ labelMap[field.name] }}</span>

    <input
      v-if="field.type === 'text'"
      class="basis-1/2"
      :value="field.value"
      @blur="handleBlur"
    />
    <n-select
      v-if="field.type === 'select'"
      :value="field.value"
      :options="options"
      :loading="isLoading"
      class="basis-1/2 grow"
      @update:value="handleSelect"
    />
  </label>
</template>

<script setup lang="ts">
import type { Field } from '@/api/customers'
import { defineProps, defineEmits, computed, ref } from 'vue'
import { NInput, NSelect, useMessage, type DataTableColumns } from 'naive-ui'
import { useQuery } from '@tanstack/vue-query'
import { selectOptions } from '@/api/nations'
import type { Selectable } from '@/api/customers'

interface Props {
  field: Field
}

const props = defineProps<Props>()

const emit = defineEmits<{
  (e: 'blur', field: Field): void
}>()

const labelMap = {
  firstName: 'First name',
  lastName: 'Last name',
  note: 'Note',
  nationId: 'Nation',
} as const

const { isLoading, data: options } = useQuery(selectOptions(props.field.query))

const handleBlur = (event: FocusEvent) => {
  const value = (event.target as HTMLInputElement).value
  emit('blur', { ...props.field, value })
}

const handleSelect = (value: number) => {
  console.log('select: ', value)
  emit('blur', { ...props.field, value })
}
</script>
