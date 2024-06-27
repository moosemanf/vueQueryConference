import { queryOptions, useQuery } from '@tanstack/vue-query'
import axios from 'axios'
import { type Ref,} from 'vue'
export type Entity = {
    id: number
    name: string
    fields: Field[]
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



export function customersOptions(enabled: Ref<boolean>) {
    return queryOptions({
        queryKey: ['customers'],
        queryFn: () => getCustomers(),
        staleTime: 15 * 1000,
        enabled,
    })
}

export const getCustomers = async (): Promise<Entity[]> =>
    (await axios.get('/api/customers')).data

export function useGetAllCustomers(enabled: Ref<boolean>) {
    return useQuery(customersOptions(enabled)) 
}   

const getCustomer = async (id: string): Promise<Entity> => {
    const { data } = await axios.get(`/api/customers/${id}`)
    return data
}

export function useGetCustomer(id: Ref<string>, enabled: Ref<boolean>) {
    return useQuery(customerOptions(id, enabled))
} 

// const { isPending, isError, error, isSuccess, mutate } = useMutation<Entity>({
//     mutationFn: (newCustomer) => axios.post('/api/customers', newCustomer),
//   })

// export function usePostCustomer(newCustomer: Entity) {
//     return mutate(newCustomer)
// }

export function customerOptions(id: Ref<string>, enabled: Ref<boolean>) {
    return queryOptions({
        queryKey: ['customers', id],
        queryFn: () => getCustomer(id.value),
        enabled,
        staleTime: 5 * 1000,
    })
}

export async function deleteCustomerQuery(id: number) {
   await axios.delete(`/api/customers/${id}`)

}

