<template>
    <div class="container">
      <div class="row">
        <div class="col">
          <h1 class="mt-3">Book Edit</h1>
        </div>
      </div>
    </div>
  </template>


<script>
export default {
  name: 'BookEdit', // need for caching in the include section of keep-alive and the same name must be in router
  data() {
    return{
      store, 
      ready: false,      
      books: {},
    }
  }, 

}
</script>