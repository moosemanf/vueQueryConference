import { ref } from 'vue'
import { defineStore } from 'pinia'

export const useDirtyStore = defineStore('dirty', () => {
  const isDirty = ref(false)
  
  function setDirtyState(payload: boolean) {
    isDirty.value = payload
  }

  return { isDirty, setDirtyState }
})
