<!-- converted to the Composition API -->
<template>
  <div class="container">
    <div class="row">
      <div class="col">
        <h1 class="mt-5">Login</h1>
        <hr>
        <form-tag @myevent="submitHandler" name="myform" event="myevent">

          <text-input
            v-model="mail"
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
import {ref, onMounted} from 'vue'
//import {useRouter} from 'vue-router'
import router from '../router/index.js'
import { store } from './store.js'
import Security from './security.js'
import notie from 'notie'

import FormTag from './forms/FormTag.vue'
import TextInput from './forms/TextInput.vue'


export default { 
  name: 'Login', 
  emits: ['error'],
  props: {},
  // declare components here  to use the the lowercase syntax with the hyphens
  // otherwise I could just use the component names as they exist
  components: {
    'form-tag': FormTag,
    'text-input': TextInput,
  },

  setup(props, ctx) {
    //const router = useRouter();
    let mail = ref("");
    let password = ref(""); 

    onMounted(() => {
      //console.log("using new component");
    })

    function submitHandler() {

      const payload = {
        email: mail.value,
        password: password.value,
      }

      //! show error if cannot connected to db 
      fetch(`${store.apiBaseURL}/users/login`, Security.requestOptions(payload))
      .then(response => response.json())
      .then((response) => {
        if (response.error) {
          ctx.emit('error', response.message)
        } else {
          store.token = response.data.token.token;
          
          store.user = {
            id: response.data.user.id,
            first_name:  response.data.user.first_name,
            last_name:  response.data.user.last_name,
            email:  response.data.user.email,
          }

          // save info to cookie 
          let date = new Date();
          let expDays = 1;
          date.setTime(date.getTime() + (expDays * 24 * 60 * 60 * 1000));
          const expires = "expires=" + date.toUTCString();

          // set the cookie
          document.cookie = "_site_data=" 
          + JSON.stringify(response.data)
          + "; " 
          + expires 
          + "; path=/; SameSite=strict; Secure;"

          router.push("/");
        }
      })
    }

    return {
      submitHandler, 
      mail, password, 
    }

  }
}


</script>