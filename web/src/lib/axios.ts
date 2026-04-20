// src/lib/axios.ts
import axios from 'axios'
import type { AxiosRequestConfig } from 'axios'

const api = axios.create({
  baseURL: '/',
  withCredentials: true,
  headers: {
    'Content-Type': 'application/json',
  },
})

api.interceptors.response.use(
  (response) => {
    return response.data
  },
  (error) => {
    if (error.response?.status === 401) {
      window.dispatchEvent(new CustomEvent('auth:unauthorized'))
    }
    return Promise.reject(error)
  },
)

const get = <T>(url: string, config?: AxiosRequestConfig): Promise<T> =>
  api.get<T, any>(url, config)

const post = <T, TData = unknown>(
  url: string,
  data?: TData,
  config?: AxiosRequestConfig,
): Promise<T> => api.post(url, data, config)

const put = <T, TData = unknown>(
  url: string,
  data?: TData,
  config?: AxiosRequestConfig,
): Promise<T> => api.put(url, data, config)

const del = <T>(url: string, config?: AxiosRequestConfig): Promise<T> =>
  api.delete(url, config)

export const apiClient = { get, post, put, del }
