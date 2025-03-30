<template>
    <div class="container">
      <div class="row">
        <div class="col">
          <h1 class="mt-3">User</h1>
          <hr>
          <form-tag @UserEditEvent="submitHandler" name="userform" event="UserEditEvent">
            <text-input
              v-model="user.first_name"
              type="text"
              required="true"
              label="First Name"
              :value="user.first_name"
              name="first-name"></text-input>
            
              <text-input
              v-model="user.last_name"
              type="text"
              required="true"
              label="Last Name"
              :value="user.last_name"
              name="last-name"></text-input>
            
              <text-input
              v-model="user.email"
              type="text"
              required="true"
              label="Email"
              :value="user.email"
              name="email"></text-input>

            <text-input
              v-if="this.user.id ===0"
              v-model="user.password"
              type="password"
              required="true"
              label="Pasword"
              :value="user.password"
              name="password"></text-input>

            <text-input
              v-else
              v-model="user.password"
              type="password"              
              label="Pasword"
              :value="user.password"
              name="password"></text-input>

            <text-input
              v-if="this.user.id ===0"
              v-model="user.password"
              type="password"
              required="true"
              label="Confirm Pasword"
              :value="user.password"
              name="confirm-password"></text-input>

            <text-input
              v-else
              v-model="user.password"
              type="password"              
              label="Confirm Pasword"
              :value="user.password"
              name="confirm-password"></text-input>

            <hr>

          </form-tag>
        </div>
      </div>
    </div>
  </template>

  
<script>
import Security from './security.js'
import FormTag from './forms/FormTag.vue'
import TextInput from './forms/TextInput.vue'
import notie from 'notie'

export default {
  beforeMount() {
    // Security.requireToken();

    if (parseInt(String(this.$route.params.userId), 10) > 0) {
      // edit an existing user
      // TODO get user from db 
    } else {
      // add a new user 
    }
  }, 
  data() {
    return {
      user: {
        id: 0,
        first_name: "",
        last_name: "",
        email: "",
        password: "",
      }
    }
  },
  components: {
    'form-tag': FormTag,
    'text-input': TextInput,
  }, 
  methods: {
    submitHandler() {

    }
  },
}     
</script>
