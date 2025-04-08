<template>
    <div class="container">
      <div class="row">
        <div class="col">
          <h1 class="mt-3">Manage Books</h1>
          <hr>

          <table v-if="this.ready" class="table table-striped table-compact">
            <thead>
              <tr>
                <th>Book</th>
                <th>Author</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="b in this.books" v-bind:key="b.id">
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


<script>
import { store } from './store.js'

export default{
  data() {
    return{
      book: {
        //author: {},
        //genres: [],
      },
      store,
      imgPath: store.imgPath,
      ready: false,
    }
  },
  activated() {
    fetch(`${store.apiBaseURL}/books`)
    .then(response => response.json())
    .then((data) => {
      if (data.error) {
        this.$emit('error', data.message)
      } else {
        this.books = data.data.books;
        this.ready = true;
      }
    })
    .catch((error) => {

    });
  },
  deactivated() {
    this.ready = false;
  },

}
</script>