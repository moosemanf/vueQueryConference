<template>
  <n-select
    :value="props.field.value"
    :options="options"
    :loading="isLoading"
    @update:value="handleSelect"
  />
</template>

<script setup lang="ts">
import type { Field } from '@/api/customers'
import { defineProps, defineEmits, ref, computed } from 'vue'
import { NSelect } from 'naive-ui'
import { useQuery } from '@tanstack/vue-query'
import { selectOptions } from '@/api/select'

interface Props {
  field: Field
}

const props = defineProps<Props>()

const emit = defineEmits<{
  (e: 'blur', field: Field): void
}>()


const { isPending: isLoading, data: options } = useQuery(selectOptions(props.field.query))

const handleSelect = (value: number) => {
  emit('blur', { ...props.field, value })
}
</script>
