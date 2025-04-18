<!-- converted to the Composition API -->
<template>
    <div class="container">
      <div class="row">
        <div class="col">
          <h1 class="mt-3">Manage Books</h1>
          <hr>

          <table v-if="ready" class="table table-striped table-compact">
            <thead>
              <tr>
                <th>Book</th>
                <th>Author</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="b in books" v-bind:key="b.id">
                <td>
                  <router-link :to="`/admin/books/${b.id}`">{{ b.title }}</router-link>
                </td>
                <td>{{ b.author.author_name }}</td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
    </div>
  </template>


<script setup>
import { ref, onMounted } from 'vue';
import { store } from './store.js'

let ready = ref(false);
let books = ref({}) // ref([]) ?? 
let book = ref({}) 

// const props = defineProps({})

const emit = defineEmits(['error'])

defineOptions({
  name: 'BooksAdmin'
});


onMounted(async() => {
  try {
    const response = await fetch(`${store.apiBaseURL}/books`);
    const data = await response.json();
    books.value = data.data.books; 
    ready.value = true;
  } catch(error) {
    emit('error', error.message || error )
  }
});

/*
onMounted(() => {
  fetch(`${store.apiBaseURL}/books`)
    .then(response => response.json())
    .then((data) => {
      if (data.error) {
        emit('error', data.message)
      } else {
        books.value = data.data.books;
        ready.value = true;
      }
    })
    .catch((error) => {
      emit('error', error.message || error )
    });
})
*/

</script>

