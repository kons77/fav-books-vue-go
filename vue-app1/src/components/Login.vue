<template>
  <div class="container">
    <div class="row">
      <div class="col">
        <h1 class="mt-5">Login</h1>
        <hr>
        <form-tag @myevent="submitHandelr" name="myform" event="myevent">

          <text-input
            v-model="email"
            label="Email"
            name="email"
            type="email"
            required="true">
          </text-input>

          <text-input
            v-model="password"
            label="Password"
            name="password"
            type="password"
            required="true">
          </text-input>

          <hr>
          <input type="submit" class="btn btn-primary" value="Login">

        </form-tag>
      </div>
    </div>
  </div>
</template>

<script>
import FormTag from './forms/FormTag.vue'
import TextInput from './forms/TextInput.vue'
import { store } from './store.js'
import router from './../router/index.js'
import notie from 'notie'

export default{
  name: 'login', 
  components: {
    FormTag,
    TextInput,
  },
  data()  {
    return {
      email: "",
      password: "",
      store,
    }
  },
  methods: {
    submitHandelr() {
      console.log("submitHandelr called - success");

      const payload = {
        email: this.email,
        password: this.password,
      }

      const requestOptions = {
        method: 'POST', 
        body: JSON.stringify(payload),
      }

      // show error if cannot connected to db 
      fetch("http://localhost:8081/users/login", requestOptions)
      .then(response => response.json())
      .then((response) => {
        if (response.error) {
          console.log("Error:", response.message);
          notie.alert({
            type: 'error',
            text: response.message,
            // stay: true, 
            // position: 'bottom',
          })
        } else {
          console.log("token:", response.data.token.token);
          store.token = response.data.token.token
          router.push("/");
        }
      })
    }
  },
}
</script>