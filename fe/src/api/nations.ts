import { queryOptions, useQuery, type UseQueryReturnType } from '@tanstack/vue-query'
import axios from 'axios'
import type { Entity, Selectable } from './customers'
import { computed } from 'vue'


async function getNations(isSelect: boolean, queryKey: string[]) {
    console.log('queryKey:', queryKey)
    const { data } = await axios.get('/api/nations')
    return isSelect ? toSelect(data) as Selectable[] : data as Entity[]
}
async function getSelect(query?: string) {
    console.log('query:', query)
    const { data } = await axios.get(`/api/${query}`)
    return toSelect(data)
}
export function useGetNations({ isSelectable = false }: { isSelectable: boolean }) {
    return useQuery(nationOptions(isSelectable))
}

const toSelect = (entities: Pick<Entity, "id" | "name">[]) =>
    entities.map(({ name: label, id: value }) => ({ label, value })) as Selectable[]

export function nationOptions(isSelectable: boolean) {
    return queryOptions({
        queryKey: ['nations'],
        queryFn: ({ queryKey }) => getSelect(queryKey),
        staleTime: 5 * 1000,
    })
}

export function selectOptions(query?: string) {
    return queryOptions({
        queryKey: [query],
        queryFn: () => getSelect(query),
        staleTime: 5 * 1000,
        enabled: !!query
    })
}
