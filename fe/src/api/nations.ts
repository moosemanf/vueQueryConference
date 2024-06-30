// import { queryOptions, useQuery, type UseQueryReturnType } from '@tanstack/vue-query'
// import axios from 'axios'
// import type { Entity, Selectable } from './customers'
// import { getSelect, toSelect } from './select'



// async function getNations(isSelect: boolean, queryKey: string[]) {
//     console.log('queryKey:', queryKey)
//     const { data } = await axios.get('/api/nations')
//     return isSelect ? toSelect(data) as Selectable[] : data as Entity[]
// }

// export function useGetNations({ isSelectable = false }: { isSelectable: boolean }) {
//     return useQuery(nationOptions(isSelectable))
// }


// export function nationOptions(isSelectable: boolean) {
//     return queryOptions({
//         queryKey: ['nations'],
//         queryFn: ({ queryKey }) => getSelect(queryKey[0]),
//         staleTime: 5 * 1000,
//     })
// }
