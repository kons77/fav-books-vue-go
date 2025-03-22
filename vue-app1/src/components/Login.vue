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
    }
  },
  methods: {
    submitHandelr() {
      console.log("submitHandelr called - success");
    }
  },
}
</script>