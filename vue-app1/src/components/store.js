import { reactive } from 'vue'

export const store = reactive({
  token: "",
  apiBaseURL: import.meta.env.VITE_API_BASE_URL
})