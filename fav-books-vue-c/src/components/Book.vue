<template>
  <div class="container">
    <div class="row">
      <div class="col md-2">
        <img v-if="this.ready" class="img-fluid img-thumbnail" :src="`${imgPath}/covers/${book.slug}.jpg`" alt="cover">
      </div>
      <div class="col md-10">
        <template v-if="this.ready">
          <h3 class="mt-3">{{ book.title }}</h3><hr>
          <p>
            <strong>Author:</strong> {{ book.author.author_name }} <br>
            <strong>Published:</strong> {{ book.publication_year }} 
          </p>
          <p>
            {{ book.description }}
          </p>
        </template>
        <p v-else>Loading</p>
      </div>
    </div>
  </div>
</template>

<script>
import { store } from "./store.js";

export default{
  name: 'Book',
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
  mounted() {
    fetch(`${store.apiBaseURL}/books/${this.$route.params.bookName}`)
    .then(response => response.json())
    .then((data) => {
      if (data.error) {
        this.$emit('error', data.message)
      } else {
        this.book = data.data;
        this.ready = true;
        //console.log("Titile is", this.book.title)
      }
    })
    .catch((error) => {

    });
  },
  //activated is actually fired any time a cached component is redisplayed and it is also called every time mounted is called
  //but  you can only do it for components that are wrapped in some kind of keep alive
  activated() {
    
  },
  deactivated() {
    //this.ready = false;
  },
}

</script>