import { reactive } from 'vue'

export const store = reactive({
  token: "",
  user: {},
  apiBaseURL: import.meta.env.VITE_API_BASE_URL,
  imgPath: import.meta.env.VITE_API_IMG_URL,
})