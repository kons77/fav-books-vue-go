<template>
  <div class="container">
    <div class="row">
      <div class="col">
        <h1 class="mt-3">All Users</h1>
      </div>
    </div>
  </div>
</template>

<script>
import Security from './security.js'
import { store } from './store.js'
import notie from 'notie'

export default {
  data(){
    return {
      users: [],
    }
  },
  beforeMount() {

    //Security.requireToken();

    fetch(`${store.apiBaseURL}/admin/users`, Security.requestOptions(""))
    .then(response => response.json())
    .then((response) => {
      if (response.error) {
        notie.alert({
          type: "error",
          text: response.message,
        })
      } else {
        this.users = response.data.users;
      }
    })
    .catch((error) => {
      notie.alert({
        type: "error",
        text: error,
      })
    });
  }
}
</script>