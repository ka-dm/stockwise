<script setup lang="ts">
import { ref, onMounted } from 'vue'

// Definir la interfaz para el tipo Stock
interface Stock {
  id: number
  ticker: string
  target_from: string
  target_to: string
  company: string
  action: string
  brokerage: string
  rating_from: string
  rating_to: string
  time: string
  created_at: string
}

// Variables reactivas
const stocks = ref<Stock[]>([])
const loading = ref(true)
const error = ref<string | null>(null)

// Función para obtener los stocks de la API
const fetchStocks = async () => {
  try {
    loading.value = true
    error.value = null
    
    const response = await fetch('http://192.168.12.149:8000/stocks')
    
    if (!response.ok) {
      throw new Error(`Error: ${response.status} ${response.statusText}`)
    }
    
    const data = await response.json()
    stocks.value = data
  } catch (err) {
    error.value = err instanceof Error ? err.message : 'Error desconocido'
    console.error('Error fetching stocks:', err)
  } finally {
    loading.value = false
  }
}

// Obtener los datos cuando el componente se monte
onMounted(() => {
  fetchStocks()
})
</script>

<template>
  <main class="bg-gray-100">
    <section class="mb-4 p-4 bg-white">
      <h1 class="text-2xl p-4 text-gray-900 mb-4">
        Ver tendencias de stocks
      </h1>

      <ul class="flex overflow-x-auto overflow-y-hidden" style="scrollbar-width: none; -ms-overflow-style: none;">
          <li class="me-2 whitespace-nowrap">
              <a href="/" class="inline-flex items-center rounded-2xl ring-0 ring-inset bg-blue-100 px-2 py-1 text-sm text-blue-600 hover:bg-blue-200 active whitespace-nowrap" aria-current="page">
                <span class="material-icons text-gray-700 pr-1" style="font-size: 18px;">stacked_line_chart</span>Índices de mercado
              </a>
          </li>
          <li class="me-2 whitespace-nowrap">
              <a href="/"  class="inline-flex items-center rounded-2xl ring-1 ring-inset ring-gray-200 px-2 py-1 text-sm text-gray-600 hover:bg-gray-100">
                <span class="material-icons text-blue-500 pr-1" style="font-size: 18px;">bar_chart</span> Mayor actividad
              </a>
          </li>
          <li class="me-2 whitespace-nowrap">
              <a href="/"  class="inline-flex items-center rounded-2xl ring-1 ring-inset ring-gray-200 px-2 py-1 text-sm text-gray-600 hover:bg-gray-100">
                <span class="material-icons text-green-500 pr-1" style="font-size: 18px;">trending_up</span> Mayores subidas
              </a>
          </li>
          <li class="me-2 whitespace-nowrap">
              <a href="/"  class="inline-flex items-center rounded-2xl ring-1 ring-inset ring-gray-200 px-2 py-1 text-sm text-gray-600 hover:bg-gray-100">
                <span class="material-icons text-red-500 pr-1" style="font-size: 18px;">trending_down</span> Mayores bajadas
              </a>
          </li>
          <li class="me-2 whitespace-nowrap">
              <a href="/"  class="inline-flex items-center rounded-2xl ring-1 ring-inset ring-gray-200 px-2 py-1 text-sm text-gray-600 hover:bg-gray-100">
                <span class="material-icons text-green-500 pr-1" style="font-size: 18px;">eco</span> Líderes en sostenibilidad
              </a>
          </li>
          <li class="me-2 whitespace-nowrap">
              <a href="/"  class="inline-flex items-center rounded-2xl ring-1 ring-inset ring-gray-200 px-2 py-1 text-sm text-gray-600 hover:bg-gray-100">
                <span class="material-icons text-yellow-500 pr-1" style="font-size: 18px;">currency_bitcoin</span> Criptomonedas
              </a>
          </li>
          <li class="me-2 whitespace-nowrap">
              <a href="/"  class="inline-flex items-center rounded-2xl ring-1 ring-inset ring-gray-200 px-2 py-1 text-sm text-gray-600 hover:bg-gray-100">
                <span class="material-icons text-yellow-500 pr-1" style="font-size: 18px;">attach_money</span> Divisas
              </a>
          </li>
      </ul>

    </section>
      
    <!-- Mostrar loading -->
    <div v-if="loading" class="text-center py-8">
      <p class="text-lg text-gray-600">Cargando stocks...</p>
    </div>
    
    <!-- Mostrar error -->
    <div v-else-if="error" class="bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded mb-4">
      <p>Error al cargar los stocks: {{ error }}</p>
      <button @click="fetchStocks" class="mt-2 bg-red-500 hover:bg-red-700 text-white font-bold py-2 px-4 rounded">
        Reintentar
      </button>
    </div>
    
    <!-- Mostrar tabla de stocks -->
    <div v-else-if="stocks.length > 0" class="overflow-x-auto">
      <table class="table-auto w-full border-collapse shadow-lg overflow-hidden">
        <thead>
          <tr class="bg-white">
            <th class="px-4 py-3 text-left font-semibold" colspan="7">Top mejores recomendaciones</th>
          </tr>
        </thead>
        <tbody>
          <tr 
            v-for="stock in stocks" 
            :key="stock.id" 
            class="bg-white even:bg-gray-100 hover:bg-yellow-50 transition border-y border-gray-200"
          >
            <td class="flex flex-col px-4 py-4 max-w-70 font-semibold">
              <span class="inline-flex items-center rounded-sm ring-1 ring-inset ring-green-400 p-1 text-sm text-green-800 bg-green-100 w-fit">
                {{ stock.ticker }}
              </span>
              <span class="text-gray-500 text-xs whitespace-nowrap">{{ stock.company }}</span>
              <span class="text-gray-700 text-sm whitespace-nowrap mt-2">{{ stock.brokerage }}</span>
            </td>
            <td class="px-4 py-4">
              <span 
                :class="{
                  'text-green-600 font-semibold': stock.action === 'Buy',
                  'text-red-600 font-semibold': stock.action === 'Sell',
                  'text-yellow-600 font-semibold': stock.action === 'Hold'
                }"
              >
                {{ stock.action }}
              </span>
            </td>
            <td class="px-4 py-4">
              {{ stock.rating_from }} 
              <span v-if="stock.rating_to && stock.rating_to !== stock.rating_from" class="text-gray-500">
                → {{ stock.rating_to }}
              </span>
            </td>
            <td class="px-4 py-4">
              {{ stock.target_from }}
              <span v-if="stock.target_to && stock.target_to !== stock.target_from" class="text-gray-500">
                → {{ stock.target_to }}
              </span>
            </td>
            <td class="px-4 py-4 text-sm text-gray-600">{{ stock.time }}</td>
          </tr>
        </tbody>
      </table>
    </div>
    
    <!-- Mostrar mensaje cuando no hay datos -->
    <div v-else class="text-center py-8">
      <p class="text-lg text-gray-600">No hay stocks disponibles</p>
      <button @click="fetchStocks" class="mt-4 bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded">
        Cargar Stocks
      </button>
    </div>
  </main>
</template>
