import type { Ref } from "vue"
import type { Entity, Field } from "./customers"
import { useQueryClient, type QueryFunctionContext } from '@tanstack/vue-query'

interface oiae {
queryKey: string[]
}
export function handleChangeAndOptimisticallyUpdateFE(field: Field, data: Ref<Entity[]>, customersOptions:oiae) {
const queryClient = useQueryClient()

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