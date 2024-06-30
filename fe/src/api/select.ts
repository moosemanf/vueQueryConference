import axios from "axios"
import { queryOptions } from "@tanstack/vue-query"
import type { Entity, Selectable } from './customers'

export function selectOptions(query?: string) {
    return queryOptions({
        queryKey: [query],
        queryFn: () => getSelect(query ?? ''),
        staleTime: 5 * 1000,
        enabled: !!query,
    })
}

export async function getSelect(query: string) {
    if (!query) return []
    const { data } = await axios.get(`/api/${query}`)
    return toSelect(data)
}

export const toSelect = (entities: Pick<Entity, "id" | "name">[]) =>
    entities.map<Selectable>(({ name: label, id: value }) => ({ label, value }))
