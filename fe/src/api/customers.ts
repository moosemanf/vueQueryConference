import { queryOptions, useMutation, useQuery, type UseQueryReturnType } from '@tanstack/vue-query'
import axios from 'axios'
import { toRef, type Ref,} from 'vue'
export type Entity = {
    id: number
    name: string
    fields: Field[]
    // firstName: string
    // lastName: string
    // note: string
    // nation: {id: number, name: string}
    // nationId: number
}

export type Field = {
    name: 'firstName' | 'lastName' | 'note' | 'nationId'
    value: string | number
    type: 'text' | 'select'
    query?: string
}

export type Selectable = {
    label: string
    value: number
}

export const getCustomers = async (): Promise<Entity[]> =>
    (await axios.get('/api/customers')).data

export function useGetAllCustomers() {
    return useQuery(customersOptions())

}   

const getCustomer = async (id: string): Promise<Entity> => {
    const { data } = await axios.get(`/api/customers/${id}`)
    return data
}

export function useGetCustomer(id: Ref<string>) {
    return useQuery(customerOptions(id))
} 

const { isPending, isError, error, isSuccess, mutate } = useMutation<Entity>({
    mutationFn: (newCustomer) => axios.post('/api/customers', newCustomer),
  })

export function usePostCustomer(newCustomer: Entity) {
    return mutate(newCustomer)
}

export function customerOptions(id: Ref<string>) {
    return queryOptions({
        queryKey: ['customers', id],
        queryFn: () => getCustomer(id.value),
        staleTime: 5 * 1000,
    })
}

export function customersOptions() {
    return queryOptions({
        queryKey: ['customers'],
        queryFn: () => getCustomers(),
        staleTime: 15 * 1000,
    })
}
